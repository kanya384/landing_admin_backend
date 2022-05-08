// Code generated by go-swagger; DO NOT EDIT.

package advantages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// PostAdvantagePhotoOrdersOKCode is the HTTP code returned for type PostAdvantagePhotoOrdersOK
const PostAdvantagePhotoOrdersOKCode int = 200

/*PostAdvantagePhotoOrdersOK положения успешно обновленны

swagger:response postAdvantagePhotoOrdersOK
*/
type PostAdvantagePhotoOrdersOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPostAdvantagePhotoOrdersOK creates PostAdvantagePhotoOrdersOK with default headers values
func NewPostAdvantagePhotoOrdersOK() *PostAdvantagePhotoOrdersOK {

	return &PostAdvantagePhotoOrdersOK{}
}

// WithPayload adds the payload to the post advantage photo orders o k response
func (o *PostAdvantagePhotoOrdersOK) WithPayload(payload *models.ResultResponse) *PostAdvantagePhotoOrdersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post advantage photo orders o k response
func (o *PostAdvantagePhotoOrdersOK) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAdvantagePhotoOrdersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAdvantagePhotoOrdersBadRequestCode is the HTTP code returned for type PostAdvantagePhotoOrdersBadRequest
const PostAdvantagePhotoOrdersBadRequestCode int = 400

/*PostAdvantagePhotoOrdersBadRequest Bad request

swagger:response postAdvantagePhotoOrdersBadRequest
*/
type PostAdvantagePhotoOrdersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPostAdvantagePhotoOrdersBadRequest creates PostAdvantagePhotoOrdersBadRequest with default headers values
func NewPostAdvantagePhotoOrdersBadRequest() *PostAdvantagePhotoOrdersBadRequest {

	return &PostAdvantagePhotoOrdersBadRequest{}
}

// WithPayload adds the payload to the post advantage photo orders bad request response
func (o *PostAdvantagePhotoOrdersBadRequest) WithPayload(payload *models.ResultResponse) *PostAdvantagePhotoOrdersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post advantage photo orders bad request response
func (o *PostAdvantagePhotoOrdersBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAdvantagePhotoOrdersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAdvantagePhotoOrdersInternalServerErrorCode is the HTTP code returned for type PostAdvantagePhotoOrdersInternalServerError
const PostAdvantagePhotoOrdersInternalServerErrorCode int = 500

/*PostAdvantagePhotoOrdersInternalServerError Internal Server Error

swagger:response postAdvantagePhotoOrdersInternalServerError
*/
type PostAdvantagePhotoOrdersInternalServerError struct {
}

// NewPostAdvantagePhotoOrdersInternalServerError creates PostAdvantagePhotoOrdersInternalServerError with default headers values
func NewPostAdvantagePhotoOrdersInternalServerError() *PostAdvantagePhotoOrdersInternalServerError {

	return &PostAdvantagePhotoOrdersInternalServerError{}
}

// WriteResponse to the client
func (o *PostAdvantagePhotoOrdersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}