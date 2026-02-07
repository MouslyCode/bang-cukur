package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type JwtClaims struct {
	UserId uuid.UUID
	RoleId uint
}

func GenerateJWT(userId uuid.UUID, roleId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId.String(),
		"role_id": roleId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func VerifyJwt(tokenString string) (*JwtClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	userID, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return nil, err
	}

	roleRaw, ok := claims["role_id"]
	if !ok {
		return nil, errors.New("role_id not found")
	}

	return &JwtClaims{
		UserId: userID,
		RoleId: uint(roleRaw.(float64)),
	}, nil
}
