package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/model"
	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type authController struct {
	authUsecase *usecase.AuthUsecase
}

func NewAuthController(authUsecase *usecase.AuthUsecase) authController {
	return authController{authUsecase: authUsecase}
}

func (a *authController) Register(ctx *gin.Context) {
	var request model.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "corpo JSON inválido"})
		return
	}

	user, err := a.authUsecase.Register(request.Username, request.Password)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) {
			ctx.JSON(http.StatusBadRequest, model.Response{Message: "username e password válidos são obrigatórios (senha >= 8 caracteres)"})
			return
		}
		if errors.Is(err, usecase.ErrUserAlreadyExists) {
			ctx.JSON(http.StatusConflict, model.Response{Message: "username já cadastrado"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "falha ao cadastrar usuário"})
		return
	}

	ctx.JSON(http.StatusCreated, model.UserResponse{ID: user.ID, Username: user.Username})
}

func (a *authController) Login(ctx *gin.Context) {
	var request model.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "corpo JSON inválido"})
		return
	}

	if strings.TrimSpace(request.Username) == "" || strings.TrimSpace(request.Password) == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "username e password são obrigatórios"})
		return
	}

	token, expiresIn, err := a.authUsecase.Authenticate(request.Username, request.Password)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) {
			ctx.JSON(http.StatusBadRequest, model.Response{Message: "username e password são obrigatórios"})
			return
		}
		if errors.Is(err, usecase.ErrInvalidCredentials) {
			ctx.JSON(http.StatusUnauthorized, model.Response{Message: "credenciais inválidas"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "falha ao autenticar usuário"})
		return
	}

	ctx.JSON(http.StatusOK, model.TokenResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   expiresIn,
	})
}
