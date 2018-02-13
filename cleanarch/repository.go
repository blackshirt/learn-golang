/////////////////////////
///// repository.go /////
/////////////////////////

package cleanarch

import (
	"database/sql"
	"fmt"
)

/////////////////////////
/// SQL repository //////
/////////////////////////

type SQLHandler struct {
	Conn *sql.DB
}

// Sql handle Constructor
func NewSQLHandler(conn *sql.DB) *SQLHandler {
	return &SQLHandler{Conn: conn}
}

// helper function
func (sqh *SQLHandler) fetch(query string, args ...interface{}) ([]*Train, error) {
	rows, err := sqh.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*Train, 0)
	for rows.Next() {
		t := new(Train)
		err := rows.Scan(&t.Id, &t.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (sqh *SQLHandler) GetById(id int) (*Train, error) {
	query := fmt.Sprintf("SELECT id, nama_diklat FROM riwayat_diklat WHERE id=%d", id)
	// query := `select id, nama_diklat from riwayat_diklat where id=?`
	/*
		res, err := sqh.fetch(query, id)
		if err != nil {
			return nil, err
		}
		train := &Train{}
		if len(rest) > 0 {
			train = rest[0]
		} else {
			return nil, err
		}
	*/
	train := new(Train)
	err := sqh.Conn.QueryRow(query).Scan(&train.Id, &train.Name)
	if err != nil {
		return nil, err
	}
	return train, nil
}

func (sqh *SQLHandler) Fetch(start, limit int) ([]*Train, error) {
	query := fmt.Sprintf("SELECT * from id, nama_diklat FROM riwayat_diklat ORDER BY id DESC LIMIT %d OFFSET %d  ", limit, start)
	res, err := sqh.fetch(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (sqh *SQLHandler) Posts(tr *Train) error {

	query := "insert into train_list(name, description, begin, end, location, address, city, jp, category, penyelenggara) values(?,?,?,?,?,?,?,?,?,?)"
	_, err := sqh.Conn.Exec(query,
		&tr.Name,
		&tr.Description,
		&tr.Begin,
		&tr.End,
		&tr.Location,
		&tr.Address,
		&tr.City,
		&tr.Jp,
		&tr.Category,
		&tr.Penyelenggara)
	if err != nil {
		return err
	}
	return nil
}
