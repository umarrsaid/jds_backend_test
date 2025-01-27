package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"be-golang/app/entity"
	"be-golang/db/seed"
)

func splitURI(uri string) (string, string) {
	// Memisahkan berdasarkan '/'
	parts := strings.Split(uri, "/")
	if len(parts) > 1 {
		// Mengambil bagian setelah '/' dan memisahkan nama database dan query string
		dbInfo := strings.Split(parts[1], "?")
		// Mengembalikan nama database dan bagian DSN tanpa nama database
		return dbInfo[0], parts[0] + "/?" + dbInfo[1]
	}
	// Jika tidak ada nama database, kembalikan string kosong untuk nama database
	return "", uri
}

func NewDB(dbURI string) (*gorm.DB, error) {
	// Koneksi ke MySQL untuk memastikan database ada
	dbName, dsnWithoutDB := splitURI(dbURI)
	sqlDB, err := sql.Open("mysql", dsnWithoutDB)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
	}
	defer sqlDB.Close()

	// Membuat query SQL untuk membuat database jika belum ada
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", dbName)
	// Jalankan query tanpa format string dalam Exec
	_, err = sqlDB.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %v", err)
	}

	log.Printf("Database %s created or already exists", dbName)

	// Koneksi ke database MySQL
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	// Melakukan auto-migration untuk entitas User
	err = db.AutoMigrate(&entity.User{}, &entity.Product{})
	if err != nil {
		return nil, err
	}

	// Seed products
	seed.SeedProducts(db)

	log.Println("Database migrated successfully")

	return db, nil
}
