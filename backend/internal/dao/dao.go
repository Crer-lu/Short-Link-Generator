package dao

import (
	"context"
	"go-svc-tpl/internal/dao/model"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBMS struct {
	*gorm.DB
}

var (
	db *gorm.DB
)

var DB = func(ctx context.Context) *DBMS {
	return &DBMS{db.WithContext(ctx)}
}

// >>>>>>>>>>>> init >>>>>>>>>>>>

type DBCfg struct {
	DSN string
}

func InitDB() {
	var cfg DBCfg
	err := viper.Sub("Database").UnmarshalExact(&cfg)
	if err != nil {
		logrus.Fatal(err)
	}

	db, err = gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	// Uncomment this if you want to use auto migrate

	if err := db.AutoMigrate(&model.User{}, &model.Link{}); err != nil {
		logrus.Fatal(err)
	}

	if viper.GetString("App.RunLevel") == "debug" {
		db = db.Debug()
	}

}
