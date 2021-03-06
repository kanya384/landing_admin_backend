// Code generated by go-swagger; DO NOT EDIT.

package docs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// GetDocOKCode is the HTTP code returned for type GetDocOK
const GetDocOKCode int = 200

/*GetDocOK returns docs list

swagger:response getDocOK
*/
type GetDocOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Doc `json:"body,omitempty"`
}

// NewGetDocOK creates GetDocOK with default headers values
func NewGetDocOK() *GetDocOK {

	return &GetDocOK{}
}

// WithPayload adds the payload to the get doc o k response
func (o *GetDocOK) WithPayload(payload []*models.Doc) *GetDocOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get doc o k response
func (o *GetDocOK) SetPayload(payload []*models.Doc) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDocOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Doc, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetDocBadRequestCode is the HTTP code returned for type GetDocBadRequest
const GetDocBadRequestCode int = 400

/*GetDocBadRequest Bad request

swagger:response getDocBadRequest
*/
type GetDocBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewGetDocBadRequest creates GetDocBadRequest with default headers values
func NewGetDocBadRequest() *GetDocBadRequest {

	return &GetDocBadRequest{}
}

// WithPayload adds the payload to the get doc bad request response
func (o *GetDocBadRequest) WithPayload(payload *models.ResultResponse) *GetDocBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get doc bad request response
func (o *GetDocBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDocBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDocInternalServerErrorCode is the HTTP code returned for type GetDocInternalServerError
const GetDocInternalServerErrorCode int = 500

/*GetDocInternalServerError Internal Server Error

swagger:response getDocInternalServerError
*/
type GetDocInternalServerError struct {
}

// NewGetDocInternalServerError creates GetDocInternalServerError with default headers values
func NewGetDocInternalServerError() *GetDocInternalServerError {

	return &GetDocInternalServerError{}
}

// WriteResponse to the client
func (o *GetDocInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
