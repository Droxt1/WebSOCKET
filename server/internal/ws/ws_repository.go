package ws

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}
type RoomRepository interface {
	CreateRoom(ctx context.Context, room *Room) (*Room, error)
	GetRoomByID(ctx context.Context, roomID int64) (*Room, error)
	// Add other room-related methods if needed
}

type roomRepository struct {
	db DBTX // Your database connection
}

func NewRoomRepository(db DBTX) RoomRepository {
	return &roomRepository{
		db: db,
	}
}

func (r *roomRepository) CreateRoom(ctx context.Context, room *Room) (*Room, error) {
	// Implement the logic to insert a new room into the database
	// Use the `db` connection to execute the SQL query

	// Example:
	query := "INSERT INTO rooms(name) VALUES ($1) RETURNING id"
	var roomID int64
	err := r.db.QueryRowContext(ctx, query, room.Name).Scan(&roomID)
	if err != nil {
		return nil, err
	}

	room.ID = strconv.FormatInt(roomID, 10)
	return room, nil
}

func (r *roomRepository) GetRoomByID(ctx context.Context, roomID int64) (*Room, error) {
	// Implement the logic to retrieve a room by its ID from the database
	// Use the `db` connection to execute the SQL query

	// Example:
	query := "SELECT id, name FROM rooms WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, roomID)

	room := &Room{}
	err := row.Scan(&room.ID, &room.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("room not found")
		}
		return nil, err
	}

	return room, nil
}

// Add other room-related methods if needed
