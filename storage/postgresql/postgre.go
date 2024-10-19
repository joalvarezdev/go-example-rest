package postgresql

import (
	"database/sql"

	"github.com/joalvarezdev/go-gpt/internal/config"
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