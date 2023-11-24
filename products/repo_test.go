package products

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	testcontainers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestProductRepository(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:16-alpine"),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)

	repo, err := NewRepository(ctx, connStr)
	assert.NoError(t, err)

	err = repo.CreateProductsTable(ctx)
	assert.NoError(t, err)

	c, err := repo.CreateProduct(ctx, Product{
		Name: "iPhone 14",
	})
	assert.NoError(t, err)
	assert.NotNil(t, c)

	product, err := repo.GetProductById(ctx, c.Id)
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "iPhone 14", product.Name)
}
