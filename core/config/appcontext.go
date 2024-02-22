package config

import (
	"fmt"
	"log"
	"task-test/swagger"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type appContext struct {
	db               *gorm.DB
	cfg              *Configuration
	requestValidator *validator.Validate
}

var appCtx appContext

func Init() error {

	cfg, err := InitConfig()
	if err != nil {
		return err
	}

	db, err := ConnectDataBase(cfg.Postgres)
	if err != nil {
		return err
	}

	swagger.SwaggerInfo.Host = cfg.Swagger.SwaggerHost
	swagger.SwaggerInfo.Schemes = []string{"http", "https"}

	host := fmt.Sprintf("\x1b[1;34mhttp://%s/task-test/swagger/index.html\x1b[0m\n", cfg.ServiceHost)
	log.Println(host)

	appCtx = appContext{
		db:  db,
		cfg: cfg,
	}

	return nil
}

func RequestValidator() *validator.Validate {
	return appCtx.requestValidator
}

func DB() *gorm.DB {
	return appCtx.db
}

func Config() Configuration {
	return *appCtx.cfg
}
