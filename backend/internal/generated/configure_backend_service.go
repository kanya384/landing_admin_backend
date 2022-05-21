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
	"landing_admin_backend/internal/generated/operations/advantages"
	"landing_admin_backend/internal/generated/operations/docs"
	"landing_admin_backend/internal/generated/operations/hod"
	"landing_admin_backend/internal/generated/operations/leads"
	"landing_admin_backend/internal/generated/operations/plans"
	"landing_admin_backend/internal/generated/operations/posters"
	"landing_admin_backend/internal/generated/operations/video"
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

	/* Posters handlers */
	api.PutUsersHandler = operations.PutUsersHandlerFunc(handlers.Users.Create)
	api.PostersPutPostersHandler = posters.PutPostersHandlerFunc(handlers.Poster.Create)
	api.PostersGetPostersHandler = posters.GetPostersHandlerFunc(handlers.Poster.Get)
	api.PostersPatchPostersHandler = posters.PatchPostersHandlerFunc(handlers.Poster.Update)
	api.PostersGetPostersPosterIDHandler = posters.GetPostersPosterIDHandlerFunc(handlers.Poster.GetPosterById)
	api.PostersDeletePostersPosterIDHandler = posters.DeletePostersPosterIDHandlerFunc(handlers.Poster.Delete)
	api.PostersPostPostersOrdersHandler = posters.PostPostersOrdersHandlerFunc(handlers.Poster.PostersOrdersChange)

	/* hod handlers*/
	/* years*/
	api.HodGetYearsHandler = hod.GetYearsHandlerFunc(handlers.Years.Get)
	api.HodPutYearsHandler = hod.PutYearsHandlerFunc(handlers.Years.Create)
	api.HodPatchYearsHandler = hod.PatchYearsHandlerFunc(handlers.Years.Update)
	api.HodDeleteYearsIDHandler = hod.DeleteYearsIDHandlerFunc(handlers.Years.Delete)
	/*months*/
	api.HodGetMonthsIDHandler = hod.GetMonthsIDHandlerFunc(handlers.Months.Get)
	api.HodPutMonthsHandler = hod.PutMonthsHandlerFunc(handlers.Months.Create)
	api.HodPatchMonthsHandler = hod.PatchMonthsHandlerFunc(handlers.Months.Update)
	api.HodDeleteMonthsIDHandler = hod.DeleteMonthsIDHandlerFunc(handlers.Months.Delete)
	/*photos*/
	api.HodGetPhotosIDHandler = hod.GetPhotosIDHandlerFunc(handlers.HodPhotos.Get)
	api.HodPutPhotosHandler = hod.PutPhotosHandlerFunc(handlers.HodPhotos.Create)
	api.HodPostPhotosOrdersHandler = hod.PostPhotosOrdersHandlerFunc(handlers.HodPhotos.ChangePhotosOrders)
	api.HodDeletePhotosIDHandler = hod.DeletePhotosIDHandlerFunc(handlers.HodPhotos.Delete)

	/* advantages*/
	api.AdvantagesGetAdvantagesHandler = advantages.GetAdvantagesHandlerFunc(handlers.Advantages.Get)
	api.AdvantagesGetAdvantagesIDHandler = advantages.GetAdvantagesIDHandlerFunc(handlers.Advantages.GetByID)
	api.AdvantagesPutAdvantagesHandler = advantages.PutAdvantagesHandlerFunc(handlers.Advantages.Create)
	api.AdvantagesPatchAdvantagesHandler = advantages.PatchAdvantagesHandlerFunc(handlers.Advantages.Update)
	api.AdvantagesDeleteAdvantagesIDHandler = advantages.DeleteAdvantagesIDHandlerFunc(handlers.Advantages.Delete)
	api.AdvantagesPostAdvantagesOrdersHandler = advantages.PostAdvantagesOrdersHandlerFunc(handlers.Advantages.ChangeAdvantageOrders)

	/* advantages photos*/
	api.AdvantagesGetAdvantagePhotoIDHandler = advantages.GetAdvantagePhotoIDHandlerFunc(handlers.AdvantagesPhoto.Get)
	api.AdvantagesPutAdvantagePhotoHandler = advantages.PutAdvantagePhotoHandlerFunc(handlers.AdvantagesPhoto.Create)
	api.AdvantagesDeleteAdvantagePhotoIDHandler = advantages.DeleteAdvantagePhotoIDHandlerFunc(handlers.AdvantagesPhoto.Delete)
	api.AdvantagesPostAdvantagePhotoOrdersHandler = advantages.PostAdvantagePhotoOrdersHandlerFunc(handlers.AdvantagesPhoto.ChangeAdvantagePhotosOrders)

	/* plans */
	api.PlansPutPlansHandler = plans.PutPlansHandlerFunc(handlers.Plans.ProcessPlans)
	api.PlansGetPlansHandler = plans.GetPlansHandlerFunc(handlers.Plans.GetPlans)
	api.PlansPatchPlansHandler = plans.PatchPlansHandlerFunc(handlers.Plans.UpdatePlan)
	api.PlansPostPlansHandler = plans.PostPlansHandlerFunc(handlers.Plans.UpdatePlansActivity)

	/* docs */
	api.DocsGetDocHandler = docs.GetDocHandlerFunc(handlers.Docs.Get)
	api.DocsPutDocHandler = docs.PutDocHandlerFunc(handlers.Docs.Create)
	api.DocsPatchDocHandler = docs.PatchDocHandlerFunc(handlers.Docs.Update)
	api.DocsDeleteDocIDHandler = docs.DeleteDocIDHandlerFunc(handlers.Docs.Delete)

	/* videos */
	api.VideoGetVideoHandler = video.GetVideoHandlerFunc(handlers.Video.Get)
	api.VideoPutVideoHandler = video.PutVideoHandlerFunc(handlers.Video.Create)
	api.VideoDeleteVideoIDHandler = video.DeleteVideoIDHandlerFunc(handlers.Video.Delete)
	api.VideoPatchVideoHandler = video.PatchVideoHandlerFunc(handlers.Video.Update)

	/* leads */
	api.LeadsGetLeadHandler = leads.GetLeadHandlerFunc(handlers.Leads.Get)
	api.LeadsPutLeadHandler = leads.PutLeadHandlerFunc(handlers.Leads.Create)
	api.LeadsDeleteLeadIDHandler = leads.DeleteLeadIDHandlerFunc(handlers.Leads.Delete)
	api.LeadsGetLeadAnalyticsHandler = leads.GetLeadAnalyticsHandlerFunc(handlers.Leads.GetAnalytics)

	http.FileServer(http.Dir("file_store"))

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
	return FileServerMiddleware(handler)
}

func FileServerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Index(r.URL.Path, "/store/") == 0 {
			http.StripPrefix("/store/", http.FileServer(http.Dir("file_store"))).ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
