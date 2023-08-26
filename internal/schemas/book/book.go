package book

type CreateBookSchema struct {
	Title         string   `json:"title" validate:"required"`
	Description   string   `json:"description" validate:"required"`
	Price         float64  `json:"price" validate:"required"`
	StockQuantity int      `json:"stock_quantity" validate:"required"`
	Genre         string   `json:"genre,omitempty"`
	FullTextURL   string   `json:"full_text_url,omitempty"`
	CoverImages   []string `json:"cover_images,omitempty"`
	SampleURL     string   `json:"sample_url,omitempty"`
	SellerID      int      `json:"seller_id" validate:"required"`
}

type UpdateBookSchema struct {
	Title         string   `json:"title,omitempty"`
	Description   string   `json:"description,omitempty"`
	Price         float64  `json:"price,omitempty"`
	StockQuantity int      `json:"stock_quantity,omitempty"`
	Genre         string   `json:"genre,omitempty"`
	FullTextURL   string   `json:"full_text_url,omitempty"`
	CoverImages   []string `json:"cover_images,omitempty"`
	SampleURL     string   `json:"sample_url,omitempty"`
	SellerID      int      `json:"seller_id,omitempty"`
}
