package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User .
type User struct {
	Model
	Email, AuthKey, PasswordHash string
	Username                     string `json:"username"`
	OrganizationID               uint
	Password                     string `gorm:"-" json:"password"`
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

	if user.Username == "" || user.Password == "" {
		http.Error(w, "Email or password field can't be empty.", http.StatusBadRequest)
		return
	}

	Db.First(&user, "username = ?", user.Username)
	if user.ID == 0 || !user.CheckPasswordHash() {
		http.Error(w, "Wrong email and password combination", http.StatusBadRequest)
		return
	}

	keyJSON, _ := json.Marshal(struct {
		Key string `json:"key"`
	}{user.AuthKey})
	fmt.Fprint(w, string(keyJSON))
}
