/*
 * Swagger Gandalf
 *
 * This is a sample Petstore server.  You can find  out more about Swagger at  [http://swagger.io](http://swagger.io) or on  [irc.freenode.net, #swagger](http://swagger.io/irc/).
 *
 * API version: 1.0.0-oas3
 * Contact: romain.fairant@orness.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ditrit/gandalf/core/models"
	"github.com/golang-jwt/jwt/v4"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	apimodels "github.com/ditrit/gandalf/core/aggregator/api/models"
	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		var user models.User
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		user.Password = models.HashAndSaltPassword(user.Password)

		if err := dao.CreateUser(database, user); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, user)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["userId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err := dao.DeleteUser(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["userId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}
		var user models.User
		if user, err = dao.ReadUser(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, user)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func GetUserByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["userName"]
	var err error
	var user models.User
	if user, err = dao.ReadUserByName(utils.DatabaseConnection.GetTenantDatabaseClient(), name); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}

func ListUser(w http.ResponseWriter, r *http.Request) {
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		users, err := dao.ListUser(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, users)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
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
		if user, err = dao.ReadUserByEmail(database, ruser.Email); err != nil {

			//var resp = map[string]interface{}{"status": false, "message": "Username not found"}
			//return resp
			utils.RespondWithError(w, http.StatusNotFound, "Email not found")
			return
		}

		errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ruser.Password))
		if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			//var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
			//return resp
			//return http.StatusUnauthorized, "", "Invalid login credentials. Please try again"
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid login credentials. Please try again")
			return
		}

		expirationTime := time.Now().Add(time.Hour * 1)

		claims := &apimodels.Claims{
			UserID: user.ID,
			Name:   user.Name,
			Email:  user.Email,
			StandardClaims: &jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(utils.GetJwtKey())
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"accessToken": tokenString, "user": user})
		return
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		var ruser models.User
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&ruser); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		ruser.Password = models.HashAndSaltPassword(ruser.Password)

		if err := dao.CreateUser(database, ruser); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		user := models.User{}
		var err error
		if user, err = dao.ReadUserByEmail(database, ruser.Email); err != nil {

			//var resp = map[string]interface{}{"status": false, "message": "Username not found"}
			//return resp
			utils.RespondWithError(w, http.StatusNotFound, "Email not found")
			return
		}

		expirationTime := time.Now().Add(time.Hour * 1)

		claims := &apimodels.Claims{
			UserID: user.ID,
			Name:   user.Name,
			Email:  user.Email,
			StandardClaims: &jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(utils.GetJwtKey())
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"accessToken": tokenString})
		return
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("USER LOGOUT")

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["userId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var user models.User
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		user.ID = uint(id)

		if err := dao.UpdateUser(database, user); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, user)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
