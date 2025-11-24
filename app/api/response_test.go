package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sampleResponse struct {
	Message string `json:"message"`
}

func TestHttpResponder_Ok(t *testing.T) {
	type args struct {
		data any
	}
	type want struct {
		statusCode  int
		contentType string
		body        string
		hasError    bool
	}
	tests := []struct {
		name  string
		args  args
		want  want
		setup func(t *testing.T, args args) http.ResponseWriter
	}{
		{
			name: "succesful http200 json response",
			args: args{
				data: sampleResponse{Message: "Success"},
			},
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				body:        `{"message":"Success"}`,
				hasError:    false,
			},
			setup: func(t *testing.T, args args) http.ResponseWriter {
				return httptest.NewRecorder()
			},
		},
		{
			name: "handle encoding error",
			args: args{
				data: sampleResponse{Message: "Success"},
			},
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json",
				hasError:    true,
			},
			setup: func(t *testing.T, args args) http.ResponseWriter {
				return &errorResponseWriter{
					recorder: httptest.NewRecorder(),
					writeErr: errors.New("write error"),
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			responder := NewHttpResponder()
			w := test.setup(t, test.args)

			responder.Ok(w, test.args.data)

			a := assert.New(t)

			if errWriter, ok := w.(*errorResponseWriter); ok {
				a.Equal(test.want.statusCode, errWriter.recorder.Code, "Expected status code to match")
				a.Equal(test.want.contentType, errWriter.recorder.Header().Get("Content-Type"), "Expected Content-Type to match")
			} else if recorder, ok := w.(*httptest.ResponseRecorder); ok {
				a.Equal(test.want.statusCode, recorder.Code, "Expected status code to match")
				a.Equal(test.want.contentType, recorder.Header().Get("Content-Type"), "Expected Content-Type to match")
				if test.want.body != "" {
					a.JSONEq(test.want.body, recorder.Body.String(), "Response body does not match expected")
				}
			}
		})
	}
}

func TestHttpResponder_Error(t *testing.T) {
	type args struct {
		status  int
		message string
	}
	type want struct {
		statusCode  int
		contentType string
		body        string
		hasError    bool
	}
	tests := []struct {
		name  string
		args  args
		want  want
		setup func(t *testing.T, args args) http.ResponseWriter
	}{
		{
			name: "json response for a given http status code",
			args: args{
				status:  http.StatusInternalServerError,
				message: "Some error occurred",
			},
			want: want{
				statusCode:  http.StatusInternalServerError,
				contentType: "application/json",
				body:        `{"error":"Some error occurred"}`,
				hasError:    false,
			},
			setup: func(t *testing.T, args args) http.ResponseWriter {
				return httptest.NewRecorder()
			},
		},
		{
			name: "handle encoding error",
			args: args{
				status:  http.StatusBadRequest,
				message: "Bad request",
			},
			want: want{
				statusCode:  http.StatusBadRequest,
				contentType: "application/json",
				hasError:    true,
			},
			setup: func(t *testing.T, args args) http.ResponseWriter {
				return &errorResponseWriter{
					recorder: httptest.NewRecorder(),
					writeErr: errors.New("write error"),
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			responder := NewHttpResponder()
			w := test.setup(t, test.args)

			responder.Error(w, test.args.status, test.args.message)

			a := assert.New(t)

			if errWriter, ok := w.(*errorResponseWriter); ok {
				a.Equal(test.want.statusCode, errWriter.recorder.Code, "Expected status code to match")
				a.Equal(test.want.contentType, errWriter.recorder.Header().Get("Content-Type"), "Expected Content-Type to match")
			} else if recorder, ok := w.(*httptest.ResponseRecorder); ok {
				a.Equal(test.want.statusCode, recorder.Code, "Expected status code to match")
				a.Equal(test.want.contentType, recorder.Header().Get("Content-Type"), "Expected Content-Type to match")
				if test.want.body != "" {
					a.JSONEq(test.want.body, recorder.Body.String(), "Response body does not match expected")
				}
			}
		})
	}
}

type errorResponseWriter struct {
	recorder *httptest.ResponseRecorder
	writeErr error
	header   http.Header
}

func (e *errorResponseWriter) Header() http.Header {
	if e.header == nil {
		e.header = e.recorder.Header()
	}
	return e.header
}

func (e *errorResponseWriter) Write(b []byte) (int, error) {
	return 0, e.writeErr
}

func (e *errorResponseWriter) WriteHeader(statusCode int) {
	e.recorder.WriteHeader(statusCode)
}
