package link

import (
	"go-api/internal/stat"
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET Null;"`
}

func NewLink(url string) *Link {

	link := &Link{
		Url: url,
	}

	link.Hash = link.GenerateHash()

	return link

}

func (link *Link) GenerateHash() string {
	return randHash(6)
}

var runes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randHash(n int) string {
	runesSlice := make([]byte, 0, n)

	for range n {
		runesSlice = append(runesSlice, runes[rand.Intn(len(runes))])
	}
	return string(runesSlice)
}
