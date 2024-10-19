package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/joalvarezdev/go-gpt/internal/config"
	"github.com/joalvarezdev/go-gpt/internal/types"
	_ "github.com/lib/pq"
)

type Postgre struct {
  Db *sql.DB
}

func New(config *config.Config) (*Postgre, error){

	 conn, err := sql.Open("postgres", config.StoragePath)

   if err != nil {
    return nil, err
   }

   return &Postgre{
    Db: conn,
   }, nil
}

func (p *Postgre) CreateProduct(name string, description string, price float64) (int64, error) {

  tx, err := p.Db.Begin()
  if err != nil {
    return 0, err
  }

  stmt, err := tx.Prepare("INSERT INTO products (name, description, price) VALUES ($1, $2, $3)")
  if err != nil {
    return 0, err
  }

  defer stmt.Close()

  var productId int64

  err = stmt.QueryRow(name, description, price).Scan(&productId)
  if err != nil {
    tx.Rollback()
    return 0, err
  }

  if err := tx.Commit(); err != nil {
    return 0, err
  }

  return productId, nil
}

func (p *Postgre) GetByIdProduct(id int64) (types.Product, error) {
  stmt, err := p.Db.Prepare("SELECT id, name, description, price FROM products WHERE id = $1")
  if err != nil {
    return types.Product{}, err
  }

  defer stmt.Close()

  var product types.Product

  err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Description, &product.Price)
  if err != nil {
    if err == sql.ErrNoRows {
      return types.Product{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
    }

    return types.Product{}, fmt.Errorf("query error: %w", err)
  }

  return product, nil
}