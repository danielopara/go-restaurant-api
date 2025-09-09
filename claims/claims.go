package claims

import (
	"os"
	"time"

	"github.com/danielopara/restaurant-api/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserId uint   `json:"user_id"`
	Email  string `json:"email"`
	Role models.Role `json:"role"`

	jwt.RegisteredClaims
}


var JwtSecret []byte

func InitJwt(){
	secret := os.Getenv("JWT_SECRET")
	JwtSecret = []byte(secret)
}

func HashPassword(password string)(string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(password, hash string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password) )
	return err == nil
}

func GenerateToken(user models.User)(string, error){
	claims := Claims{
		UserId: user.ID,
		Email: user.Email,
		Role: user.Role,
		
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 *time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "restaurant-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func ParseToken(tokenString string)(*Claims, error){
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error){
		return JwtSecret, nil
	})

	if err != nil{
		return nil, err
	}

	if claims, ok := token.Claims.(* Claims); ok && token.Valid{
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}