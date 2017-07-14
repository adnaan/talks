//S1 OMIT
package user

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	mgo "gopkg.in/mgo.v2"
)

// Service interface
type Service interface {
	// invoked from a http request and returns the role of the user.
	HandleGetRoleRequest(uid string) string
	HandleGetAccessKey(uid string) (string, error)
}

// Config for the service
type Config struct {
	RedisPool *redis.Pool
	Msession  *mgo.Session
}

//S2 OMIT
// configurable service which implements the Service interface
type service struct {
	c *Config
}

func (s *service) HandleGetRoleRequest(uid string) string {
	return getRole(s.c.Msession.Copy(), uid)
}

func (s *service) HandleGetAccessKey(uid string) (string, error) {
	if s.c.RedisPool == nil {
		return "", fmt.Errorf("redis pool is not initalized")
	}
	return getAccessKey(s.c.RedisPool.Get(), uid)
}

// NewService configures the service
func NewService(defaultConfig Config, configOptions ...func(*Config)) Service {
	for _, option := range configOptions {
		option(&defaultConfig)
	}
	return &service{c: &defaultConfig}
}

//S3 OMIT
func PrivilegedMode(RedisPool *redis.Pool) func(*Config) {
	return func(c *Config) {
		c.RedisPool = RedisPool
	}
}

func getRole(session *mgo.Session, uid string) string {
	defer session.Close()
	// fetch role
	return "admin"
}

func getAccessKey(conn redis.Conn, uid string) (string, error) {
	defer conn.Close()
	//fetch accesskey
	return "xyz123", nil
}

//S4 OMIT
