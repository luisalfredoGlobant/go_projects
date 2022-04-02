package product

import (
	"context"
	"database/sql"
	"errors"
	"github.com/globant/crud_project/internal/entity"
	"github.com/globant/crud_project/pkg/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

var errCRUD = errors.New("error crud")

func TestCreateProductRequest_Validate(t *testing.T) {
	tests := []struct {
		name      string
		model     CreateProductRequest
		wantError bool
	}{
		{"success", CreateProductRequest{Name: "test"}, false},
		{"required", CreateProductRequest{Name: ""}, true},
		{"too long", CreateProductRequest{Name: "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.Validate()
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func TestUpdateProductRequest_Validate(t *testing.T) {
	tests := []struct {
		name      string
		model     UpdateProductRequest
		wantError bool
	}{
		{"success", UpdateProductRequest{Name: "test"}, false},
		{"required", UpdateProductRequest{Name: ""}, true},
		{"too long", UpdateProductRequest{Name: "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.Validate()
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func Test_service_CRUD(t *testing.T) {
	logger, _ := log.NewForTest()
	s := NewService(&mockRepository{}, logger)

	ctx := context.Background()

	// initial count
	count, _ := s.Count(ctx)
	assert.Equal(t, 0, count)

	// successful creation
	product, err := s.Create(ctx, CreateProductRequest{Name: "test"})
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
	id := product.ID
	assert.Equal(t, "test", product.Name)
	count, _ = s.Count(ctx)
	assert.Equal(t, 1, count)

	// validation error in creation
	_, err = s.Create(ctx, CreateProductRequest{Name: ""})
	assert.NotNil(t, err)
	count, _ = s.Count(ctx)
	assert.Equal(t, 1, count)

	_, _ = s.Create(ctx, CreateProductRequest{Name: "test2"})

	// update
	product, err = s.Update(ctx, id, UpdateProductRequest{Name: "test updated"})
	assert.Nil(t, err)
	assert.Equal(t, "test updated", product.Name)
	_, err = s.Update(ctx, id, UpdateProductRequest{Name: "test updated"})
	assert.NotNil(t, err)

	// validation error in update
	_, err = s.Update(ctx, id, UpdateProductRequest{Name: ""})
	assert.NotNil(t, err)
	count, _ = s.Count(ctx)
	assert.Equal(t, 1, count)

	// unexpected error in update
	_, err = s.Update(ctx, id, UpdateProductRequest{Name: "error"})
	assert.Equal(t, errCRUD, err)
	count, _ = s.Count(ctx)
	assert.Equal(t, 2, count)

	// get
	_, err = s.Get(ctx, id)
	assert.NotNil(t, err)
	product, err = s.Get(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, "test updated", product.Name)
	assert.Equal(t, id, product.ID)

	// query
	products, _ := s.Query(ctx, 0, 0)
	assert.Equal(t, 2, len(products))

	// delete
	_, err = s.Delete(ctx, 0)
	assert.NotNil(t, err)
	product, err = s.Delete(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, id, product.ID)
	count, _ = s.Count(ctx)
	assert.Equal(t, 1, count)
}

type mockRepository struct {
	items []entity.Product
}

func (m mockRepository) Get(ctx context.Context, id uint32) (entity.Product, error) {
	for _, item := range m.items {
		if item.ID == id {
			return item, nil
		}
	}
	return entity.Product{}, sql.ErrNoRows
}

func (m mockRepository) Count(ctx context.Context) (int, error) {
	return len(m.items), nil
}

func (m mockRepository) Query(ctx context.Context, offset, limit int) ([]entity.Product, error) {
	return m.items, nil
}

func (m *mockRepository) Create(ctx context.Context, product entity.Product) error {
	if product.Name == "error" {
		return errCRUD
	}
	m.items = append(m.items, product)
	return nil
}

func (m *mockRepository) Update(ctx context.Context, product entity.Product) error {
	if product.Name == "error" {
		return errCRUD
	}
	for i, item := range m.items {
		if item.ID == product.ID {
			m.items[i] = product
			break
		}
	}
	return nil
}

func (m *mockRepository) Delete(ctx context.Context, id uint32) error {
	for i, item := range m.items {
		if item.ID == id {
			m.items[i] = m.items[len(m.items)-1]
			m.items = m.items[:len(m.items)-1]
			break
		}
	}
	return nil
}