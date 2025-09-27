package stat

type StatHandler struct {
	StatRepository StatRepositoryInterface
}

type StatHandlerDeps struct {
	StatRepository StatRepositoryInterface
}

// func NewStatHandler(deps StatHandlerDeps) StatHandlerInterface {
// 	return &StatHandler{
// 		StatRepository: deps.StatRepository,
// 	}
// }

// func (h *StatHandler) AddClick(linkId uint) {
// 	go h.StatRepository.AddClick(linkId)
// }
