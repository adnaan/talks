package polls

import (
	"fmt"
	"github.com/drone/routes"
	"labix.org/v2/mgo"
	"net"
	"net/http"
)

//The server type holds instances of all components
//TYPEINIT OMIT
type Server struct {
	listener   net.Listener
	httpServer *http.Server
	route      *routes.RouteMux
	db         *Db
}

//mongodb
type Db struct {
	Url        string
	MgoSession *mgo.Session
}

//get new db session
func (d *Db) GetSession() *mgo.Session {
	if d.MgoSession == nil {
		var err error
		d.MgoSession, err = mgo.Dial(d.Url)
		if err != nil {
			panic(err) // no, not really
		}
	}
	return d.MgoSession.Clone()
}

//TYPEEND OMIT

//dbUrl:"mongodb://gopher:gopher@localhost:27017/gotalk"
//serverUrl:fmt.Sprintf("%s:%d", address, port)
//test db: mongodb://gopher:gopher@ds041228.mongolab.com:41228/gotalk
//to create user in mongodb: db.addUser( { user: "gopher",pwd: "gopher",roles: [ "userAdminAnyDatabase" ] } )
//creates a new server
//NEWINIT OMIT
func NewServer() *Server {

	//initialize server
	r := routes.New()
	s := &Server{
		httpServer: &http.Server{Addr: fmt.Sprintf(":%d", 9947), Handler: r},
		route:      r,
		db:         &Db{Url: "mongodb://gopher:gopher@localhost:27017/gotalk"},
	}

	s.addHandlers()

	return s

}

//NEWEND OMIT

//HINIT OMIT

func (s *Server) addHandlers() {
	s.route.Post("/api/v1/q/", s.createQuestion)
	s.route.Put("/api/v1/q/:id", s.updateQuestion)
	s.route.Get("/api/v1/q/:id", s.getQuestion)
	s.route.Del("/api/v1/q/:id", s.deleteQuestion)

}

//HEND OMIT

//listen and serve a fastcgi server
//LINIT OMIT
func (s *Server) ListenAndServe() error {

	listener, err := net.Listen("tcp", s.httpServer.Addr)
	if err != nil {
		fmt.Printf("error listening: %v \n", err)
		return err
	}
	s.listener = listener

	go s.httpServer.Serve(s.listener)

	//fmt.Printf("Poll Server is now listening on http://localhost%s\n", s.httpServer.Addr)

	return nil
}

//LEND OMIT

//SINIT OMIT
// stops the server.
func (s *Server) Shutdown() error {

	if s.listener != nil {
		// Then stop the server.
		err := s.listener.Close()
		s.listener = nil
		if err != nil {
			return err
		}
	}

	return nil
}

//SEND OMIT
