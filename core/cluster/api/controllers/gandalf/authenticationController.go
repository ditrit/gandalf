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

type AuthenticationController struct {
	gandalfDatabase *gorm.DB
}

func NewAuthenticationController(gandalfDatabase *gorm.DB) (authenticationController *AuthenticationController) {
	authenticationController = new(AuthenticationController)
	authenticationController.gandalfDatabase = gandalfDatabase

	return
}

func (ac AuthenticationController) Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		fmt.Println(err)
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	fmt.Println("FIND")
	fmt.Println(user.Email)
	fmt.Println(user.Password)
	resp := ac.FindOne(user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)
}

//TODO REVOIR
func (ac AuthenticationController) FindOne(email, password string) map[string]interface{} {
	fmt.Println("FINDONE")
	fmt.Println(ac.gandalfDatabase)
	user := models.User{}
	var err error
	if user, err = dao.ReadUserByEmail(ac.gandalfDatabase, email); err != nil {
		fmt.Println(err)
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	fmt.Println("FIND1")
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		fmt.Println(errf)
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}
	fmt.Println("FIND2")
	tk := &apimodels.Claims{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Role:   user.Role.Name,
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
