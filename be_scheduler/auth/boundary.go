package auth

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Role     string        `json:"role"`
	Password string        `json:"password"`
	Email    string        `json:"email"`
	IsActive bool          `json:"is_active" bson:"is_active,omitempty"`
}

func GetToken(email, role string) (string, error) {
	return getToken(email, role)
}
func ValidateToken(tokenString string) (string, string, bool) {
	return validateToken(tokenString)
}
func Login(phoneNumber, password string) (User, error) {
	return login(phoneNumber, password)
}
func GetPasswordHash(password string) (string, error) {
	return getPasswordHash(password)
}
