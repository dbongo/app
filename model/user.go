package model

import (
	"errors"

	"code.google.com/p/go.crypto/bcrypt"

	"github.com/dbongo/app/db"
	"github.com/mholt/binding"
	"gopkg.in/mgo.v2/bson"
)

// User ...
type User struct {
	ID            bson.ObjectId `bson:"_id" json:"-"`
	Username      string        `bson:"username" json:"username"`
	Password      string        `bson:"password,omitempty" json:"-"`
	Email         string        `bson:"email,omitempty" json:"email,omitempty"`
	VerifiedEmail bool          `bson:"mailV,omitempty" json:"verifiedEmail,omitempty"`
}

// FieldMap ...
func (u *User) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&u.Email: binding.Field{
			Form:     "email",
			Required: true,
		},
		&u.Username: binding.Field{
			Form:     "username",
			Required: true,
		},
		&u.Password: binding.Field{
			Form:     "password",
			Required: true,
		},
	}
}

// AuthUser ...
func AuthUser(username, password string) (*User, error) {
	conn, err := db.Conn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	user := &User{}
	err = conn.Users().Find(bson.M{"username": username}).One(user)
	if err != nil {
		return nil, err
	}
	if user.ID == "" {
		return nil, errors.New("No user found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("Incorrect password")
	}
	return user, nil
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

// FindUser ...
func FindUser(id bson.ObjectId) (*User, error) {
	var u User
	conn, err := db.Conn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	err = conn.Users().Find(bson.M{"_id": id}).One(&u)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &u, nil
}

// FindUserByEmail ...
func FindUserByEmail(email string) (*User, error) {
	var u User
	conn, err := db.Conn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	err = conn.Users().Find(bson.M{"email": email}).One(&u)
	if err != nil {
		return nil, errors.New("user not found")
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
	pwHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(pwHash)
	u.ID = bson.NewObjectId()
	u.VerifiedEmail = false
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
