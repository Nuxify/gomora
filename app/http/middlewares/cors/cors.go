package cors

import (
	"github.com/go-chi/cors"
)

func Init() *cors.Cors {
	//list of allowed origins
	allowedOriginsArr := []string{
		"https://v1.arbitrage.ph",
		"https://arbitrage.ph",
		"https://vyndue.com",
		"https://dev-v1.arbitrage.ph",
		"http://dev.vyndue.com",
		"https://dev-v1.arbitrage.local",
	}

	//list of allowed methods
	allowedMethodsArr := []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"OPTIONS",
	}

	//list of allowed headers
	allowedHeadersArr := []string{
		"Accept",
		"Authorization",
		"Content-Type",
		"X-CSRF-Token",
	}

	//list of exposed headers
	exposedHeadersArr := []string{"Link"}

	//set allow credentials
	allowCredentials := true

	//set max age
	maxAge := 300

	//instantiate cors rule
	cors := cors.New(cors.Options{
		AllowedOrigins:   allowedOriginsArr,
		AllowedMethods:   allowedMethodsArr,
		AllowedHeaders:   allowedHeadersArr,
		ExposedHeaders:   exposedHeadersArr,
		AllowCredentials: allowCredentials,
		MaxAge:           maxAge, // Maximum value not ignored by any of major browsers
	})

	return cors
}
