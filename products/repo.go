package products

import (
	"context"
	"fmt"
	pgx "github.com/jackc/pgx/v5"
	"os"
)

type Product struct {
	Id   int
	Name string
}

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(ctx context.Context, connStr string) (*Repository, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	return &Repository{
		conn: conn,
	}, nil
}

func (r Repository) CreateProductsTable(ctx context.Context) error {
	query := "CREATE TABLE IF NOT EXISTS products (id serial, name varchar(255))"
	_, err := r.conn.Exec(ctx, query)
	return err
}

func (r Repository) CreateProduct(ctx context.Context, product Product) (Product, error) {
	query := "INSERT INTO products (name) VALUES ($1) RETURNING id"
	err := r.conn.QueryRow(ctx, query, product.Name).Scan(&product.Id)
	return product, err
}

func (r Repository) GetProductById(ctx context.Context, id int) (*Product, error) {
	var product Product
	query := "SELECT id, name FROM products WHERE id = $1"
	err := r.conn.QueryRow(ctx, query, id).Scan(&product.Id, &product.Name)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
