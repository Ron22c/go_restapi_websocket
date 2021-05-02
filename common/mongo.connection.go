package common

import (
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"sync"
	"crypto/tls"
	"net"
	"fmt"
)

type Mongo struct {
	*mgo.Session
}

var mongoDB *Mongo
var once sync.Once

func CreateDBInstance() *Mongo{
	once.Do(func() {
		mongoDB = CreateNewConnection();
	})
	return mongoDB;
}

func CreateNewConnection() *Mongo{
	url := "PROVIDE_YOUR_URL"
	dialInfo, err := mgo.ParseURL(url);
	if err != nil {
		panic(err);
	}

	tlsConfig := new(tls.Config);
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		con, err := tls.Dial("tcp", addr.String(), tlsConfig);
		return con, err;
	}
	session, err := mgo.DialWithInfo(dialInfo);

	if err != nil {
		fmt.Println("ERROR", err)
		panic(err);
	}

	return &Mongo{
		Session: session,
	}

}