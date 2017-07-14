package main

import (
	"github.com/adnaan/talks/dependency_injection_july2017/user"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
)

//S1 OMIT
type AppHandler struct {
	UserService user.Service
	// ... more services.
}

func main() {

	msession := &mgo.Session{} //dummy
	defaultConfig := user.Config{Msession: msession}
	redisPool := &redis.Pool{} //dummy

	myUserService := user.NewService(defaultConfig)
	myPrivilegedUserService := user.NewService(defaultConfig, user.PrivilegedMode(redisPool))

	appHandler := &AppHandler{UserService: myUserService}
	appHandler2 := &AppHandler{UserService: myPrivilegedUserService}

	// If necessaryimplement functional config options for AppHandler too
	// ...
	// register appHandler to the http server.

}

//S2 OMIT
