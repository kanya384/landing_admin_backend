// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"landing_admin_backend/internal/generated/operations/advantages"
	"landing_admin_backend/internal/generated/operations/hod"
	"landing_admin_backend/internal/generated/operations/plans"
	"landing_admin_backend/internal/generated/operations/posters"
)

// NewBackendServiceAPI creates a new BackendService instance
func NewBackendServiceAPI(spec *loads.Document) *BackendServiceAPI {
	return &BackendServiceAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		JSONProducer: runtime.JSONProducer(),

		AdvantagesDeleteAdvantagePhotoIDHandler: advantages.DeleteAdvantagePhotoIDHandlerFunc(func(params advantages.DeleteAdvantagePhotoIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation advantages.DeleteAdvantagePhotoID has not yet been implemented")
		}),
		AdvantagesDeleteAdvantagesIDHandler: advantages.DeleteAdvantagesIDHandlerFunc(func(params advantages.DeleteAdvantagesIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation advantages.DeleteAdvantagesID has not yet been implemented")
		}),
		HodDeleteMonthsIDHandler: hod.DeleteMonthsIDHandlerFunc(func(params hod.DeleteMonthsIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hod.DeleteMonthsID has not yet been implemented")
		}),
		HodDeletePhotosIDHandler: hod.DeletePhotosIDHandlerFunc(func(params hod.DeletePhotosIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hod.DeletePhotosID has not yet been implemented")
		}),
		PostersDeletePostersPosterIDHandler: posters.DeletePostersPosterIDHandlerFunc(func(params posters.DeletePostersPosterIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation posters.DeletePostersPosterID has not yet been implemented")
		}),
		HodDeleteYearsIDHandler: hod.DeleteYearsIDHandlerFunc(func(params hod.DeleteYearsIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hod.DeleteYearsID has not yet been implemented")
		}),
		AdvantagesGetAdvantagePhotoIDHandler: advantages.GetAdvantagePhotoIDHandlerFunc(func(params advantages.GetAdvantagePhotoIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation advantages.GetAdvantagePhotoID has not yet been implemented")
		}),
		AdvantagesGetAdvantagesHandler: advantages.GetAdvantagesHandlerFunc(func(params advantages.GetAdvantagesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation advantages.GetAdvantages has not yet been implemented")
		}),
		AdvantagesGetAdvantagesIDHandler: advantages.GetAdvantagesIDHandlerFunc(func(params advantages.GetAdvantagesIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation advantages.GetAdvantagesID has not yet been implemented")
		}),
		HodGetMonthsIDHandler: hod.GetMonthsIDHandlerFunc(func(params hod.GetMonthsIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hod.GetMonthsID has not yet been implemented")
		}),
		HodGetPhotosIDHandler: hod.GetPhotosIDHandlerFunc(func(params hod.GetPhotosIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hod.GetPhotosID has not yet been implemented")
		}),
		GetPingHandler: GetPingHandlerFunc(func(params GetPingParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetPing has not yet been implemented")
		}),
		PlansGetPlansHandler: plans.GetPlansHandlerFunc(func(params plans.GetPlansParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation plans.GetPlans has not yet been implemented")
		}),
		PostersGetPostersHandler: posters.GetPostersHandlerFunc(func(params posters.GetPostersParams) middleware.Responder {
			return middleware.NotImplemented("operation posters.GetPosters has not yet been implemented")
		}),
		PostersGetPostersPosterIDHandler: posters.GetPostersPosterIDHandlerFunc(func(params posters.GetPostersPosterIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation posters.GetPostersPosterID has not yet been implemented")
		}),
		HodGetYearsHandler: hod.GetYearsHandlerFunc(func(params hod.GetYearsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hod.GetYears has not yet been implemented")
		}),
		AdvantagesPatchAdvantagesHandler: advantages.PatchAdvantagesHandlerFunc(func(params advantages.PatchAdvantagesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation advantages.PatchAdvantages has not yet been implemented")
		}),
		HodPatchMonthsHandler: hod.PatchMonthsHandlerFunc(func(params hod.PatchMonthsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hod.PatchMonths has not yet been implemented")
		}),
		PlansPatchPlansHandler: plans.PatchPlansHandlerFunc(func(params plans.PatchPlansParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation plans.PatchPlans has not yet been implemented")
		}),
		PostersPatchPostersHandler: posters.PatchPostersHandlerFunc(func(params posters.PatchPostersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation posters.PatchPosters has not yet been implemented")
		}),
		HodPatchYearsHandler: hod.PatchYearsHandlerFunc(func(params hod.PatchYearsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hod.PatchYears has not yet been implemented")
		}),
		AdvantagesPostAdvantagePhotoOrdersHandler: advantages.PostAdvantagePhotoOrdersHandlerFunc(func(params advantages.PostAdvantagePhotoOrdersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation advantages.PostAdvantagePhotoOrders has not yet been implemented")
		}),
		AdvantagesPostAdvantagesOrdersHandler: advantages.PostAdvantagesOrdersHandlerFunc(func(params advantages.PostAdvantagesOrdersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation advantages.PostAdvantagesOrders has not yet been implemented")
		}),
		PostLoginHandler: PostLoginHandlerFunc(func(params PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation PostLogin has not yet been implemented")
		}),
		HodPostPhotosOrdersHandler: hod.PostPhotosOrdersHandlerFunc(func(params hod.PostPhotosOrdersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hod.PostPhotosOrders has not yet been implemented")
		}),
		PlansPostPlansHandler: plans.PostPlansHandlerFunc(func(params plans.PostPlansParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation plans.PostPlans has not yet been implemented")
		}),
		PostersPostPostersOrdersHandler: posters.PostPostersOrdersHandlerFunc(func(params posters.PostPostersOrdersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation posters.PostPostersOrders has not yet been implemented")
		}),
		AdvantagesPutAdvantagePhotoHandler: advantages.PutAdvantagePhotoHandlerFunc(func(params advantages.PutAdvantagePhotoParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation advantages.PutAdvantagePhoto has not yet been implemented")
		}),
		AdvantagesPutAdvantagesHandler: advantages.PutAdvantagesHandlerFunc(func(params advantages.PutAdvantagesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation advantages.PutAdvantages has not yet been implemented")
		}),
		HodPutMonthsHandler: hod.PutMonthsHandlerFunc(func(params hod.PutMonthsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hod.PutMonths has not yet been implemented")
		}),
		HodPutPhotosHandler: hod.PutPhotosHandlerFunc(func(params hod.PutPhotosParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hod.PutPhotos has not yet been implemented")
		}),
		PlansPutPlansHandler: plans.PutPlansHandlerFunc(func(params plans.PutPlansParams) middleware.Responder {
			return middleware.NotImplemented("operation plans.PutPlans has not yet been implemented")
		}),
		PostersPutPostersHandler: posters.PutPostersHandlerFunc(func(params posters.PutPostersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation posters.PutPosters has not yet been implemented")
		}),
		PutUsersHandler: PutUsersHandlerFunc(func(params PutUsersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PutUsers has not yet been implemented")
		}),
		HodPutYearsHandler: hod.PutYearsHandlerFunc(func(params hod.PutYearsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation hod.PutYears has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		TokenAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (Token) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*BackendServiceAPI the backend service API */
type BackendServiceAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// TokenAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	TokenAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// AdvantagesDeleteAdvantagePhotoIDHandler sets the operation handler for the delete advantage photo ID operation
	AdvantagesDeleteAdvantagePhotoIDHandler advantages.DeleteAdvantagePhotoIDHandler
	// AdvantagesDeleteAdvantagesIDHandler sets the operation handler for the delete advantages ID operation
	AdvantagesDeleteAdvantagesIDHandler advantages.DeleteAdvantagesIDHandler
	// HodDeleteMonthsIDHandler sets the operation handler for the delete months ID operation
	HodDeleteMonthsIDHandler hod.DeleteMonthsIDHandler
	// HodDeletePhotosIDHandler sets the operation handler for the delete photos ID operation
	HodDeletePhotosIDHandler hod.DeletePhotosIDHandler
	// PostersDeletePostersPosterIDHandler sets the operation handler for the delete posters poster ID operation
	PostersDeletePostersPosterIDHandler posters.DeletePostersPosterIDHandler
	// HodDeleteYearsIDHandler sets the operation handler for the delete years ID operation
	HodDeleteYearsIDHandler hod.DeleteYearsIDHandler
	// AdvantagesGetAdvantagePhotoIDHandler sets the operation handler for the get advantage photo ID operation
	AdvantagesGetAdvantagePhotoIDHandler advantages.GetAdvantagePhotoIDHandler
	// AdvantagesGetAdvantagesHandler sets the operation handler for the get advantages operation
	AdvantagesGetAdvantagesHandler advantages.GetAdvantagesHandler
	// AdvantagesGetAdvantagesIDHandler sets the operation handler for the get advantages ID operation
	AdvantagesGetAdvantagesIDHandler advantages.GetAdvantagesIDHandler
	// HodGetMonthsIDHandler sets the operation handler for the get months ID operation
	HodGetMonthsIDHandler hod.GetMonthsIDHandler
	// HodGetPhotosIDHandler sets the operation handler for the get photos ID operation
	HodGetPhotosIDHandler hod.GetPhotosIDHandler
	// GetPingHandler sets the operation handler for the get ping operation
	GetPingHandler GetPingHandler
	// PlansGetPlansHandler sets the operation handler for the get plans operation
	PlansGetPlansHandler plans.GetPlansHandler
	// PostersGetPostersHandler sets the operation handler for the get posters operation
	PostersGetPostersHandler posters.GetPostersHandler
	// PostersGetPostersPosterIDHandler sets the operation handler for the get posters poster ID operation
	PostersGetPostersPosterIDHandler posters.GetPostersPosterIDHandler
	// HodGetYearsHandler sets the operation handler for the get years operation
	HodGetYearsHandler hod.GetYearsHandler
	// AdvantagesPatchAdvantagesHandler sets the operation handler for the patch advantages operation
	AdvantagesPatchAdvantagesHandler advantages.PatchAdvantagesHandler
	// HodPatchMonthsHandler sets the operation handler for the patch months operation
	HodPatchMonthsHandler hod.PatchMonthsHandler
	// PlansPatchPlansHandler sets the operation handler for the patch plans operation
	PlansPatchPlansHandler plans.PatchPlansHandler
	// PostersPatchPostersHandler sets the operation handler for the patch posters operation
	PostersPatchPostersHandler posters.PatchPostersHandler
	// HodPatchYearsHandler sets the operation handler for the patch years operation
	HodPatchYearsHandler hod.PatchYearsHandler
	// AdvantagesPostAdvantagePhotoOrdersHandler sets the operation handler for the post advantage photo orders operation
	AdvantagesPostAdvantagePhotoOrdersHandler advantages.PostAdvantagePhotoOrdersHandler
	// AdvantagesPostAdvantagesOrdersHandler sets the operation handler for the post advantages orders operation
	AdvantagesPostAdvantagesOrdersHandler advantages.PostAdvantagesOrdersHandler
	// PostLoginHandler sets the operation handler for the post login operation
	PostLoginHandler PostLoginHandler
	// HodPostPhotosOrdersHandler sets the operation handler for the post photos orders operation
	HodPostPhotosOrdersHandler hod.PostPhotosOrdersHandler
	// PlansPostPlansHandler sets the operation handler for the post plans operation
	PlansPostPlansHandler plans.PostPlansHandler
	// PostersPostPostersOrdersHandler sets the operation handler for the post posters orders operation
	PostersPostPostersOrdersHandler posters.PostPostersOrdersHandler
	// AdvantagesPutAdvantagePhotoHandler sets the operation handler for the put advantage photo operation
	AdvantagesPutAdvantagePhotoHandler advantages.PutAdvantagePhotoHandler
	// AdvantagesPutAdvantagesHandler sets the operation handler for the put advantages operation
	AdvantagesPutAdvantagesHandler advantages.PutAdvantagesHandler
	// HodPutMonthsHandler sets the operation handler for the put months operation
	HodPutMonthsHandler hod.PutMonthsHandler
	// HodPutPhotosHandler sets the operation handler for the put photos operation
	HodPutPhotosHandler hod.PutPhotosHandler
	// PlansPutPlansHandler sets the operation handler for the put plans operation
	PlansPutPlansHandler plans.PutPlansHandler
	// PostersPutPostersHandler sets the operation handler for the put posters operation
	PostersPutPostersHandler posters.PutPostersHandler
	// PutUsersHandler sets the operation handler for the put users operation
	PutUsersHandler PutUsersHandler
	// HodPutYearsHandler sets the operation handler for the put years operation
	HodPutYearsHandler hod.PutYearsHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *BackendServiceAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *BackendServiceAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *BackendServiceAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *BackendServiceAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *BackendServiceAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *BackendServiceAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *BackendServiceAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *BackendServiceAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *BackendServiceAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the BackendServiceAPI
func (o *BackendServiceAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.TokenAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.AdvantagesDeleteAdvantagePhotoIDHandler == nil {
		unregistered = append(unregistered, "advantages.DeleteAdvantagePhotoIDHandler")
	}
	if o.AdvantagesDeleteAdvantagesIDHandler == nil {
		unregistered = append(unregistered, "advantages.DeleteAdvantagesIDHandler")
	}
	if o.HodDeleteMonthsIDHandler == nil {
		unregistered = append(unregistered, "hod.DeleteMonthsIDHandler")
	}
	if o.HodDeletePhotosIDHandler == nil {
		unregistered = append(unregistered, "hod.DeletePhotosIDHandler")
	}
	if o.PostersDeletePostersPosterIDHandler == nil {
		unregistered = append(unregistered, "posters.DeletePostersPosterIDHandler")
	}
	if o.HodDeleteYearsIDHandler == nil {
		unregistered = append(unregistered, "hod.DeleteYearsIDHandler")
	}
	if o.AdvantagesGetAdvantagePhotoIDHandler == nil {
		unregistered = append(unregistered, "advantages.GetAdvantagePhotoIDHandler")
	}
	if o.AdvantagesGetAdvantagesHandler == nil {
		unregistered = append(unregistered, "advantages.GetAdvantagesHandler")
	}
	if o.AdvantagesGetAdvantagesIDHandler == nil {
		unregistered = append(unregistered, "advantages.GetAdvantagesIDHandler")
	}
	if o.HodGetMonthsIDHandler == nil {
		unregistered = append(unregistered, "hod.GetMonthsIDHandler")
	}
	if o.HodGetPhotosIDHandler == nil {
		unregistered = append(unregistered, "hod.GetPhotosIDHandler")
	}
	if o.GetPingHandler == nil {
		unregistered = append(unregistered, "GetPingHandler")
	}
	if o.PlansGetPlansHandler == nil {
		unregistered = append(unregistered, "plans.GetPlansHandler")
	}
	if o.PostersGetPostersHandler == nil {
		unregistered = append(unregistered, "posters.GetPostersHandler")
	}
	if o.PostersGetPostersPosterIDHandler == nil {
		unregistered = append(unregistered, "posters.GetPostersPosterIDHandler")
	}
	if o.HodGetYearsHandler == nil {
		unregistered = append(unregistered, "hod.GetYearsHandler")
	}
	if o.AdvantagesPatchAdvantagesHandler == nil {
		unregistered = append(unregistered, "advantages.PatchAdvantagesHandler")
	}
	if o.HodPatchMonthsHandler == nil {
		unregistered = append(unregistered, "hod.PatchMonthsHandler")
	}
	if o.PlansPatchPlansHandler == nil {
		unregistered = append(unregistered, "plans.PatchPlansHandler")
	}
	if o.PostersPatchPostersHandler == nil {
		unregistered = append(unregistered, "posters.PatchPostersHandler")
	}
	if o.HodPatchYearsHandler == nil {
		unregistered = append(unregistered, "hod.PatchYearsHandler")
	}
	if o.AdvantagesPostAdvantagePhotoOrdersHandler == nil {
		unregistered = append(unregistered, "advantages.PostAdvantagePhotoOrdersHandler")
	}
	if o.AdvantagesPostAdvantagesOrdersHandler == nil {
		unregistered = append(unregistered, "advantages.PostAdvantagesOrdersHandler")
	}
	if o.PostLoginHandler == nil {
		unregistered = append(unregistered, "PostLoginHandler")
	}
	if o.HodPostPhotosOrdersHandler == nil {
		unregistered = append(unregistered, "hod.PostPhotosOrdersHandler")
	}
	if o.PlansPostPlansHandler == nil {
		unregistered = append(unregistered, "plans.PostPlansHandler")
	}
	if o.PostersPostPostersOrdersHandler == nil {
		unregistered = append(unregistered, "posters.PostPostersOrdersHandler")
	}
	if o.AdvantagesPutAdvantagePhotoHandler == nil {
		unregistered = append(unregistered, "advantages.PutAdvantagePhotoHandler")
	}
	if o.AdvantagesPutAdvantagesHandler == nil {
		unregistered = append(unregistered, "advantages.PutAdvantagesHandler")
	}
	if o.HodPutMonthsHandler == nil {
		unregistered = append(unregistered, "hod.PutMonthsHandler")
	}
	if o.HodPutPhotosHandler == nil {
		unregistered = append(unregistered, "hod.PutPhotosHandler")
	}
	if o.PlansPutPlansHandler == nil {
		unregistered = append(unregistered, "plans.PutPlansHandler")
	}
	if o.PostersPutPostersHandler == nil {
		unregistered = append(unregistered, "posters.PutPostersHandler")
	}
	if o.PutUsersHandler == nil {
		unregistered = append(unregistered, "PutUsersHandler")
	}
	if o.HodPutYearsHandler == nil {
		unregistered = append(unregistered, "hod.PutYearsHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *BackendServiceAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *BackendServiceAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "Token":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.TokenAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *BackendServiceAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *BackendServiceAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *BackendServiceAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *BackendServiceAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the backend service API
func (o *BackendServiceAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *BackendServiceAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/advantage_photo/{id}"] = advantages.NewDeleteAdvantagePhotoID(o.context, o.AdvantagesDeleteAdvantagePhotoIDHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/advantages/{id}"] = advantages.NewDeleteAdvantagesID(o.context, o.AdvantagesDeleteAdvantagesIDHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/months/{id}"] = hod.NewDeleteMonthsID(o.context, o.HodDeleteMonthsIDHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/photos/{id}"] = hod.NewDeletePhotosID(o.context, o.HodDeletePhotosIDHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/posters/{posterID}"] = posters.NewDeletePostersPosterID(o.context, o.PostersDeletePostersPosterIDHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/years/{id}"] = hod.NewDeleteYearsID(o.context, o.HodDeleteYearsIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/advantage_photo/{id}"] = advantages.NewGetAdvantagePhotoID(o.context, o.AdvantagesGetAdvantagePhotoIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/advantages"] = advantages.NewGetAdvantages(o.context, o.AdvantagesGetAdvantagesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/advantages/{id}"] = advantages.NewGetAdvantagesID(o.context, o.AdvantagesGetAdvantagesIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/months/{id}"] = hod.NewGetMonthsID(o.context, o.HodGetMonthsIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/photos/{id}"] = hod.NewGetPhotosID(o.context, o.HodGetPhotosIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ping"] = NewGetPing(o.context, o.GetPingHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/plans"] = plans.NewGetPlans(o.context, o.PlansGetPlansHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/posters"] = posters.NewGetPosters(o.context, o.PostersGetPostersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/posters/{posterID}"] = posters.NewGetPostersPosterID(o.context, o.PostersGetPostersPosterIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/years"] = hod.NewGetYears(o.context, o.HodGetYearsHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/advantages"] = advantages.NewPatchAdvantages(o.context, o.AdvantagesPatchAdvantagesHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/months"] = hod.NewPatchMonths(o.context, o.HodPatchMonthsHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/plans"] = plans.NewPatchPlans(o.context, o.PlansPatchPlansHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/posters"] = posters.NewPatchPosters(o.context, o.PostersPatchPostersHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/years"] = hod.NewPatchYears(o.context, o.HodPatchYearsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/advantage_photo/orders"] = advantages.NewPostAdvantagePhotoOrders(o.context, o.AdvantagesPostAdvantagePhotoOrdersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/advantages/orders"] = advantages.NewPostAdvantagesOrders(o.context, o.AdvantagesPostAdvantagesOrdersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = NewPostLogin(o.context, o.PostLoginHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/photos/orders"] = hod.NewPostPhotosOrders(o.context, o.HodPostPhotosOrdersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/plans"] = plans.NewPostPlans(o.context, o.PlansPostPlansHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/posters/orders"] = posters.NewPostPostersOrders(o.context, o.PostersPostPostersOrdersHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/advantage_photo"] = advantages.NewPutAdvantagePhoto(o.context, o.AdvantagesPutAdvantagePhotoHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/advantages"] = advantages.NewPutAdvantages(o.context, o.AdvantagesPutAdvantagesHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/months"] = hod.NewPutMonths(o.context, o.HodPutMonthsHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/photos"] = hod.NewPutPhotos(o.context, o.HodPutPhotosHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/plans"] = plans.NewPutPlans(o.context, o.PlansPutPlansHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/posters"] = posters.NewPutPosters(o.context, o.PostersPutPostersHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users"] = NewPutUsers(o.context, o.PutUsersHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/years"] = hod.NewPutYears(o.context, o.HodPutYearsHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *BackendServiceAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *BackendServiceAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *BackendServiceAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *BackendServiceAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *BackendServiceAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
