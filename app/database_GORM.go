package app

// to do: to try / learn gorm
import (
	"fmt"
	"tobialbertino/portfolio-be/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBGORM(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone==%s`,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.SSL,
		cfg.DB.TIMEZONE,
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	return db
}
