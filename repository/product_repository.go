package repository

import (
	"database/sql"

	"github.com/joalvarezdev/go-gpt/database"
	"github.com/joalvarezdev/go-gpt/model"
)

func GetAllProducts() ([]model.Product, error) {
	rows, err := database.DB.Query("select id, name, description, price FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func GetProductById(id string) (model.Product, error) {
	var product model.Product
	err := database.DB.QueryRow("SELECT id, name, description, price FROM products WHERE id = $1", id).Scan(
		&product.Id, &product.Name, &product.Description, &product.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return product, nil
		}
		return product, err
	}

	return product, nil
}

func CreateProduct(product model.Product) error {
	_, err := database.DB.Exec("INSERT INTO products (id, name, description, price) VALUES ($1, $2, $3, $4)",
				product.Id, product.Name, product.Description, product.Price)

	return err
}

func UpdateProduct(product model.Product) error {
	_, err := database.DB.Exec("UPDATE products SET name = %2, description = $3, price = $4 WHERE id = $1",
				product.Id, product.Name, product.Description, product.Price)

	return err
}

func DeleteProduct(id string) error {
	_, err := database.DB.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}