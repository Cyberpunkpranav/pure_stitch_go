package controllers

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	utils "pure_stitch/utils/images"
	"strconv"
)

// Resize resizes the image to the given width and height using high-quality interpolation.
func Product_types_image(w http.ResponseWriter, r *http.Request) {
	// Parse parameters from query
	widthStr := r.URL.Query().Get("width")
	heightStr := r.URL.Query().Get("height")
	qualityStr := r.URL.Query().Get("quality")
	format := r.URL.Query().Get("format")
	image_name := r.URL.Query().Get("image")
	// Set default values if parameters are not provided
	fmt.Println(widthStr, heightStr, image_name, format)

	width, err := strconv.Atoi(widthStr)
	if err != nil || width <= 0 {
		width = 800 // default width
	}
	height, err := strconv.Atoi(heightStr)
	if err != nil || height <= 0 {
		height = 600 // default height
	}
	quality, err := strconv.Atoi(qualityStr)
	if err != nil || quality < 1 || quality > 100 {
		quality = 80 // default JPEG quality
	}
	if format == "" {
		format = "png" // default format
	}

	// Open the image file
	file, err := os.Open("assets/arrivals/" + image_name) // Specify your input image path
	if err != nil {
		http.Error(w, "Failed to open image", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Decode the image
	src, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Failed to decode image", http.StatusInternalServerError)
		return
	}

	// Resize the image based on the width and height parameters
	resized := utils.Resize(src, width, height)

	// Set appropriate headers and serve the resized image
	if format == "jpeg" || format == "jpg" {
		w.Header().Set("Content-Type", "image/jpeg")
		err = jpeg.Encode(w, resized, &jpeg.Options{Quality: quality})
	} else if format == "png" {
		w.Header().Set("Content-Type", "image/png")
		encoder := png.Encoder{CompressionLevel: png.BestCompression}
		err = encoder.Encode(w, resized)
	} else {
		http.Error(w, "Un	supported format", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Failed to encode image", http.StatusInternalServerError)
		return
	}
}
