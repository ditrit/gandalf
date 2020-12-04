package gandalf

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ditrit/gandalf/core/cluster/api/dao"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	apimodels "github.com/ditrit/gandalf/core/cluster/api/models"
	"github.com/ditrit/gandalf/core/models"

	"github.com/jinzhu/gorm"
)

// AuthenticationController :
type AuthenticationController struct {
	gandalfDatabase *gorm.DB
}

// NewAuthenticationController :
func NewAuthenticationController(gandalfDatabase *gorm.DB) (authenticationController *AuthenticationController) {
	authenticationController = new(AuthenticationController)
	authenticationController.gandalfDatabase = gandalfDatabase

	return
}

// Login :
func (ac AuthenticationController) Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {

		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := ac.FindOne(user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)
}

// FindOne :
func (ac AuthenticationController) FindOne(email, password string) map[string]interface{} {
	user := models.User{}
	var err error
	if user, err = dao.ReadUserByEmail(ac.gandalfDatabase, email); err != nil {
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
		Tenant: "gandalf",
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
	//resp["user"] = user
	return resp
}
