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

func (service *HistoryService) Insert(entry *models.SearchHistoryEntry) error {
	return service.repo.Insert(entry)
}
