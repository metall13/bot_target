package config

import (
	"fmt"
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
	"os"
)

type Tg struct {
	Api_key string `default:"1155763164:AAFfDGqNcjMCSntZiorkla8mYlgiifr6zHk" env:"TOKEN" flag:"api_key"`
}

type Postgres struct {
	UrlDB      string `default:"postgres://postgres:postgres@localhost/postgres?sslmode=disable" env:"URlDB" `
	URlDBDEBAG string `default:"postgres://postgres:postgres@192.168.0.200:5400/postgres?sslmode=disable" env:"URlDBDEBAG" `
}

type Config struct {
	Tg       `env:"TG"`
	Postgres `env:"DB"`
	Debag    bool `default:"true" env:"DEBAG"`
}

var cfg *Config

func ConfigLoader(cfg interface{}) *aconfig.Loader {
	cfgPath := ""
	if len(os.Args) > 2 {
		cfgPath = os.Args[2]
	}
	files := []string{cfgPath, ".env"}

	return aconfig.LoaderFor(cfg, aconfig.Config{
		AllowUnknownEnvs: true,
		Files:            files,
		FileDecoders: map[string]aconfig.FileDecoder{
			".env": aconfigdotenv.New(),
		},
	})
}
func Load(cfg *Config) error {
	loader := ConfigLoader(cfg)
	if err := loader.Load(); err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	return nil
}
