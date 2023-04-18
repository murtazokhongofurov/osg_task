package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/osg_task/api/handler/v1/tokens"
	"github.com/osg_task/internal/controller/storage"
	"github.com/osg_task/internal/pkg/config"
	"github.com/osg_task/internal/pkg/logger"
	"github.com/spf13/cast"
)

type handlerV1 struct {
	log        logger.Logger
	storage    storage.StorageI
	cfg        config.Config
	jwthandler tokens.JWTHandler
}

type HandlerV1Config struct {
	Logger     logger.Logger
	Storage    storage.StorageI
	Cfg        config.Config
	JwtHandler tokens.JWTHandler
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:        c.Logger,
		storage:    c.Storage,
		cfg:        c.Cfg,
		jwthandler: c.JwtHandler,
	}
}

func GetClaims(h handlerV1, c *gin.Context) (*tokens.CustomClaims, error) {
	var (
		claims = tokens.CustomClaims{}
	)

	strToken := c.GetHeader("Authorization")
	token, err := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) { return []byte(h.cfg.SigninKey), nil })
	if err != nil {
		h.log.Error("invalid access token")
	}
	rawClaims := token.Claims.(jwt.MapClaims)

	claims.Sub = rawClaims["sub"].(string)
	claims.Exp = rawClaims["exp"].(float64)
	aud := cast.ToStringSlice(rawClaims["aud"])
	claims.Aud = aud
	claims.Role = rawClaims["role"].(string)
	claims.Token = token

	return &claims, nil
}
