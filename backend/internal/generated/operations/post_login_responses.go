// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// PostLoginOKCode is the HTTP code returned for type PostLoginOK
const PostLoginOKCode int = 200

/*PostLoginOK регистрация успешна

swagger:response postLoginOK
*/
type PostLoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.AuthenticateResponse `json:"body,omitempty"`
}

// NewPostLoginOK creates PostLoginOK with default headers values
func NewPostLoginOK() *PostLoginOK {

	return &PostLoginOK{}
}

// WithPayload adds the payload to the post login o k response
func (o *PostLoginOK) WithPayload(payload *models.AuthenticateResponse) *PostLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post login o k response
func (o *PostLoginOK) SetPayload(payload *models.AuthenticateResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostLoginBadRequestCode is the HTTP code returned for type PostLoginBadRequest
const PostLoginBadRequestCode int = 400

/*PostLoginBadRequest Bad request

swagger:response postLoginBadRequest
*/
type PostLoginBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.AuthenticateResponse `json:"body,omitempty"`
}

// NewPostLoginBadRequest creates PostLoginBadRequest with default headers values
func NewPostLoginBadRequest() *PostLoginBadRequest {

	return &PostLoginBadRequest{}
}

// WithPayload adds the payload to the post login bad request response
func (o *PostLoginBadRequest) WithPayload(payload *models.AuthenticateResponse) *PostLoginBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post login bad request response
func (o *PostLoginBadRequest) SetPayload(payload *models.AuthenticateResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoginBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostLoginForbiddenCode is the HTTP code returned for type PostLoginForbidden
const PostLoginForbiddenCode int = 403

/*PostLoginForbidden Authentication Fail

swagger:response postLoginForbidden
*/
type PostLoginForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.AuthenticateResponse `json:"body,omitempty"`
}

// NewPostLoginForbidden creates PostLoginForbidden with default headers values
func NewPostLoginForbidden() *PostLoginForbidden {

	return &PostLoginForbidden{}
}

// WithPayload adds the payload to the post login forbidden response
func (o *PostLoginForbidden) WithPayload(payload *models.AuthenticateResponse) *PostLoginForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post login forbidden response
func (o *PostLoginForbidden) SetPayload(payload *models.AuthenticateResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoginForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostLoginInternalServerErrorCode is the HTTP code returned for type PostLoginInternalServerError
const PostLoginInternalServerErrorCode int = 500

/*PostLoginInternalServerError Internal Server Error

swagger:response postLoginInternalServerError
*/
type PostLoginInternalServerError struct {
}

// NewPostLoginInternalServerError creates PostLoginInternalServerError with default headers values
func NewPostLoginInternalServerError() *PostLoginInternalServerError {

	return &PostLoginInternalServerError{}
}

// WriteResponse to the client
func (o *PostLoginInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
