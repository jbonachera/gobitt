package models

import (
	"gopkg.in/mgo.v2"
)

type ApplicationContext struct {
	Session *mgo.Session
}
