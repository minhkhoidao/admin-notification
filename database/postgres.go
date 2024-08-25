package database

import (
	"admin-backend/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type service struct {
	db *gorm.DB
}

type Service interface {
	Close() error
	Migrate() error
	GetDB() *gorm.DB
}

var (
	dbConfig = struct {
		name     string
		password string
		username string
		port     string
		host     string
		schema   string
	}{
		name:     "postgres",
		password: "123456",
		username: "khoidaoo",
		port:     "5432",
		host:     "localhost",
		schema:   "system_admin",
	}
	dbInstance *service
)

func NewConnection() *service {
	if dbInstance != nil {
		return dbInstance
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbConfig.host, dbConfig.username, dbConfig.password, dbConfig.name, dbConfig.port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // Enable color
		},
	)

	dbPostGres, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN: dsn,
		},
	), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: dbConfig.schema + ".",
		},
		Logger: newLogger,
	})

	log.Println("Connecting to database: ", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	dbPostGres.Debug()
	dbInstance = &service{db: dbPostGres}
	return dbInstance
}

func (s *service) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	log.Printf("Disconnected from database: %s", dbConfig.name)
	return sqlDB.Close()
}

func (s *service) Migrate() error {
	err := s.db.AutoMigrate(&models.User{}, &models.Notification{}, &models.Campaigns{})
	if err != nil {
		return fmt.Errorf("failed to auto migrate: %v", err)
	}
	log.Println("Migration completed")
	return nil
}

func (s *service) GetDB() *gorm.DB {
	return dbInstance.db
}
