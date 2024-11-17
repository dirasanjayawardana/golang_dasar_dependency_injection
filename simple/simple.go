package simple

import "errors"

// Untuk melakukan Dependecy Injection, perlu membuat function Constructor pada struct
// Untuk penamaannya biasanya dengan `NewNamaStructnya()`
// Dalam Google wire, ini disebut Provider

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("failed create service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}
}
