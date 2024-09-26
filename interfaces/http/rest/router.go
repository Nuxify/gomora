/*
|--------------------------------------------------------------------------
| Router
|--------------------------------------------------------------------------
|
| This file contains the routes mapping and groupings of your REST API calls.
| See README.md for the routes UI server.
|
*/
package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"gomora/interfaces"
	"gomora/interfaces/http/rest/middlewares/cors"
	"gomora/interfaces/http/rest/viewmodels"
)

// ChiRouterInterface declares methods for the chi router
type ChiRouterInterface interface {
	InitRouter() *chi.Mux
	Serve(port int)
}

type router struct{}

var (
	m          *router
	routerOnce sync.Once
)

// InitRouter initializes main routes
func (router *router) InitRouter() *chi.Mux {
	// DI assignment
	recordCommandController := interfaces.ServiceContainer().RegisterRecordRESTCommandController()
	recordQueryController := interfaces.ServiceContainer().RegisterRecordRESTQueryController()

	// create router
	r := chi.NewRouter()

	// global and recommended middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(cors.Init().Handler)
	r.Use(middleware.Recoverer)

	// default route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := viewmodels.HTTPResponseVM{
			Status:  http.StatusOK,
			Success: true,
			Message: "alive",
			Data:    map[string]string{"version": os.Getenv("API_VERSION")},
		}

		response.JSON(w)
	})

	// docs routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth(os.Getenv("API_NAME"), map[string]string{
			"sudo": os.Getenv("OPENAPI_DOCS_PASSWORD"),
		}))

		workDir, _ := os.Getwd()
		docsDir := http.Dir(filepath.Join(workDir, "docs"))
		FileServer(r, "/docs", docsDir)
	})

	// API routes
	r.Group(func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			// record module
			r.Route("/record", func(r chi.Router) {
				r.Post("/", recordCommandController.CreateRecord)
				r.Get("/{id}", recordQueryController.GetRecordByID)
			})
		})
	})

	return r
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func (router *router) Serve(port int) {
	log.Printf("[SERVER] REST server running on :%d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router.InitRouter())
	if err != nil {
		log.Fatalf("[SERVER] REST server failed %v", err)
	}
}

func registerHandlers() {}

// ChiRouter export instantiated chi router once
func ChiRouter() ChiRouterInterface {
	if m == nil {
		routerOnce.Do(func() {
			// register http handlers
			registerHandlers()

			m = &router{}
		})
	}

	return m
}
