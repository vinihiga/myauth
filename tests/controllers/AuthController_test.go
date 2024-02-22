package controllers_test

import (
	"bytes"
	"myauth/internal/controllers"
	mocks_test "myauth/tests/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthHandler(t *testing.T) {
	// Given
	var buffer bytes.Buffer = bytes.Buffer{}
	buffer.WriteString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIifQ.efyqXgAxh_VBmy2TKDcVmbmC6MZ0YFK9PqYe3uEllnA")

	var req = httptest.NewRequest(http.MethodPost, "/auth", &buffer)
	var writer = httptest.NewRecorder()

	var providerMock mocks_test.ProviderMock = mocks_test.ProviderMock{}

	var sut controllers.AuthController = controllers.AuthController{
		Provider:   &providerMock,
		PrivateKey: []byte("test123"),
	}

	// When
	sut.AuthHandler(writer, req)

	res := writer.Result()
	defer res.Body.Close()

	// Then
	if res.StatusCode != http.StatusOK {
		t.Fail()
	}
}
