package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HelloHandler)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    var response Response
    err = json.NewDecoder(rr.Body).Decode(&response)
    if err != nil {
        t.Fatal(err)
    }

    expected := "Hello, Docker CI/CD!"
    if response.Message != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            response.Message, expected)
    }
}
