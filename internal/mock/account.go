package mock

import (
	"github.com/google/uuid"
	"time"
	"transaction-manager/database/postgres/repository"
)

type AccountRepo struct {
}

func (ar AccountRepo) Save(a repository.AccountEntity) error {
	return nil
}
func (ar AccountRepo) FindOne(ID uuid.UUID) (repository.AccountEntity, error){
	a := repository.AccountEntity{
		ID:             uuid.New(),
		DocumentNumber: "09365523957",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	return a, nil
}