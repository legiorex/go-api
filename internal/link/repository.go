package link

import "go-api/pkg/db"

type LinkRepository struct {
	Database db.DatabaseInterface
}

func NewLinkRepository(database db.DatabaseInterface) LinkRepositoryInterface {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) error {
	return repo.Database.GetDB().Create(link).Error
}
