package api

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ptiple/chicken-svc/mongodatabase"
	"ptiple/util"
)

type Api struct {
	DB    *mongodatabase.MongoDatabase
	Redis *redis.Client
}

func Start(_mongodb *mongodatabase.MongoDatabase, _redisClient *redis.Client) {
	var api = Api{_mongodb, _redisClient}

	router := api.setupHandlers()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type"})
	// TODO remove unsafe cors origins when building for containers that will run in the same domain
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodOptions})
	log.Fatal(http.ListenAndServe(":8082", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func (api Api) setupHandlers() *mux.Router {
	r := mux.NewRouter()
	r.Use(util.JwtAuthentication)

	r.HandleFunc("/chickens", api.getChickensOfUserHandler).Methods("GET")
	r.HandleFunc("/chickens/buy/{barnId}", api.buyChickenHandler).Methods("GET")
	r.HandleFunc("/chickens/{chickenId}/feed", api.feedChickenHandler).Methods("GET")
	r.HandleFunc("/chickens/bulkFeed", api.bulkFeedChickenHandler).Methods("POST")
	r.HandleFunc("/chickens/{chickenId}/sell", api.sellChickenHandler).Methods("GET")

	http.Handle("/", r)

	return r
}
