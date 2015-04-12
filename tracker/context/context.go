package context

import "github.com/jbonachera/gobitt/tracker/plugin"

type ApplicationContext struct {
	Database plugin.DatabasePlugin
}
