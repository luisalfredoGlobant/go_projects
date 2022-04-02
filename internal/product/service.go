package product

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/globant/crud_project/internal/entity"
	"github.com/globant/crud_project/pkg/log"
)

type Service interface {
	Get(ctx context.Context, id uint32) (Product, error)
	Query(ctx context.Context, offset, limit int) ([]Product, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateProductRequest) (Product, error)
	Update(ctx context.Context, id uint32, input UpdateProductRequest) (Product, error)
	Delete(ctx context.Context, id uint32) (Product, error)
}

// Product represents the data about an product.
type Product struct {
	entity.Product
}

// CreateProductRequest represents an product creation request.
type CreateProductRequest struct {
	Name 			string 	  `json:"name"`
	SupplierID 		uint32 	  `json:"supplier_id"`
	CategoryID 		uint32 	  `json:"category_id"`
	UnitsInStock	uint32	  `json:"units_in_stock"`
	UnitPrice		float64   `json:"unit_price"`
	Discontinued    bool	  `json:"discontinued"`
}

// Validate validates the CreateAlbumRequest fields.
func (m CreateProductRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateProductRequest represents an product update request.
type UpdateProductRequest struct {
	Name 			string 	  `json:"name"`
	SupplierID 		uint32 	  `json:"supplier_id"`
	CategoryID 		uint32 	  `json:"category_id"`
	UnitsInStock	uint32	  `json:"units_in_stock"`
	UnitPrice		float64   `json:"unit_price"`
	Discontinued    bool	  `json:"discontinued"`
}

// Validate validates the UpdateProductRequest fields.
func (m UpdateProductRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new product service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the product with the specified the product ID.
func (s service) Get(ctx context.Context, id uint32) (Product, error) {
	album, err := s.repo.Get(ctx, id)
	if err != nil {
		return Product{}, err
	}
	return Product{album}, nil
}

func (s service) Create(ctx context.Context, req CreateProductRequest) (Product, error) {
	if err := req.Validate(); err != nil {
		return Product{}, err
	}
	id := entity.GenerateUint32ID()

	s.repo.Create(ctx, entity.Product{
		ID:				id,
		Name:      		req.Name,
	    SupplierID: 	req.SupplierID,
		CategoryID: 	req.CategoryID,
		UnitsInStock: 	req.UnitsInStock,
		UnitPrice: 		req.UnitPrice,
		Discontinued: 	req.Discontinued,
	})

	return s.Get(ctx, uint32(id))
}

// Update updates the album with the specified ID.
func (s service) Update(ctx context.Context, id uint32, req UpdateProductRequest) (Product, error) {
	if err := req.Validate(); err != nil {
		return Product{}, err
	}

	product, err := s.Get(ctx, id)
	if err != nil {
		return product, err
	}
	product.Name = req.Name
	product.SupplierID = req.SupplierID
	product.CategoryID = req.CategoryID
	product.UnitsInStock = req.UnitsInStock
	product.UnitPrice =	req.UnitPrice
	product.Discontinued = req.Discontinued

	if err := s.repo.Update(ctx, product.Product); err != nil {
		return product, err
	}
	return product, nil
}

// Delete deletes the album with the specified ID.
func (s service) Delete(ctx context.Context, id uint32) (Product, error) {
	product, err := s.Get(ctx, id)
	if err != nil {
		return Product{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Product{}, err
	}
	return product, nil
}

// Count returns the number of albums.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the albums with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Product, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Product{}
	for _, item := range items {
		result = append(result, Product{item})
	}
	return result, nil
}