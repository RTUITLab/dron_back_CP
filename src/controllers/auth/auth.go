package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/0B1t322/CP-Rosseti-Back/ent"
	"github.com/0B1t322/CP-Rosseti-Back/ent/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"

	e "github.com/0B1t322/CP-Rosseti-Back/models/err"
	log "github.com/sirupsen/logrus"
)

func init() {
	_ = e.Error{}
}

type AuthController struct {
	AccessSecret		string
	RefreshSecret		string

	Client *ent.Client
}

func New(
	client 			*ent.Client,
	AccessSecret	string,
	RefreshSecret	string,
	) *AuthController {
	return &AuthController{
		Client: client,
		AccessSecret: AccessSecret,
		RefreshSecret: RefreshSecret,
	}
}

type LoginReq struct {
	Login		string	`json:"login"`
	Password	string	`json:"password"`
}

type LoginResp struct {
	AccessToken		string	`json:"accesstoken"`
	RefreshToken	string	`json:"refreshToken"`
}

// Login
// 
// @Tags auth
// 
// @Summary login
// 
// @Description login
// 
// @Router /v1/auth/login [post]
// 
// @Param userInfo body auth.LoginReq true "user info"
// 
// @Accept json
// 
// @Produce json
// 
// @Success 200 {object} auth.LoginResp
// 
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 500 {object} e.Error "internal"
// 
// @Failure 401 {object} e.Error "not auth"
func (a AuthController) Login(c *gin.Context) {
	var req LoginReq
	{
		if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("Unexpected body"))
			c.Abort()
			return
		}
	}

	getUser, err := a.Client.User.Query().Where(user.Login(req.Login)).WithRole().Only(c)
	if ent.IsNotFound(err) {
		c.JSON(http.StatusUnauthorized, e.FromString("User not found"))
		c.Abort()
		return
	} else if err != nil {
		log.WithFields(
			log.Fields{
				"pkg": "controllers/user",
				"func": "Login",
				"err": err,
			},
		).Error("Failed create user")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to login"))
		c.Abort()
		return
	}

	if getUser.Password != req.Password {
		c.JSON(http.StatusUnauthorized, e.FromString("Invalid password"))
		c.Abort()
		return
	}

	td, err := a.CreateToken(getUser.ID, getUser.Edges.Role.Role)
	if err != nil {
		log.WithFields(
			log.Fields{
				"pkg": "controllers/user",
				"func": "Login",
				"err": err,
			},
		).Error("Failed create user")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to login"))
		c.Abort()
	}

	c.JSON(http.StatusOK, LoginResp{AccessToken: td.AccessToken, RefreshToken: td.RefreshToken})

}

type RefreshReq struct {
	RefreshToken	string	`json:"refreshToken"`
}

// Куакуыр
// 
// @Tags auth
// 
// @Summary refresh
// 
// @Description refresh
// 
// @Router /v1/auth/refresh [post]
// 
// @Param token body auth.RefreshReq true "token"
// 
// @Accept json
// 
// @Produce json
// 
// @Success 200 {object} auth.LoginResp
// 
// @Failure 400 {object} e.Error "some user error"
//
// @Failure 500 {object} e.Error "internal"
// 
// @Failure 401 {object} e.Error "not auth"
func (a AuthController) Refresh(c *gin.Context) {
	var req RefreshReq
	{
		if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
			c.JSON(http.StatusBadRequest, e.FromString("Unexpected body"))
			c.Abort()
			return
		}
	}

	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		   return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.RefreshSecret), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, e.FromString("Refresh token expired"))
		return
	}
	 //is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, e.FromError(err))
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	uid := claims["user_id"].(float64)

	user, err := a.Client.User.Query().WithRole().Where(user.ID(int(uid))).Only(c)
	if err != nil {
		log.WithFields(
			log.Fields{
				"pkg": "controllers/user",
				"func": "Refresh",
				"err": err,
			},
		).Error("Failed create user")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to refresh"))
		c.Abort()
		return
	}

	td, err := a.CreateToken(user.ID, user.Edges.Role.Role)
	if err != nil {
		log.WithFields(
			log.Fields{
				"pkg": "controllers/user",
				"func": "Refresh",
				"err": err,
			},
		).Error("Failed create user")
		c.JSON(http.StatusInternalServerError, e.FromString("Failed to refresh"))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, LoginResp{AccessToken: td.AccessToken, RefreshToken: td.RefreshToken})
}

type TokensDetails struct {
	AccessToken		string
	RefreshToken	string
	AccessUuid		string
	RefreshUuid		string
	AtExpires		int64
	RtExpires		int64
}

func (a AuthController) CreateToken(userid int, role string) (*TokensDetails, error) {
	td := &TokensDetails{}
	td.AtExpires = time.Now().Add(time.Hour * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["role"] = role
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(a.AccessSecret))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rtClaims["role"] = role
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(a.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (a AuthController) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := a.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   //Make sure that the token method conform to "SigningMethodHMAC"
	   if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		  return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	   }
	   return []byte(a.AccessSecret), nil
	})
	if err != nil {
	   return nil, err
	}
	return token, nil
}

func (a AuthController) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
	   return strArr[1]
	}
	return ""
}

func (a AuthController) TokenValid(c *gin.Context, r *http.Request) error {
	token, err := a.VerifyToken(r)
	if err != nil {
	   return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
	   return err
	}
	
	claims := token.Claims.(jwt.MapClaims)
	user_id := int(claims["user_id"].(float64))
	c.Set("user_id", user_id)

	return nil
}

func (a AuthController) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := a.TokenValid(c, c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, e.FromError(err))
			c.Abort()
			return
		}
		c.Next()
	}
}