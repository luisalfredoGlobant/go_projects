package entity

// Product represents an product.
type Product struct {
	ID        		uint32    `json:"id"`
	Name      		string    `json:"name"`
	SupplierID 		uint32 	  `json:"supplier_id"`
	CategoryID 		uint32 	  `json:"category_id"`
	UnitsInStock	uint32	  `json:"units_in_stock"`
	UnitPrice		float64   `json:"unit_price"`
	Discontinued    bool	  `json:"discontinued"`
}