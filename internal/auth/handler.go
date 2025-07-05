package auth

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type TokenRequest struct {
	AccessToken  string
	RefreshToken string
}

type AuthHandler interface {
	CreateToken(ctx context.Context, userUid string, role int32) (*TokenRequest, error)
	ValidateAccessToken(ctx context.Context, accessToken string) (*TokenClaims, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
	GenerateEmailValidationToken(ctx context.Context, email string) (string error)
	GeneratePasswordResetToken(ctx context.Context, email string) (string, error)
}

type authHandler struct {
	jwtSecret []byte
}


func NewAuthHandler(jwtSecret []byte) AuthHandler {
	return &authHandler{jwtSecret: jwtSecret}
}

func (h *authHandler) CreateToken(ctx context.Context, userUid string, role int32) (*TokenRequest, error) {
	accessToken, err := GenerateToken(h.jwtSecret, userUid, role, 15*time.Minute)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token")
	}

	refreshToken, err := GenerateToken(h.jwtSecret, userUid, role, 7*24*time.Hour)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token")
	}

	return &TokenRequest{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (h *authHandler) ValidateAccessToken(ctx context.Context, token string) (*TokenClaims, error) {
	claims, err := ParseToken(h.jwtSecret, token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid access token")
	}
	return claims, nil
}

func (h *authHandler) RefreshToken(ctx context.Context, token string) (string, error) {
	claims, err := ParseToken(h.jwtSecret, token)
	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, "invalid refresh token")
	}

	newAccessToken, err := GenerateToken(h.jwtSecret, claims.UserUid, claims.Role, 15*time.Minute)
	if err != nil {
		return "", status.Errorf(codes.Internal, "failed to generate access token")
	}

	return newAccessToken, nil
}

func (h *authHandler) GenerateEmailValidationToken(ctx context.Context, email string) (string error) {
	panic("unimplemented")
}

// GeneratePasswordResetToken implements AuthHandler.
func (h *authHandler) GeneratePasswordResetToken(ctx context.Context, email string) (string, error) {
	panic("unimplemented")
}
