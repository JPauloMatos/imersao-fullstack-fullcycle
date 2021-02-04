package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/jpaulomatos/imersao/codepix-go/application/usecase"
	"github.com/jpaulomatos/imersao/codepix-go/infrastructure/repository"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         pixRepository,
	}
	return transactionUseCase
}
