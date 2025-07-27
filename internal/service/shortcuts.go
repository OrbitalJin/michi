package service

import (
	"github.com/OrbitalJin/qmuxr/internal/models"
	"github.com/OrbitalJin/qmuxr/internal/repository"
)

type ShortcutServiceIface interface {
	Insert(shortcut *models.Shortcut) error
	GetFromAlias(alias string) (*models.Shortcut, error)
	Delete(id int) error
}

type ShortcutService struct {
	repo repository.ShortcutsRepoIface
}

func NewShortcutService(repo repository.ShortcutsRepoIface) *ShortcutService {
	return &ShortcutService{
		repo: repo,
	}
}

func (service *ShortcutService) Insert(shortcut *models.Shortcut) error {
	return service.repo.Insert(shortcut)
}

func (service *ShortcutService) GetFromAlias(alias string) (*models.Shortcut, error) {
	return service.repo.GetFromAlias(alias)
}

func (service *ShortcutService) Delete(id int) error {
	return service.repo.Delete(id)
}
