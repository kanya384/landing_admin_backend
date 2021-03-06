// Code generated by go-swagger; DO NOT EDIT.

package posters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// PostPostersOrdersOKCode is the HTTP code returned for type PostPostersOrdersOK
const PostPostersOrdersOKCode int = 200

/*PostPostersOrdersOK положения успешно обновленны

swagger:response postPostersOrdersOK
*/
type PostPostersOrdersOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPostPostersOrdersOK creates PostPostersOrdersOK with default headers values
func NewPostPostersOrdersOK() *PostPostersOrdersOK {

	return &PostPostersOrdersOK{}
}

// WithPayload adds the payload to the post posters orders o k response
func (o *PostPostersOrdersOK) WithPayload(payload *models.ResultResponse) *PostPostersOrdersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post posters orders o k response
func (o *PostPostersOrdersOK) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPostersOrdersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostPostersOrdersBadRequestCode is the HTTP code returned for type PostPostersOrdersBadRequest
const PostPostersOrdersBadRequestCode int = 400

/*PostPostersOrdersBadRequest Bad request

swagger:response postPostersOrdersBadRequest
*/
type PostPostersOrdersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPostPostersOrdersBadRequest creates PostPostersOrdersBadRequest with default headers values
func NewPostPostersOrdersBadRequest() *PostPostersOrdersBadRequest {

	return &PostPostersOrdersBadRequest{}
}

// WithPayload adds the payload to the post posters orders bad request response
func (o *PostPostersOrdersBadRequest) WithPayload(payload *models.ResultResponse) *PostPostersOrdersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post posters orders bad request response
func (o *PostPostersOrdersBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPostersOrdersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostPostersOrdersInternalServerErrorCode is the HTTP code returned for type PostPostersOrdersInternalServerError
const PostPostersOrdersInternalServerErrorCode int = 500

/*PostPostersOrdersInternalServerError Internal Server Error

swagger:response postPostersOrdersInternalServerError
*/
type PostPostersOrdersInternalServerError struct {
}

// NewPostPostersOrdersInternalServerError creates PostPostersOrdersInternalServerError with default headers values
func NewPostPostersOrdersInternalServerError() *PostPostersOrdersInternalServerError {

	return &PostPostersOrdersInternalServerError{}
}

// WriteResponse to the client
func (o *PostPostersOrdersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
