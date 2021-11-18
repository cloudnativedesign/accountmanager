package accountProvider

type linkedinProvider struct {
}

func (l *linkedinProvider) connect() string {
	return "connected to LinkedIn"
}

func (l *linkedinProvider) disconnect() string {
	return "disconnected from Linkedin"
}

func (l *linkedinProvider) authenticate() string {
	return "Authenticated with Linkedin. Received and stored access token"
}

// GET Methods
func (l *linkedinProvider) retrieveProfile() (*userProfile, error) {
	return &userProfile{
		firstName:   "Frank",
		lastName:    "Winkler",
		username:    "frank.winkler@lkas.com",
		hometown:    "Dubai",
		homecountry: "United Arab Emirates",
	}, nil
}

func (l *linkedinProvider) retrievePosts() (*[]post, error) {
	posts := make([]post, 0)
	return &posts, nil
}

func (l *linkedinProvider) retrieveArticles() (*[]article, error) {
	articles := make([]article, 0)
	return &articles, nil
}

func (l *linkedinProvider) retrieveFollowers() (*[]socialMediaAccount, error) {
	followers := make([]socialMediaAccount, 0)
	return &followers, nil
}

func (l *linkedinProvider) updateProfile(profileInput map[string]interface{}) (*userProfile, error) {
	profile := userProfile{}
	return &profile, nil
}
