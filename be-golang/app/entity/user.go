package entity

import (
	"be-golang/config"
	"regexp"
	"time"
	"unicode"

	"github.com/google/uuid"
	// "golang.org/x/crypto/bcrypt"
	"crypto/sha512"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uuid.UUID `gorm:"type:char(36)" json:"id"`                                                                                                     // Primary key
	NIK       string    `gorm:"type:char(16);unique" json:"nik,omitempty" validate:"required,len=16"`                                                        // NIK wajib, panjang 16 karakter, unik
	Password  string    `json:"password,omitempty" validate:"required,min=6"`                                                                                // Minimal 6 karakter
	Role      string    `gorm:"type:enum('superadmin','admin','user');default:'user'" json:"role,omitempty" validate:"required,oneof=superadmin admin user"` // Role hanya superadmin, admin, atau user
	CreatedAt time.Time `json:"-"`                                                                                                                           // Timestamp saat dibuat
	UpdatedAt time.Time `json:"-"`                                                                                                                           // Timestamp saat diperbarui
}

var cfg *config.Config

func init() {
	cfg, _ = config.NewConfig()
}

// Fungsi validasi kustom untuk password
func passwordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	// Regex untuk validasi password
	var passwordRegex = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*]{8,}$`)
	if !passwordRegex.MatchString(password) {
		return false
	}
	hasLowercase := false
	hasUppercase := false
	hasNumber := false
	hasSpecialChar := false
	for _, c := range password {
		if unicode.IsLower(c) {
			hasLowercase = true
		} else if unicode.IsUpper(c) {
			hasUppercase = true
		} else if unicode.IsNumber(c) {
			hasNumber = true
		} else if regexp.MustCompile(`[!@#$%^&*]`).MatchString(string(c)) {
			hasSpecialChar = true
		}
	}
	return hasLowercase && hasUppercase && hasNumber && hasSpecialChar
}

// Fungsi untuk melakukan validasi pada struct User
func (u *User) Validate() error {
	validate := validator.New()

	// Daftarkan validasi kustom password
	validate.RegisterValidation("passwordCheck", passwordValidation)

	// Lakukan validasi
	if err := validate.Struct(u); err != nil {
		return err
	}

	// Tambahkan validasi password jika ada
	if u.Password != "" {
		if err := validate.Var(u.Password, "passwordCheck"); err != nil {
			return err
		}
	}

	return nil
}

// Fungsi hashing password dengan bcrypt menggunakan BYCRIPT_KEY dari config
func hashPassword(password string) (string, error) {
	// Hash password menggunakan secret yang diatur
	iterations := 1000
	keyLength := 64
	hash := pbkdf2.Key([]byte(password), []byte(cfg.SECRET), iterations, keyLength, sha512.New)
	return hex.EncodeToString(hash), nil
}

// Callback BeforeCreate untuk hashing password sebelum membuat user baru
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return hash(u)
}

// Callback BeforeUpdate untuk hashing password sebelum update
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	return hash(u)
}

func hash(u *User) (err error) {
	if u.Password != "" {
		hashedPassword, err := hashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hashedPassword
	}
	return nil
}
