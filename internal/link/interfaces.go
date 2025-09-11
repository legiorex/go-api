package link

type LinkRepositoryInterface interface {
	Create(link *Link) (*Link, error)
	GetByHash(hash string) (*Link, error)
	Update(link *Link) (*Link, error)
	Delete(id uint) error
	GetByID(id uint) (*Link, error)
	GetAll() ([]Link, error)
	GetAllPagination(limit, offset int) ([]Link, error)
	GetCount() int64
}
