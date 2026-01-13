package auth

import "time"

type Token struct {
	IDToken      string    `yaml:"id_token"`
	RefreshToken string    `yaml:"refresh_token"`
	Expiry       time.Time `yaml:"expiry"`
}

func (t *Token) Valid() bool {
	return time.Now().Before(t.Expiry.Add(-1 * time.Minute))
}
