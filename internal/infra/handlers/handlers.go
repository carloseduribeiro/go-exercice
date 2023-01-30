package handlers

import (
	"encoding/json"
	"github.com/carloseduribeiro/go-exercise/internal/application/usecase"
	"net/http"
)

type ValidatePalindromeHandler struct {
	method   string
	resource string
}

func NewValidatePalindrome(method, resource string) ValidatePalindromeHandler {
	return ValidatePalindromeHandler{method, resource}
}

func (v ValidatePalindromeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if r.URL.Path != v.resource {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	sentenceParam := r.URL.Query().Get("sentence")
	if len(sentenceParam) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	useCase := usecase.NewValidatePalindrome()
	result := useCase.Execute(sentenceParam)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
