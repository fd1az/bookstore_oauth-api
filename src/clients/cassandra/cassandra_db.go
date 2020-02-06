package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	//Connect to cassandra cluster:
	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
	fmt.Println("DB connection ready " + cluster.Keyspace + cluster.Hosts[0])
}

func GetSession() *gocql.Session {
	return session

}
