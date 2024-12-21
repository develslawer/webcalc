package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalcHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		expectedCode   int
		expectedResult *Response
		expectedError  *ErrResponse
	}{
		{
			name:           "ValidRequest",
			requestBody:    Request{Expression: "1+1"},
			expectedCode:   http.StatusOK,
			expectedResult: &Response{Result: 2.0},
		},
		{
			name:          "InvalidRequestBody",
			requestBody:   map[string]interface{}{"invalid": "data"},
			expectedCode:  http.StatusBadRequest,
			expectedError: &ErrResponse{Error: ErrInvalidExpression.Error()},
		},
		{
			name:          "UnknownChar",
			requestBody:   Request{Expression: "1+1.."},
			expectedCode:  http.StatusInternalServerError,
			expectedError: &ErrResponse{Error: ErrInternalServer.Error()},
		},
		{
			name:          "StackOverflow",
			requestBody:   Request{Expression: "1+1++"},
			expectedCode:  http.StatusBadRequest,
			expectedError: &ErrResponse{Error: ErrInvalidExpression.Error()},
		},
		{
			name:          "EmptyRequestBody",
			requestBody:   nil,
			expectedCode:  http.StatusBadRequest,
			expectedError: &ErrResponse{Error: ErrInvalidExpression.Error()},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var body []byte
			if test.requestBody != nil {
				body, _ = json.Marshal(test.requestBody)
			}

			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(body))
			rec := httptest.NewRecorder()

			CalcHandler(rec, req)

			resp := rec.Result()
			defer resp.Body.Close()

			if resp.StatusCode != test.expectedCode {
				t.Errorf("expected status %d, got %d", test.expectedCode, resp.StatusCode)
			}

			if test.expectedResult != nil {
				var actualResult Response
				err := json.Unmarshal(rec.Body.Bytes(), &actualResult)
				if err != nil {
					t.Fatalf("failed to unmarshal response body: %v", err)
				}

				if actualResult != *test.expectedResult {
					t.Errorf("expected result %+v, got %+v", *test.expectedResult, actualResult)
				}
			} else if test.expectedError != nil {
				actualErrorBytes, _ := io.ReadAll(rec.Body)
				actualError := ErrResponse{}
				fmt.Printf("%s", string(actualErrorBytes))
				json.NewDecoder(bytes.NewReader(actualErrorBytes)).Decode(&actualError)
				if *test.expectedError != actualError {
					t.Errorf("expected error %+v, got %+v", *test.expectedError, actualError)
				}
			}
		})
	}
}

func TestLoggerMiddleware(t *testing.T) {
	handler := LoggerMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/api/v1/calculate", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}
