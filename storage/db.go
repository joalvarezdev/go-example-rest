package storage

type Storage interface {
  CreateProduct(name string, description string, price float64) (int64, error)
}