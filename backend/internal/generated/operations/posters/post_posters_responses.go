// Code generated by go-swagger; DO NOT EDIT.

package posters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// PostPostersOKCode is the HTTP code returned for type PostPostersOK
const PostPostersOKCode int = 200

/*PostPostersOK постер успешно создан

swagger:response postPostersOK
*/
type PostPostersOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPostPostersOK creates PostPostersOK with default headers values
func NewPostPostersOK() *PostPostersOK {

	return &PostPostersOK{}
}

// WithPayload adds the payload to the post posters o k response
func (o *PostPostersOK) WithPayload(payload *models.ResultResponse) *PostPostersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post posters o k response
func (o *PostPostersOK) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPostersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostPostersBadRequestCode is the HTTP code returned for type PostPostersBadRequest
const PostPostersBadRequestCode int = 400

/*PostPostersBadRequest Bad request

swagger:response postPostersBadRequest
*/
type PostPostersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPostPostersBadRequest creates PostPostersBadRequest with default headers values
func NewPostPostersBadRequest() *PostPostersBadRequest {

	return &PostPostersBadRequest{}
}

// WithPayload adds the payload to the post posters bad request response
func (o *PostPostersBadRequest) WithPayload(payload *models.ResultResponse) *PostPostersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post posters bad request response
func (o *PostPostersBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPostersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostPostersInternalServerErrorCode is the HTTP code returned for type PostPostersInternalServerError
const PostPostersInternalServerErrorCode int = 500

/*PostPostersInternalServerError Internal Server Error

swagger:response postPostersInternalServerError
*/
type PostPostersInternalServerError struct {
}

// NewPostPostersInternalServerError creates PostPostersInternalServerError with default headers values
func NewPostPostersInternalServerError() *PostPostersInternalServerError {

	return &PostPostersInternalServerError{}
}

// WriteResponse to the client
func (o *PostPostersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}