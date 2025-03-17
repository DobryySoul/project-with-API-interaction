package handler

import (
	"DobryySoul/project-with-API-interaction/internal/config"
	"DobryySoul/project-with-API-interaction/internal/entity"
	"DobryySoul/project-with-API-interaction/internal/services/service"
	"DobryySoul/project-with-API-interaction/pkg/jwt"
	"DobryySoul/project-with-API-interaction/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handlers struct {
	Service *service.Service
	Cfg     *config.Config
	Logger  *logger.Logger
}

func NewHandlers(service *service.Service, cfg *config.Config, logger *logger.Logger) *Handlers {
	return &Handlers{
		Service: service,
		Cfg:     cfg,
		Logger:  logger,
	}
}

func (h *Handlers) RegisterUser(c *gin.Context) {
	var userReg entity.RegisterUser

	if err := c.ShouldBindJSON(&userReg); err != nil {
		h.Logger.L.Error("invalid request body", zap.Error(err))
		c.JSON(400, gin.H{"error": "invalid request body"})

		return
	}

	h.Logger.L.Info("Registering user", zap.String("username", userReg.Username))

	id, err := h.Service.RegisterUser(c, &userReg)
	if err != nil {
		h.Logger.L.Error("failed to register user", zap.Error(err))
		c.JSON(500, gin.H{"error": "failed to register user"})

		return
	}

	h.Logger.L.Info("User registered successfully", zap.Int("user_id", id))

	c.JSON(200, gin.H{"message": "user registered successfully"})
}

func (h *Handlers) LoginUser(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		h.Logger.L.Error("invalid request body", zap.Error(err))
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	h.Logger.L.Info("Login attempt", zap.String("email", credentials.Email))

	user, err := h.Service.AuthUser(c, credentials.Email, credentials.Password)
	if err != nil {
		h.Logger.L.Error("failed to authenticate user", zap.Error(err))
		c.JSON(401, gin.H{"error": "invalid email or password"})
		return
	}

	if user == nil {
		h.Logger.L.Error("user not found", zap.String("email", credentials.Email))
		c.JSON(401, gin.H{"error": "invalid email or password"})
		return
	}

	h.Logger.L.Info("User authenticated successfully", zap.String("email", credentials.Email))

	token, err := jwt.GenerateJWT(user, h.Cfg.JWTSecret)
	if err != nil {
		h.Logger.L.Error("failed to generate token", zap.Error(err))
		c.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"message": "user logged in", "token": token})
}
