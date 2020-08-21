package oauth2

import (
	"log"
	"net/http"

	"github.com/ditrit/gandalf/core/models"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
)

// NewOAuth2Server
func NewOAuth2Server() {
	manager := manage.NewDefaultManager()

	configtoken := &Config{DSN: "/home/romainfairant/gandalf/database/tenant1.db", DBType: "sqlite3", TableName: "token", Token: true}
	tokenstore := NewStore(configtoken, 600).(*TokenStore)
	manager.MapTokenStorage(tokenstore)

	configclient := &Config{DSN: "/home/romainfairant/gandalf/database/tenant1.db", DBType: "sqlite3", TableName: "client", Token: false}
	clientstore := NewStore(configclient, 600).(*ClientStore)

	clientstore.Set(&models.Client{
		//ID:     "222222",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientstore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})
	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	log.Fatal(http.ListenAndServe(":9096", nil))

}
