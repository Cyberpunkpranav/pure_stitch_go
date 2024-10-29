package models

type Products struct {
	Id                 int     `json:"id"`
	SizeId             int     `json:"size_id"`
	ColorId            int     `json:"color_id"`
	ProductCategoryId  int     `json:"product_category_id"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	Gender             string  `json:"gender"`
	CreatedOn          *string `json:"created_on"`
	UpdatedOn          *string `json:"updated_on"`
}

type ProductTypes struct {
	Id              int     `json:"id"`
	TypeName        string  `json:"type_name"`
	TypeDescription *string `json:"type_description"`
	Image           *string `json:"image"`
}
