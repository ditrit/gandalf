package gandalf

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/cluster/api/dao"
	"github.com/ditrit/gandalf/core/cluster/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type ClusterController struct {
	gandalfDatabase *gorm.DB
}

func NewClusterController(gandalfDatabase *gorm.DB) (clusterController *ClusterController) {
	clusterController = new(ClusterController)
	clusterController.gandalfDatabase = gandalfDatabase

	return
}

func (cc ClusterController) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LIST")
	cluster, err := dao.ListCluster(cc.gandalfDatabase)
	fmt.Println("clusters")
	fmt.Println(cluster)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, cluster)
}

func (cc ClusterController) Create(w http.ResponseWriter, r *http.Request) {
	var cluster models.Cluster
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cluster); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := dao.CreateCluster(cc.gandalfDatabase, cluster); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, cluster)
}

func (cc ClusterController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var cluster models.Cluster
	if cluster, err = dao.ReadCluster(cc.gandalfDatabase, id); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, cluster)
}

func (cc ClusterController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var cluster models.Cluster
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cluster); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	cluster.ID = uint(id)

	if err := dao.UpdateCluster(cc.gandalfDatabase, cluster); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, cluster)
}

func (cc ClusterController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := dao.DeleteCluster(cc.gandalfDatabase, id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
