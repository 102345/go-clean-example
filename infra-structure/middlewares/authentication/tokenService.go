package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// CreateToken retorna um token assinado com as permissões do usuario
func CreateToken(usuarioID uint64) (string, error) {

	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString([]byte(viper.GetString("tokenKey.secretKey"))) //Secret

}

// ValidateToken verifica se token passado pela requisição é válido
func ValidateToken(r *http.Request) error {

	tokenString := extractToken(r)

	token, erro := jwt.Parse(tokenString, returnVerificationKey)
	if erro != nil {
		//fmt.Print("Entrou no jwt.Parse do ValidateToken")
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid Token")
}

func ExtractUserID(r *http.Request) (uint64, error) {

	tokenString := extractToken(r)

	token, erro := jwt.Parse(tokenString, returnVerificationKey)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, errors.New("Unexpected user extraction error!")
		}

		return usuarioID, nil
	}

	return 0, errors.New("Invalid Token")
}

func extractToken(r *http.Request) string {

	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

		return nil, fmt.Errorf("Unexpected signature method! %v", token.Header["alg"])
	}

	return []byte(viper.GetString("tokenKey.secretKey")), nil

}
