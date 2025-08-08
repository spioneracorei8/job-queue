package handler

import (
	"auth-service/constants"
	"auth-service/models"
	"auth-service/services/register"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type registerHandler struct {
	registerUs register.RegisterUsecase
}

func NewRegisterHandlerImpl(registerUs register.RegisterUsecase) register.RegisterHandler {
	return &registerHandler{
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
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   constants.INVALID_REQUEST_BODY,
		})
		return
	}
	if err := validate.Struct(register); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   constants.INVALID_REQUEST_BODY,
		})
		return
	}
	if err := h.registerUs.RegisterUser(g.Request.Context(), register, source); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	g.Status(http.StatusCreated)
}
