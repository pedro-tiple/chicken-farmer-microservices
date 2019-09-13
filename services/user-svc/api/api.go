package api

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ptiple/user-svc/mongodatabase"
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
	log.Fatal(http.ListenAndServe(":8083", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func (api Api) setupHandlers() *mux.Router {
	r := mux.NewRouter()
	r.Use(util.JwtAuthentication)

	r.HandleFunc("/users/ws", api.wsHandler)
	r.HandleFunc("/users/login/{userId}", api.loginHandler).Methods("GET")
	r.HandleFunc("/users/getGoldEggs", api.getGoldEggsHandler).Methods("GET")
	r.HandleFunc("/users/spendGoldEggs", api.spendGoldEggsHandler).Methods("POST")

	http.Handle("/", r)

	return r
}
