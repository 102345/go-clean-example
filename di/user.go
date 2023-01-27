package di

import (
	"github.com/marc/go-clean-example/adapter/http/userservice"
	"github.com/marc/go-clean-example/adapter/postgres"
	"github.com/marc/go-clean-example/adapter/postgres/userrepository"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/usecase/userusecase"
)

// ConfiguserDI return a UserService abstraction with dependency injection configuration
func ConfigUserDI(conn postgres.PoolInterface) domain.UserService {
	userRepository := userrepository.New(conn)
	userUseCase := userusecase.New(userRepository)
	userService := userservice.New(userUseCase)

	return userService
}
