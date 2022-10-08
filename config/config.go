package config

import (
	"fmt"
	"log"

	"github.com/devdinu/simple-api/dbi"
	"github.com/kelseyhightower/envconfig"
)

var app Application

type Server struct {
	Port   int    `envconfig:"PORT"`
	Host   string `required:"true"`
	Scheme string `default:"http"`
}

func (s Server) Address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func AppAddress() string {
	return fmt.Sprintf("%s:%d", app.server.Host, app.server.Port)
}

type Application struct {
	server Server
	DB     dbi.Config
}

func MustLoad() Application {
	var errs []error
	if err := envconfig.Process("", &app.server); err != nil {
		errs = append(errs, err)
	}
	if err := envconfig.Process("DB", &app.DB); err != nil {
		errs = append(errs, err)
	}
	if len(errs) != 0 {
		log.Fatalf("Error loading configuration: %v", errs)
	}
	return app
}
