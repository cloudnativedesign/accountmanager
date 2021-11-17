package accountProvider

import (
	"fmt"
	strngs "strings"
)

type accountProvider interface {
	connect() string
	disconnect() string
	authenticate() string

	// Retrieve
	retrieveProfile() (*userProfile, error)
	retrievePosts() (*[]post, error)
	retrieveArticles() (*[]article, error)
	retrieveFollowers() (*[]socialMediaAccount, error)

	// Update
	updateProfile(profileInput map[string]interface{}) (*userProfile, error)
	// Create

	// Delete
}

type accountManager struct {
	accounts []socialMediaAccount
}

func NewAccountManager() accountManager {
	return accountManager{}
}

func (a *accountManager) AddAccount(providerType string, accountInput UserID) *accountManager {
	// Assume the following is available correctly
	var provider accountProvider
	switch strngs.ToUpper(providerType) {
	case "LINKEDIN":
		provider = new(linkedinProvider)

	}

	account := socialMediaAccount{
		provider: provider,
		UserID: UserID{
			Username: accountInput.Username,
			Password: accountInput.Password,
		},
	}
	a.accounts = append(a.accounts, account)
	return a
}

func (a *accountManager) AuthenticateAll() {
	for _, account := range a.accounts {
		if !account.isAuthenticated {
			fmt.Println(account.provider.authenticate())
		}
	}

}
