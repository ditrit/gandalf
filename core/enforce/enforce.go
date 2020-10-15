package enforce

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func Enforce(tenantDatabaseClient *gorm.DB, u models.User, d models.Domain, o models.Object, a models.Action) bool {

	objects := GetObjects(tenantDatabaseClient, o)
	domains := GetDomains(tenantDatabaseClient, d)
	//fmt.Println("objects")
	//fmt.Println(objects)
	for _, object := range objects {
		//fmt.Println("object")
		//fmt.Println(object)
		for _, odomain := range object.Domains {
			if IsDomainAncestor(tenantDatabaseClient, odomain, d) {
				//fmt.Println("IS")
				/* 	fmt.Println("domains")
				fmt.Println(domains) */
				//if InDomainsObject(d, domains) {
				//	fmt.Println("IN")
				for _, domain := range domains {
					/* 	fmt.Println("ROLE")
					fmt.Println(u)
					fmt.Println(domain) */
					roles := GetRoles(tenantDatabaseClient, u, domain)
					/* 	fmt.Println("roles")
					fmt.Println(roles) */
					for _, role := range roles {
						for _, robject := range objects {
							for _, rdomain := range domains {
								/* fmt.Println("RULE")
								fmt.Println("role")
								fmt.Println(role)
								fmt.Println("domain")
								fmt.Println(tata)
								fmt.Println("object")
								fmt.Println(toto)
								fmt.Println("action")
								fmt.Println(a) */
								rules := GetPermissions(tenantDatabaseClient, role, rdomain, robject, a)
								/* 	fmt.Println("rules")
								fmt.Println(rules) */
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
		}
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
	var authorizations []models.Authorization
	tenantDatabaseClient.Where("user_id = ? AND domain_id = ?", user.ID, domain.ID).Preload("Role").Find(&authorizations)

	for _, authorization := range authorizations {
		roles = append(roles, authorization.Role)
	}
	return roles
}

func GetPermissions(tenantDatabaseClient *gorm.DB, role models.Role, domain models.Domain, object models.Object, action models.Action) (permissions []models.Permission) {
	tenantDatabaseClient.Where("role_id = ? AND domain_id = ? AND object_id = ? AND (action_id = ? OR action_id = 1) ", role.ID, domain.ID, object.ID, action.ID).Find(&permissions)

	return
}
