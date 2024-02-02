package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO category (id, name, description) VALUES ($1, $2, $3)", id, name, description)
	if err != nil {
		return Category{}, err
	}

	return Category{db: c.db, ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
    err := c.db.QueryRow("SELECT c.id, c.name, c.description FROM category c JOIN courses co ON c.id = co.category_id WHERE co.id = $1", courseID).Scan(&c.ID, &c.Name, &c.Description)
	if err != nil {
		return Category{}, err
	}

	return Category{db: c.db, ID: c.ID, Name: c.Name, Description: c.Description}, nil
}

func (c *Category) FindByID(id string) (Category, error) {
    err := c.db.QueryRow("SELECT id, name, description FROM category WHERE id = $1", id).Scan(&c.ID, &c.Name, &c.Description)
    if err != nil {
        return Category{}, err
    }

    return Category{db: c.db, ID: c.ID, Name: c.Name, Description: c.Description}, nil
}
