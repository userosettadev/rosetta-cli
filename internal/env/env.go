package env

import (
	"os"
	"sync"
)

var (
	once     sync.Once
	instance *Env
)

const (
	EnvKeyHome   = "ROSETTA_HOME"
	EnvKeyApiKey = "ROSETTA_API_KEY"
)

type Env struct {
	apiKey string
	home   string
}

func GetInstance() *Env {

	once.Do(func() {
		instance = &Env{}
		instance.initialize()
	})

	return instance
}

func (e *Env) initialize() {

	e.home = os.Getenv(EnvKeyHome)
	if e.home == "" {
		e.home = "rosetta-ztdrjhl5kq-uc.a.run.app:443"
	}
	e.apiKey = os.Getenv(EnvKeyApiKey)
}

func (e *Env) GetHome() string {

	return e.home
}

func (e *Env) GetApiKey() string {

	return e.apiKey
}
