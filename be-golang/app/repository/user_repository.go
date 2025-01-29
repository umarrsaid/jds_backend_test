package repository

import (
	"be-golang/app/entity"
	"be-golang/config"
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"crypto/sha512"
	"encoding/hex"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/pbkdf2"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	AuthenticateUser(nik, password string) (map[string]interface{}, error)
}

type UserRepositoryImpl struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewUserRepository(db *gorm.DB, cfg *config.Config) UserRepository {
	return &UserRepositoryImpl{db: db, cfg: cfg}
}

func hashPassword(password string, secret string) (string, error) {
	// Hash password menggunakan secret yang diatur
	iterations := 1000
	keyLength := 64
	hash := pbkdf2.Key([]byte(password), []byte(secret), iterations, keyLength, sha512.New)
	return hex.EncodeToString(hash), nil
}

// Fungsi generate password random
func GeneratePassword(length int) (string, error) {
	// Define the character set for the password
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, length)

	for i := range password {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[index.Int64()]
	}

	return string(password), nil
}

func (r *UserRepositoryImpl) CreateUser(user entity.User) (entity.User, error) {
	user.ID = uuid.New()
	password, err := GeneratePassword(6)
	if err != nil {
		return user, err
	}
	user.Password = password
	err = r.db.Create(&user).Error
	fmt.Println("err:", err)
	if err != nil {
		return user, err
	}

	//tampilkan plain password
	user.Password = password
	return user, nil
}

// fungsi untuk cek existing user
func (r *UserRepositoryImpl) AuthenticateUser(nik, password string) (map[string]interface{}, error) {
	var user entity.User

	// Query user by NIK
	if err := r.db.Where("nik = ?", nik).First(&user).Error; err != nil {
		return nil, errors.New("User tidak ditemukan")
	}

	// Verify password
	hashedPassword, err := hashPassword(password, r.cfg.SECRET)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal([]byte(hashedPassword), []byte(user.Password)) {
		return nil, errors.New("invalid password")
	}

	//membuat token jwt
	secret := r.cfg.JWT_KEY_SECRET
	if secret == "" {
		return nil, fmt.Errorf("secret key is not set")
	}

	claims := jwt.MapClaims{
		"exp":  time.Now().Add(24 * time.Hour).Unix(), // Expired: 24 jam
		"data": user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"id":    user.ID,
		"nik":   user.NIK,
		"token": signedToken,
	}

	return response, nil
}
