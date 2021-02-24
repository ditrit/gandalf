module github.com/ditrit/gandalf/core

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ditrit/gandalf/libraries/goclient v0.0.0-20210216134342-40c7d10bd6c4
	github.com/ditrit/gandalf/libraries/gogrpc v0.0.0-20210216134342-40c7d10bd6c4
	github.com/ditrit/gandalf/verdeter v0.0.0-20210217103225-d2bf79e86cbd
	github.com/ditrit/shoset v0.0.0-20201026092509-225b8a4a5276
	github.com/go-oauth2/oauth2/v4 v4.2.0
	github.com/go-session/session v3.1.2+incompatible
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	github.com/spf13/viper v1.7.1
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/grpc v1.33.2
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/ditrit/gandalf/verdeter => ../verdeter
replace github.com/ditrit/gandalf/libraries/goclient => ../libraries/goclient
replace github.com/ditrit/gandalf/libraries/gogrpc => ../libraries/gogrpc
