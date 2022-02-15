package database

import (
	"fmt"

	database "yuno/echo-example/config"
	model "yuno/echo-example/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB
var err error

func generateDsn(pgConfig database.PostgresConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		pgConfig.PostgresConnection.DatabaseHost,
		pgConfig.PostgresConnection.DatabaseUser,
		pgConfig.PostgresConnection.DatabasePassword,
		pgConfig.PostgresConnection.DatabaseName,
		pgConfig.PostgresConnection.DatabasePort,
	)
}

func ConnectPostgres() error {
	dsn := generateDsn(database.GetPostgresConnection())
	instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	instance.AutoMigrate(&model.Country{}, &model.Student{})
	return nil
}

func GetConnection() *gorm.DB {
	return instance
}
