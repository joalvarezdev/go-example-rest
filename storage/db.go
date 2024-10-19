package storage

import "github.com/joalvarezdev/go-gpt/internal/types"

type Storage interface {
  CreateProduct(name string, description string, price float64) (int64, error)
  GetByIdProduct(id int64) (types.Product, error)
}