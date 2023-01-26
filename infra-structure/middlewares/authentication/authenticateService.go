package authentication

import (
	"log"
	"net/http"

	infrastructure "github.com/marc/go-clean-example/infra-structure"
)

// Logger escreve informações da requisição no terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}

}

// Authenticate verifica se o usuário faz a requisição autenticado
func Authenticate(nextFunction http.HandlerFunc, isAuthenticated bool) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if isAuthenticated {
			if erro := ValidateToken(r); erro != nil {
				infrastructure.Erro(w, http.StatusUnauthorized, erro)
				return
			}
		}

		nextFunction(w, r)
	}

}
