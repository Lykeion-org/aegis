package api

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	auth "github.com/lykeion-org/aegis/internal/auth"
)

type RestHandler struct {
	handler auth.AuthHandler
}

type token struct{
	Token string `json:"token"`
}


func(h *RestHandler)GenerateTokens(c *gin.Context){
	var req auth.TokenClaims

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	tokens, err := h.handler.CreateToken(context.Background(),req.UserUid, req.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate tokens"})
	}
	c.JSON(http.StatusOK, tokens)
}

func(h *RestHandler)RefreshTokens(c *gin.Context){
	var req token

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := h.handler.RefreshToken(context.Background(),req.Token)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate tokens"})
	}

	c.JSON(http.StatusOK, token)
}

func(h *RestHandler)Validate(c *gin.Context){
	var req token

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	claims, err := h.handler.ValidateAccessToken(context.Background(),req.Token)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate tokens"})
	}

	c.JSON(http.StatusOK, claims)
}



