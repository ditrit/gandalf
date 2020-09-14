package tenants

import (
	"encoding/json"
	"fmt"
	"gandalf/core/cluster/api/dao"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/ditrit/gandalf/core/cluster/api/utils"
	"github.com/ditrit/gandalf/core/models"
	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
)

type AuthenticationController struct {
	mapDatabase  map[string]*gorm.DB
	databasePath string
}

func NewAuthenticationController(mapDatabase map[string]*gorm.DB, databasePath string) (aggregatorController *AggregatorController) {
	aggregatorController = new(AggregatorController)
	aggregatorController.mapDatabase = mapDatabase
	aggregatorController.databasePath = databasePath

	return
}

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
	resp := FindOne(database, user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)
}

//TODO REVOIR
func (ac AuthenticationController) FindOne(database *gorm.DB, email, password string) map[string]interface{} {
	user := &models.User{}
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

	tk := &models.Token{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}
