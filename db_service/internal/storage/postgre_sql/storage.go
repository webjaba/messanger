package postgresql

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDatabase struct {
	Port     string `env:"PORT" env-default:"5432"`
	Host     string `env:"HOST" env-default:"localhost"`
	Name     string `env:"NAME" env-default:"messanger_db"`
	User     string `env:"USER" env-default:"postgres"`
	Password string `env:"PASSWORD"`
}

func ConnectDB() *gorm.DB {

	DBConfig := ConfigDatabase{}

	err := cleanenv.ReadConfig(".env", &DBConfig)

	if err != nil {
		panic(fmt.Sprintf("Unable to load db_config: %v", err.Error()))
	}

	if DBConfig.Password == "" {
		panic("password must be specified")
	}

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		DBConfig.Host,
		DBConfig.User,
		DBConfig.Password,
		DBConfig.Name,
		DBConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
