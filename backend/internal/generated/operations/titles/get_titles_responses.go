// Code generated by go-swagger; DO NOT EDIT.

package titles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// GetTitlesOKCode is the HTTP code returned for type GetTitlesOK
const GetTitlesOKCode int = 200

/*GetTitlesOK returns titles list

swagger:response getTitlesOK
*/
type GetTitlesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Title `json:"body,omitempty"`
}

// NewGetTitlesOK creates GetTitlesOK with default headers values
func NewGetTitlesOK() *GetTitlesOK {

	return &GetTitlesOK{}
}

// WithPayload adds the payload to the get titles o k response
func (o *GetTitlesOK) WithPayload(payload []*models.Title) *GetTitlesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get titles o k response
func (o *GetTitlesOK) SetPayload(payload []*models.Title) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTitlesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Title, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetTitlesBadRequestCode is the HTTP code returned for type GetTitlesBadRequest
const GetTitlesBadRequestCode int = 400

/*GetTitlesBadRequest Bad request

swagger:response getTitlesBadRequest
*/
type GetTitlesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewGetTitlesBadRequest creates GetTitlesBadRequest with default headers values
func NewGetTitlesBadRequest() *GetTitlesBadRequest {

	return &GetTitlesBadRequest{}
}

// WithPayload adds the payload to the get titles bad request response
func (o *GetTitlesBadRequest) WithPayload(payload *models.ResultResponse) *GetTitlesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get titles bad request response
func (o *GetTitlesBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTitlesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTitlesInternalServerErrorCode is the HTTP code returned for type GetTitlesInternalServerError
const GetTitlesInternalServerErrorCode int = 500

/*GetTitlesInternalServerError Internal Server Error

swagger:response getTitlesInternalServerError
*/
type GetTitlesInternalServerError struct {
}

// NewGetTitlesInternalServerError creates GetTitlesInternalServerError with default headers values
func NewGetTitlesInternalServerError() *GetTitlesInternalServerError {

	return &GetTitlesInternalServerError{}
}

// WriteResponse to the client
func (o *GetTitlesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
