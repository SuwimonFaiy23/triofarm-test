package db

import (
	"fmt"
	"log"
	"net/url"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DBConfig struct {
	Server   string
	Port     int
	User     string
	Password string
	Name     string
}

// load config database
func LoadDBConfig() DBConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	return DBConfig{
		Server:   viper.GetString("database.server"),
		Port:     viper.GetInt("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Name:     viper.GetString("database.name"),
	}
}

// connect database
func Connect() *gorm.DB {
	dbConfig := LoadDBConfig()
	passwordEscaped := url.QueryEscape(dbConfig.Password)
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=true",
		dbConfig.User, passwordEscaped, dbConfig.Server, dbConfig.Port, dbConfig.Name)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open connection: %v", err)
	}

	// ดึง sql.DB ออกมา
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	// db.Migrator().DropTable(&model.Item{})
	// db.AutoMigrate(&model.Menu{})
	// db.AutoMigrate(&model.Item{})

	log.Println("✅ Connected to Azure SQL successfully!")

	return db
}
