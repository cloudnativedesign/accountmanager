package main

import "github.com/cloudnativedesign/accountmanager/accountProvider"

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

	//

	// Collect list of personal tasks
	// myArticles := linkedinProvider.GetMyArticles()
	// myPosts := linkedinProvider.GetOwnPosts(
	// 	fromDate: "2021-12-09:1930"
	// 	toDate: "2021-12-13:1940"
	// )
	// myProfile := linkedinProvider.GetOwnProfile()

	// // Store all data to the accountManager
	// accountManager.addAccount({
	// 	account
	// })
}
