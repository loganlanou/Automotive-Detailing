package auth

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/labstack/echo/v4"
)

// Context keys for storing user info
type contextKey string

const (
	UserIDKey    contextKey = "clerk_user_id"
	SessionIDKey contextKey = "clerk_session_id"
	UserKey      contextKey = "clerk_user"
)

// ClerkConfig holds Clerk configuration
type ClerkConfig struct {
	SecretKey      string
	PublishableKey string
}

// LoadConfig loads Clerk configuration from environment
func LoadConfig() *ClerkConfig {
	return &ClerkConfig{
		SecretKey:      os.Getenv("CLERK_SECRET_KEY"),
		PublishableKey: os.Getenv("CLERK_PUBLISHABLE_KEY"),
	}
}

// Init initializes the Clerk SDK with the secret key
func Init(secretKey string) {
	clerk.SetKey(secretKey)
}

// RequireAuth middleware ensures the user is authenticated
// Redirects to sign-in page if not authenticated
func RequireAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sessionToken := extractSessionToken(c.Request())
			if sessionToken == "" {
				// Redirect to sign-in for browser requests
				if isHTMLRequest(c.Request()) {
					return c.Redirect(http.StatusTemporaryRedirect, "/sign-in?redirect_url="+c.Request().URL.Path)
				}
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Authentication required",
				})
			}

			// Verify the session token
			claims, err := jwt.Verify(c.Request().Context(), &jwt.VerifyParams{
				Token: sessionToken,
			})
			if err != nil {
				if isHTMLRequest(c.Request()) {
					return c.Redirect(http.StatusTemporaryRedirect, "/sign-in?redirect_url="+c.Request().URL.Path)
				}
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid session",
				})
			}

			// Store user info in context
			ctx := context.WithValue(c.Request().Context(), UserIDKey, claims.Subject)
			ctx = context.WithValue(ctx, SessionIDKey, claims.SessionID)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

// OptionalAuth middleware extracts user info if available but doesn't require it
func OptionalAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sessionToken := extractSessionToken(c.Request())
			if sessionToken == "" {
				return next(c)
			}

			// Try to verify the session token
			claims, err := jwt.Verify(c.Request().Context(), &jwt.VerifyParams{
				Token: sessionToken,
			})
			if err != nil {
				// Invalid token, continue without user
				return next(c)
			}

			// Store user info in context
			ctx := context.WithValue(c.Request().Context(), UserIDKey, claims.Subject)
			ctx = context.WithValue(ctx, SessionIDKey, claims.SessionID)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

// GetUserID extracts the user ID from the request context
func GetUserID(ctx context.Context) string {
	if userID, ok := ctx.Value(UserIDKey).(string); ok {
		return userID
	}
	return ""
}

// GetSessionID extracts the session ID from the request context
func GetSessionID(ctx context.Context) string {
	if sessionID, ok := ctx.Value(SessionIDKey).(string); ok {
		return sessionID
	}
	return ""
}

// IsAuthenticated checks if the current request has a valid user
func IsAuthenticated(ctx context.Context) bool {
	return GetUserID(ctx) != ""
}

// GetUser fetches the full user object from Clerk
func GetUser(ctx context.Context) (*clerk.User, error) {
	userID := GetUserID(ctx)
	if userID == "" {
		return nil, nil
	}
	return user.Get(ctx, userID)
}

// extractSessionToken extracts the session token from cookies or Authorization header
func extractSessionToken(r *http.Request) string {
	// Try to get from __session cookie (Clerk's default)
	if cookie, err := r.Cookie("__session"); err == nil && cookie.Value != "" {
		return cookie.Value
	}

	// Try to get from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}

	return ""
}

// isHTMLRequest checks if the request expects HTML response
func isHTMLRequest(r *http.Request) bool {
	accept := r.Header.Get("Accept")
	return strings.Contains(accept, "text/html") || accept == "" || accept == "*/*"
}

// UserInfo represents basic user information for templates
type UserInfo struct {
	ID             string
	Email          string
	FirstName      string
	LastName       string
	FullName       string
	ImageURL       string
	IsAdmin        bool
}

// GetUserInfo fetches user info and returns a template-friendly struct
func GetUserInfo(ctx context.Context) *UserInfo {
	userID := GetUserID(ctx)
	if userID == "" {
		return nil
	}

	u, err := user.Get(ctx, userID)
	if err != nil {
		return nil
	}

	firstName := ""
	lastName := ""
	if u.FirstName != nil {
		firstName = *u.FirstName
	}
	if u.LastName != nil {
		lastName = *u.LastName
	}

	email := ""
	if len(u.EmailAddresses) > 0 {
		email = u.EmailAddresses[0].EmailAddress
	}

	fullName := strings.TrimSpace(firstName + " " + lastName)
	if fullName == "" {
		fullName = email
	}

	return &UserInfo{
		ID:        u.ID,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		FullName:  fullName,
		ImageURL:  *u.ImageURL,
		IsAdmin:   false, // Can be enhanced with role checks
	}
}
