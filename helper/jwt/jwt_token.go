package jwt

import "github.com/golang-jwt/jwt/v5"

type JwtToken interface {
	GenerateToken(id string) (*string, error)
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
