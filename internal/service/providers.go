package service

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/OrbitalJin/qmuxr/internal/cache"
	"github.com/OrbitalJin/qmuxr/internal/models"
	"github.com/OrbitalJin/qmuxr/internal/parser"
	"github.com/OrbitalJin/qmuxr/internal/repository"
)

type SearchProviderService struct {
	parser *parser.Parser
	repo   *repository.ProviderRepo
	cache  *cache.Cache[string, *models.SearchProvider]
	config *Config
}

func NewSearchProviderService(p *parser.Parser, r *repository.ProviderRepo, config *Config) *SearchProviderService {
	return &SearchProviderService{
		parser: p,
		repo:   r,
		cache:  cache.New[string, *models.SearchProvider](),
		config: config,
	}
}

func (service *SearchProviderService) GetByTag(t string) (*models.SearchProvider, error) {
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

func (service *SearchProviderService) Collect(v string) (*[]*models.SearchProvider, error) {
	result, err := service.parser.Collect(v)

	if err != nil {
		return nil, err
	}

	if len(result.Matches) == 0 {
		return nil, nil
	}

	var sps []*models.SearchProvider

	for _, tag := range result.Matches {
		p, err := service.GetByTag(tag)

		if err != nil {
			continue
		}

		sps = append(sps, p)
	}

	return &sps, nil
}

func (service *SearchProviderService) CollectAndRank(v string) (*parser.Result, *models.SearchProvider, error) {
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

func (service *SearchProviderService) Rank(result *parser.Result) *models.SearchProvider {
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

func (service *SearchProviderService) Resolve(query string, provider *models.SearchProvider) (*models.SearchProvider, *string, error) {
	if provider == nil {
		return nil, nil, fmt.Errorf("provider cannot be nil.")
	}

	encoded := url.QueryEscape(query)
	result := strings.Replace(provider.URL, "{{{s}}}", encoded, 1)
	return provider, &result, nil
}

func (service *SearchProviderService) ResolveWithFallback(query string, provider *models.SearchProvider) (*models.SearchProvider, *string, error) {
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

func (service *SearchProviderService) GetCfg() *Config {
	return service.config
}
