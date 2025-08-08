package auth

import (
	"errors"
	"latihan/config"
	"latihan/internal/user"
	"latihan/pkg/hash"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Login(email, password string) (string, error) {
	var u user.User
	if err := config.DB.Where("email = ? AND password is not null", email).First(&u).Error; err != nil {
		return "", errors.New("user tidak ditemukan")
	}

	if err := hash.Validate(password, u.Password); !err {
		return "", errors.New("password salah")
	}

	claims := jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
