package service

import (
	"fmt"
	"log"
	"net/url"

	"github.com/OrbitalJin/michi/internal/cache"
	"github.com/OrbitalJin/michi/internal/models"
	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/OrbitalJin/michi/internal/repository"
)

type SPServiceIface interface {
	GetByTag(t string) (*models.SearchProvider, error)
	GetAll() ([]models.SearchProvider, error)
	Collect(v string) ([]models.SearchProvider, error)
	CollectAndRank(v string) (*parser.Result, *models.SearchProvider, error)
	Rank(result *parser.Result) *models.SearchProvider
	Resolve(query string, provider *models.SearchProvider) (*models.SearchProvider, *string, error)
	ResolveAndFallback(query string, provider *models.SearchProvider) (*models.SearchProvider, *string, error)
	ResolveWithFallback(query string) (*models.SearchProvider, *string, error)
	Delete(id int) error
	GetCfg() *Config
}

type SPService struct {
	repo   repository.ProviderRepoIface
	parser *parser.Parser
	cache  *cache.Cache[string, *models.SearchProvider]
	config *Config
}

func NewSearchProviderService(
	p *parser.Parser,
	r repository.ProviderRepoIface,
	config *Config,
) *SPService {

	return &SPService{
		parser: p,
		repo:   r,
		cache:  cache.New[string, *models.SearchProvider](),
		config: config,
	}
}

func (service *SPService) GetByTag(t string) (*models.SearchProvider, error) {
	provider, ok := service.cache.Load(t)

	if ok {
		return provider, nil
	}

	provider, err := service.repo.GetByTag(t)

	if err == nil && provider != nil {
		service.cache.Store(t, provider)
	}

	return provider, err
}

func (service *SPService) Collect(v string) ([]models.SearchProvider, error) {
	result, err := service.parser.Collect(v)

	if err != nil {
		return nil, err
	}

	if len(result.Matches) == 0 {
		return nil, nil
	}

	var sps []models.SearchProvider

	for _, tag := range result.Matches {
		p, err := service.GetByTag(tag)

		if err != nil || p == nil {
			continue
		}

		sps = append(sps, *p)
	}

	return sps, nil
}

func (service *SPService) CollectAndRank(v string) (
	*parser.Result,
	*models.SearchProvider,
	error,
) {
	result, err := service.parser.Collect(v)

	if err != nil {
		return nil, nil, err
	}

	if len(result.Matches) == 0 {
		return result, nil, nil
	}

	best := service.Rank(result)

	return result, best, nil
}

func (service *SPService) Rank(result *parser.Result) *models.SearchProvider {
	if result == nil {
		return nil
	}

	var best *models.SearchProvider
	bestRank := -1

	for _, tag := range result.Matches {
		p, err := service.GetByTag(tag)

		if err != nil || p == nil {
			continue
		}

		if p.Rank > bestRank {
			best = p
			bestRank = p.Rank
		}
	}

	return best
}

func (service *SPService) Resolve(
	query string,
	provider *models.SearchProvider,
) (*models.SearchProvider, *string, error) {

	if provider == nil {
		return nil, nil, fmt.Errorf("provider cannot be nil.")
	}

	// replace %s with the encoded query
	encoded := url.QueryEscape(query)
	url := fmt.Sprintf(provider.URL, encoded)
	return provider, &url, nil
}

func (service *SPService) ResolveAndFallback(
	query string,
	provider *models.SearchProvider,
) (*models.SearchProvider, *string, error) {

	if provider != nil {
		return service.Resolve(query, provider)
	}

	p, err := service.GetByTag(service.config.defaultProvider)

	if err != nil {
		log.Println("no default provider available.")
		return nil, nil, err
	}

	return service.Resolve(query, p)
}

func (service *SPService) ResolveWithFallback(query string) (*models.SearchProvider, *string, error) {
	p, err := service.GetByTag(service.config.defaultProvider)

	if err != nil {
		log.Println("no default provider available.")
		return nil, nil, err
	}

	return service.Resolve(query, p)
}

func (service *SPService) GetCfg() *Config {
	return service.config
}

func (service *SPService) GetAll() ([]models.SearchProvider, error) {
	return service.repo.GetAll()
}

func (service *SPService) Delete(id int) error {
	return service.repo.Delete(id)
}
