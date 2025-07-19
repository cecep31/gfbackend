package middleware

import (
	"context"
	"errors"
	"gfbackend/internal/consts"
	"gfbackend/internal/model"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims represents the JWT claims
// Add any custom claims you need
// jwt.RegisteredClaims contains standard claims like "exp" (expiration time)
type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var (
	// JWT signing key (in production, use a secure key from environment variables or config)
	jwtSecret = []byte("your-secret-key")

	// ErrMissingToken is returned when the request is missing the JWT token
	ErrMissingToken = errors.New("authorization token is required")

	// ErrInvalidToken is returned when the token is invalid or expired
	ErrInvalidToken = errors.New("invalid or expired token")
)

// GenerateToken generates a new JWT token for the given user
func GenerateToken(user *model.User) (string, error) {
	claims := CustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// Set token to expire in 24 hours
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWT is a middleware function that validates JWT tokens
func JWT() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			r.Response.WriteStatus(http.StatusUnauthorized)
			r.Response.WriteJson(g.Map{
				"code":    http.StatusUnauthorized,
				"message": ErrMissingToken.Error(),
			})
			r.Exit()
			return
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			r.Response.WriteStatus(http.StatusUnauthorized)
			r.Response.WriteJson(g.Map{
				"code":    http.StatusUnauthorized,
				"message": "invalid token format, expected 'Bearer {token}'",
			})
			r.Exit()
			return
		}

		// Parse and validate the token
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil {
			r.Response.WriteStatus(http.StatusUnauthorized)
			r.Response.WriteJson(g.Map{
				"code":    http.StatusUnauthorized,
				"message": ErrInvalidToken.Error(),
				"error":   err.Error(),
			})
			r.Exit()
			return
		}

		// Check if token is valid
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			// Add user information to the context for use in handlers
			ctx := r.Context()
			ctx = context.WithValue(ctx, consts.CtxUserID, claims.UserID)
			ctx = context.WithValue(ctx, consts.CtxUsername, claims.Username)
			r.SetCtx(ctx)
			r.Middleware.Next()
		} else {
			r.Response.WriteStatus(http.StatusUnauthorized)
			r.Response.WriteJson(g.Map{
				"code":    http.StatusUnauthorized,
				"message": ErrInvalidToken.Error(),
			})
			r.Exit()
		}
	}
}
