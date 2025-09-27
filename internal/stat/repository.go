package stat

import (
	"go-api/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	Database db.DatabaseInterface
}

func NewStatRepository(database db.DatabaseInterface) StatRepositoryInterface {
	return &StatRepository{
		Database: database,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat

	currentDate := datatypes.Date(time.Now())

	repo.Database.GetDB().Find(
		&stat,
		"link_id = ? and date = ?",
		linkId,
		currentDate,
	)

	if stat.ID == 0 {
		repo.Database.GetDB().Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks += 1
		repo.Database.GetDB().Save(&stat)
	}

}
