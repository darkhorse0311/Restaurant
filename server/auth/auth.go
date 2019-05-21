package auth

import (
	"context"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/reynld/carbtographer/server/models"
)

// Protected middleware
func Protected(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// We can obtain the session token from the requests cookies, which come with every request
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tknStr := c.Value

		// Initialize a new instance of `Claims`
		claims := &models.Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return getJWTKey(), nil
		})
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, string("username"), claims.Username)
		ctx = context.WithValue(ctx, string("id"), claims.ID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GenerateToken creates JWT
func GenerateToken(username string, id int) (models.JWTResponse, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &models.Claims{
		Username: username,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(getJWTKey())
	if err != nil {
		return models.JWTResponse{}, err
	}
	return models.JWTResponse{Token: tokenString, Time: expirationTime}, nil
}

// Refresh JWT handler
func Refresh(w http.ResponseWriter, r *http.Request) {
	// (BEGIN) The code uptil this point is the same as the first part of the `Welcome` route
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &models.Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return getJWTKey(), nil
	})
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// (END) The code up-till this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJWTKey())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

//getJWTKey returns JWTiKey
func getJWTKey() []byte {
	return []byte(os.Getenv("JWT_KEY"))
}
