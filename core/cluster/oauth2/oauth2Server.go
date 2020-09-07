package oauth2

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ditrit/gandalf/core/models"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-session/session"
	"github.com/jinzhu/gorm"
)

// NewOAuth2Server
func NewOAuth2Server() {
	manager := manage.NewDefaultManager()

	configtoken := &Config{DSN: "/home/romainfairant/gandalf/database/tenant1.db", DBType: "sqlite3", TableName: "token", Token: true}
	tokenstore := NewStore(configtoken, 600).(*TokenStore)
	manager.MapTokenStorage(tokenstore)

	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))

	configclient := &Config{DSN: "/home/romainfairant/gandalf/database/tenant1.db", DBType: "sqlite3", TableName: "client", Token: false}
	clientstore := NewStore(configclient, 600).(*ClientStore)

	clientstore.Set(&models.Client{
		//ID:     "222222",
		Secret: "999999",
		Domain: "http://localhost:9094",
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

	srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	srv.SetPasswordAuthorizationHandler(func(email, password string) (userID string, err error) {
		tenantDatabaseClient, err := gorm.Open("sqlite3", "/home/romainfairant/gandalf/database/tenant1.db")

		var user models.User
		tenantDatabaseClient.Where("email = ?", email).First(&user)
		fmt.Println(user)

		if models.CompareHashAndPassword(user.Password, password) {
			userID = user.Email
		}
		/*
			if username == "User1" && password == "User1" {
				userID = "User1"
			} */
		return
	})

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/auth", authHandler)

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		store, err := session.Start(r.Context(), w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var form url.Values
		if v, ok := store.Get("ReturnUri"); ok {
			form = v.(url.Values)
		}
		r.Form = form

		store.Delete("ReturnUri")
		store.Save()

		err = srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		token, err := srv.ValidationBearerToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		/* 	fmt.Println("token")
		fmt.Println(token)
		jwttoken, err := jwt.ParseWithClaims(token.GetAccess(), &generates.JWTAccessClaims{}, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("parse error")
			}
			return []byte("00000000"), nil
		})
		if err != nil {
			// panic(err)
		}
		fmt.Println("jwttoken")
		fmt.Println(jwttoken)

		claims, ok := jwttoken.Claims.(*generates.JWTAccessClaims)
		if !ok || !jwttoken.Valid {
			// panic("invalid token")
		}
		fmt.Println("claims")
		fmt.Println(claims)
		*/

		//TODO ADD VALIDATION

		data := map[string]interface{}{
			"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
			"client_id":  token.GetClientID(),
			"user_id":    token.GetUserID(),
		}
		e := json.NewEncoder(w)
		e.SetIndent("", "  ")
		e.Encode(data)
	})

	log.Fatal(http.ListenAndServe(":9096", nil))
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}

//TODO REVOIR
func loginHandler(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		if r.Form == nil {
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		//VERIFICATION EMAIL PASSWORD
		tenantDatabaseClient, _ := gorm.Open("sqlite3", "/home/romainfairant/gandalf/database/tenant1.db")

		var user models.User
		tenantDatabaseClient.Where("email = ?", r.Form.Get("email")).First(&user)
		fmt.Println(user)

		if models.CompareHashAndPassword(user.Password, r.Form.Get("password")) {
			store.Set("LoggedInUserID", r.Form.Get("email"))
			store.Save()

			w.Header().Set("Location", "/auth")
			w.WriteHeader(http.StatusFound)
		} else {
			w.Header().Set("Location", "/login")
			w.WriteHeader(http.StatusFound)
		}

		return
	}
	outputHTML(w, r, "./oauth2/static/login.html")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	outputHTML(w, r, "./oauth2/static/auth.html")
}

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
