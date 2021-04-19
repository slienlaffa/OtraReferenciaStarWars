package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerTopSecret(t *testing.T) {
	message := []byte(`{
		"satellites": [
		{
		"name": "kenobi",
		"distance": 100.0,
		"message": ["este", "", "", "mensaje", ""]
		},
		{
		"name": "skywalker",
		"distance": 115.5,
		"message": ["", "es", "", "", "secreto"]
		},
		{
		"name": "sato",
		"distance": 142.7,
		"message": ["este", "", "un", "", ""]
		}
		]}`)
	req, err := http.NewRequest("POST", "/topsecret", bytes.NewBuffer(message))
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()

	handlerTopSecret(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
	}

	var data Response
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		t.Fatalf("could not unmarshall response %v", err)
	}

	responseExpected := Response{
		Coordinates{-487.29, 1557.01},
		"este es un mensaje secreto",
	}
	if data != responseExpected {
		t.Errorf("expected %v, got: %v", responseExpected, data)
	}
}

func TestHandlerTopSecretEmpty(t *testing.T) {
	message := []byte(`{
		"satellites": [
		{
		"name": "kenobi",
		"distance": 100.0,
		"message": ["", "", "", "", ""]
		},
		{
		"name": "skywalker",
		"distance": 115.5,
		"message": ["", "", "", "", ""]
		},
		{
		"name": "sato",
		"distance": 142.7,
		"message": ["", "", "", "", ""]
		}
		]}`)
	req, err := http.NewRequest("POST", "/topsecret", bytes.NewBuffer(message))
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()

	handlerTopSecret(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected %d, got: %d", http.StatusNotFound, res.StatusCode)
	}

}

func TestHandlerTopSecretSplitEmpty(t *testing.T) {
	req, err := http.NewRequest("GET", "/topsecret_split/", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()

	handlerTopSecretSplit(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected %d, got: %d", http.StatusNotFound, res.StatusCode)
	}
}

func TestHandlerTopSecretSplitPost(t *testing.T) {
	message := []byte(`{
		"distance": 100.0,
		"message": ["este", "", "", "mensaje", ""]
		}`)
	req, err := http.NewRequest("POST", "/topsecret_split/kenobi", bytes.NewBuffer(message))
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()

	handlerTopSecretSplit(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
	}
}

func TestHandlerTopSecretSplitGetCorrect(t *testing.T) {
	message1 := []byte(`{
		"distance": 100.0,
		"message": ["este", "", "", "mensaje", ""]
		}`)
	message2 := []byte(`{
		"distance": 115.5,
		"message": ["", "es", "", "", "secreto"]
		}`)
	message3 := []byte(`{
		"distance": 142.7,
		"message": ["este", "", "un", "", ""]
		}`)
	req1, _ := http.NewRequest("POST", "/topsecret_split/kenobi", bytes.NewBuffer(message1))
	req2, _ := http.NewRequest("POST", "/topsecret_split/skywalker", bytes.NewBuffer(message2))
	req3, _ := http.NewRequest("POST", "/topsecret_split/sato", bytes.NewBuffer(message3))
	rec := httptest.NewRecorder()

	handlerTopSecretSplit(rec, req1)
	handlerTopSecretSplit(rec, req2)
	handlerTopSecretSplit(rec, req3)

	req, _ := http.NewRequest("GET", "/topsecret_split/", nil)
	rec = httptest.NewRecorder() // Limpiar
	handlerTopSecretSplit(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
	}

	var data Response
	err := json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		t.Fatalf("could not unmarshall response %v", err)
	}

	responseExpected := Response{
		Coordinates{-487.29, 1557.01},
		"este es un mensaje secreto",
	}
	if data != responseExpected {
		t.Errorf("expected %v, got: %v", responseExpected, data)
	}
}
