package migrations

import (
	"fmt"

	"go-relation/relasi-gorm/databases"
	"go-relation/relasi-gorm/models"

	"golang.org/x/crypto/bcrypt"
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

	user := new(models.UserStore)

	user.Name = "Superadmin"
	user.Email = "superadmin@gmail.com"
	user.Password = "password"
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic("Create superadmin error!")
	}

	user.Password = string(hashedPassword)

	databases.DB.Create(&user)

	fmt.Println("migrated.")
}
