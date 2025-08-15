package service

import (
	"github.com/OrbitalJin/michi/internal/cache"
	"github.com/OrbitalJin/michi/internal/models"
	"github.com/OrbitalJin/michi/internal/repository"
)

type ShortcutServiceIface interface {
	Insert(shortcut *models.Shortcut) error
	GetFromAlias(alias string) (*models.Shortcut, error)
	GetAll() ([]models.Shortcut, error)
	Delete(id int) error
}

type ShortcutService struct {
	repo  repository.ShortcutsRepoIface
	cache *cache.Cache[string, *models.Shortcut]
}

func NewShortcutService(repo repository.ShortcutsRepoIface) *ShortcutService {
	return &ShortcutService{
		repo:  repo,
		cache: cache.New[string, *models.Shortcut](),
	}
}

func (service *ShortcutService) Insert(shortcut *models.Shortcut) error {
	if err := service.repo.Insert(shortcut); err != nil {
		return err
	}

	service.cache.Store(shortcut.Alias, shortcut)
	return nil
}

func (service *ShortcutService) GetFromAlias(alias string) (*models.Shortcut, error) {
	shortcut, ok := service.cache.Load(alias)

	if ok {
		return shortcut, nil
	}

	shortcut, err := service.repo.GetFromAlias(alias)

	if err != nil {
		return nil, err
	}

	service.cache.Store(alias, shortcut)

	return shortcut, nil
}

func (service *ShortcutService) GetAll() ([]models.Shortcut, error) {
	return service.repo.GetAll()
}

func (service *ShortcutService) Delete(id int) error {
	return service.repo.Delete(id)
}
