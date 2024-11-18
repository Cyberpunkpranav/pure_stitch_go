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

func New_arrivals(w http.ResponseWriter, r *http.Request) {
	db := sql.DB
	product_type_id := r.URL.Query().Get("product_type_id")
	query := "Select * FROM new_arrivals WHERE product_type_id = ?"
	arrivalsRows, err := db.Query(query, product_type_id)
	if err != nil {
		log.Fatal(err)
	}
	var Arrivals []models.Arrivals
	for arrivalsRows.Next() {
		var arrival models.Arrivals
		err := arrivalsRows.Scan(
			&arrival.Id,
			&arrival.ProductId,
			&arrival.ProductTypeId,
			&arrival.DateOfLaunch,
			&arrival.File,
			&arrival.OriginalHeight,
			&arrival.OriginalWidth,
			&arrival.Format,
			&arrival.BroadcastHours,
			&arrival.Active,
			&arrival.CreatedOn,
			&arrival.UpdatedOn,
		)
		if err != nil {
			log.Fatal(err)
		}
		query := "SELECT * FROM products WHERE id=?"
		productRows, err := db.Query(query, arrival.ProductId)
		if err != nil {
			log.Fatal(err)
		}
		var product models.Products
		for productRows.Next() {
			err := productRows.Scan(
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
		}
		var Media []models.Media
		query2 := "SELECT * FROM products_media WHERE product_id = ?"
		mediaRows, err := db.Query(query2, product.Id)
		if err != nil {
			log.Fatal(err)
		}
		for mediaRows.Next() {
			var media models.Media
			mediaRows.Scan(
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
			Media = append(Media, media)
		}
		product.Media = Media
		arrival.Product = product
		Arrivals = append(Arrivals, arrival)
	}
	var Payload = config.Payload{
		Message: "arrivals",
		Status:  http.StatusText(200),
		Data:    Arrivals,
	}
	json, err := json.Marshal(Payload)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	// db.Close()

	w.Write(json)
}

func Arrival_post(w http.ResponseWriter, r *http.Request) {
	widthStr := r.URL.Query().Get("width")
	heightStr := r.URL.Query().Get("height")
	qualityStr := r.URL.Query().Get("quality")
	format := r.URL.Query().Get("format")
	image_name := r.URL.Query().Get("image")
	category := r.URL.Query().Get("category")
	path := "assets/arrivals/post/" + category + "/"
	utils.Optimizing_image(widthStr, heightStr, qualityStr, format, image_name, path, w)
}
