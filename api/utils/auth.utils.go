package utils

import (
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var secretHash = jwt.NewHS256([]byte("eyJhbGciOiJIUzI1NiJ9.ew0KICAic3ViIjogIjEyMzQ1Njc4OTAiLA0KICAibmFtZSI6ICJUaW1lIENhc3NldHRlIiwNCiAgImlhdCI6IDE1MTYyMzkwMjINCn0.7ZReMaMMSjM5nMHHYlf3CM5rH6nKYE34MGb914VF5XU"))

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}
	return string(bytes)
}

type JwtPayload struct {
	jwt.Payload
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

func GenerateToken(email string) (string, error) {

	now := time.Now()
	payload := JwtPayload{
		Payload: jwt.Payload{
			Issuer:         "timecassette",
			Subject:        "timecassette",
			Audience:       jwt.Audience{"https://timecassette.com", "https://timecassette.io"},
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          uuid.Must(uuid.NewRandom()).String(),
		},
		Email: email,
	}

	token, err := jwt.Sign(payload, secretHash)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func VerifyToken(token string, payload JwtPayload) (jwt.Header, error) {
	hd, err := jwt.Verify([]byte(token), secretHash, payload)
	return hd, err
}
