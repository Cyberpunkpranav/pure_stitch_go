package models

type Products struct {
	Id                 int     `json:"id"`
	ProductCategoryId  int     `json:"product_category_id"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	Gender             string  `json:"gender"`
	CreatedOn          *string `json:"created_on"`
	UpdatedOn          *string `json:"updated_on"`
	Media              []Media `json:"media"`
}

type ProductTypes struct {
	Id              int     `json:"id"`
	TypeName        string  `json:"type_name"`
	TypeDescription *string `json:"type_description"`
	Image           *string `json:"image"`
}

type Media struct {
	Id        int     `json:"id"`
	ProductId int     `json:"product_id"`
	MediaType string  `json:"media_type"`
	Category  string  `json:"category"`
	Format    string  `json:"format"`
	File      string  `json:"file"`
	CreatedOn *string `json:"created_on"`
	UpdatedOn *string `json:"updated_on"`
}
