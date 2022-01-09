package app

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type StorageMock struct {
	mock.Mock
}

func (sm *StorageMock) AddURL(id string, url string) {}
func (sm *StorageMock) GetURL(id string) string      { return "" }

type URLServiceMock struct {
	mock.Mock
}

func (usm *URLServiceMock) CreateShortURL(url string) string {
	return "testID"
}
func (usm *URLServiceMock) GetLongURLByID(id string) (string, error) {
	if id == "not_exist_id" {
		return "", errors.New("")
	}
	return "http://rp21sh.yandex/avshz7qrl/h4ululwnp7bow", nil
}

func TestURLShortenerHandler_HandleShortURL(t *testing.T) {
	type want struct {
		contentType    string
		headerLocation string
		statusCode     int
		result         string
	}
	tests := []struct {
		name        string
		method      string
		request     string
		want        want
		requestBody string
	}{
		{
			name:   "POST endpoint positive test #1",
			method: http.MethodPost,
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusCreated,
				result:      "testID",
			},
			request:     "/",
			requestBody: "http://rp21sh.yandex/avshz7qrl/h4ululwnp7bow",
		},

		{
			name:   "POST endpoint bad request test #1",
			method: http.MethodPost,
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusBadRequest,
				result:      "StatusBadRequest\n",
			},
			request:     "/",
			requestBody: "",
		},

		{
			name:   "POST endpoint bad request test #2",
			method: http.MethodPost,
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusBadRequest,
				result:      "StatusBadRequest\n",
			},
			request:     "/",
			requestBody: "1234567",
		},

		{
			name:   "GET endpoint positive #1",
			method: http.MethodGet,
			want: want{
				contentType:    "text/plain; charset=utf-8",
				headerLocation: "http://rp21sh.yandex/avshz7qrl/h4ululwnp7bow",
				statusCode:     http.StatusTemporaryRedirect,
				result:         "",
			},
			request:     "/testID",
			requestBody: "",
		},

		{
			name:   "GET endpoint bad request test #1",
			method: http.MethodGet,
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusNotFound,
				result:      "Not found",
			},
			request:     "/not_exist_id",
			requestBody: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, tt.request, bytes.NewBuffer([]byte(tt.requestBody)))
			w := httptest.NewRecorder()
			handler := NewURLShortenerHandler(new(URLServiceMock))

			h := http.HandlerFunc(handler.HandleShortURL)
			h.ServeHTTP(w, request)
			result := w.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)

			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			if tt.method == http.MethodGet {
				assert.Equal(t, tt.want.headerLocation, w.Header().Get("Location"))
				return
			}

			if tt.method == http.MethodPost {
				response, err := ioutil.ReadAll(result.Body)
				require.NoError(t, err)

				err = result.Body.Close()
				require.NoError(t, err)

				data := strings.Split(string(response), "/")
				id := data[len(data)-1]

				assert.Equal(t, tt.want.result, id)
				return
			}
		})
	}
}
