package repository

import "database/sql"

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) FetchProductByID(id int64) (Product, error) {
	//TODO: You must implement this function fot fetch product by id
	var product Product

	row := p.db.QueryRow(`SELECT * FROM products WHERE id = ?`, id)

	err := row.Scan(&product.ID, &product.ProductName, &product.Category, &product.Price, &product.Quantity)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) FetchProductByName(productName string) (Product, error) {
	// TODO: You must implement this function for fetch product by name
	var product Product

	row := p.db.QueryRow(`SELECT * FROM products WHERE product_name = ?`, productName)

	err := row.Scan(&product.ID, &product.ProductName, &product.Category, &product.Price, &product.Quantity)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) FetchProducts() ([]Product, error) {
	// TODO: You must implement this function for fetch all products
	var products []Product

	rows, err := p.db.Query(`SELECT * FROM products`)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.ProductName, &product.Category, &product.Price, &product.Quantity)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	return products, nil
}
