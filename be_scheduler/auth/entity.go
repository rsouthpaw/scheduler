package auth

import (
	"log"

	"../base"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func checkIfUserExistsEntity(email string) (bool, error) {

	session, err := mgo.Dial(base.MONGO_BASE_URL)
	if err != nil {
		log.Println("Error:", err)
		return false, err
	}
	defer session.Close()

	c := session.DB(base.DB_ACTYV).C(base.COL_USERS)
	var user User
	query := bson.M{
		"email": email,
	}
	err = c.Find(query).One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return false, nil
		} else {
			log.Println("Error:", err)
			return false, err
		}
	} else {
		return true, nil
	}
}

func getUserEntity(email string) (User, error) {

	session, err := mgo.Dial(base.MONGO_BASE_URL)
	if err != nil {
		log.Println("Error:", err)
		return User{}, err
	}
	defer session.Close()

	c := session.DB(base.DB_ACTYV).C(base.COL_USERS)
	var user User
	query := bson.M{
		"email": email,
	}
	err = c.Find(query).One(&user)
	if err != nil {
		log.Println("ERROR:", err)
	}
	return user, err

}
