package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api/configs"
	"go-api/internal/auth"
	"go-api/pkg/jwt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginSuccess(t *testing.T) {

	ts := httptest.NewServer(App())
	defer ts.Close()

	// w := httptest.NewRecorder()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "s@d.ru",
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

	jwtTest := jwt.NewJWT(config.Auth.Secret)

	token, _ := jwtTest.Create(&jwt.JWTData{Email: "s@d.ru"})

	dataBody, _ := io.ReadAll(res.Body)

	var userResponse auth.LoginPayload

	err = json.Unmarshal(dataBody, &userResponse)

	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Printf("userResponse: %+v\n", userResponse)

	if userResponse.Token != token {
		t.Fatalf("Expected token %s got %s", token, userResponse.Token)
	}
}
