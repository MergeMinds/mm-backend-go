package session

import (
	"errors"
	"net/http"

	"github.com/InTeam-Russia/go-backend-template/internal/apierr"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const COOKIE_NAME = "SESSION_ID"

func CheckHTTPReq(c *gin.Context, sessionRepo Repo, logger *zap.Logger) (*Model, error) {
	cookie, err := c.Cookie(COOKIE_NAME)
	if err != nil {
		c.JSON(http.StatusUnauthorized, apierr.CookieNotExists)
		return nil, err
	}

	cookieIdUUID, err := uuid.Parse(cookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, apierr.CookieNotExists)
		return nil, err
	}

	session, err := sessionRepo.GetById(cookieIdUUID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, apierr.InternalServer)
		logger.Error(err.Error())
		return nil, err
	}

	if session == nil {
		c.JSON(http.StatusUnauthorized, apierr.SessionNotFound)
		return nil, errors.New("Session not found")
	}

	if session.IsExpired() {
		c.JSON(http.StatusUnauthorized, apierr.SessionExpired)
		return nil, errors.New("Session is expired")
	}

	return session, nil
}
