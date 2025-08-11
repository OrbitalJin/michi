package service

type Services struct {
	providersService SPServiceIface
	historyService   HistoryServiceIface
	sessionService   SessionServiceIface
	shortcutService  ShortcutServiceIface
}

func NewServices(
	providersService SPServiceIface,
	historyService HistoryServiceIface,
	sessionService SessionServiceIface,
	shortcutService ShortcutServiceIface,
) *Services {
	return &Services{
		providersService: providersService,
		historyService:   historyService,
		sessionService:   sessionService,
		shortcutService:  shortcutService,
	}
}

func (s *Services) GetProvidersService() SPServiceIface {
	return s.providersService
}

func (s *Services) GetHistoryService() HistoryServiceIface {
	return s.historyService
}

func (s *Services) GetSessionService() SessionServiceIface {
	return s.sessionService
}

func (s *Services) GetShortcutService() ShortcutServiceIface {
	return s.shortcutService
}
