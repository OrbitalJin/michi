package service

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/OrbitalJin/pow/internal/parser"
	"github.com/OrbitalJin/pow/internal/store"
)

type ProviderService struct {
	parser *parser.Parser
	store  *store.Store
}

func NewProviderService(p *parser.Parser, s *store.Store) *ProviderService {
	return &ProviderService{
		parser: p,
		store:  s,
	}
}

func (svc *ProviderService) GetByTag(tag string) (*store.SearchProvider, error) {
	return svc.store.GetProviderByTag(tag)
}

func (svc *ProviderService) CollectAll(value string) (*[]*store.SearchProvider, error) {
	result, err := svc.parser.Collect(value)

	if err != nil {
		return nil, err
	}

	if len(result.Matches) == 0 {
		return nil, nil
	}

	var sps []*store.SearchProvider

	for _, tag := range result.Matches {
		p, err := svc.GetByTag(tag)

		if err != nil {
			continue
		}

		sps = append(sps, p)
	}

	return &sps, nil
}

func (svc *ProviderService) CollectAndRank(value string) (*parser.Result, *store.SearchProvider, error) {
	result, err := svc.parser.Collect(value)

	if err != nil {
		return nil, nil, err
	}

	if len(result.Matches) == 0 {
		return result, nil, nil
	}

	var best *store.SearchProvider
	bestRank := -1

	for _, tag := range result.Matches {
		p, err := svc.GetByTag(tag)

		if err != nil {
			continue
		}

		if p.Rank > bestRank {
			best = p
			bestRank = p.Rank
		}
	}

	return result, best, nil
}

func (svc *ProviderService) Resolve(query string, provider *store.SearchProvider) (*string, error) {
	if provider == nil {
		log.Println("provider cannot be nil.")
		return nil, fmt.Errorf("provider cannot be nil.")
	}

	encoded := url.QueryEscape(query)
	result := strings.Replace(provider.URL, "{{{s}}}", encoded, 1)
	return &result, nil
}


func (svc *ProviderService) ResolveWithFallback(query string, provider *store.SearchProvider) (*string, error) {
	if provider != nil {
		return svc.Resolve(query, provider)
	}

	log.Println("no provider provided, falling back to default.")

	defaultProvider := svc.store.GetCfg().GetDefaultProvider()
	p, err := svc.store.GetProviderByTag(defaultProvider)

	if err != nil {
		log.Println("no default provider available.")
		return nil, err
	}

	return svc.Resolve(query, p)
}

