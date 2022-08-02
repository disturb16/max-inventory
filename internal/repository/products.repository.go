package repository

import (
	context "context"

	entity "github.com/disturb/max-inventory/internal/entity"
)

const (
	qryInsertProduct = `
		insert into PRODUCTS (name, description, price, created_by) values (?, ?, ?, ?);`

	qryGetAllProducts = `
		select
			id
			name,
			description
			price
			created_by
		from PRODUCTS;`

	qryGetProductByID = `
		select
			id
			name,
			description
			price
			created_by
		from PRODUCTS
		where id = ?;`
)

func (r *repo) SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error {
	_, err := r.db.ExecContext(ctx, qryInsertProduct, name, description, price, createdBy)
	return err
}

func (r *repo) GetProducts(ctx context.Context) ([]entity.Product, error) {
	pp := []entity.Product{}

	err := r.db.SelectContext(ctx, &pp, qryGetAllProducts)
	if err != nil {
		return nil, err
	}

	return pp, nil
}

func (r *repo) GetProduct(ctx context.Context, id int64) (*entity.Product, error) {
	p := &entity.Product{}

	err := r.db.GetContext(ctx, p, qryGetProductByID, id)
	if err != nil {
		return nil, err
	}

	return p, nil
}
