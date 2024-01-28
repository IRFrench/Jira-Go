package cfg

type Config struct {
	Auth           Authentication `yaml:"authentication"`
	URL            string         `yaml:"jira_domain"`
	SearchRequests []string       `yaml:"search_requests"`
}

type Authentication struct {
	Email string `yaml:"email"`
	Token string `yaml:"token"`
}
