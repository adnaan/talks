//S1 OMIT
package mock

import "github.com/adnaan/talks/dependency_injection_july2017/user"

// configurable service which implements the Service interface
type service struct {
	c *user.Config
}

// NewService configures the service
func NewService(defaultConfig user.Config, configOptions ...func(*user.Config)) user.Service {
	for _, option := range configOptions {
		option(&defaultConfig)
	}
	return &service{c: &defaultConfig}
}

//...
//S2 OMIT
