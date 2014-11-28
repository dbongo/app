package model

import (
	stderrors "errors"
	"time"

	"github.com/dbongo/app/db"
	"gopkg.in/mgo.v2/bson"
)

// User ...
type User struct {
	ID       string `bson:"_id,omitempty"`
	Created  time.Time
	Username string
	Password string
	Email    string
	Posts    int
}

// ListUsers list all users registred in tsuru
func ListUsers() ([]User, error) {
	conn, err := db.Conn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	var users []User
	err = conn.Users().Find(nil).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByEmail ...
func GetUserByEmail(email string) (*User, error) {
	var u User
	conn, err := db.Conn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	err = conn.Users().Find(bson.M{"email": email}).One(&u)
	if err != nil {
		return nil, stderrors.New("user not found")
	}
	return &u, nil
}

// Create ...
func (u *User) Create() error {
	conn, err := db.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.Users().Insert(u)
}

// Delete ...
func (u *User) Delete() error {
	conn, err := db.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.Users().Remove(bson.M{"email": u.Email})
}

// Update ...
func (u *User) Update() error {
	conn, err := db.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.Users().Update(bson.M{"email": u.Email}, u)
}
