package stat

import "go-api/pkg/event"

type ServiceStatDeps struct {
	EventBus       event.EventBusInterface
	StatRepository StatRepositoryInterface
}

type StatService struct {
	EventBus       event.EventBusInterface
	StatRepository StatRepositoryInterface
}

func NewServiceStat(deps *ServiceStatDeps) StatServiceInterface {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (s *StatService) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.TypeLinkVisitedEvent {
			// return msg.Data.(uint)
			s.StatRepository.AddClick(msg.Data.(uint))
		}
	}
}
