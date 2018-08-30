package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User .
type User struct {
	gorm.Model
	Username, Email, AuthKey, PasswordHash string
	OrganizationID                         uint
	Password                               string `gorm:"-"`
}

// HashPassword .
func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.PasswordHash = string(bytes)
	return err
}

// CheckPasswordHash .
func (u *User) CheckPasswordHash() bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	return err == nil
}

// GenerateAuthKey .
func (u *User) GenerateAuthKey() {
	h := md5.New()
	io.WriteString(h, time.Now().String())
	u.AuthKey = fmt.Sprintf("%x", h.Sum(nil))
}

// UserAuth .
func UserAuth(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	decoder.Decode(&user)

	if user.Email == "" || user.Password == "" {
		http.Error(w, "Email or password field can't be empty.", http.StatusBadRequest)
		return
	}

	Db.First(&user, "email = ?", user.Email)
	if user.ID == 0 || user.CheckPasswordHash() {
		http.Error(w, "Wrong email and password combination", http.StatusBadRequest)
		return
	}

	keyJSON, _ := json.Marshal(struct{ Key string }{user.AuthKey})
	fmt.Println(keyJSON)
	fmt.Fprint(w, keyJSON)
}
