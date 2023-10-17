package jwt_token

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken interface {
	GenerateToken(id string) (*string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtToken struct {
}

var SECRET_KEY = []byte("ini sangat rahasia")

func NewJwtToken() *jwtToken {
	return &jwtToken{}
}

func (s *jwtToken) GenerateToken(userId string) (*string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	// claims["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}

func (s *jwtToken) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
