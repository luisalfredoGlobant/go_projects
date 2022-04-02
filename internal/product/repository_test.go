package product

import (
	"context"
	"database/sql"
	"github.com/globant/crud_project/internal/entity"
	"github.com/globant/crud_project/internal/test"
	"github.com/globant/crud_project/pkg/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository(t *testing.T) {
	logger, _ := log.NewForTest()
	db := test.DB(t)
	test.ResetTables(t, db, "product")
	repo := NewRepository(db, logger)

	ctx := context.Background()
	// initial count
	count, err := repo.Count(ctx)
	assert.Nil(t, err)

	// create
	err = repo.Create(ctx, entity.Product{
		ID:				5,
		Name:      		"product1",
	    SupplierID: 	20,
		CategoryID: 	19,
		UnitsInStock: 	2,
		UnitPrice: 		12.09,
		Discontinued: 	false,
	})
	assert.Nil(t, err)
	count2, _ := repo.Count(ctx)
	assert.Equal(t, 1, count2-count)

	// get
	product, err := repo.Get(ctx, 5)
	assert.Nil(t, err)
	assert.Equal(t, "product1", product.Name)
	_, err = repo.Get(ctx, 0)
	assert.Equal(t, sql.ErrNoRows, err)

	// update
	err = repo.Update(ctx, entity.Product{
		ID:        		5,
		Name:      		"product1 updated",
		SupplierID: 	20,
		CategoryID: 	19,
		UnitsInStock: 	2,
		UnitPrice: 		12.09,
		Discontinued: 	false,
	})
	assert.Nil(t, err)
	product, _ = repo.Get(ctx, 5)
	assert.Equal(t, "product1 updated", product.Name)

	// query
	products, err := repo.Query(ctx, 0, count2)
	assert.Nil(t, err)
	assert.Equal(t, count2, len(products))

	// delete
	err = repo.Delete(ctx, 5)
	assert.Nil(t, err)
	_, err = repo.Get(ctx, 5)
	assert.Equal(t, sql.ErrNoRows, err)
	err = repo.Delete(ctx, 5)
	assert.Equal(t, sql.ErrNoRows, err)
}
