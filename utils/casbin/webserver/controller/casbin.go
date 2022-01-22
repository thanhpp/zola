package controller

import (
	"fmt"
	"net/http"

	"bitbucket.org/tysud/gt-casbin/core"
	"github.com/gin-gonic/gin"
)

var (
	ErrPermissionIsAdded   = fmt.Errorf("Permission is already added to user/role")
	ErrPermissionIsDeleted = fmt.Errorf("Permission is already deleted to user/role")
	ErrUserHasRoleAlready  = fmt.Errorf("User has this role already")
	ErrUserDoesntHaveRole  = fmt.Errorf("User doesn't have this role")
	ErrNoUser              = fmt.Errorf("User doesn't exist")
	ErrNilUser             = fmt.Errorf("Nil user")
	ErrNilRole             = fmt.Errorf("Nil role")
	ErrNoRole              = fmt.Errorf("Role doesn't exist")
)

type CasbinHandler struct{}

func (handlers *CasbinHandler) CheckAuthorization(c *gin.Context) {
	type DTO struct {
		User   string `json:"user,omitempty"`
		Domain string `json:"domain,omitempty"`
		Object string `json:"object,omitempty"`
		Action string `json:"action,omitempty"`
	}

	var req DTO
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}
	if req.User == "" {
		c.Error(ErrNilUser)
		c.Status(http.StatusBadRequest)
		return
	}

	ok, err := core.Enforce(
		req.User,
		req.Domain,
		req.Object,
		req.Action,
	)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if ok {
		c.Status(http.StatusOK)
	} else {
		// add default role
		defaultRole := core.GetConfig().DefaultRole
		for i := 0; i < len(defaultRole); i++ {
			_, err := core.AddRoleForUser(req.User, defaultRole[i], req.Domain)
			if err != nil {
				c.Error(err)
				c.Status(http.StatusInternalServerError)
				return
			}
		}
		ok, err := core.Enforce(
			req.User,
			req.Domain,
			req.Object,
			req.Action,
		)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}

		if ok {
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusUnauthorized)
		}
	}
}

