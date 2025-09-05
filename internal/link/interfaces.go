package link

type LinkRepositoryInterface interface {
	Create(link *Link) (*Link, error)
	GetByHash(hash string) (*Link, error)
	Update(id uint, url string) (*Link, error)
	// GetByID(id uint) (*Link, error)
	// GetAll() ([]Link, error)
	// Delete(id uint) error
}
