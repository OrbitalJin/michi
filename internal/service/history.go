package service

import (
	"github.com/OrbitalJin/qmuxr/internal/models"
	"github.com/OrbitalJin/qmuxr/internal/repository"
)

type HistoryService struct {
	repo *repository.HistoryRepo
}

func NewHistoryService(repo *repository.HistoryRepo) *HistoryService {
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

func (service *HistoryService) DeleteEntry(id int) error {
	return service.repo.DeleteEntry(id)
}
