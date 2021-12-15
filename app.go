package main

import (
	"example/mvc/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

func GetRouters() *mux.Router {
	var result *mux.Router
	var services []Route

	result = mux.NewRouter().StrictSlash(true)
	services = append(
		services,
		Route{`GetStudentServ`, `GET`, `/students`, service.GetStudentServ},
		Route{`AddStudentServ`, `POST`, `/students`, service.AddStudentServ},
		Route{`EditStudentServ`, `PATCH`, `/students`, service.EditStudentServ},
		Route{`DeleteStudentServ`, `DELETE`, `/students`, service.DeleteStudentServ},
		Route{`GetSubjectServ`, `GET`, `/subjects`, service.GetSubjectServ},
		Route{`AddSubjectServ`, `POST`, `/subjects`, service.AddSubjectServ},
		Route{`EditSubjectServ`, `PATCH`, `/subjects`, service.EditSubjectServ},
		Route{`DeleteSubjectServ`, `DELETE`, `/subjects`, service.DeleteSubjectServ},
	)

	for _, serv := range services {
		result.
			Methods(serv.Method).
			Path(serv.Pattern).
			Name(serv.Name).
			HandlerFunc(serv.HandleFunc)
	}

	return result
}
