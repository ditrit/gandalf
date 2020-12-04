package tenants

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ditrit/gandalf/core/cluster/api/dao"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	apimodels "github.com/ditrit/gandalf/core/cluster/api/models"
	"github.com/ditrit/gandalf/core/cluster/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
)

// AuthenticationController :
type AuthenticationController struct {
	mapDatabase  map[string]*gorm.DB
	databasePath string
}

//NewAuthenticationController :
func NewAuthenticationController(mapDatabase map[string]*gorm.DB, databasePath string) (authenticationController *AuthenticationController) {
	authenticationController = new(AuthenticationController)
	authenticationController.mapDatabase = mapDatabase
	authenticationController.databasePath = databasePath

	return
}

// Login :
func (ac AuthenticationController) Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(ac.mapDatabase, ac.databasePath, tenant)

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := ac.FindOne(database, user.Email, user.Password, tenant)
	json.NewEncoder(w).Encode(resp)
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

	tokenString, error := token.SignedString([]byte("tenants"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}
