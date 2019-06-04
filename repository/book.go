package repository

import (
	"github.com/globalsign/mgo"
	"log"
	"os"
)

var Infos, Chapters *mgo.Collection

func init() {
	mgoURI := os.Getenv("mgo_url")
	if mgoURI == "" {
		log.Fatalf("connect to mongo failed %v", mgoURI)
	}
	session, err := mgo.Dial(mgoURI)
	if err != nil {
		log.Fatalf("connect to mongo failed %v", err)
	}

	database := session.DB("book")
	Infos = database.C("info")
	Chapters = database.C("chapters")

}
