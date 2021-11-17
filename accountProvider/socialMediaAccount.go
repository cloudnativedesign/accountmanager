package accountProvider

import "fmt"

type UserID struct {
	Username string
	Password string
	Token    *string
}

type userProfile struct {
	firstName   string
	lastName    string
	username    string
	hometown    string
	homecountry string
}

// Applicable for the owned account that is updatable as well as retrieved accounts for followers, commenters, etc..
// isOwned is used to select which accounts can be edited and which are read only
type socialMediaAccount struct {
	UserID
	profile         userProfile
	articles        []article
	posts           []post
	provider        accountProvider
	isOwned         bool
	isAuthenticated bool
}

// GET
func (s *socialMediaAccount) GetPosts() ([]post, error) {
	if s.posts == nil {
		posts, err := s.provider.retrievePosts()
		if err != nil {
			return nil, err
		}
		s.posts = *posts
	}
	return s.posts, nil
}

func (s *socialMediaAccount) GetProfile() (userProfile, error) {
	// Check if profile is set
	if s.profile == (userProfile{}) {
		profile, err := s.provider.retrieveProfile()
		if err != nil {
			return userProfile{}, err
		}
		s.profile = *profile
	}
	return s.profile, nil
}

// UPDATE
func (s *socialMediaAccount) UpdateProfile(profileInput map[string]interface{}) (userProfile, error) {
	if s.isOwned {
		// Check if provided data suits to update the profile
		// Iterate over all
		profile, err := s.provider.updateProfile(profileInput)
		if err != nil {
			return userProfile{}, err
		}
		s.profile = *profile
	}
	return userProfile{}, fmt.Errorf("Account is not authorized with the Account Manager. It likely is a foreign user account which is immutable.")
}
