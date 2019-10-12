package auth

import (
	"errors"
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	jwt_password  = "xaWf7a3Cv5OpB79DJOQ7bigxomjXm6Pe"
	minBcryptCost = 10
)

func getToken(email, role string) (string, error) {

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
		"time":  now.String(),
	})
	str := []byte(jwt_password)
	tokenString, err := token.SignedString(str)
	if err != nil {
		log.Println("ERROR:", err)
	}
	return tokenString, err

}
func validateToken(tokenString string) (string, string, bool) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		hmacSampleSecret := []byte(jwt_password)
		return hmacSampleSecret, nil
	})
	if err != nil {
		log.Println("ERROR:", err)
		return "", "", false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userVerificationFromToken(claims["email"].(string), claims["time"].(string)) {
			return claims["email"].(string), claims["role"].(string), true
		} else {
			return claims["email"].(string), claims["role"].(string), false
		}
	} else {
		log.Println("ERROR:", err)
		return claims["email"].(string), claims["role"].(string), false
	}

}
func login(email string, passwordString string) (User, error) {

	user, err := getUserEntity(email)
	if err != nil {
		log.Println("no user found for email:", email, "password:", passwordString)
		return User{}, err
	}
	//log.Println("%+v", user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordString)); err != nil {
		log.Println("Wrong password entered for", email)
		log.Println("ERROR:", err)
		return User{}, err
	} else {
		if user.IsActive {
			return user, nil
		} else {
			err = errors.New("user not allowed to access! accound disabled")
			log.Println("ERROR:", err)
			return User{}, err
		}
	}
}
func getPasswordHash(password string) (string, error) {

	if password == "" {
		return "", errors.New("blank password")
	} else if pwd, err := bcrypt.GenerateFromPassword([]byte(password), minBcryptCost); err != nil {
		return "", err
	} else {
		hash := string(pwd)
		return string(hash), nil
	}
}
func userVerificationFromToken(email, timeString string) bool {
	available, err := checkIfUserExistsEntity(email)
	if err != nil {
		log.Println("ERROR:", err)
	}
	return available
}
