package lokyantra

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"net/http"
	"time"
)

//TYPEINIT OMIT
type Option struct {
	OptionId string `bson:"option_id" json:"option_id"`
	Name     string `bson:"name" json:"name"`
	Vote     int    `bson:"vote" json:"vote"`
}

type Question struct {
	Id      bson.ObjectId `bson:"_id" json:"id"`
	Text    string        `bson:"text" json:"text"`
	Options []Option      `bson:"options" json:"options"`
	Created time.Time     `bson:"created" json:"created"`
}

//TYPEEND OMIT

func (s *Server) createQuestion(w http.ResponseWriter, r *http.Request) {

	//JSONINIT OMIT
	question := &Question{}

	if err := readJson(question, r, w); err != nil {
		fmt.Println("malformed question json")
		return
	}

	question.Id = bson.NewObjectId()
	question.Created = time.Now()
	//JSONEND OMIT

	//INSERINIT OMIT
	//get collection
	session := s.db.GetSession()
	defer session.Close()
	c := session.DB("").C("lokyantra")
	//insert
	if err := c.Insert(question); err != nil {

		fmt.Println("error inserting")
		return

	}
	serveJson(w, &question)
	//INSERTEND OMIT

}

func (s *Server) updateQuestion(w http.ResponseWriter, r *http.Request) {

	var vote map[string]string
	if err := readJson(&vote, r, w); err != nil {
		fmt.Println("malformed vote json")
		return
	}
	//UPDATEINIT OMIT
	question_id := r.URL.Query().Get(":id")

	//get collection
	session := s.db.GetSession()
	defer session.Close()
	c := session.DB("").C("lokyantra")

	if err := c.Update(bson.M{"_id": bson.ObjectIdHex(question_id), "options.option_id": vote["option"]}, bson.M{"$inc": bson.M{"options.$.vote": 1}}); err != nil {
		fmt.Printf("error updating%v\n", err)
	}
	//UPDATEEND OMIT

	question := Question{}
	if err := c.FindId(bson.ObjectIdHex(question_id)).One(&question); err != nil {
		fmt.Printf("error finding%v\n", err)
		return

	}
	serveJson(w, &question)

}

func (s *Server) getQuestion(w http.ResponseWriter, r *http.Request) {
	question_id := r.URL.Query().Get(":id")
	//get collection
	session := s.db.GetSession()
	defer session.Close()
	c := session.DB("").C("lokyantra")

	//find
	//GETINIT OMIT
	question := Question{}
	if err := c.FindId(bson.ObjectIdHex(question_id)).One(&question); err != nil {
		fmt.Printf("error finding%v\n", err)
		return

	}
	serveJson(w, &question)
	//GETEND OMIT

}

func (s *Server) deleteQuestion(w http.ResponseWriter, r *http.Request) {
	question_id := r.URL.Query().Get(":id")
	//get collection
	session := s.db.GetSession()
	defer session.Close()
	c := session.DB("").C("lokyantra")

	//DELETEINIT OMIT
	//delete
	if err := c.RemoveId(bson.ObjectIdHex(question_id)); err != nil {
		fmt.Printf("error deleting%v\n", err)
		return

	}

	serveJson(w, struct {
		Status string `json:"status"`
	}{"ok"})

	//DELETEEND OMIT
}
