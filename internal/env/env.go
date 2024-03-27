package env

import (
	"os"
	"sync"
)

var (
	once     sync.Once
	instance *Env
)

const EnvHomeKey = "ROSETTA_HOME"
const EnvTenantKey = "ROSETTA_TENANT"

type Env struct {
	tenant string
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

	e.home = os.Getenv(EnvHomeKey)
	if e.home == "" {
		e.home = "rosetta-ztdrjhl5kq-uc.a.run.app:443"
	}
	e.tenant = os.Getenv(EnvTenantKey)
}

func (e *Env) GetHome() string {

	return e.home
}

func (e *Env) GetTenant() string {

	return e.tenant
}
