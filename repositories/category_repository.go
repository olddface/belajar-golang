package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetAll() ([]models.Category, error) {
	query := "SELECT id, name, description FROM categories"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]models.Category, 0)
	for rows.Next() {
		var c models.Category
		err := rows.Scan(&c.ID, &c.Name, &c.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (repo *CategoryRepository) GetCategoryWithProductsByName(name string) (*models.CategoryDetail, error) {
	query := `
		SELECT c.id, c.name, c.description, p.id, p.name, p.price, p.stock, p.category_id
		FROM categories c
		LEFT JOIN products p ON c.id = p.category_id
		WHERE c.name = $1`

	rows, err := repo.db.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categoryDetail models.CategoryDetail
	var products []models.Product
	var foundCategory bool

	for rows.Next() {
		var cID int
		var cName string
		var cDescription string
		var pID sql.NullInt64
		var pName sql.NullString
		var pPrice sql.NullInt64
		var pStock sql.NullInt64
		var pCategoryID sql.NullInt64

		err := rows.Scan(&cID, &cName, &cDescription, &pID, &pName, &pPrice, &pStock, &pCategoryID)
		if err != nil {
			return nil, err
		}

		if !foundCategory {
			categoryDetail.ID = cID
			categoryDetail.Name = cName
			categoryDetail.Description = cDescription
			foundCategory = true
		}

		if pID.Valid {
			products = append(products, models.Product{
				ID:         int(pID.Int64),
				Name:       pName.String,
				Price:      int(pPrice.Int64),
				Stock:      int(pStock.Int64),
				CategoryID: int(pCategoryID.Int64),
			})
		}
	}

	if !foundCategory {
		return nil, errors.New("category not found")
	}

	categoryDetail.Products = products
	return &categoryDetail, nil
}

func (repo *CategoryRepository) Create(category *models.Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"
	err := repo.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
	return err
}

func (repo *CategoryRepository) Update(category *models.Category) error {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"
	result, err := repo.db.Exec(query, category.Name, category.Description, category.ID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("category not found")
	}
	return nil
}

func (repo *CategoryRepository) Delete(id int) error {
	// Note: This might fail if there are foreign keys. Assuming database handles cascade or user handles it.
	query := "DELETE FROM categories WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("category not found")
	}
	return nil
}

func (repo *CategoryRepository) GetByID(id int) (*models.Category, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"
	var c models.Category
	err := repo.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Description)
	if err == sql.ErrNoRows {
		return nil, errors.New("category not found")
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}
