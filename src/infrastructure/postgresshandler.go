package infrastructure

import (
	"database/sql"
	"interfaces"
)

type PostgressHandler struct {
	Conn *sql.DB
}

func (handler *PostgressHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *PostgressHandler) Query(statement string) (interfaces.Row, error) {
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		return nil, err
	}

	row := new(PostgressRow)
	row.Rows = rows
	return row, nil
}

type PostgressRow struct {
	Rows *sql.Rows
}

func (row *PostgressRow) Next() bool {
	return ow.Rows.Next()
}

func (row *PostgressRow) Scan(dest ...interface{}) {
	row.Rows.Scan(dest...)
}

func NewPostgressHandler(dbFile string) (*PostgressHandler, error) {
	conn, err := sql.Open("postgress", dbFile)
	if err != nil {
		return nil, err
	}

	postgressHandler := new(PostgressHandler)
	postgressHandler.Conn = conn
	return postgressHandler, nil
}
