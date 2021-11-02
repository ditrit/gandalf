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

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	"github.com/ditrit/gandalf/core/models"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/gorilla/mux"
)

func CreateDomain(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		parentDomainId, err := uuid.Parse(vars["domainId"])

		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		var parentDomain models.Domain
		if parentDomain, err = dao.ReadDomain(database, parentDomainId); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		var domain *models.Domain
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&domain); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		defer r.Body.Close()
		if err := dao.CreateDomain(database, domain, parentDomainId); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Library
		for _, library := range parentDomain.Libraries {
			dao.AddDomainLibrary(database, *domain, library)
		}

		// Env
		for _, environment := range parentDomain.Environments {
			dao.AddDomainEnvironment(database, *domain, environment)
		}

		// Tag
		for _, tag := range parentDomain.Tags {
			dao.AddDomainTag(database, *domain, tag)
		}

		// Members
		for _, authorization := range parentDomain.Authorizations {
			currentAuthorization := new(models.Authorization)
			currentAuthorization.User = authorization.User
			currentAuthorization.Role = authorization.Role
			currentAuthorization.Domain = *domain
			dao.CreateAuthorization(database, currentAuthorization)
		}

		utils.RespondWithJSON(w, http.StatusCreated, domain)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func DeleteDomain(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := uuid.Parse(vars["domainId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		//TEST
		var domains []models.Domain
		database.Unscoped().Find(&domains)
		fmt.Println("domains")
		for _, domain := range domains {
			fmt.Println(domain)

		}

		if err := dao.DeleteDomain(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func GetDomainTree(w http.ResponseWriter, r *http.Request) {
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		domain, err := dao.TreeDomain(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, domain)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func GetDomainById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := uuid.Parse(vars["domainId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var domain models.Domain
		if domain, err = dao.ReadDomain(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, domain)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func GetDomainByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["domainName"]

	var domain models.Domain
	var err error
	if domain, err = dao.ReadDomainByName(utils.DatabaseConnection.GetTenantDatabaseClient(), name); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, domain)
}

func ListDomain(w http.ResponseWriter, r *http.Request) {
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		domains, err := dao.ListDomain(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, domains)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func UpdateDomain(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := uuid.Parse(vars["domainId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var domain models.Domain
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&domain); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		domain.ID = id

		if err := dao.UpdateDomain(database, domain); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, domain)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func ListDomainTag(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		idDomain, err := uuid.Parse(vars["domainId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var domain models.Domain
		if domain, err = dao.ReadDomain(database, idDomain); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		tags, err := dao.ListDomainTag(database, domain)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, tags)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func CreateDomainTag(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		idDomain, err := uuid.Parse(vars["domainId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var domain models.Domain
		if domain, err = dao.ReadDomain(database, idDomain); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		idTag, err := uuid.Parse(vars["tagId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var tag models.Tag
		if tag, err = dao.ReadTag(database, idTag); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		if err := dao.AddDomainTag(database, domain, tag); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func DeleteDomainTag(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		idDomain, err := uuid.Parse(vars["domainId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var domain models.Domain
		if domain, err = dao.ReadDomain(database, idDomain); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		idTag, err := uuid.Parse(vars["tagId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var tag models.Tag
		if tag, err = dao.ReadTag(database, idTag); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		if err := dao.RemoveDomainTag(database, domain, tag); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

///

func ListDomainEnvironment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		idDomain, err := uuid.Parse(vars["domainId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var domain models.Domain
		if domain, err = dao.ReadDomain(database, idDomain); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		environments, err := dao.ListDomainEnvironment(database, domain)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, environments)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func CreateDomainEnvironment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		idDomain, err := uuid.Parse(vars["domainId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var domain models.Domain
		if domain, err = dao.ReadDomain(database, idDomain); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		idEnvironment, err := uuid.Parse(vars["environmentId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var environment models.Environment
		if environment, err = dao.ReadEnvironment(database, idEnvironment); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		if err := dao.AddDomainEnvironment(database, domain, environment); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func DeleteDomainEnvironment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		idDomain, err := uuid.Parse(vars["domainId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var domain models.Domain
		if domain, err = dao.ReadDomain(database, idDomain); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		idEnvironment, err := uuid.Parse(vars["environmentId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var environment models.Environment
		if environment, err = dao.ReadEnvironment(database, idEnvironment); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		if err := dao.RemoveDomainEnvironment(database, domain, environment); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

///

func ListDomainLibrary(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		idDomain, err := uuid.Parse(vars["domainId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var domain models.Domain
		if domain, err = dao.ReadDomain(database, idDomain); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		libraries, err := dao.ListDomainLibrary(database, domain)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, libraries)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func CreateDomainLibrary(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		idDomain, err := uuid.Parse(vars["domainId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var domain models.Domain
		if domain, err = dao.ReadDomain(database, idDomain); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		idLibrary, err := uuid.Parse(vars["libraryId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var library models.Library
		if library, err = dao.ReadLibrary(database, idLibrary); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		if err := dao.AddDomainLibrary(database, domain, library); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func DeleteDomainLibrary(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		idDomain, err := uuid.Parse(vars["domainId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var domain models.Domain
		if domain, err = dao.ReadDomain(database, idDomain); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		idLibrary, err := uuid.Parse(vars["libraryId"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		var library models.Library
		if library, err = dao.ReadLibrary(database, idLibrary); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		if err := dao.RemoveDomainLibrary(database, domain, library); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
