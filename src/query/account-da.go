package query

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	mongourl string
	Database string
)

type BankAccount struct {
	Name    string
	Balance int
}

func writeAccount(s BankAccount) {
	session := GetSession()
	defer session.Close()

	c := session.DB(Database).C("bankAccounts")
	err := c.Insert(s)
	if err != nil {
		log.Fatal(err)
	}
}

func findAccount(id int) BankAccount {
	session := GetSession()
	defer session.Close()

	c := session.DB(Database).C("bankAccounts")

	var result BankAccount
	c.Find(bson.M{"id": id}).One(&result)

	return result
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
