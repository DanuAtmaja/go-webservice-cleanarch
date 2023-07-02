package infrastructure

import (
	"cleanGoProjectCRUD/exception"
	"github.com/joho/godotenv"
	"os"
)

type IConfig interface {
	Get(key string) string
}

type config struct {
}

func (config *config) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) IConfig {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &config{}
}
