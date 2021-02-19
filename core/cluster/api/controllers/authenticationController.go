package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/cluster/api/dao"
	"github.com/ditrit/gandalf/core/cluster/api/utils"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	apimodels "github.com/ditrit/gandalf/core/cluster/api/models"
	"github.com/ditrit/gandalf/core/models"
)

// AuthenticationController :
type AuthenticationController struct {
	databaseConnection *database.DatabaseConnection
}

// NewAuthenticationController :
func NewAuthenticationController(databaseConnection *database.DatabaseConnection) (authenticationController *AuthenticationController) {
	authenticationController = new(AuthenticationController)
	authenticationController.databaseConnection = databaseConnection

	return
}

// Login :
func (ac AuthenticationController) Login(w http.ResponseWriter, r *http.Request) {
	ruser := &models.User{}
	err := json.NewDecoder(r.Body).Decode(ruser)
	if err != nil {

		utils.RespondWithError(w, http.StatusBadRequest, err.Error())

		//var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		//json.NewEncoder(w).Encode(resp)
		return
	}
	user := models.User{}
	if user, err = dao.ReadUserByName(ac.databaseConnection.GetGandalfDatabaseClient(), ruser.Name); err != nil {

		//var resp = map[string]interface{}{"status": false, "message": "Username not found"}
		//return resp
		fmt.Println("LOGIN3")
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

	tokenString, err := token.SignedString([]byte("gandalf"))
	if err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, tokenString)
	return

}

/* // FindOne :
func (ac AuthenticationController) FindOne(name, password string) (int, string, string) {
	user := models.User{}
	var err error
	if user, err = dao.ReadUserByName(ac.gandalfDatabase, name); err != nil {

		//var resp = map[string]interface{}{"status": false, "message": "Username not found"}
		//return resp
		return http.StatusNotFound, "", "Username not found"
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		//var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		//return resp
		return http.StatusUnauthorized, "", "Invalid login credentials. Please try again"
	}
	tk := &apimodels.Claims{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Tenant: "gandalf",
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, err := token.SignedString([]byte("gandalf"))
	if err != nil {
		fmt.Println(err)
	}

	//var resp = map[string]interface{}{"status": true, "message": "logged in"}
	//resp["token"] = tokenString //Store the token in the response
	//resp["user"] = user
	//return resp
	return http.StatusOK, tokenString, ""
}
*/
