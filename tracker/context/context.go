package context

import "github.com/jbonachera/gobitt/tracker/plugin"
import "github.com/jbonachera/gobitt/tracker/config"

type ApplicationContext struct {
	Database plugin.DatabasePlugin
	Confdir  string
	Config   config.Config
}
