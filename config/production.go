// +build heroku

package config

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gorelic"
	"os"
)

const DEBUG = false
const APPNAME = "BOILERPLATE"

func Initialize(m *martini.ClassicMartini) {
	fmt.Println("Running in production environment")

	gorelic.InitNewrelicAgent(os.Getenv("NEW_RELIC_LICENSE_KEY"), APPNAME, true)
	m.Use(gorelic.Handler)
}
