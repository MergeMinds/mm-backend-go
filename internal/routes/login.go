package routes

import (
	"net/http"

	"github.com/MergeMinds/mm-backend-go/internal/apierr"
	"github.com/MergeMinds/mm-backend-go/internal/auth/cookie"
	"github.com/MergeMinds/mm-backend-go/internal/auth/password"
	"github.com/MergeMinds/mm-backend-go/internal/auth/session"
	"github.com/MergeMinds/mm-backend-go/internal/auth/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type LoginModel struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterModel struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func Login(c *gin.Context, userRepo user.Repo,
	sessionRepo session.Repo,
	logger *zap.Logger,
	cookieConfig *cookie.CookieConfig) {
	var loginJson LoginModel
	if err := c.ShouldBindBodyWithJSON(&loginJson); err != nil {
		c.JSON(http.StatusBadRequest, apierr.InvalidJSON)
		return
	}

	user, err := userRepo.GetByUsername(loginJson.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, apierr.InternalServer)
		logger.Error(err.Error())
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, apierr.NotFound)
		return
	}

	if !password.Valid(loginJson.Password, user.PasswordHash, user.PasswordSalt) {
		c.JSON(http.StatusUnauthorized, apierr.WrongCredentials)
		return
	}

	s, err := sessionRepo.Create(user.Id, cookieConfig.SessionLifetime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apierr.InternalServer)
		logger.Error(err.Error())
		return
	}

	c.SetCookie(
		session.COOKIE_NAME,
		s.Id.String(),
		cookieConfig.SessionLifetime,
		cookieConfig.Path,
		cookieConfig.Domain,
		cookieConfig.Secure,
		cookieConfig.HttpOnly,
	)

	c.JSON(http.StatusCreated, gin.H{
		"status": "OK",
	})
}

func Register(c *gin.Context, userRepo user.Repo,
	sessionRepo session.Repo,
	logger *zap.Logger,
	cookieConfig *cookie.CookieConfig) {
	var registerJson RegisterModel
	if err := c.ShouldBindBodyWithJSON(&registerJson); err != nil {
		c.JSON(http.StatusBadRequest, apierr.InvalidJSON)
		return
	}

	createUser := user.CreateModel{
		FirstName: registerJson.FirstName,
		LastName:  registerJson.LastName,
		Username:  registerJson.Username,
		Email:     registerJson.Email,
		Password:  registerJson.Password,
		Role:      "USER",
	}

	u, err := userRepo.Create(&createUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apierr.InternalServer)
		logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusCreated, mapUserToUserOut(u))
}

func Logout(c *gin.Context, userRepo user.Repo,
	sessionRepo session.Repo,
	logger *zap.Logger,
	cookieConfig *cookie.CookieConfig) {
	cookie, err := c.Cookie(session.COOKIE_NAME)
	if err != nil {
		c.JSON(http.StatusUnauthorized, apierr.CookieNotExists)
		return
	}

	cookieIdUUID, err := uuid.Parse(cookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, apierr.CookieNotExists)
		return
	}

	err = sessionRepo.DeleteById(cookieIdUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apierr.InternalServer)
		logger.Error(err.Error())
		return
	}

	c.SetCookie(session.COOKIE_NAME, "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusCreated, gin.H{
		"status": "OK",
	})
}

func Session(c *gin.Context, userRepo user.Repo,
	sessionRepo session.Repo,
	logger *zap.Logger,
	cookieConfig *cookie.CookieConfig) {
	session, err := session.CheckHTTPReq(c, sessionRepo, logger)
	if err != nil {
		return
	}

	u, err := userRepo.GetById(session.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apierr.InternalServer)
		logger.Error(err.Error())
		return
	}

	if u == nil {
		c.JSON(http.StatusUnauthorized, apierr.UserNotFound)
		return
	}

	c.JSON(http.StatusCreated, mapUserToUserOut(u))
}

func mapUserToUserOut(u *user.Model) *user.OutModel {
	return &user.OutModel{
		Id:        u.Id,
		CreatedAt: u.CreatedAt,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
	}
}
