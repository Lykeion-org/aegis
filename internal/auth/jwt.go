package auth

import (
	"time"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	AccessTokenType  TokenType = "access"
	RefreshTokenType TokenType = "refresh"
)

type TokenClaims struct {
	UserUid string `json:"userUid"`
	Role int32 `json:"role"`
	TokenType TokenType `json:"tokenType"`
	jwt.RegisteredClaims
}

func GenerateToken(secret []byte, userUid string, role int32, tokenType TokenType, duration time.Duration) (string, error){
	claims := TokenClaims{
        UserUid: userUid,
        Role:   role,
        TokenType: tokenType,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Subject:   userUid,
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secret)
}

// ParseToken only accepts tokens of the expected type, so a refresh token
// can't be used where an access token is required (and the other way round).
func ParseToken(secret []byte, tokenStr string, expectedType TokenType) (*TokenClaims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
        return secret, nil
    }, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

    if err != nil || !token.Valid {
        return nil, errors.New("invalid token")
    }

    claims, ok := token.Claims.(*TokenClaims)
    if !ok {
        return nil, errors.New("could not parse claims")
    }

    if claims.TokenType != expectedType {
        return nil, errors.New("wrong token type")
    }

    return claims, nil
}
