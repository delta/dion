package repository

import (
	"fmt"

	"delta.nitt.edu/dion/config"
	"delta.nitt.edu/dion/models"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Params struct {
	fx.In
	Conf   *config.Config
	Logger *zap.Logger
}

func New(p Params) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		p.Conf.Db.Host,
		p.Conf.Db.User,
		p.Conf.Db.Password,
		p.Conf.Db.DbName,
		p.Conf.Db.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.Project{})
	return err
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)
