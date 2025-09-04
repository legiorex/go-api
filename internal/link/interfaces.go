package link

type LinkRepositoryInterface interface {
	Create(link *Link) error
	// GetByID(id uint) (*Link, error)
	// GetByHash(hash string) (*Link, error)
	// GetAll() ([]Link, error)
	// Update(id uint, link *Link) error
	// Delete(id uint) error
}
