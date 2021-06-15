package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "8000"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"upload",
		"POST",
		"/api/upload",
		upload,
	},
	Route{
		"login",
		"POST",
		"/api/login",
		login,
	},
	Route{
		"logout",
		"GET",
		"/api/logout",
		logout,
	},
	Route{
		"auth",
		"GET",
		"/api/auth",
		auth,
	},
	Route{
		"readRecords",
		"POST",
		"/api/readRecords",
		readRecords,
	},
	Route{
		"updateRecord",
		"POST",
		"/api/updateRecord",
		updateRecord,
	},
	Route{
		"createRecord",
		"POST",
		"/api/createRecord",
		createRecord,
	},
	Route{
		"deleteRecord",
		"POST",
		"/api/deleteRecord",
		isAuthorized(deleteRecord),
	},
}

func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

// https://pkg.go.dev/github.com/gorilla/mux#readme-serving-single-page-applications
type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)

	router := AddRoutes(muxRouter)

	spa := spaHandler{staticPath: "dist", indexPath: "index.html"}

	router.HandleFunc("/public/upload", upload).Methods("POST", "OPTIONS")
	router.HandleFunc("/public/status", status).Methods("GET", "OPTIONS")
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	// Handles all the requests that comes to /static/*
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// Handles everything exsept requests coming to the API
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      router,
		Addr:         CONN_HOST + ":" + CONN_PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
