package services

import (
	"fmt"

	"github.com/kasfulk/orders-backend/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DBConn *gorm.DB
)

func InitDB(config configs.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.Database.DBHost, config.Database.Username, config.Database.Password, config.Database.DBName, config.Database.DBPort)
	DBConn, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Database.SchemaName + ".",
			SingularTable: false,
		},
		Logger: logger.Default.LogMode(logger.LogLevel(config.Database.LogLevel)),
	})
	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Database Connected")
}
