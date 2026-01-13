package config

type Credentials struct {
	Firebase FirebaseCredentials `yaml:"firebase"`
}

type FirebaseCredentials struct {
	APIKey   string `yaml:"api_key"`
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}
