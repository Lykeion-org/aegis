package auth

import (
	"time"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)



type TokenClaims struct {
	UserUid string `json:"userUid"`
	Role int32 `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(secret []byte, userUid string, role int32, duration time.Duration) (string, error){
	claims := TokenClaims{
        UserUid: userUid,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Subject:   userUid,
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secret)
}

func ParseToken(secret []byte, tokenStr string) (*TokenClaims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
        return secret, nil
    })

    if err != nil || !token.Valid {
        return nil, errors.New("invalid token")
    }

    claims, ok := token.Claims.(*TokenClaims)
    if !ok {
        return nil, errors.New("could not parse claims")
    }

    return claims, nil
}