package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"test-dans/config"
	"test-dans/internal/app/delivery"
	deliveryHttp "test-dans/internal/app/delivery/http"
	repoAuth "test-dans/internal/app/repository/authentication/mysql"
	repoJob "test-dans/internal/app/repository/jobs/http"
	usecaseAuth "test-dans/internal/app/usecase/authentication/handler"
	usecaseJob "test-dans/internal/app/usecase/jobs/handler"
	"test-dans/pkg/connection/mysql"
)

func main() {
	// Init config
	config, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	// Init database
	dbUser, err := mysql.InitConnection(config.Database)
	defer dbUser.Close()
	if err != nil {
		panic(err)
	}

	// Init Repository
	authRepo := repoAuth.New(dbUser, config.Database.TableName, config.Database.TimeoutMS)
	jobsRepo := repoJob.New(config.EndPoint.URL, config.EndPoint.TimeoutMS)

	// Init Usecase
	authUsecase := usecaseAuth.New(authRepo)
	jobsUsecase := usecaseJob.New(jobsRepo)

	// Init Delivery
	deliveryHTTP := deliveryHttp.New(authUsecase, jobsUsecase, config.JWTKey, config.HTTPServer.TimeoutMS)

	router := newRouter(deliveryHTTP)

	log.Println("Serve at port 9999")
	http.ListenAndServe(":9999", router)

}

func newRouter(deliveryHTTP delivery.Delivery) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/register", wrapHandler(deliveryHTTP.Register)).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", wrapHandler(deliveryHTTP.Login)).Methods("POST", "OPTIONS")
	r.HandleFunc("/jobs/list", wrapHandler(deliveryHTTP.Register)).Methods("GET", "OPTIONS")
	r.HandleFunc("/jobs/details/:id", wrapHandler(deliveryHTTP.Register)).Methods("GET", "OPTIONS")

	return r
}

func wrapHandler(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if (*r).Method == "OPTIONS" {
			return
		}
		handler(w, r)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	(*w).Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	(*w).Header().Set("Access-Control-Expose-Headers", "Authorization")
}
