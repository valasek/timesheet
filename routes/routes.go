package routes

import (
	"net/http"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"text/tabwriter"

	"github.com/gorilla/mux"
	"github.com/valasek/timesheet/api"
	"github.com/valasek/timesheet/auth"
	"github.com/urfave/negroni"
)

var w *tabwriter.Writer

// NewRoutes builds the routes for the api
func NewRoutes(api *api.API) *mux.Router {

	mux := mux.NewRouter()

	// client static files
	mux.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods("GET")
	mux.PathPrefix("/js").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./client/dist/js/"))))
	mux.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./client/dist/css/"))))

	// api
	a := mux.PathPrefix("/api").Subrouter()

	// app settings
	a.HandleFunc("/settings", api.AppSettings).Methods("GET")

	// users
	u := a.PathPrefix("/user").Subrouter()
	u.HandleFunc("/signup", api.UserSignup).Methods("POST")
	u.HandleFunc("/login", api.UserLogin).Methods("POST")
	u.Handle("/info", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.UserInfo)),
	))

	// consultants
	a.HandleFunc("/consultants", api.ConsultantList).Methods("GET")

	// projects
	a.HandleFunc("/projects", api.ProjectsGetAll).Methods("GET")

	// rates
	a.HandleFunc("/rates", api.RatesGetAll).Methods("GET")

	// holidays
	a.HandleFunc("/holidays", api.HolidaysGetAll).Methods("GET")

	// reported records
	a.HandleFunc("/reported", api.ReportedRecordsAddRecord).Methods("POST")
	a.HandleFunc("/reported", api.ReportedRecordsGetAll).Methods("GET")
	a.HandleFunc("/reported/summary", api.ReportedRecordsSummary).Methods("GET")
	a.HandleFunc("/reported/year/{year}/month/{month}", api.ReportedRecordsInMonth).Methods("GET")
	a.HandleFunc("/reported/{id}", api.ReportedRecordDelete).Methods("DELETE")
	a.HandleFunc("/reported/{id}", api.ReportedRecordUpdate).Methods("PUT")

	// quotes
	// q := a.PathPrefix("/quote").Subrouter()
	// q.HandleFunc("/random", api.Quote).Methods("GET")
	// q.Handle("/protected/random", negroni.New(
	// 	negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
	// 	negroni.Wrap(http.HandlerFunc(api.SecretQuote)),
	// ))

	return mux
}

// PrintRoutes prints all set routes
func PrintRoutes (routes *mux.Router) {
	fmt.Println("Available routes")
	w = new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, "[URI]\t[Methods]\t[Handler]")
	err := routes.Walk(gorillaWalkFn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Fprintln(w)
	w.Flush()
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
	fmt.Fprintln(w, path, "\t", methods, "\t", handler)
    return nil
}
