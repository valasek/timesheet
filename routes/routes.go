package routes

import (
	"net/http"
	"fmt"
	"os"
	"reflect"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/valasek/time-sheet/api"
	"github.com/valasek/time-sheet/auth"
	"github.com/urfave/negroni"
)

// NewRoutes builds the routes for the api
func NewRoutes(api *api.API) *mux.Router {

	mux := mux.NewRouter()

	// client static files
	mux.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods("GET")
	mux.PathPrefix("/js").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./client/dist/js/"))))
	mux.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./client/dist/css/"))))

	// api
	a := mux.PathPrefix("/api").Subrouter()

	// users
	u := a.PathPrefix("/user").Subrouter()
	u.HandleFunc("/signup", api.UserSignup).Methods("POST")
	u.HandleFunc("/login", api.UserLogin).Methods("POST")
	u.Handle("/info", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.UserInfo)),
	))

	// consultants
	c := a.PathPrefix("/consultants").Subrouter()
	c.HandleFunc("/list", api.ConsultantList).Methods("GET")

	// projects
	p := a.PathPrefix("/projects").Subrouter()
	p.HandleFunc("/list", api.ProjectsGetAll).Methods("GET")

	// rates
	rates := a.PathPrefix("/rates").Subrouter()
	rates.HandleFunc("/list", api.RatesGetAll).Methods("GET")

	// reported records
	r := a.PathPrefix("/reported").Subrouter()
	r.HandleFunc("/all", api.ReportedRecordsGetAll).Methods("GET")
	r.HandleFunc("/month/{month}", api.ReportedRecordsInMonth).Methods("GET")
	r.HandleFunc("/{id}", api.ReportedRecordDelete).Methods("DELETE")
	r.HandleFunc("/{id}", api.ReportedRecordUpdate).Methods("PUT")

	// quotes
	q := a.PathPrefix("/quote").Subrouter()
	q.HandleFunc("/random", api.Quote).Methods("GET")
	q.Handle("/protected/random", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.SecretQuote)),
	))

	return mux
}

// PrintRoutes prints all set routes
func PrintRoutes (routes *mux.Router) {
	fmt.Println("Available routes")
	fmt.Println("[URI] [Methods] [Handler]")
	err := routes.Walk(gorillaWalkFn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func gorillaWalkFn(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	path, err := route.GetPathTemplate()
	if err != nil {
		fmt.Println(err)
		return err
	}
	methods, _ := route.GetMethods()
	reflectValue := reflect.ValueOf(route.GetHandler())
	handler := "sub-router"
	if reflectValue.IsValid() {
		handler = runtime.FuncForPC(reflectValue.Pointer()).Name()
	}
	fmt.Println(path, methods, handler)
    return nil
}
