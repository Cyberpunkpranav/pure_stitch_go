package models


type Products struct{
	Id int `json:"id"`
	ProductName string `json:"product_name"`
	ProductDescription string `json:"product_description"`
}