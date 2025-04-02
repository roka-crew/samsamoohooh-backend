package postgres

import (
	"fmt"
	"time"

	"github.com/roka-crew/samsamoohooh-backend/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgres struct {
	*gorm.DB
}

func New(
	config *config.Config,
) (*Postgres, error) {
	format := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"

	dsn := fmt.Sprintf(format,
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.DBname,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent),
		NowFunc:        func() time.Time { return time.Now().UTC() },
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(config.Postgres.Options.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Postgres.Options.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(config.Postgres.Options.ConnMaxLifetime)

	return &Postgres{DB: db}, nil
}
