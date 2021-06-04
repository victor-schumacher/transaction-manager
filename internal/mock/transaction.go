package mock

import "transaction-manager/database/postgres/repository"

type TransactionRepo struct {
}

func(tr TransactionRepo) Save(transaction repository.TransactionEntity) error {
	return nil
}
