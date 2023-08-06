package migrations

import (
	"fmt"

	"go-relation/relasi-gorm/databases"
	"go-relation/relasi-gorm/models"
)

func Migration() {
	err := databases.DB.AutoMigrate(
		&models.User{},
		&models.Locker{},
		&models.Post{},
		&models.Tag{},
	)

	if err != nil {
		fmt.Println("can't running migration")
	}

	fmt.Println("migrated.")
}
