package command

import (
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

type BankEvent struct {
	Occurred  time.Time
	UserId    int
	Data      string
	AccountId int
}

var (
	mongourl string
	Database string
)

func write(s BankEvent) {
	session := GetSession()
	defer session.Close()

	c := session.DB(Database).C("bankEvents")
	err := c.Insert(s)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSession() *mgo.Session {
	session, err := mgo.Dial(mongourl)
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}

func init() {

	if mongourl = os.Getenv("MONGO_URL"); mongourl == "" {
		mongourl = "mongodb://localhost"
	}

	Database = "eventsourcing"
}
