package enforce

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func enforce(u models.User, d models.Domain, o models.Object, a models.Action) bool {
	result = false

	tenantDatabaseClient, _ := gorm.Open("sqlite3", "/home/romainfairant/gandalf/database/tenant1.db")

	objects := GetObjects(tenantDatabaseClient, o)
	for _, object := range objects {
		if InDomainsObject(d, object.Domain) {
			for _, domain := range object.Domain {
				if IsDomainAncestor(tenantDatabaseClient, domain, d) {
					roles := GetRoles(tenantDatabaseClient, u, domain)
					for _, role := range roles {
						rules := GetRules(tenantDatabaseClient, role, domain, object, a)
						for _, rule := range rules {
							if rule.Allow {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}

func InDomainsObject(d models.Domain, domains []models.Domain) bool {
	for _, domain := range domains {
		if d == domain {
			return true
		}
	}
	return false
}

func IsDomainAncestor(tenantDatabaseClient, domain models.Domain, d models.Domain) bool {
	ancestors := models.GetDomainAncestors(tenantDatabaseClient, d.ID)
	for _, ancestor := range ancestors {
		if ancestor == domain {
			return true
		}
	}
	return false
}

func GetObjects(tenantDatabaseClient *gorm.DB, object models.Object) []models.Object {
	return models.GetObjectAncestors(tenantDatabaseClient)
}

func GetRoles(tenantDatabaseClient *gorm.DB, user models.User, domain models.Domain) (roles []models.Role) {
	var perimeters []models.Perimeter
	tenantDatabaseClient.Where("user.id = ? AND domain.id = ?", user.ID, domain.ID).Find(&perimeters)

	for _, perimeter := range perimeters {
		roles = append(roles, perimeter.Role)
	}
	return roles
}

func GetRules(tenantDatabaseClient *gorm.DB, role models.Role, domain models.Domain, object models.Object, action models.Action) (rules []models.Rule) {
	return tenantDatabaseClient.Where("role.id = ? AND domain.id = ? AND object.id = ? AND action.id = ?", role.ID, domain.ID, object.ID, action.ID).Find(&rules)
}
