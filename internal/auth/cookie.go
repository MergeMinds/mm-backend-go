package auth

type Seconds = int

type CookieConfig struct {
	SessionLifetime Seconds
	Secure          bool
	Path            string
	HttpOnly        bool
	Domain          string
}

func DefaultCookieConfig() *CookieConfig {
	return &CookieConfig{
		SessionLifetime: 604800, // 7 weeks
		Secure:          true,
		Path:            "/",
		HttpOnly:        true,
		Domain:          "",
	}
}
