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

func (repo *StatRepository) GetStats(data StatData) ([]StatPeriod, error) {

	var period []StatPeriod

	var periodRequest string

	if data.Period == "month" {
		periodRequest = "YYYY-MM"
	} else {
		periodRequest = "YYYY-MM-DD"
	}

	result := repo.Database.GetDB().
		Table("stats").
		Select("TO_CHAR(date, ?) AS period, sum(clicks)", periodRequest).
		Where("date BETWEEN ? AND ?", data.From, data.To).
		Group("period").
		Order("period").
		Find(&period)

	if result.Error != nil {
		return nil, result.Error

	}
	return period, nil
}
