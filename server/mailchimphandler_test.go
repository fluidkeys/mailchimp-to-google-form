package server

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"strings"
	"testing"
)

func TestPingEndpoint(t *testing.T) {
	t.Run("test webhook endpoint", func(t *testing.T) {
		formData := url.Values{}
		formData.Set("data[email]", "example@example.com")
		formData.Set("type", "subscribe")

		mockResponse := postForm(t, formData)
		assertStatusCode(t, http.StatusNoContent, mockResponse.Code)
	})
}

func postForm(t *testing.T, formData url.Values) *httptest.ResponseRecorder {
	// Create a request to pass to our handler. We don't have any query
	// parameters for now, so we'll pass 'nil' as the third parameter.
	t.Helper()

	req, err := http.NewRequest("POST", "/fake_secret", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//     r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	// create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	return recorder
}

func assertStatusCode(t *testing.T, expected int, got int) {
	t.Helper()
	if expected != got {
		t.Fatalf("expected HTTP %d, got HTTP %d", expected, got)
	}
}
