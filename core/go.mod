module github.com/ditrit/gandalf/core

go 1.14

require (
	github.com/ditrit/gandalf/libraries/goclient v0.0.0-20211213144417-a91cc4cf66b4
	github.com/ditrit/gandalf/libraries/gogrpc v0.0.0-20211213144417-a91cc4cf66b4
	github.com/ditrit/gandalf/verdeter v0.0.0-20211213144417-a91cc4cf66b4
	github.com/ditrit/shoset v0.0.0-20210903074700-5ef969b431ba
	github.com/golang-jwt/jwt/v4 v4.2.0
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.10.4 // indirect
	github.com/rs/cors v1.8.0
	github.com/spf13/cobra v1.2.1 // indirect
	github.com/spf13/viper v1.10.0
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b
	golang.org/x/net v0.0.0-20211209124913-491a49abca63 // indirect
	golang.org/x/sys v0.0.0-20211213223007-03aa0b5f6827 // indirect
	google.golang.org/grpc v1.42.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/ditrit/gandalf/verdeter => ../verdeter

replace github.com/ditrit/gandalf/libraries/goclient => ../libraries/goclient

replace github.com/ditrit/gandalf/libraries/gogrpc => ../libraries/gogrpc
