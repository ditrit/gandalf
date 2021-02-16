package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ditrit/gandalf/core/cluster/api/utils"

	apimodels "github.com/ditrit/gandalf/core/cluster/api/models"

	"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/mux"
)

// CommonMiddleware :
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}

// GandalfJwtVerify :
func GandalfJwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//var header = r.Header.Get("x-access-token") //Grab the token from the header
		header := utils.ExtractToken(r)

		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(apimodels.Exception{Message: "Missing auth token"})
			return
		}
		tk := &apimodels.Claims{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("gandalf"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(apimodels.Exception{Message: err.Error()})
			return
		}

		if tk.Tenant != "gandalf" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(apimodels.Exception{Message: "Wrong tenant"})
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TenantsJwtVerify :
func TenantsJwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tenant := vars["tenant"]
		//var header = r.Header.Get("x-access-token") //Grab the token from the header
		header := utils.ExtractToken(r)

		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(apimodels.Exception{Message: "Missing auth token"})
			return
		}
		tk := &apimodels.Claims{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("gandalf"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(apimodels.Exception{Message: err.Error()})
			return
		}

		if tk.Tenant != tenant && tk.Tenant != "gandalf" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(apimodels.Exception{Message: "Wrong tenant"})
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
