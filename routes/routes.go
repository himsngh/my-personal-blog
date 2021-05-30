package routes

import (
	"github.com/himsngh/my-personal-blog/database"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Server struct {
	dbStore database.Store
	template 	*template.Template
}

var CurrentWorkingDirectory, _ = os.Getwd()


func NewServer() (*Server, error) {

	tpl := template.Must(template.ParseGlob(CurrentWorkingDirectory + "/templates/*.html"))

	dbStore, err := database.NewDatabaseStore()
	if err != nil {
		log.Println("Error creating the database store: err ", err.Error())
		return nil, err
	}

	return &Server{
		dbStore: dbStore,
		template: tpl,
	}, nil
}

func (s *Server) ServeRoutes() (http.Handler, error) {

	handler := http.NewServeMux()

	fs := http.FileServer(http.Dir(CurrentWorkingDirectory + "/static"))
	handler.Handle("/static/", http.StripPrefix("/static", fs))

	// Register all the routes
	handler.HandleFunc("/signup", func(writer http.ResponseWriter, request *http.Request) {
		s.handleSignUp(writer, request)
	})
	handler.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		s.handleLogin(writer, request)
	})
	handler.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		s.handleHome(writer, request)
	})
	handler.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		s.handleAbout(writer, request)
	})

	return handler, nil
}

func (s *Server) handleSignUp(w http.ResponseWriter, r *http.Request) {
	// TODO IMPLEMENT ME
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	// TODO IMPLEMENT ME
}

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {

	postList, err := s.dbStore.ListPost()
	if err != nil {
		return
	}

	if err := s.template.ExecuteTemplate(w, "home.html", postList); err != nil {
		log.Println("Error executing home template : ", err.Error())
		return
	}
}

func (s *Server) handleAbout(w http.ResponseWriter, r *http.Request) {

	if err := s.template.ExecuteTemplate(w,"about.html", nil); err != nil {
		log.Println("Error executing about template : ", err.Error())
		return
	}
}
