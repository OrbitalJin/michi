package service

import (
	"github.com/OrbitalJin/michi/internal/models"
	"github.com/OrbitalJin/michi/internal/repository"
)

type HistoryServiceIface interface {
	Insert(entry *models.SearchHistoryEvent) error
	GetRecentHistory(limit int) ([]models.SearchHistoryEvent, error)
	GetAllHistory() ([]models.SearchHistoryEvent, error)
	DeleteEntry(id int) error
}

type HistoryService struct {
	repo repository.HistoryRepoIface
}

func NewHistoryService(repo repository.HistoryRepoIface) *HistoryService {
	return &HistoryService{
		repo: repo,
	}
}

func (service *HistoryService) Insert(entry *models.SearchHistoryEvent) error {
	return service.repo.Insert(entry)
}

func (service *HistoryService) GetRecentHistory(limit int) ([]models.SearchHistoryEvent, error) {
	return service.repo.GetRecentHistory(limit)
}

func (service *HistoryService) GetAllHistory() ([]models.SearchHistoryEvent, error) {
	return service.repo.GetAllHistory()
}

func (service *HistoryService) DeleteEntry(id int) error {
	return service.repo.DeleteEntry(id)
}
