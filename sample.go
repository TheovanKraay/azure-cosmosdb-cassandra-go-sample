package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("ACCOUNTNAME.cassandra.cosmos.azure.com")
	cluster.Port = 10350
	var sslOptions = new(gocql.SslOptions)
	sslOptions.EnableHostVerification = true
	sslOptions.CaPath = "path/to/.pem"
	// sslOptions.ServerName = `endpoint` // This should fix client side ssl verification, TODO: figure out how to do this right.
	cluster.SslOpts = sslOptions
	cluster.ProtoVersion = 4
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "ACCOUNTNAME", Password: "pwd"}
	session, _ := cluster.CreateSession()
	defer session.Close()

	// Create Keyspace
	if err := session.Query(`CREATE KEYSPACE IF NOT EXISTS kspc1 WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'} AND durable_writes = true;`).Exec(); err != nil {
		log.Fatal(err)
	}

	// Create Table
	if err := session.Query(`CREATE TABLE IF NOT EXISTS kspc1.tbl1 (pk int, ck int, col1 text, col2 bigint, PRIMARY KEY (pk, ck));`).Exec(); err != nil {
		log.Fatal(err)
	}

	// Insert Row
	if err := session.Query(`INSERT INTO kspc1.tbl1 (pk, ck, col2) VALUES (1, 1, 1);`).Exec(); err != nil {
		log.Fatal(err)
	}

	var pk int32
	var ck int32
	var col1 string
	var col2 int64

	// SELECT
	if err := session.Query(`SELECT * FROM kspc1.tbl1;`).Scan(&pk, &ck, &col1, &col2); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Row:", pk, ck, col1, col2)
}
