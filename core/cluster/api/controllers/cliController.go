package controllers

import (
	"net/http"

	"github.com/ditrit/gandalf/core/cluster/api/utils"
)

// CliController :
type CliController struct {
}

// NewCliController :
func NewCliController() (cliController *CliController) {

	return
}

// Login :
func (cc CliController) Cli(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, "cluster")
	return

}
