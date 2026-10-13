package auth

import (
	"context"
	"testing"
	"time"
)

func TestTokenTypesAreNotInterchangeable(t *testing.T) {
	ctx := context.Background()
	h := NewAuthHandler([]byte("test-secret"))

	tokens, err := h.CreateToken(ctx, "user-1", 3)
	if err != nil {
		t.Fatalf("CreateToken: %v", err)
	}

	claims, err := h.ValidateAccessToken(ctx, tokens.AccessToken)
	if err != nil {
		t.Fatalf("access token rejected as access token: %v", err)
	}
	if claims.UserUid != "user-1" || claims.Role != 3 {
		t.Fatalf("unexpected claims: %+v", claims)
	}

	if _, err := h.ValidateAccessToken(ctx, tokens.RefreshToken); err == nil {
		t.Fatal("refresh token accepted as access token")
	}

	if _, err := h.RefreshToken(ctx, tokens.AccessToken); err == nil {
		t.Fatal("access token accepted as refresh token")
	}

	newAccess, err := h.RefreshToken(ctx, tokens.RefreshToken)
	if err != nil {
		t.Fatalf("refresh token rejected: %v", err)
	}
	if _, err := h.ValidateAccessToken(ctx, newAccess); err != nil {
		t.Fatalf("refreshed access token rejected: %v", err)
	}
}

func TestTokenWithoutTypeIsRejected(t *testing.T) {
	secret := []byte("test-secret")
	legacy, err := GenerateToken(secret, "user-1", 3, "", time.Hour)
	if err != nil {
		t.Fatalf("GenerateToken: %v", err)
	}
	if _, err := ParseToken(secret, legacy, AccessTokenType); err == nil {
		t.Fatal("untyped token accepted as access token")
	}
}
