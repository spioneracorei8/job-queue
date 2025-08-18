package handler

import (
	"auth-service/constants"
	"auth-service/logger"
	"auth-service/models"
	"auth-service/services/register"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type registerHandler struct {
	logger     *logger.Logger
	registerUs register.RegisterUsecase
}

func NewRegisterHandlerImpl(logger *logger.Logger, registerUs register.RegisterUsecase) register.RegisterHandler {
	return &registerHandler{
		logger:     logger,
		registerUs: registerUs,
	}
}

func (h *registerHandler) RegisterUser(g *gin.Context) {
	var (
		source   = g.GetHeader("source")
		register = new(models.Register)
		validate = validator.New()
	)
	if err := g.ShouldBindJSON(&register); err != nil {
		h.logger.Info(g, err.Error(), map[string]any{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
			"source": source,
		})
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   constants.INVALID_REQUEST_BODY,
		})
		return
	}
	if err := validate.Struct(register); err != nil {
		h.logger.Info(g, err.Error(), map[string]any{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
			"source": source,
		})
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   constants.INVALID_REQUEST_BODY,
		})
		return
	}
	if err := h.registerUs.RegisterUser(g.Request.Context(), register, source); err != nil {
		h.logger.Error(g, err.Error(), map[string]any{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
			"source": source,
		})
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.logger.Info(g, constants.SUCCESSFUL_MAG, map[string]any{
		"status": http.StatusCreated,
	})
	g.Status(http.StatusCreated)
}
