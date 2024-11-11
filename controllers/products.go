package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"pure_stitch/config"
	"pure_stitch/models"
	sql "pure_stitch/sql"
	utils "pure_stitch/utils/images"
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
		query := "SELECT * FROM products_media WHERE product_id =? "
		var Media []models.Media
		mediaRows, eror := db.Query(query, product.Id)
		for mediaRows.Next() {
			var media models.Media
			err := mediaRows.Scan(
				&media.Id,
				&media.ProductId,
				&media.MediaType,
				&media.Category,
				&media.Format,
				&media.OriginalHeight,
				&media.OriginalWidth,
				&media.File,
				&media.CreatedOn,
				&media.UpdatedOn,
			)
			if eror != nil {
				log.Fatal(err)
			}
			Media = append(Media, media)
		}
		product.Media = Media
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
func Product_images(w http.ResponseWriter, r *http.Request) {
	widthStr := r.URL.Query().Get("width")
	heightStr := r.URL.Query().Get("height")
	qualityStr := r.URL.Query().Get("quality")
	format := r.URL.Query().Get("format")
	image_name := r.URL.Query().Get("image")
	category := r.URL.Query().Get("category")
	path := "assets/posts/" + category + "/"
	utils.Optimizing_image(widthStr, heightStr, qualityStr, format, image_name, path, w)
}
func Product_types_images(w http.ResponseWriter, r *http.Request) {
	widthStr := r.URL.Query().Get("width")
	heightStr := r.URL.Query().Get("height")
	qualityStr := r.URL.Query().Get("quality")
	format := r.URL.Query().Get("format")
	image_name := r.URL.Query().Get("image")
	path := "assets/arrivals/"
	utils.Optimizing_image(widthStr, heightStr, qualityStr, format, image_name, path, w)
}
