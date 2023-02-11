package productvalidator

import (
	"errors"
	"strings"

	"github.com/marc/go-clean-example/core/dto"
)

func ValidateInsert(productRequest *dto.CreateProductRequest) error {

	if strings.Trim(productRequest.Description, " ") == "" {
		return errors.New("Product description is required")
	}

	if strings.Trim(productRequest.Name, "") == "" {
		return errors.New("Product name is required")
	}
	//price, _ := strconv.ParseFloat(productRequest.Price, 64)
	if productRequest.Price <= 0 {
		return errors.New("Product price is required")
	}

	return nil

}

func ValidateUpdate(productRequest *dto.UpdateProductRequest) error {

	if productRequest.ID <= 0 {
		return errors.New("Product ID is required")
	}

	if strings.Trim(productRequest.Description, " ") == "" {
		return errors.New("Product description is required")
	}

	if strings.Trim(productRequest.Name, "") == "" {
		return errors.New("Product name is required")
	}

	//price, _ := strconv.ParseFloat(productRequest.Price, 64)
	if productRequest.Price <= 0 {
		return errors.New("Product price is required")
	}

	return nil

}
