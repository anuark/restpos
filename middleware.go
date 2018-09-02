package main

import (
	"log"
	"net/http"
	userkey "restpos/pkg/userKey"
	"strconv"
	"strings"
	"time"

	Context "github.com/gorilla/context"
)

// Middleware .
type Middleware func(f http.HandlerFunc) http.HandlerFunc

// Log .
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var id string
		if user, ok := Context.GetOk(r, userkey.Key); ok {
			uid := user.(User).ID
			id = strconv.Itoa(int(uid))
		} else {
			id = "-"
		}
		defer func() { log.Printf("%s [%s] %s %s", time.Since(start), id, r.Method, r.URL.Path) }()
		next.ServeHTTP(w, r)
	})
}

// Skip urls that don't need authentication
var skipUrls = []string{
	"/",
	"/auth",
}

// Auth .
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isImagePath := len(r.URL.Path) > 7 && r.URL.Path[:7] == "/static"
		for _, v := range skipUrls {
			if r.URL.Path == v || isImagePath {
				next.ServeHTTP(w, r)
				return
			}
		}

		authKey := r.Header["Authorization"]
		if authKey == nil {
			http.Error(w, "No Bearer Token.", http.StatusForbidden)
			return
		}

		var key []string
		if key = strings.Split(authKey[0], " "); len(key) < 2 {
			http.Error(w, "Wrong Bearer token format.", http.StatusBadRequest)
			return
		}

		var user User
		Db.First(&user, "auth_key = ?", key[1])

		if user.ID == 0 {
			http.Error(w, "Wrong authentication token.", http.StatusForbidden)
			return
		}

		Context.Set(r, userkey.Key, user)
		next.ServeHTTP(w, r)
	})
}

// Cors .
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		w.Header().Add("Access-Control-Expose-Headers", "X-Total-Count")
		if r.Method == "OPTIONS" {
			w.Header().Add("Access-Control-Allow-Methods", "POST, DELETE, PUT, PATCH")
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// JSONContentType All responses will be json Content-Type
func JSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isImagePath := len(r.URL.Path) > 7 && r.URL.Path[:7] == "/static"
		defer next.ServeHTTP(w, r)

		if isImagePath {
			return
		}

		w.Header().Add("Content-Type", "application/json")
	})
}
