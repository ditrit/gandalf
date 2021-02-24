package controllers

import (
	"net/http"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
)

// CliController :
type CliController struct {
}

// NewCliController :
func NewCliController() (cliController *CliController) {
	cliController = new(CliController)

	return
}

// Login :
func (cc CliController) Cli(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, "aggregator")
	return

}
