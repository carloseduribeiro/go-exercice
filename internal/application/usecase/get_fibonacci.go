package usecase

import "github.com/carloseduribeiro/go-exercice/internal/domain/entity"

type GetFibonacciUseCase struct {
}

func NewGetFibonacci() GetFibonacciUseCase {
	return GetFibonacciUseCase{}
}

func (g GetFibonacciUseCase) Execute(input int) GetFibonacciOutputDto {
	sequenceResult := entity.Fibonacci(input)
	return GetFibonacciOutputDto{
		Sequence: sequenceResult,
	}
}

type GetFibonacciOutputDto struct {
	Sequence int
}
