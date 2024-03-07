package config

import (
	"fmt"
	"github.com/yzaimoglu/mitocho/data/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"os"
	"strings"
)

type Database struct {
	Gorm *gorm.DB
}

func GetDatabaseURL() string {
	return os.Getenv("MYSQL_URL")
}

func convertDSN(originalDSN string) (string, error) {
	u, err := url.Parse(originalDSN)
	if err != nil {
		return "", err
	}

	user := u.User.Username()
	pass, _ := u.User.Password()

	host := u.Host
	dbName := strings.TrimPrefix(u.Path, "/")

	newDSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, dbName)
	return newDSN, nil
}

func NewDB() *Database {
	db := &Database{}
	db.Init()
	return db
}

func (db *Database) Init() {
	db.Connect()
	if err := db.Migrate(); err != nil {
		panic("Failed to migrate database")
	}
}

func (db *Database) Connect() {
	dsn, err := convertDSN(GetDatabaseURL())
	if err != nil {
		panic("Could not convert dsn to mysql format")
	}
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.Gorm = database
	fmt.Println("Connected to database")
}

func (db *Database) Close() error {
	sqlDB, err := db.Gorm.DB()
	if err != nil {
		panic("Failed to close database")
	}
	return sqlDB.Close()
}

func (db *Database) Migrate() error {
	models := []interface{}{
		&types.Role{},
		&types.User{},
		&types.UserRole{},
		&types.Setting{},
		&types.Site{},
		&types.SiteUser{},
		&types.SiteRole{},
	}

	for _, model := range models {
		if err := db.Gorm.AutoMigrate(model); err != nil {
			return err
		}
	}
	return nil
}
