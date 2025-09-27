package stat

type StatRepositoryInterface interface {
	AddClick(link uint)
}

// type StatHandlerInterface interface {
// 	AddClick(link uint)
// }

type StatServiceInterface interface {
	AddClick()
}
