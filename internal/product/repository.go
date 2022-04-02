package product

import (
	"context"
	"github.com/globant/crud_project/internal/entity"
	"github.com/globant/crud_project/pkg/dbcontext"
	"github.com/globant/crud_project/pkg/log"
)

type Repository interface {
	// Get returns the album with the specified album ID.
	Get(ctx context.Context, id uint32) (entity.Product, error)
	// Count returns the number of albums.
	Count(ctx context.Context) (int, error)
	// Query returns the list of albums with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Product, error)
	// Create saves a new album in the storage.
	Create(ctx context.Context, album entity.Product) error
	// Update updates the album with given ID in the storage.
	Update(ctx context.Context, album entity.Product) error
	// Delete removes the album with given ID from the storage.
	Delete(ctx context.Context, id uint32) error
}

// repository persists albums in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the product with the specified ID from the database.
func (r repository) Get(ctx context.Context, id uint32) (entity.Product, error) {
	var product entity.Product
	err := r.db.With(ctx).Select().Model(id, &product)
	return product, err
}

func (r repository) Create(ctx context.Context, product entity.Product) error {
	return r.db.With(ctx).Model(&product).Insert()
}

func (r repository) Update(ctx context.Context, product entity.Product) error {
	return r.db.With(ctx).Model(&product).Update()
}

func (r repository) Delete(ctx context.Context, id uint32) error {
	product, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&product).Delete()
}

func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("product").Row(&count)
	return count, err
}

func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&products)
	return products, err
}