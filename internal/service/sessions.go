package service

import (
	"github.com/OrbitalJin/qmuxr/internal/cache"
	"github.com/OrbitalJin/qmuxr/internal/models"
	"github.com/OrbitalJin/qmuxr/internal/repository"
)

type SessionServiceIface interface {
	GetCfg() *Config
	Insert(session *models.Session) error
	GetFromAlias(alias string) (*models.Session, error)
	GetAll() ([]models.Session, error)
	Update(session *models.Session) error
	Delete(id int) error
}

type SessionService struct {
	repo  repository.SessionsRepoIface
	cache *cache.Cache[string, *models.Session]
}

func NewSessionService(repo repository.SessionsRepoIface) *SessionService {
	return &SessionService{
		repo:  repo,
		cache: cache.New[string, *models.Session](),
	}
}

func (s *SessionService) Insert(session *models.Session) error {

	err := s.repo.Insert(session)

	if err != nil {
		return err
	}

	s.cache.Store(session.Alias, session)

	return nil
}

func (s *SessionService) GetFromAlias(alias string) (*models.Session, error) {
	shortcut, ok := s.cache.Load(alias)

	if ok {
		return shortcut, nil
	}

	shortcut, err := s.repo.GetFromAlias(alias)

	if err != nil {
		return nil, err
	}

	s.cache.Store(alias, shortcut)

	return shortcut, nil
}

func (s *SessionService) GetAll() ([]models.Session, error) {
	return s.repo.GetAll()
}

func (s *SessionService) Update(session *models.Session) error {
	return s.repo.Update(session)
}
