package repository

import (
	"app/internal"
	"database/sql"
	"fmt"
)

// NewRepositoryProductStore creates a new repository for products.
func NewRepositoryProductStore(db *sql.DB) (r *RepositoryProductStore) {
	r = &RepositoryProductStore{
		db: db,
	}
	return
}

// RepositoryProductStore is a repository for products.
type RepositoryProductStore struct {
	// st is the underlying store.
	db *sql.DB
}

// FindById finds a product by id.
func (r *RepositoryProductStore) FindById(id int) (p internal.Product, err error) {
	query := `SELECT p.id, p.name, p.quantity, p.code_value, p.is_published, p.expiration, p.price 
              FROM products p WHERE id = ?`

    row := r.db.QueryRow(query, id)

	err = row.Scan(&p.Id, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price)
	if err != nil {
        if err == sql.ErrNoRows {
            return p, internal.ErrRepositoryProductNotFound
        }
        return p, fmt.Errorf("could not retrieve product: %v", err)
    }

	return 
}

// Save saves a product.
func (r *RepositoryProductStore) Save(p *internal.Product) (product *internal.Product, err error) {
	var maxId int
    
    err = r.db.QueryRow("SELECT IFNULL(MAX(id), 0) FROM products").Scan(&maxId)
    if err != nil {
        return nil, fmt.Errorf("could not get max ID: %v", err)
    }

    p.Id = maxId + 1

    query := `INSERT INTO products (id, name, quantity, code_value, is_published, expiration, price) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`

    _, err = r.db.Exec(query, p.Id, p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price)
    if err != nil {
        return nil, fmt.Errorf("could not create product: %v", err)
    }
	product = p
    return
}

// UpdateOrSave updates or saves a product.
func (r *RepositoryProductStore) UpdateOrSave(p *internal.Product) (product *internal.Product, err error) {
	if p.Id == 0 {
		product, err = r.Save(p)
	} else {
		product, err = r.Update(p)
	}

	return
}

// Update updates a product.
func (r *RepositoryProductStore) Update(p *internal.Product) (product *internal.Product, err error) {
	query := `UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?`

	_, err = r.db.Exec(query, p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.Id)
	if err != nil {
		return nil, fmt.Errorf("could not execute update: %v", err)
	}

	product = p

	return
}

// Delete deletes a product.
func (r *RepositoryProductStore) Delete(id int) (err error) {
	query := `DELETE FROM products WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not execute delete: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not retrieve affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return internal.ErrRepositoryProductNotFound
	}

	return
}