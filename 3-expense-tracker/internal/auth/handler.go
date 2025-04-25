package auth

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/amreshpro/go-backend/3-expense-tracker/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)


type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtKey = []byte(config.GetEnv("JWT_SECRET_KEY",""))


func SignupHandler(w http.ResponseWriter, r *http.Request){
w.Write([] byte("signup handler"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Yeh hardcoded hai abhi, DB se verify karna production mein
    if user.Username != "amresh" || user.Password != "1234" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Token expiry
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: user.Username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Response
    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
