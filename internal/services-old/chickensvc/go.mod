module ptiple/chickensvc

go 1.18

require (
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/gorilla/mux v1.7.3
	go.mongodb.org/mongo-driver v1.5.1
	ptiple/util v0.0.0
)

require (
	github.com/aws/aws-sdk-go v1.34.28 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/compress v1.9.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073 // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/text v0.3.5 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace ptiple/util v0.0.0 => ./../util
