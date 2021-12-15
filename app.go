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
		Route{`GetStudentServ`, `GET`, `/student`, service.GetStudentServ},
		Route{`AddStudentServ`, `POST`, `/student`, service.AddStudentServ},
		Route{`EditStudentServ`, `PATCH`, `/student`, service.EditStudentServ},
		Route{`DeleteStudentServ`, `DELETE`, `/student`, service.DeleteStudentServ},
		Route{`GetSubjectServ`, `GET`, `/subject`, service.GetSubjectServ},
		Route{`AddSubjectServ`, `POST`, `/subject`, service.AddSubjectServ},
		Route{`EditSubjectServ`, `PATCH`, `/subject`, service.EditSubjectServ},
		Route{`DeleteSubjectServ`, `DELETE`, `/subject`, service.DeleteSubjectServ},
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