func (handlers *CasbinHandler) AddPermissionToRoleOnDomain(c *gin.Context) {
	type DTO struct {
		Role   string `json:"role,omitempty"`
		Domain string `json:"domain,omitempty"`
		Object string `json:"object,omitempty"`
		Action string `json:"action,omitempty"`
	}

	var req DTO
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}
	if req.Role == "" {
		c.Error(ErrNilRole)
		c.Status(http.StatusBadRequest)
		return
	}

	ok, err := core.AddPermissionToRoleOnDomain(
		req.Role,
		req.Domain,
		req.Object,
		req.Action,
	)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if !ok {
		c.Error(ErrPermissionIsAdded)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (handlers *CasbinHandler) DeletePermissionOfRoleOnDomain(c *gin.Context) {
	type DTO struct {
		Role   string `json:"role,omitempty"`
		Domain string `json:"domain,omitempty"`
		Object string `json:"object,omitempty"`
		Action string `json:"action,omitempty"`
	}

	var req DTO
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}
	if req.Role == "" {
		c.Error(ErrNilRole)
		c.Status(http.StatusBadRequest)
		return
	}

	ok, err := core.DeletePermissionOfRoleOnDomain(
		req.Role,
		req.Domain,
		req.Object,
		req.Action,
	)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if !ok {
		c.Error(ErrPermissionIsDeleted)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (handlers *CasbinHandler) AddRoleToUser(c *gin.Context) {
	type DTO struct {
		Role   string `json:"role,omitempty"`
		User   string `json:"user,omitempty"`
		Domain string `json:"domain,omitempty"`
	}

	var req DTO
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}
	if req.User == "" {
		c.Error(ErrNilUser)
		c.Status(http.StatusBadRequest)
		return
	}
	if req.Role == "" {
		c.Error(ErrNilRole)
		c.Status(http.StatusBadRequest)
		return
	}

	ok, err := core.AddRoleForUser(req.User, req.Role, req.Domain)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if !ok {
		c.Error(ErrUserHasRoleAlready)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (handlers *CasbinHandler) DeleteRoleOfUser(c *gin.Context) {
	type DTO struct {
		Role   string `json:"role,omitempty"`
		User   string `json:"user,omitempty"`
		Domain string `json:"domain,omitempty"`
	}

	var req DTO
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}
	if req.User == "" {
		c.Error(ErrNilUser)
		c.Status(http.StatusBadRequest)
		return
	}
	if req.Role == "" {
		c.Error(ErrNilRole)
		c.Status(http.StatusBadRequest)
		return
	}

	ok, err := core.DeleteRoleForUser(req.User, req.Role, req.Domain)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if !ok {
		c.Error(ErrUserDoesntHaveRole)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (handlers *CasbinHandler) AddPermissionToUserOnDomain(c *gin.Context) {
	type DTO struct {
		User   string `json:"user,omitempty"`
		Domain string `json:"domain,omitempty"`
		Object string `json:"object,omitempty"`
		Action string `json:"action,omitempty"`
	}

	var req DTO
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}
	if req.User == "" {
		c.Error(ErrNilUser)
		c.Status(http.StatusBadRequest)
		return
	}

	ok, err := core.AddPermissionToUserOnDomain(
		req.User,
		req.Domain,
		req.Object,
		req.Action,
	)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if !ok {
		c.Error(ErrPermissionIsAdded)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (handlers *CasbinHandler) DeletePermissionOfUserOnDomain(c *gin.Context) {
	type DTO struct {
		User   string `json:"user,omitempty"`
		Domain string `json:"domain,omitempty"`
		Object string `json:"object,omitempty"`
		Action string `json:"action,omitempty"`
	}

	var req DTO
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}
	if req.User == "" {
		c.Error(ErrNilUser)
		c.Status(http.StatusBadRequest)
		return
	}

	ok, err := core.DeletePermissionOfUserOnDomain(
		req.User,
		req.Domain,
		req.Object,
		req.Action,
	)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if !ok {
		c.Error(ErrPermissionIsDeleted)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (handlers *CasbinHandler) DeleteUser(c *gin.Context) {
	type DTO struct {
		User string `json:"user,omitempty"`
	}

	var req DTO
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}
	if req.User == "" {
		c.Error(ErrNilUser)
		c.Status(http.StatusBadRequest)
		return
	}

	ok, err := core.DeleteUser(req.User)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if !ok {
		c.Error(ErrNoUser)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (handlers *CasbinHandler) DeleteRole(c *gin.Context) {
	type DTO struct {
		Role string `json:"role,omitempty"`
	}

	var req DTO
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}
	if req.Role == "" {
		c.Error(ErrNilRole)
		c.Status(http.StatusBadRequest)
		return
	}

	ok, err := core.DeleteRole(req.Role)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if !ok {
		c.Error(ErrNoRole)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (handlers *CasbinHandler) GetUsersForRoleInDomain(c *gin.Context) {
	role := c.Query("role")
	domain := c.Query("domain")

	if role == "" {
		c.Error(ErrNilRole)
		c.Status(http.StatusBadRequest)
		return
	}

	users := core.GetUsersForRoleInDomain(role, domain)
	c.JSON(http.StatusOK, users)
}

func (handlers *CasbinHandler) GetRolesForUserInDomain(c *gin.Context) {
	user := c.Query("user")
	domain := c.Query("domain")
	if user == "" {
		c.Error(ErrNilUser)
		c.Status(http.StatusBadRequest)
		return
	}

	roles := core.GetRolesForUserInDomain(user, domain)
	c.JSON(http.StatusOK, roles)
}

func (handlers *CasbinHandler) GetPermissionsForUserInDomain(c *gin.Context) {
	user := c.Query("user")
	domain := c.Query("domain")
	if user == "" {
		c.Error(ErrNilUser)
		c.Status(http.StatusBadRequest)
		return
	}

	permission := core.GetPermissionsForUserInDomain(user, domain)
	c.JSON(http.StatusOK, permission)
}

func (handlers *CasbinHandler) GetAllUsersByDomain(c *gin.Context) {
	domain := c.Query("domain")

	user := core.GetAllUsersByDomain(domain)
	c.JSON(http.StatusOK, user)
}
