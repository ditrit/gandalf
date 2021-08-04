package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ditrit/gandalf/core/aggregator/database"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	apimodels "github.com/ditrit/gandalf/core/aggregator/api/models"
	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/ditrit/gandalf/core/models"
)

// AuthenticationController :
type AuthenticationController struct {
	databaseConnection *database.DatabaseConnection
}

//NewAuthenticationController :
func NewAuthenticationController(databaseConnection *database.DatabaseConnection) (authenticationController *AuthenticationController) {
	authenticationController = new(AuthenticationController)
	authenticationController.databaseConnection = databaseConnection

	return
}

// Login :
func (ac AuthenticationController) Login(w http.ResponseWriter, r *http.Request) {
	database := ac.databaseConnection.GetTenantDatabaseClient()
	if database != nil {

		ruser := &models.User{}
		fmt.Println("LOGIN")
		err := json.NewDecoder(r.Body).Decode(ruser)
		if err != nil {

			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")

			//var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
			//json.NewEncoder(w).Encode(resp)
			return
		}
		user := models.User{}
		if user, err = dao.ReadAdminByName(database, ruser.Name); err != nil {

			//var resp = map[string]interface{}{"status": false, "message": "Username not found"}
			//return resp
			utils.RespondWithError(w, http.StatusNotFound, "Username not found")
			return
		}
		expiresAt := time.Now().Add(time.Minute * 100000).Unix()

		errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ruser.Password))
		if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			//var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
			//return resp
			//return http.StatusUnauthorized, "", "Invalid login credentials. Please try again"
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid login credentials. Please try again")
			return
		}
		tk := &apimodels.Claims{
			UserID: user.ID,
			Name:   user.Name,
			Email:  user.Email,
			StandardClaims: &jwt.StandardClaims{
				ExpiresAt: expiresAt,
			},
		}

		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

		tokenString, err := token.SignedString([]byte("aggregator"))
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, tokenString)
		return
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

/*
func (ac AuthenticationController) Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabaseClientByTenant(ac.mapDatabase, tenant)
	if database != nil {

		user := &models.User{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		resp := ac.FindOne(database, user.Email, user.Password, tenant)
		json.NewEncoder(w).Encode(resp)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// FindOne :
func (ac AuthenticationController) FindOne(database *gorm.DB, email, password, tenant string) map[string]interface{} {
	user := models.User{}
	var err error
	if user, err = dao.ReadUserByEmail(database, email); err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &apimodels.Claims{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Tenant: tenant,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("gandalf"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}
*/
