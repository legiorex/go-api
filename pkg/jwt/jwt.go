package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Email string
}

type JWT struct {
	secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		secret: secret,
	}
}

func (j *JWT) Create(data *JWTData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})

	tokenString, err := token.SignedString([]byte(j.secret))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) Parse(token string) (bool, *JWTData) {

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

		return []byte(j.secret), nil
	})

	if err != nil {
		return false, nil
	}
	email := t.Claims.(jwt.MapClaims)["email"]

	return t.Valid, &JWTData{Email: email.(string)}

}
