module ptiple/barn-svc

go 1.12

require (
	github.com/alicebob/gopher-json v0.0.0-20180125190556-5a6b3ba71ee6 // indirect
	github.com/alicebob/miniredis v2.5.0+incompatible
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/mock v1.3.1
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/tidwall/pretty v1.0.0 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	github.com/yuin/gopher-lua v0.0.0-20190514113301-1cd887cd7036 // indirect
	go.mongodb.org/mongo-driver v1.1.1
	golang.org/x/crypto v0.0.0-20190829043050-9756ffdc2472 // indirect
	golang.org/x/text v0.3.2 // indirect
	ptiple/util v0.0.0
)

replace ptiple/util v0.0.0 => ./../util
