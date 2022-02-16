package routes

import (
	"assignment/database/helper"
	"assignment/handler"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ContextKeys string

const (
	userContext ContextKeys = "__userContext"
)

type Server struct {
	chi.Router
}

//func middleware(handle http.HandlerFunc) http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		log.Println(request.URL.Path)
//		handle(writer, request)
//	}
//}

func Auth(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		apikey := request.Header.Get("x-api-key")
		fmt.Println(apikey)
		user, err := helper.GetSession(apikey)
		if err != nil || user == nil {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte(fmt.Sprintf("Please Login")))
			panic(err)
		}
		fmt.Println(user)
		ctx := context.WithValue(request.Context(), userContext, user)
		handle.ServeHTTP(writer, request.WithContext(ctx))
	}
}

func Route() *Server {
	router := chi.NewRouter()
	router.Route("/assignment", func(assignment chi.Router) {
		assignment.Post("/signup", handler.Signup)
		assignment.Put("/update", handler.Update)
		assignment.Delete("/delete", handler.Delete)
		assignment.Post("/login", handler.Login)
		assignment.Get("/home", Auth(handler.Home))
		assignment.Delete("/logout", Auth(handler.Logout))
	})
	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
