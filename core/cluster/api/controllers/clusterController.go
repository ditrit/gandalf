package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/cluster/api/dao"
	"github.com/ditrit/gandalf/core/cluster/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/gorilla/mux"
)

// ClusterController :
type ClusterController struct {
	databaseConnection *database.DatabaseConnection
}

// NewClusterController :
func NewClusterController(databaseConnection *database.DatabaseConnection) (clusterController *ClusterController) {
	clusterController = new(ClusterController)
	clusterController.databaseConnection = databaseConnection

	return
}

// List :
func (cc ClusterController) List(w http.ResponseWriter, r *http.Request) {
	cluster, err := dao.ListCluster(cc.databaseConnection.GetGandalfDatabaseClient())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, cluster)
}

// Create :
func (cc ClusterController) Create(w http.ResponseWriter, r *http.Request) {
	var cluster models.Cluster
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cluster); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := dao.CreateCluster(cc.databaseConnection.GetGandalfDatabaseClient(), cluster); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, cluster)
}

// DeclareMember :
func (cc ClusterController) DeclareMember(w http.ResponseWriter, r *http.Request) {

	cluster, err := dao.ReadFirstCluster(cc.databaseConnection.GetGandalfDatabaseClient())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	var newCluster models.Cluster
	newCluster.LogicalName = cluster.LogicalName
	newCluster.Secret = utils.GenerateHash()

	if err := dao.CreateCluster(cc.databaseConnection.GetGandalfDatabaseClient(), newCluster); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, newCluster)
}

// Read :
func (cc ClusterController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var cluster models.Cluster
	if cluster, err = dao.ReadCluster(cc.databaseConnection.GetGandalfDatabaseClient(), id); err != nil {
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

// Update :
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

	if err := dao.UpdateCluster(cc.databaseConnection.GetGandalfDatabaseClient(), cluster); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, cluster)
}

// Delete :
func (cc ClusterController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := dao.DeleteCluster(cc.databaseConnection.GetGandalfDatabaseClient(), id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
