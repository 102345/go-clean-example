package security

import "golang.org/x/crypto/bcrypt"

//Hash recebe uma string e coloca um hash nela
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerificarSenha compara uma senha e um hash e retorna se elas são iguais
func VerifyPassword(passwordHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}
