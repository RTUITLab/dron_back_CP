package user

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/0B1t322/CP-Rosseti-Back/ent"
	"github.com/0B1t322/CP-Rosseti-Back/ent/role"
	"github.com/0B1t322/CP-Rosseti-Back/ent/user"
	e "github.com/0B1t322/CP-Rosseti-Back/models/err"
	rolemodel "github.com/0B1t322/CP-Rosseti-Back/models/role"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	_ = e.Error{}
}

type UserController struct {
	Client *ent.Client
}

func New(client *ent.Client) *UserController {
	return &UserController{
		Client: client,
	}
}


type CreateUserReq struct {
	Login		string		`json:"login"`
	Password	string		`json:"password"`
	Role		string		`json:"role" enums:"student,teacher,admin"`
}

type CreateUserResp struct {
	ID			int			`json:"int"`
	Login		string		`json:"login"`
	Role		string		`json:"role"`
}

// CreateUser
// 
// @Tags user
// 
// @Summary create user
// 
// @Description create user with role
// 
// @Description that can do only admin
// 
// @Router /v1/user [post]
// 
// @Security ApiKeyAuth
// 
// @Param user body user.CreateUserReq true "User info"
// 
// @Accept json
// 
// @Produce json
// 
// @Success 201 {object} user.CreateUserResp
// 
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 500 {object} e.Error "internal"
// 
// @Failure 401 {object} e.Error "not auth"
func (u UserController) CreateUser(c *gin.Context) {
	var req CreateUserReq
	{
		if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("Unexpected body"))
			c.Abort()
			return
		}
	}

	// Check role is Exist
	roleId, err := u.getRoleID(c, req.Role)
	if ent.IsNotFound(err) {
		c.JSON(http.StatusBadRequest, e.FromString("Role not found"))
		c.Abort()
		return
	} else if err == rolemodel.ErrRoleNotValid {
		c.JSON(http.StatusBadRequest, e.FromString("Not valid role"))
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg": "controllers/user",
				"func": "CreateUser",
				"err": err,
			},
		).Error("Failed create user")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to create user"))
		c.Abort()
		return
	}

	created, err := u.Client.User.Create().
		SetLogin(req.Login).
		SetPassword(req.Password).
		SetRoleID(roleId).Save(c)
	
	if ent.IsConstraintError(err) {
		c.JSON(http.StatusBadRequest, e.FromString("User with this login exist"))
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg": "controllers/user",
				"func": "CreateUser",
				"err": err,
			},
		).Error("Failed create user")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to create user"))
		c.Abort()
		return
	}

	c.JSON(
		http.StatusCreated, 
		CreateUserResp{
			ID: created.ID,
			Login: created.Login,
			Role: req.Role,
		},
	)
}

// DeleteUser
// 
// @Tags user
// 
// @Summary delete user
// 
// @Description delete user
// 
// @Description that can do only admin
// 
// @Router /v1/user/{id} [delete]
// 
// @Security ApiKeyAuth
// 
// @Param id path integer true "User id"
// 
// @Produce json
// 
// @Success 200
// 
// @Failure 400 {object} e.Error "some user error"
// 
// @Failure 404 {object} e.Error "user not found"
//
// @Failure 500 {object} e.Error "internal"
// 
// @Failure 401 {object} e.Error "not auth"
func (u UserController) DeleteUser(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, e.FromString("ID is not integer"))
		c.Abort()
		return
	}

	if err := u.Client.User.DeleteOneID(id).Exec(c); ent.IsNotFound(err) {
		c.JSON(http.StatusNotFound, e.FromString("User not found"))
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg": "controllers/user",
				"func": "DeleteUser",
				"err": err,
			},
		).Error("Failed to delete user")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to delete user"))
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}

type UpdateUserRequest struct {
	Password	*string		`json:"password",omitempty`
}

