package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"go-api/configs"
	"go-api/internal/link"
	"go-api/pkg/db"
)

func main() {
	count := flag.Int("count", 20, "Количество ссылок для создания")
	flag.Parse()

	config := configs.LoadConfig()
	database := db.NewDb(config)

	seedLinks(database, *count)
}

func seedLinks(database *db.Db, count int) {
	rand.Seed(time.Now().UnixNano())

	domains := []string{
		"example.com", "test.org", "demo.net", "sample.io",
		"placeholder.co", "mock.dev", "fake.site", "dummy.app",
	}

	paths := []string{
		"/home", "/about", "/contact", "/products", "/services",
		"/blog", "/news", "/help", "/support", "/api/v1",
	}

	fmt.Printf("Создаю %d тестовых ссылок...\n", count)

	for i := 0; i < count; i++ {
		// Генерируем случайный URL
		domain := domains[rand.Intn(len(domains))]
		path := paths[rand.Intn(len(paths))]
		url := fmt.Sprintf("https://%s%s?id=%d", domain, path, rand.Intn(1000))

		// Создаем ссылку
		newLink := link.NewLink(url)

		// Проверяем уникальность хеша
		for {
			var existingLink link.Link
			result := database.GetDB().Where("hash = ?", newLink.Hash).First(&existingLink)
			if result.Error != nil {
				break
			}
			newLink.Hash = newLink.GenerateHash()
		}

		// Сохраняем
		if result := database.GetDB().Create(newLink); result.Error != nil {
			fmt.Printf("Ошибка: %v\n", result.Error)
		} else {
			fmt.Printf("%d. %s -> %s\n", i+1, newLink.Hash, newLink.Url)
		}
	}

	fmt.Println("Готово!")
}
