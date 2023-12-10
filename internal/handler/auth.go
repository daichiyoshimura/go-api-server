package handler

import (
	"awsomeapp/internal/module/auth"
	"awsomeapp/internal/server"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
)

const (
	msgFailedToSignin string = "invalid id or password"
)

const (
	tokenExpirationLimit int = 30
)

type AuthHandler struct {
	signingKey []byte
}

func NewAuthHandler(siginingKey []byte) *AuthHandler {
	return &AuthHandler{
		signingKey: siginingKey,
	}
}

func (h *AuthHandler) PostSignin(ctx echo.Context) error {
	var req server.SigninRequest
	if err := ctx.Bind(&req); err != nil {
		writeLog(ctx, err)

		return ctx.JSON(http.StatusBadRequest, &server.Error{
			Message: err.Error(),
		})
	}

	in := &auth.SigninInput{
		ID:   req.Id,
		Pass: req.Pass,
	}
	ok, err := auth.NewAuthUsecase().Signin(in)
	if !ok {
		writeLog(ctx, err)
		return ctx.JSON(http.StatusUnauthorized, &server.Error{
			Message: msgFailedToSignin,
		})
	}
	if err != nil {
		writeLog(ctx, err)

		return ctx.JSON(http.StatusInternalServerError, &server.Error{
			Message: err.Error(),
		})
	}

	// TODO: save session

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * time.Duration(tokenExpirationLimit)).Unix(),
		Id:        req.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(h.signingKey)
	if err != nil {
		writeLog(ctx, err)
		return ctx.JSON(http.StatusInternalServerError, &server.Error{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &server.SigninResponse{
		Token: t,
	})
}
