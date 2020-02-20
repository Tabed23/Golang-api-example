package database

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)
	//error in this area, could not load my database
func GetMongo()(*mgo.Database, error){
	host := os.Getenv("MONGO_HOST")
	if len(host) <= 0{
		host = "mongodb://localhost:27017"

	}
	session , err :=  mgo.Dial(host)
	if err != nil{
		return nil,err
	}
	db :=  session.DB("company")
	return db,nil
}
