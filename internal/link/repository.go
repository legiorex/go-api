package link

import (
	"go-api/pkg/db"

	"gorm.io/gorm/clause"
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

func (repo *LinkRepository) Update(link *Link) (*Link, error) {

	result := repo.Database.GetDB().Clauses(clause.Returning{}).Updates(link)

	if result.Error != nil {
		return nil, result.Error

	}

	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {
	result := repo.Database.GetDB().Delete(&Link{}, id)

	if result.Error != nil {
		return result.Error

	}

	return nil
}

func (repo *LinkRepository) GetByID(id uint) (*Link, error) {

	var link Link

	result := repo.Database.GetDB().First(link, id)

	if result.Error != nil {
		return nil, result.Error

	}

	return &link, nil
}

func (repo *LinkRepository) GetAll() ([]Link, error) {

	var links []Link

	result := repo.Database.GetDB().Find(&links)

	if result.Error != nil {
		return nil, result.Error

	}

	return links, nil
}

func (repo *LinkRepository) GetAllPagination(limit, offset int) ([]Link, error) {

	var links []Link

	// repo.Database.GetDB().
	// 	Table("links").
	// 	Where("deleted_at is null").
	// 	Order("id").
	// 	Limit(limit).
	// 	Offset(offset).
	// 	Scan(&links)

	result := repo.Database.GetDB().
		Order("id").
		Limit(limit).
		Offset(offset).
		Find(&links)

	if result.Error != nil {
		return nil, result.Error

	}

	return links, nil
}

func (repo *LinkRepository) GetCount() int64 {
	var count int64
	repo.Database.GetDB().
		Table("links").
		Where("deleted_at is null").
		Count(&count)
	return count
}

// func (repo *LinkRepository) Update(id uint, url string) (*Link, error) {

// 	db := repo.Database.GetDB()

// 	var link Link

// 	finedLink := db.Find(&link, "id = ?", id)

// 	if finedLink.Error != nil {
// 		return nil, finedLink.Error

// 	}

// 	result := db.Model(&link).Update("url", url)

// 	if result.Error != nil {
// 		return nil, result.Error

// 	}

// 	return &link, nil
// }
