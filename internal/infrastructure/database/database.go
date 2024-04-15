package database

import (
	"fmt"
	"go-challenge/configs"
	"go-challenge/internal/infrastructure/dto"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbMy *gorm.DB
	err  error
)

func Init() {
	openMySQL()
}

func Close() {
	my, _ := dbMy.DB()
	my.Close()
}

func openMySQL() {
	cnf := configs.GetConfig()

	database := cnf.MySQLDatabase

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cnf.MySQLUser,
		cnf.MySQLPassword,
		cnf.MySQLHost,
		cnf.MySQLPort,
		database,
	)

	dbMy, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return dbMy
}

func MigrateExecution() error {
	err := dbMy.AutoMigrate(dto.Location{}, dto.EVSE{})
	if err != nil {
		return fmt.Errorf("migration execution failed: %w", err)
	}
	fmt.Println("migration was successfully performed")
	return nil
}
