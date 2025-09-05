package link

type LinkRepositoryInterface interface {
	Create(link *Link) (*Link, error)
	GetByHash(hash string) (*Link, error)
	// GetByID(id uint) (*Link, error)
	// GetAll() ([]Link, error)
	// Update(id uint, link *Link) error
	// Delete(id uint) error
}
