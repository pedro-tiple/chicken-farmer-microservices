package api

import (
	"log"
	"net/http"
	"ptiple/barnsvc/mongodatabase"
	"ptiple/util"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceCalls struct {
	SpendGoldEggs  func(_farmerId primitive.ObjectID, _amount uint) error
	AddFreeChicken func(_barnId primitive.ObjectID, _barnOwnerId primitive.ObjectID) error
}

type Api struct {
	DB           mongodatabase.IMongoDatabase
	Redis        *redis.Client
	ServiceCalls ServiceCalls
}

func Start(_mongodb mongodatabase.IMongoDatabase, _redisClient *redis.Client) {
	var api = Api{
		_mongodb,
		_redisClient,
		ServiceCalls{
			SpendGoldEggs:  util.SpendGoldEggs,
			AddFreeChicken: util.AddFreeChicken,
		},
	}

	_ = api.setupHandlers()
	// TODO remove unsafe cors origins when building for containers that will run in the same domain
	//router := api.setupHandlers()
	//headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization"})
	//originsOk := handlers.AllowedOrigins([]string{"*"})
	//methodsOk := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodOptions})
	//log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (api *Api) setupHandlers() *mux.Router {
	r := mux.NewRouter()
	r.Use(util.JwtAuthentication)

	r.HandleFunc("/", api.getBarnsHandler).Methods("GET")
	r.HandleFunc("/buy", api.buyBarnHandler).Methods("GET")
	r.HandleFunc("/spendFeed", api.spendFeedHandler).Methods("POST")
	r.HandleFunc("/{barnId}", api.getBarnHandler).Methods("GET")
	r.HandleFunc("/{barnId}/buy/feed", api.buyFeedHandler).Methods("GET")
	r.HandleFunc("/{barnId}/buy/autoFeeder", api.buyAutoFeederHandler).Methods("GET")

	http.Handle("/", r)

	return r
}
