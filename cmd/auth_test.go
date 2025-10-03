package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api/configs"
	"go-api/internal/auth"
	"go-api/internal/user"
	"go-api/pkg/jwt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv(("DSN"))), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

var userEmail = "s@d2.ru"

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    userEmail,
		Password: "$2a$10$l4sm.W4LRQB.hzdtqSHcCehA.7kZxCabvdmdk3dCyxmYH3IeJZFca",
		Name:     "Jon",
	})
}

func TestLoginSuccess(t *testing.T) {
	db := initDb()
	initData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    userEmail,
		Password: "12345",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))

	if err != nil {
		t.Fatal()
	}

	if res.StatusCode != 200 {
		t.Fatalf("Expected 200 got %d", res.StatusCode)
	}

	config := configs.LoadConfig()

	jwtInit := jwt.NewJWT(config.Auth.Secret)

	token, _ := jwtInit.Create(&jwt.JWTData{Email: userEmail})

	dataBody, _ := io.ReadAll(res.Body)

	var userResponse auth.LoginPayload

	json.Unmarshal(dataBody, &userResponse)

	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Printf("userResponse: %+v\n", userResponse)

	if userResponse.Token != token {
		t.Fatalf("Expected token %s got %s", token, userResponse.Token)
	}
}

func TestLoginFail(t *testing.T) {

	db := initDb()
	initData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    userEmail,
		Password: "1234",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))

	if err != nil {
		t.Fatal()
	}

	if res.StatusCode != 400 {
		t.Fatalf("Expected 400 got %d", res.StatusCode)
	}

}
