package controller

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
)

type ApiController struct {
	beego.Controller
}

func init() {
	gob.Register(auth.Claims{})
}

func GetUserName(user *auth.User) string {
	if user == nil {
		return ""
	}

	return user.Name
}

func (c *ApiController) GetSessionClaims() *auth.Claims {
	s := c.GetSession("user")
	if s == nil {
		return nil
	}

	claims := s.(auth.Claims)
	return &claims
}

func (c *ApiController) SetSessionClaims(claims *auth.Claims) {
	if claims == nil {
		c.DelSession("user")
		return
	}

	c.SetSession("user", *claims)
}

func (c *ApiController) GetSessionUser() *auth.User {
	claims := c.GetSessionClaims()
	if claims == nil {
		return nil
	}

	return &claims.User
}

func (c *ApiController) SetSessionUser(user *auth.User) {
	if user == nil {
		c.DelSession("user")
		return
	}

	claims := c.GetSessionClaims()
	if claims != nil {
		claims.User = *user
		c.SetSessionClaims(claims)
	}
}

func (c *ApiController) GetSessionUsername() string {
	user := c.GetSessionUser()
	if user == nil {
		return ""
	}

	return GetUserName(user)
}

func (c *ApiController) Signin() {
	code := c.Input().Get("code")
	state := c.Input().Get("state")

	token, err := auth.GetOAuthToken(code, state)
	if err != nil {
		//c.ResponseError(err.Error())
		return
	}

	claims, err := auth.ParseJwtToken(token.AccessToken)
	if err != nil {
		//c.ResponseError(err.Error())
		return
	}
	/*
		affected, err := object.UpdateMemberOnlineStatus(&claims.User, true, util.GetCurrentTime())
		if err != nil {
			//c.ResponseError(err.Error())
			return
		}
	*/
	claims.AccessToken = token.AccessToken
	c.SetSessionClaims(claims)

	//c.ResponseOk(claims, affected)
}
