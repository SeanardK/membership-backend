package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
)

const ClaimsKey = "claims"

type OIDCAuth struct {
	verifier *oidc.IDTokenVerifier
}

func New(ctx context.Context, issuer, clientID string) (*OIDCAuth, error) {
	if issuer == "" || clientID == "" {
		return nil, errors.New("issuer and clientID are required")
	}

	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, err
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID:          clientID,
		SkipClientIDCheck: true,
	})

	return &OIDCAuth{verifier: verifier}, nil
}

func (a *OIDCAuth) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		tokenString, err := extractBearerToken(header)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization header"})
			return
		}

		ctx := c.Request.Context()
		idToken, err := a.verifier.Verify(ctx, tokenString)
		if err != nil {
			log.Printf("token verification failed: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err != nil {
			log.Printf("failed to parse token claims: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}

		c.Set(ClaimsKey, claims)
		c.Next()
	}
}

func GetClaims(c *gin.Context) (map[string]interface{}, bool) {
	raw, exists := c.Get(ClaimsKey)
	if !exists {
		return nil, false
	}
	claims, ok := raw.(map[string]interface{})
	return claims, ok
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("no auth header")
	}
	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", errors.New("invalid auth header")
	}
	return parts[1], nil
}
