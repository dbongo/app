package main

import "gopkg.in/mgo.v2"

type mongodb struct {
	session *mgo.Session
}

func createDB() *mongodb {
	mgodb := &mongodb{}
	return mgodb
}
