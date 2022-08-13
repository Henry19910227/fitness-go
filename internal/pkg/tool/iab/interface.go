package iab

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/iab"
	"time"
)

type Tool interface {
	GenerateGoogleOAuth2Token(duration time.Duration) (string, error)
	APIGetGooglePlayToken(oauthToken string) (string, error)
	APIGetProducts(productID string, purchaseToken string, token string) (*model.IABProductAPIResponse, error)
	APIGetSubscription(productID string, purchaseToken string, token string) (*model.IABSubscriptionAPIResponse, error)
}
