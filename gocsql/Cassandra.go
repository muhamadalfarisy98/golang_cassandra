package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error
	// connect to local adress
	cluster := gocql.NewCluster("127.0.0.1")

	// keyspace we created previously
	cluster.Keyspace = "restfulapi"

	// specifed and create a new session
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra well initialized")
}
