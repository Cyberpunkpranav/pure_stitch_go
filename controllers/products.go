package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"pure_stitch/config"
	"pure_stitch/models"
	sql "pure_stitch/sql"
)

func Get_products(w http.ResponseWriter, r *http.Request) {
	db := sql.DB
	query := "SELECT * FROM products "

	data, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var Products []models.Products
	for data.Next() {
		var product models.Products
		err := data.Scan(
			&product.Id,
			&product.SizeId,
			&product.ColorId,
			&product.ProductCategoryId,
			&product.ProductName,
			&product.ProductDescription,
			&product.Gender,
			&product.CreatedOn,
			&product.UpdatedOn,
		)
		if err != nil {
			log.Fatal(err)
		}
		Products = append(Products, product)
	}
	var Payload = config.Payload{
		Message: "produts",
		Status:  http.StatusText(200),
		Data:    Products,
	}
	json, err := json.Marshal(Payload)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	// db.Close()

	w.Write(json)

}
func Product_types(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id,type_name,type_description,image FROM product_types"
	db := sql.DB
	data, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var product_types []models.ProductTypes
	for data.Next() {
		var product_type models.ProductTypes
		err := data.Scan(
			&product_type.Id,
			&product_type.TypeName,
			&product_type.TypeDescription,
			&product_type.Image,
		)
		if err != nil {
			log.Fatal(err)
		}
		product_types = append(product_types, product_type)
	}
	var Payload = config.Payload{
		Message: "Product types",
		Status:  http.StatusText(200),
		Data:    product_types,
	}
	json, err := json.Marshal(Payload)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(json)

}
