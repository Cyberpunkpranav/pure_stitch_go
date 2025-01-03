package models

type Arrivals struct {
	Id             int      `json:"id"`
	ProductId      int      `json:"product_id"`
	ProductTypeId  int      `json:"product_type_id"`
	DateOfLaunch   string   `json:"date_of_lunch"`
	File           *string  `json:"file"`
	OriginalHeight int      `json:"original_height"`
	OriginalWidth  int      `json:"original_width"`
	Format         string   `json:"format"`
	BroadcastHours int      `json:"broacast_hours"`
	Active         bool     `json:"active"`
	CreatedOn      *string  `json:"created_on"`
	UpdatedOn      *string  `json:"updated_on"`
	Product        Products `json:"product"`
}
