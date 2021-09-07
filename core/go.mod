module github.com/ditrit/gandalf/core

go 1.14

require (
	github.com/ditrit/gandalf/libraries/goclient v0.0.0-20210216134342-40c7d10bd6c4
	github.com/ditrit/gandalf/libraries/gogrpc v0.0.0-20210216134342-40c7d10bd6c4
	github.com/ditrit/gandalf/verdeter v0.0.0-20210217103225-d2bf79e86cbd
	github.com/ditrit/shoset v0.0.0-20210903074700-5ef969b431ba
	github.com/golang-jwt/jwt/v4 v4.0.0
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	github.com/rs/cors v1.8.0
	github.com/spf13/viper v1.8.0
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/grpc v1.39.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/ditrit/gandalf/verdeter => ../verdeter

replace github.com/ditrit/gandalf/libraries/goclient => ../libraries/goclient

replace github.com/ditrit/gandalf/libraries/gogrpc => ../libraries/gogrpc
