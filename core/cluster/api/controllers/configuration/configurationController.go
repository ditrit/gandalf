package configuration

import (
	"github.com/jinzhu/gorm"
)

// ConfigurationController :
type ConfigurationController struct {
	gandalfDatabase *gorm.DB
}

// NewConfigurationController :
func NewConfigurationController(gandalfDatabase *gorm.DB) (configurationController *ConfigurationController) {
	configurationController = new(ConfigurationController)
	configurationController.gandalfDatabase = gandalfDatabase

	return
}

/* // Create :
func (cc ConfigurationController) Create(w http.ResponseWriter, r *http.Request) {
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

// Read :
func (cc ConfigurationController) Read(w http.ResponseWriter, r *http.Request) {
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
*/
