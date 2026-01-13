package auth

func ResolveToken(apiKey string, login func() (*Token, error)) (*Token, error) {
	// 1️⃣ Load cached token
	if t, err := LoadToken(); err == nil {
		if t.Valid() {
			return t, nil
		}

		// 2️⃣ Try refresh
		if t.RefreshToken != "" {
			if nt, err := RefreshToken(apiKey, t.RefreshToken); err == nil {
				_ = SaveToken(nt)
				return nt, nil
			}
		}
	}

	// 3️⃣ Full login
	t, err := login()
	if err != nil {
		return nil, err
	}

	_ = SaveToken(t)
	return t, nil
}
