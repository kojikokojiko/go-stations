package service

import (
	"context"
	"database/sql"
	"log"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	stmt, err := s.db.PrepareContext(ctx, insert)
	if err != nil {
		return nil, err
	}

	res, err := stmt.ExecContext(ctx, subject, description)
	if err != nil {
		return nil, err
	}

	id, _ := res.LastInsertId()
	var todo model.TODO
	todo.ID = id
	err = s.db.QueryRowContext(ctx, confirm, id).Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)

	return &todo, nil
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	var todos []*model.TODO
	todos = []*model.TODO{}

	if prevID == 0 {

		rows, _ := s.db.QueryContext(ctx, read, size)

		for rows.Next() {
			// todo := model.TODO{}
			var todo model.TODO
			// todo = model.TODO{}
			// log.Println(todo.ID)
			err := rows.Scan(&todo.ID, &todo.Subject,
				&todo.Description,
				&todo.UpdatedAt, &todo.UpdatedAt)
			if err != nil {
				log.Println("EEEE")
			} else {
				todos = append(todos, &todo)

			}
		}
	} else {
		log.Println(prevID, size)
		rows, _ := s.db.QueryContext(ctx, readWithID, prevID, size)
		for rows.Next() {
			// todo := model.TODO{}
			var todo model.TODO
			err := rows.Scan(&todo.ID, &todo.Subject,
				&todo.Description,
				&todo.UpdatedAt, &todo.UpdatedAt)

			log.Println(todo.ID)
			log.Println(todo.Subject)
			if err != nil {
				log.Println("EEEE")
			} else {
				todos = append(todos, &todo)
				log.Println(todos[0].ID)
			}
		}

	}

	// rows, err := s.db.QueryContext(ctx, read, size)
	// aarows, err := s.db.QueryContext(ctx, readWithID, prevID, size)

	// .Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)

	return todos, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	stmt, err := s.db.PrepareContext(ctx, update)
	if err != nil {
		return nil, err
	}

	res, err := stmt.ExecContext(ctx, subject, description, id)
	if err != nil {
		return nil, err
	}
	num, err := res.RowsAffected()
	if num == 0 {
		return nil, &model.ErrNotFound{}
	}
	// log.Println(id)

	var todo model.TODO
	todo.ID = id
	err = s.db.QueryRowContext(ctx, confirm, id).Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)

	return &todo, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
