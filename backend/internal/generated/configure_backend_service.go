// This file is safe to edit. Once it exists it will not be overwritten

package generated

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/generated/operations"
	"landing_admin_backend/internal/generated/operations/posters"
	"landing_admin_backend/internal/handlers"
	mng "landing_admin_backend/internal/repository/mongo"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/pkg/logger"
	"landing_admin_backend/pkg/token_manager"
)

//go:generate swagger generate server --target ../../../landing_admin_backend --name BackendService --spec ../../api/openapi/openapi.yml --server-package ./internal/generated --principal interface{}

func configureFlags(api *operations.BackendServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BackendServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	cfg, err := config.InitConfig("APP")
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}

	file, err := os.OpenFile(cfg.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening log file config %s", err))
	}

	logger, err := logger.NewLogger(cfg.ServiceName, cfg.LogLevel, file)
	if err != nil {
		panic(fmt.Sprintf("error initializing logger %s", err))
	}
	api.Logger = logger.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	clientOptions := options.Client().ApplyURI(cfg.DSN)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Panic("error connection to mongo", err, nil)
	}

	repository := mng.Setup(client.Database("public"))
	services := services.Setup(cfg, repository)
	handlers := handlers.NewHandlers(services, cfg)
	api.PostLoginHandler = operations.PostLoginHandlerFunc(handlers.Auth.Authenticate)
	api.GetPingHandler = operations.GetPingHandlerFunc(handlers.Auth.Ping)
	api.TokenAuth = func(token string) (claims interface{}, err error) {
		claims, err = token_manager.CheckToken(token, cfg.TokenSecret)
		return
	}

	api.PutUsersHandler = operations.PutUsersHandlerFunc(handlers.Users.Create)
	api.PostersPutPostersHandler = posters.PutPostersHandlerFunc(handlers.Poster.Create)
	api.PostersGetPostersHandler = posters.GetPostersHandlerFunc(handlers.Poster.Get)
	api.PostersPostPostersHandler = posters.PostPostersHandlerFunc(handlers.Poster.Update)
	http.StripPrefix("/spec", http.FileServer(http.Dir("api/openapi")))

	api.PreServerShutdown = func() {
	}

	api.ServerShutdown = func() {
		fmt.Print("gracefull shutdown\n")
		file.Close()
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return uiMiddleware(handler)
}

func uiMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Shortcut helpers for swagger-ui
		if r.URL.Path == "/swagger-ui" || r.URL.Path == "/api/help" {
			http.Redirect(w, r, "/swagger-ui/", http.StatusFound)
			return
		}
		// Serving ./swagger-ui/
		if strings.Index(r.URL.Path, "/swagger-ui/") == 0 {
			http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("swaggerui"))).ServeHTTP(w, r)
			return
		}
		if strings.Index(r.URL.Path, "/spec/") == 0 {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			http.StripPrefix("/spec/", http.FileServer(http.Dir("api/openapi"))).ServeHTTP(w, r)
		}
		handler.ServeHTTP(w, r)
	})
}
