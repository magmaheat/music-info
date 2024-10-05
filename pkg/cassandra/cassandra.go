package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
)

type Cassandra struct {
	db *gocql.Session
}

func New() (*Cassandra, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "music-library"
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("csdb - New - CreateSession: %v", err)
	}
	defer session.Close()

	return &Cassandra{db: session}, nil
}
