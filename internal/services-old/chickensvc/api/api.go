package api

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ptiple/chickensvc/mongodatabase"
	"ptiple/util"
)

type Api struct {
	DB    *mongodatabase.MongoDatabase
	Redis *redis.Client
}

func Start(_mongodb *mongodatabase.MongoDatabase, _redisClient *redis.Client) {
	var api = Api{_mongodb, _redisClient}

	_ = api.setupHandlers()
	// TODO remove unsafe cors origins when building for containers that will run in the same domain
	//router := api.setupHandlers()
	//headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization"})
	//originsOk := handlers.AllowedOrigins([]string{"*"})
	//methodsOk := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodOptions})
	//log.Fatal(http.ListenAndServe(":8082", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (api Api) setupHandlers() *mux.Router {
	r := mux.NewRouter()
	r.Use(util.JwtAuthentication)

	r.HandleFunc("/", api.getChickensOfFarmerHandler).Methods("GET")
	r.HandleFunc("/buy/{barnId}", api.buyChickenHandler).Methods("GET")
	r.HandleFunc("/{chickenId}/feed", api.feedChickenHandler).Methods("GET")
	r.HandleFunc("/bulkFeed", api.bulkFeedChickenHandler).Methods("POST")
	r.HandleFunc("/{chickenId}/sell", api.sellChickenHandler).Methods("GET")

	http.Handle("/", r)

	return r
}
