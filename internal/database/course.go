package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name string, description string, categoryID string) (Course, error) {
	id := uuid.New().String()
	stmt, err := c.db.Prepare("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)")
	_, err = stmt.Exec(id, name, description, categoryID)
	if err != nil {
		return Course{}, err
	}

	return Course{db: c.db, ID: id, Name: name, Description: description, CategoryID: categoryID}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	stmt, err := c.db.Prepare("SELECT id, name, description, category_id FROM courses WHERE category_id = $1")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}
