package enforce

import (
	"fmt"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func Enforce(tenantDatabaseClient *gorm.DB, u models.User, d models.Domain, o models.Object, a models.Action) bool {

	objects := GetObjects(tenantDatabaseClient, o)
	fmt.Println(objects)
	for _, object := range objects {
		fmt.Println("object")
		fmt.Println(object)
		//if InDomainsObject(d, object.Domain) {
		//fmt.Println("IN")
		for _, odomain := range object.Domain {
			if IsDomainAncestor(tenantDatabaseClient, odomain, d) {
				fmt.Println("IS")
				domains := GetDomains(tenantDatabaseClient, odomain)
				fmt.Println("domains")
				fmt.Println(domains)
				for _, domain := range domains {
					fmt.Println("ROLE")
					fmt.Println(u)
					fmt.Println(domain)
					roles := GetRoles(tenantDatabaseClient, u, domain)
					fmt.Println("roles")
					fmt.Println(roles)
					for _, role := range roles {
						fmt.Println("RULE")
						fmt.Println(role)
						fmt.Println(domain)
						fmt.Println(object)
						fmt.Println(a)
						rules := GetRules(tenantDatabaseClient, role, domain, object, a)
						fmt.Println("rules")
						fmt.Println(rules)
						for _, rule := range rules {
							if rule.Allow {
								return true
							}
						}
					}
				}
			}
		}
		//}
	}
	return false
}

func InDomainsObject(d models.Domain, domains []models.Domain) bool {
	for _, domain := range domains {
		if d.Name == domain.Name {
			return true
		}
	}
	return false
}

func IsDomainAncestor(tenantDatabaseClient *gorm.DB, domain models.Domain, d models.Domain) bool {
	ancestors := models.GetDomainAncestors(tenantDatabaseClient, d.ID)
	for _, ancestor := range ancestors {
		if ancestor.Name == domain.Name {
			return true
		}
	}
	return false
}

func GetObjects(tenantDatabaseClient *gorm.DB, object models.Object) []models.Object {
	return models.GetObjectAncestors(tenantDatabaseClient, object.ID)
}

func GetDomains(tenantDatabaseClient *gorm.DB, domain models.Domain) []models.Domain {
	return models.GetDomainAncestors(tenantDatabaseClient, domain.ID)
}

func GetRoles(tenantDatabaseClient *gorm.DB, user models.User, domain models.Domain) (roles []models.Role) {
	var perimeters []models.Perimeter
	tenantDatabaseClient.Where("user_id = ? AND domain_id = ?", user.ID, domain.ID).Preload("Role").Find(&perimeters)

	for _, perimeter := range perimeters {
		roles = append(roles, perimeter.Role)
	}
	return roles
}

func GetRules(tenantDatabaseClient *gorm.DB, role models.Role, domain models.Domain, object models.Object, action models.Action) (rules []models.Rule) {
	tenantDatabaseClient.Where("role_id = ? AND domain_id = ? AND object_id = ? AND action_id = ?", role.ID, domain.ID, object.ID, action.ID).Find(&rules)

	return
}
