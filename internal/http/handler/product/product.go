package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joalvarezdev/go-gpt/internal/types"
	"github.com/joalvarezdev/go-gpt/internal/utils/response"
	"github.com/joalvarezdev/go-gpt/storage"
)

func Create(storage storage.Storage) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    slog.Info("product created")

    var product types.Product

    err := json.NewDecoder(r.Body).Decode(&product)

    if errors.Is(err, io.EOF) {
      response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
      return
    }

    if err != nil {
      response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
      return
    }

    if err := validator.New().Struct(product); err != nil {
      validateErrors := err.(validator.ValidationErrors)
      response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrors))
      return
    }

    productId, err := storage.CreateProduct(
      product.Name,
      product.Description,
      product.Price,
    )

    if err != nil {
      slog.Error(err.Error())
      response.WriteJson(w, http.StatusInternalServerError, err)
      return
    }

    response.WriteJson(w, http.StatusCreated, map[string] interface{}{
      "id": productId,
      "name": product.Name,
      "description": product.Description,
      "price": product.Price,
    })
  }
}

func GetById(storage storage.Storage) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")

    slog.Info("getting a product", slog.String("id", id))

    intId, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
      response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
      return
    }

    product, err := storage.GetByIdProduct(intId)
    if err != nil {
      slog.Error(err.Error())
      response.WriteJson(w, http.StatusBadRequest, err)
      return
    }

    response.WriteJson(w, http.StatusOK, product)
  }
}

func GetAll(storage storage.Storage) http.HandlerFunc {
  return func(w http.ResponseWriter, r*http.Request) {
    slog.Info("getting all products")

    products, err := storage.GetAllProducts()
    if err != nil {
      response.WriteJson(w, http.StatusInternalServerError, err)
      return
    }

    response.WriteJson(w, http.StatusOK, products)
  }
}