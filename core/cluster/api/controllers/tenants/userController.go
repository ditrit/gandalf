package tenants

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/cluster/api/dao"
	"github.com/ditrit/gandalf/core/cluster/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// UserController :
type UserController struct {
	mapDatabase  map[string]*gorm.DB
	databasePath string
}

// NewUserController :
func NewUserController(mapDatabase map[string]*gorm.DB, databasePath string) (userController *UserController) {
	userController = new(UserController)
	userController.mapDatabase = mapDatabase
	userController.databasePath = databasePath

	return
}

// List :
func (uc UserController) List(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(uc.mapDatabase, uc.databasePath, tenant)

	users, err := dao.ListUser(database)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, users)
}

// Create :
func (uc UserController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(uc.mapDatabase, uc.databasePath, tenant)

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := dao.CreateUser(database, user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}

// Read :
func (uc UserController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(uc.mapDatabase, uc.databasePath, tenant)

	id, err := strconv.Atoi(vars["id"])
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
}

// Update :
func (uc UserController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(uc.mapDatabase, uc.databasePath, tenant)

	id, err := strconv.Atoi(vars["id"])
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
}

// Delete :
func (uc UserController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(uc.mapDatabase, uc.databasePath, tenant)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := dao.DeleteUser(database, id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
