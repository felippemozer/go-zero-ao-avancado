package credential

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(token string, ctx context.Context) (string, error) {
	token = strings.Replace(token, "Bearer ", "", 1)
	provider, err := oidc.NewProvider(ctx, os.Getenv("KEYCLOAK"))
	if err != nil {
		return "", errors.New("error to connect to the provider")
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})
	_, err = verifier.Verify(ctx, token)
	if err != nil {
		return "", errors.New("invalid token")
	}

	parsedToken, _ := jwt.Parse(token, nil)
	claims := parsedToken.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	return email, nil
}
