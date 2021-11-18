package accountProvider

type tag struct {
	value    string
	category string
}

type reaction struct {
	value     string
	sentiment int
}

type share struct {
	articleID string
	userID    string
}

type article struct {
	id      string
	title   string
	content string
	tags    []tag
}

type post struct {
	id        string
	content   string
	tags      []tag
	reactions []reaction
	reshares  []share
}
