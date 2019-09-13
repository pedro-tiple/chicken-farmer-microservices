package api

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ptiple/barn-svc/mongodatabase"
	"ptiple/util"
)

type Api struct {
	DB    *mongodatabase.MongoDatabase
	Redis *redis.Client
}

func Start(_mongodb *mongodatabase.MongoDatabase, _redisClient *redis.Client) {
	var api = Api{_mongodb, _redisClient}

	router := api.setupHandlers()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization"})
	// TODO remove unsafe cors origins when building for containers that will run in the same domain
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodOptions})
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func (api Api) setupHandlers() *mux.Router {
	r := mux.NewRouter()
	r.Use(util.JwtAuthentication)

	r.HandleFunc("/barns", api.getBarnsHandler).Methods("GET")
	r.HandleFunc("/barns/buy", api.buyBarnHandler).Methods("GET")
	r.HandleFunc("/barns/spendFeed", api.spendFeedHandler).Methods("POST")
	r.HandleFunc("/barns/{barnId}", api.getBarnHandler).Methods("GET")
	r.HandleFunc("/barns/{barnId}/buy/feed", api.buyFeedHandler).Methods("GET")
	r.HandleFunc("/barns/{barnId}/buy/autoFeeder", api.buyAutoFeederHandler).Methods("GET")

	http.Handle("/", r)

	return r
}
