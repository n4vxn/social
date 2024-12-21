package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"

	"github.com/n4vxn/social/internal/auth"
	"github.com/n4vxn/social/internal/store"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/swaggo/swag/example/basic/docs"
)

type application struct {
	config        config
	store         store.Storage
	logger        *zap.SugaredLogger
	authenticator auth.Authenticator
}

type config struct {
	addr   string
	db     dbConfig
	env    string
	apiURL string
	auth   authConfig
}

type authConfig struct {
	basic basicConfig
	token tokenConfig
}

type tokenConfig struct {
	secret string
	exp    time.Duration
	iss    string
}

type basicConfig struct {
	user string
	pass string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	//	@title			Social API
	//	@description	This is a social network.
	//	@termsOfService	http://swagger.io/terms/

	//	@contact.name	API Support
	//	@contact.url	http://www.swagger.io/support
	//	@contact.email	support@swagger.io

	//	@license.name	Apache 2.0
	//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

	//	@BasePath					/v1
	//
	//	@securityDefinitions.apikey	ApiKeyAuth
	//	@in							header
	//	@name						Authorization
	//	@description

	r.Route("/v1", func(r chi.Router) {
		// Health check route
		r.With(app.BasicAuthMiddleware()).Get("/health", app.healthCheckerHandler)

		docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.addr)
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))
		// Posts-related routes
		r.Route("/posts", func(r chi.Router) {
			r.Use(app.AuthTokenMiddleware)
			r.Post("/", app.createPostHandler)

			r.Route("/{postID}", func(r chi.Router) {
				r.Use(app.postsContextMiddleware)

				r.Get("/", app.getPostHandler)
				r.Delete("/", app.deletePostHandler)
				r.Patch("/", app.updatePostHandler)
			})
		})

		// Users-related routes
		r.Route("/users", func(r chi.Router) {
			// User-specific routes with userID
			r.Route("/{userID}", func(r chi.Router) {
				r.Use(app.AuthTokenMiddleware)

				r.Get("/", app.getUserHandler)
				r.Put("/follow", app.followUserHandler)
				r.Put("/unfollow", app.unfollowedUserHandler)
			})

			// Feed route
			r.Group(func(r chi.Router) {
				r.Get("/feed", app.getUserFeedHandler)
			})
		})

		r.Route("/authentication", func(r chi.Router) {
			r.Post("/user", app.registerUserHandler)
			r.Post("/token", app.createTokenHandler)
		})
	})

	return r
}

func (app *application) run(mux *chi.Mux) error {
	//Docs
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.apiURL
	docs.SwaggerInfo.BasePath = "/v1"
	srv := http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	app.logger.Infow("server has started", "addr", app.config.addr, "env", app.config.env)
	return srv.ListenAndServe()
}
