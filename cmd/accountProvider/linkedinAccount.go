package accountProvider

type linkedinAccount struct {
	socialMediaAccount
}

// Internal helper functions
func (l *linkedinAccount) connect() string {
	// Authenticates the account using the password stored in the account
	return "authenticated"
}

func (l *linkedinAccount) authenticate() string {
	return "authenticated with LinkedIn"
}

func (l *linkedinAccount) disconnect() string {
	// Logs out and closes connection to the service
	return "logged off"
}

// PUBLIC INTERFACE
func (l *linkedinAccount) GetMyArticles() string {
	// Returns the list of
	return "My articles"
}

func (l *linkedinAccount) GetArticlesFromUser(username string) []article {
	return make([]article, 0)
}