// UpdateUser
// 
// @Tags user
// 
// @Summary update user
// 
// @Description update user
// 
// @Description it can do this user or admin
// 
// @Router /v1/user/{id} [put]
// 
// @Security ApiKeyAuth
// 
// @Param id path integer true "User id"
// 
// @Produce json
// 
// @Success 200
// 
// @Failure 400 {object} e.Error "some user error"
// 
// @Failure 404 {object} e.Error "user not found"
//
// @Failure 500 {object} e.Error "internal"
// 
// @Failure 401 {object} e.Error "not auth"
func (u UserController) UpdateUser(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, e.FromString("ID is not integer"))
		c.Abort()
		return
	}

	var req UpdateUserRequest
	{
		if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("Unexpected body"))
			c.Abort()
			return
		}
	}

	builder := u.Client.User.UpdateOneID(id)
	{
		if req.Password != nil {
			builder.SetPassword(*req.Password)
		}
	}

	if err := builder.Exec(c); ent.IsNotFound(err) {
		c.JSON(http.StatusNotFound, e.FromString("User not found"))
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg": "controllers/user",
				"func": "UpdateUser",
				"err": err,
			},
		).Error("Failed to delete user")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to delete user"))
		c.Abort()
		return
	}
}

type GetUserResponce struct {
	ID		int		`json:"id"`
	Login	string	`json:"login"`
	Role	string	`json:"role"`
}

type GetUsersResponce struct {
	Users []GetUserResponce	`json:"users"`
}

// GetUsers
// 
// @Tags user
// 
// @Summary get users
// 
// @Description get users
// 
// @Router /v1/user [get]
// 
// @Security ApiKeyAuth
// 
// @Param login query string false "match user on login field"
// 
// @Param role query string false "match user on role"
// 
// @Param limit query int false "limit"
// 
// @Param offset query int false "offset"
// 
// @Produce json
// 
// @Success 200 {object} user.GetUsersResponce "users"
// 
// @Failure 400 {object} e.Error "some user error"
// 
// @Failure 404 {object} e.Error "user not found"
//
// @Failure 500 {object} e.Error "internal"
// 
// @Failure 401 {object} e.Error "not auth"
func (u UserController) GetUsers(c *gin.Context) {
	var (
		loginMatch 	= c.Query("login")
		roleMatch 	= c.Query("role")
		limitStr	= c.Query("limit")
		offsetStr	= c.Query("offset")
	)

	builder := u.Client.User.Query().WithRole()
	
	if loginMatch != "" {
		builder.Where(user.LoginContains(loginMatch))
	}

	log.Info("roleMatch:", roleMatch)

	if roleMatch != "" {
		builder.Where(user.HasRoleWith(role.Role(roleMatch)))
	}

	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			// Ignore
		} else {
			builder.Offset(offset)
		}
	}

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			// Ignore
		} else {
			builder.Limit(limit)
		}
	}

	var users GetUsersResponce
	{
		get, err := builder.All(c)
		if err != nil {
			log.WithFields(
				log.Fields{
					"pkg": "controllers/user",
					"func": "GetUsers",
					"err": err,
				},
			).Error("Failed to get users")
			c.JSON(http.StatusInternalServerError, e.FromString("Failed to get users"))
			c.Abort()
			return
		}

		for _, user := range get {
			users.Users = append(users.Users, GetUserResponce{ID: user.ID,Login: user.Login, Role: user.Edges.Role.Role})
		}
	}

	c.JSON(http.StatusOK, users)
}

func (u UserController) getRoleID(ctx context.Context, roleStr string) (int, error) {
	if err := rolemodel.ValidRole(roleStr); err != nil {
		return -1, err
	}

	return u.Client.Role.Query().Where(role.Role(roleStr)).OnlyID(ctx)
}

func (u UserController) CreateUserOnStartUp() error {
	role, err := u.Client.Role.Query().Where(role.Role("admin")).Only(context.Background())
	if err != nil {
		return err
	}

	if _, err := u.Client.User.Create().SetRole(role).SetPassword("root").SetLogin("root").
		Save(context.Background()); ent.IsConstraintError(err) {
			// Pass
		} else if ent.IsNotFound(err) {
			// Passs
		} else if err != nil {
			return err
		}

	return nil
}