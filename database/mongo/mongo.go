package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

//TYPEINIT OMIT

//mongodb
type Db struct {
	Url        string
	MgoSession *mgo.Session
}

type Person struct {
	Name string
	Age  int
}

//TYPEEND OMIT

//SESSIONINIT OMIT
//get new db session
func (d *Db) GetSession() *mgo.Session {
	if d.MgoSession == nil {
		var err error
		d.MgoSession, err = mgo.Dial(d.Url)
		if err != nil {
			panic(err)
		}
	}
	return d.MgoSession.Clone()

}

//SESSIONEND OMIT

func main() {

	//MINIT OMIT
	db := &Db{Url: "mongodb://gopher:gopher@ds041228.mongolab.com:41228/gotalk"}

	session := db.GetSession()
	defer session.Close()
	c := session.DB("").C("People")

	//insert
	err := c.Insert(&Person{"Gopher", 4},
		&Person{"SuperGopher", 100})
	if err != nil {
		panic(err)
	}

	//find
	result := Person{}
	err = c.Find(bson.M{"name": "Gopher"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("result: %v\n", result)
	//MEND OMIT
}
