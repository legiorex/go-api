package link

import (
	"go-api/pkg/db"
)

type LinkRepository struct {
	Database db.DatabaseInterface
}

func NewLinkRepository(database db.DatabaseInterface) LinkRepositoryInterface {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {

	result := repo.Database.GetDB().Create(link)

	if result.Error != nil {
		return nil, result.Error

	}

	return link, nil

}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {

	var link Link

	result := repo.Database.GetDB().Take(&link, "hash = ?", hash)

	if result.Error != nil {
		return nil, result.Error

	}

	return &link, nil

}
