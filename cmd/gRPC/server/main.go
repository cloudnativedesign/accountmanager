package main

import (
	"github.com/cloudnativedesign/accountmanager/cmd/accountProvider"
)

func main() {
	// Create a common account manager object
	accounts := accountProvider.NewAccountManager()

	// Link the social media accounts available
	accounts.AddAccount("linkedin", accountProvider.UserID{
		Username: "frank@thewinklerway.com",
		Password: "alksdjf@Q%#@%",
	})

	// Authenticate all accounts
	accounts.AuthenticateAll()

	// Check out
}
