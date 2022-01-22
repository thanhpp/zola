package core

import (
	"sync"

	"github.com/casbin/casbin/v2"
)

var enforcer *casbin.Enforcer
var enforcerLock sync.RWMutex
var fileMode bool

func insertCasbinRuleFromFile(rawdata [][]string) (err error) {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	for i := 0; i < len(rawdata); i++ {
		if rawdata[i][0] == "p" {
			_, err = enforcer.AddNamedPolicy("p", rawdata[i][1], rawdata[i][2], rawdata[i][3], rawdata[i][4])
			if err != nil {
				return err
			}
		} else if rawdata[i][0] == "g" {
			_, err = enforcer.AddGroupingPolicy(rawdata[i][1], rawdata[i][2], rawdata[i][3])
			if err != nil {
				return err
			}
		}
	}
	enforcer.LoadPolicy()
	return nil
}

func SetFileMode(_fileMode bool) {
	fileMode = _fileMode
}
func SetCasbinEnforcer(e *casbin.Enforcer) {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	enforcer = e
	enforcer.LoadPolicy()
}

func ReloadCasbinEnforcer() {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	enforcer.LoadPolicy()
}

func AddPermissionToUserOnDomain(user, domain, object, act string) (bool, error) {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	//ok, err := enforcer.AddPermissionForUser(user, domain+"/"+object, act)
	ok, err := enforcer.AddPermissionForUser(user, domain, object, act)
	if err != nil {
		return ok, err
	}
	// if change then save policy
	if ok && !fileMode {
		enforcer.SavePolicy()
	}
	return ok, err
}

func DeletePermissionOfUserOnDomain(user, domain, object, act string) (bool, error) {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	ok, err := enforcer.DeletePermissionForUser(user, domain, object, act)
	if err != nil {
		return ok, err
	}
	// if change then save policy
	if ok && !fileMode {
		enforcer.SavePolicy()
	}
	return ok, err
}

func AddPermissionToRoleOnDomain(role, domain, object, act string) (bool, error) {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	ok, err := enforcer.AddPermissionForUser(role, domain, object, act)
	if err != nil {
		return ok, err
	}
	// if change then save policy
	if ok && !fileMode {
		enforcer.SavePolicy()
	}
	return ok, err
}

func DeletePermissionOfRoleOnDomain(role, domain, object, act string) (bool, error) {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	ok, err := enforcer.DeletePermissionForUser(role, domain, object, act)
	if err != nil {
		return ok, err
	}
	// if change then save policy
	if ok && !fileMode {
		enforcer.SavePolicy()
	}
	return ok, err
}

func AddRoleForUser(user, role, domain string) (bool, error) {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	ok, err := enforcer.AddRoleForUserInDomain(user, role, domain)
	if err != nil {
		return ok, err
	}
	// if change then save policy
	if ok && !fileMode {
		enforcer.SavePolicy()
	}
	return ok, err
}

func DeleteRoleForUser(user, role, domain string) (bool, error) {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	ok, err := enforcer.DeleteRoleForUserInDomain(user, role, domain)
	if err != nil {
		return ok, err
	}
	// if change then save policy
	if ok && !fileMode {
		enforcer.SavePolicy()
	}
	return ok, err
}

func DeleteUser(user string) (bool, error) {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	ok, err := enforcer.DeleteUser(user)
	if err != nil {
		return ok, err
	}
	// if change then save policy
	if ok && !fileMode {
		enforcer.SavePolicy()
	}
	return ok, err
}

func DeleteRole(role string) (bool, error) {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	ok, err := enforcer.DeleteRole(role)
	if err != nil {
		return ok, err
	}
	// if change then save policy
	if ok && !fileMode {
		enforcer.SavePolicy()
	}
	return ok, err
}

/*
func DeleteDomain(domain string) (bool, error) {
	enforcerLock.Lock()
	defer enforcerLock.Unlock()
	ok, err := enforcer.DeleteDomains(role)
	if err != nil {
		return ok, err
	}
	// if change then save policy
	if ok && !fileMode{
		enforcer.SavePolicy()
	}
	return ok, err

}
*/

func Enforce(sub, dom, obj, act string) (ok bool, err error) {
	enforcerLock.RLock()
	defer enforcerLock.RUnlock()

	ok, err = enforcer.Enforce(sub, dom, obj, act)
	return
}

func GetUsersForRoleInDomain(name, domain string) (users []string) {
	enforcerLock.RLock()
	defer enforcerLock.RUnlock()

	users = enforcer.GetUsersForRoleInDomain(name, domain)
	return
}

func GetRolesForUserInDomain(name, domain string) (roles []string) {
	enforcerLock.RLock()
	defer enforcerLock.RUnlock()

	roles = enforcer.GetRolesForUserInDomain(name, domain)
	return
}

func GetPermissionsForUserInDomain(name, domain string) (permission [][]string) {
	enforcerLock.RLock()
	defer enforcerLock.RUnlock()

	permission = enforcer.GetPermissionsForUserInDomain(name, domain)
	return
}

func GetAllUsersByDomain(domain string) (users []string) {
	enforcerLock.RLock()
	defer enforcerLock.RUnlock()

	users = enforcer.GetAllUsersByDomain(domain)
	return
}
