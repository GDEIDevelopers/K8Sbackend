package auth

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	RefreshTokenExpired = 24 * time.Hour * 3
	AccessTokenExpired  = 2 * time.Hour
)

// 是否登录 godoc
// @Summary 是否登录
// @Schemes
// @Description 是否登录
// @Tags auth
// @Accept json
// @Produce json
// @Param   token  header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[model.GetUserResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /isvalid [get]
func (s *Auth) IsValidSession(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	prefix := "Bearer "
	token := ""

	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}

	if token == "" {
		apputils.Throw(c, errhandle.TokenError)
		return
	}

	info, ok := s.Token.Verify(context.Background(), token)
	if !ok {
		apputils.Throw(c, errhandle.PermissionDenied)
		return
	}

	var usr model.User
	err := s.DB.Table("users").
		Where("id = ?", info.UserID).
		First(&usr).Error

	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var getResponse model.GetUserResponse

	apputils.IgnoreStructCopy(&getResponse, &usr, "")

	if usr.Class != 0 {
		classname, ok := s.Class.GetClassNameByID(usr.Class)
		if ok {
			getResponse.Class = classname
		}
	}

	apputils.OK[model.GetUserResponse](c, getResponse)
}

// 登录 godoc
// @Summary 登录
// @Schemes
// @Description 登录
// @Tags auth
// @Accept json
// @Produce json
// @Param   password  query    string  true   "密码"
// @Param   userid    query    int     false  "用户ID"
// @Param   name      query    string  false  "用户名"
// @Param   email     query    string  false  "用户邮箱"
// @Success 200 {object} model.CommonResponse[model.TokenResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /login [post]
func (s *Auth) UserLogin(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	var loginReq model.UserLoginRequest
	json.Unmarshal(b, &loginReq)

	tx := apputils.BuildQuerySQL(s.DB.Table("users"), &loginReq.QueryRequest)
	if tx == nil {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}

	var user model.User
	err = tx.First(&user).Error
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(loginReq.Password),
	)

	if err != nil {
		apputils.Throw(c, errhandle.PasswordInvalid)
		return
	}

	accessToken, err := s.Token.Token(user.ID, user.Role, user.Name, AccessTokenExpired)
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	refreshToken, err := s.Token.Token(user.ID, user.Role, user.Name, AccessTokenExpired)
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	apputils.OK[model.TokenResponse](c, model.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Scope:        user.Role,
		ExpiredAt:    time.Now().Add(AccessTokenExpired).Unix(),
	})
}

// 刷新登录令牌 godoc
// @Summary 刷新登录令牌
// @Schemes
// @Description 刷新登录令牌
// @Tags auth
// @Accept json
// @Produce json
// @Param   refreshToken    header    int     false  "用户Refresh Token"
// @Success 200 {object} model.CommonResponse[model.TokenResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /refresh [post]
func (s *Auth) UserLoginRefresh(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	prefix := "Bearer "
	token := ""

	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}

	if token == "" {
		apputils.Throw(c, errhandle.TokenError)
		return
	}

	userinfo, ok := s.Token.Verify(context.Background(), token)
	if !ok {
		apputils.Throw(c, errhandle.PermissionDenied)
		return
	}

	accessToken, err := s.Token.Token(userinfo.UserID, userinfo.Role, userinfo.Name, AccessTokenExpired)
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	apputils.OK[model.TokenResponse](c, model.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: token,
		Scope:        userinfo.Role,
		ExpiredAt:    time.Now().Add(AccessTokenExpired).Unix(),
	})
}
