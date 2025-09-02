package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: randHash(6),
	}
}

var runes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randHash(n int) string {
	runesSlice := make([]byte, n)

	for range n {
		runesSlice = append(runesSlice, runes[rand.Intn(len(runes))])
	}
	return string(runesSlice)
}
