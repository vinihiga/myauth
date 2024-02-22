package controllers_test

import (
	"bytes"
	"encoding/json"
	"myauth/internal/controllers"
	mocks_test "myauth/tests/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

type login_mock struct {
	Username string
	Password string
}

func TestLoginHandler(t *testing.T) {
	// Given
	var mock = login_mock{
		Username: "test",
		Password: "test",
	}

	var buffer bytes.Buffer
	_ = json.NewEncoder(&buffer).Encode(mock)

	var req = httptest.NewRequest(http.MethodPost, "/login", &buffer)
	var writer = httptest.NewRecorder()

	var providerMock mocks_test.ProviderMock = mocks_test.ProviderMock{}

	var sut controllers.LoginController = controllers.LoginController{
		Provider:   &providerMock,
		PrivateKey: []byte("test123"),
	}

	// When
	sut.LoginHandler(writer, req)

	res := writer.Result()
	defer res.Body.Close()

	// Then
	if res.StatusCode != http.StatusOK {
		t.Fail()
	}
}
