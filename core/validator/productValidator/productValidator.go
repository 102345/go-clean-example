package productvalidator

import (
	"errors"
	"strings"

	"github.com/marc/go-clean-example/core/dto"
)

func ValidateInsert(productRequest *dto.CreateProductRequest) error {

	if strings.Trim(productRequest.Description, " ") == "" {
		return errors.New("Descrição do produto é obrigatório")
	}

	if strings.Trim(productRequest.Name, "") == "" {
		return errors.New("Nome do produto é obrigatório")
	}

	if productRequest.Price <= 0 {
		return errors.New("O preço do produto é obrigatório")
	}

	return nil

}

func ValidateUpdate(productRequest *dto.UpdateProductRequest) error {

	if productRequest.ID <= 0 {
		return errors.New("O ID do produto é obrigatório")
	}

	if strings.Trim(productRequest.Description, " ") == "" {
		return errors.New("Descrição do produto é obrigatório")
	}

	if strings.Trim(productRequest.Name, "") == "" {
		return errors.New("Nome do produto é obrigatório")
	}

	if productRequest.Price <= 0 {
		return errors.New("O preço do produto é obrigatório")
	}

	return nil

}
