package cors

// Config holds the CORS configurations
type Config struct{}

// AllowCredentials return a boolean expression whether credentials are allowed
func (c *Config) AllowCredentials() bool {
	return true
}

// AllowedHeaders returns list of allowed headers
func (c *Config) AllowedHeaders() []string {
	return []string{
		"Accept",
		"Authorization",
		"Content-Type",
		"X-CSRF-Token",
	}
}

// AllowedOrigins returns list of allowed origins
func (c *Config) AllowedOrigins() []string {
	return []string{"*"}
}

// AllowedMethods returns list of allowed methods
func (c *Config) AllowedMethods() []string {
	return []string{
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"OPTIONS",
	}
}

// ExposedHeaders returns list of exposed headers
func (c *Config) ExposedHeaders() []string {
	return []string{"Link"}
}

// MaxAge returns the maximum number of age in browser in seconds
func (c *Config) MaxAge() int {
	return 300
}
