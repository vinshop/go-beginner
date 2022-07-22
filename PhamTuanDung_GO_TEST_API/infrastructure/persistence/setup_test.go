package persistence

import (
	"fmt"
	"github.com/dungbk10t/test_api/domain/entity"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DBConn() (*gorm.DB, error) {
	if _, err := os.Stat("./../../.env"); !os.IsNotExist(err) {
		var err error
		err = godotenv.Load(os.ExpandEnv("./../../.env"))
		if err != nil {
			log.Fatalf("Error getting env %v\n", err)
		}
		return LocalDatabase()
	}
	return CIBuild()
}

//Circle CI DB
func CIBuild() (*gorm.DB, error) {
	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "dev", "123456789aA@", "127.0.0.1", "3306", "user_04")
	conn, err := gorm.Open("mysql", DBURL)
	if err != nil {
		log.Fatal("This is the error:", err)
	}
	return conn, err
}

//local DB
func LocalDatabase() (*gorm.DB, error) {
	err_env := godotenv.Load(".env")
	if err_env != nil {
		log.Fatalf("Error loading .env file")
	}
	dbdriver := os.Getenv("TEST_DB_DRIVER")
	dbhost := os.Getenv("TEST_DB_HOST")
	dbpassword := os.Getenv("TEST_DB_PASSWORD")
	dbuser := os.Getenv("TEST_DB_USER")
	dbname := os.Getenv("TEST_DB_NAME")
	dbport := os.Getenv("TEST_DB_PORT")

	//dbdriver := "mysql"
	//dbhost := "127.0.0.1"
	//dbpassword := "123456789aA@"
	//dbuser := "dev"
	//dbname := "user_04"
	//dbport := "3306"

	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbport, dbname)
	conn, err := gorm.Open(dbdriver, DBURL)
	if err != nil {
		return nil, err
	} else {
		log.Println("CONNECT TO: ", dbdriver)
	}

	err = conn.DropTableIfExists(&entity.User{}).Error
	if err != nil {
		return nil, err
	}
	err = conn.Debug().AutoMigrate(entity.User{}).Error
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func seedUser(db *gorm.DB) (*entity.User, error) {
	user := &entity.User{
		ID:        100,
		Name:      "dung100",
		Email:     "dung100@gmail.com",
		Password:  "123456",
		DeletedAt: nil,
	}
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func seedUsers(db *gorm.DB) ([]entity.User, error) {
	users := []entity.User{
		{
			ID:        1,
			Name:      "dung01",
			Email:     "dung01@example.com",
			Password:  "123456",
			DeletedAt: nil,
		},
		{
			ID:        2,
			Name:      "dung02",
			Email:     "dung02@example.com",
			Password:  "123456",
			DeletedAt: nil,
		},
		{
			ID:        3,
			Name:      "dung03",
			Email:     "dung03@example.com",
			Password:  "123456",
			DeletedAt: nil,
		},
	}
	for _, v := range users {
		err := db.Create(&v).Error
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}
