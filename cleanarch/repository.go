package cleanarch

import "database/sql"

/////////////////////////
/// SQL repository ////
/////////////////////////

type SQLHandler struct {
	Conn *sql.DB
}

func NewSQLHandler(conn *sql.DB) *SQLHandler {
	return &SQLHandler{Conn: conn}
}

func (tr *SQLHandler) fetch(query string, args ...interface{}) ([]*Train, error) {
	rows, err := tr.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*Train, 0)
	for rows.Next() {
		t := new(Train)
		err := rows.Scan(&t.Id, &t.Nama)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (tr *SQLHandler) GetById(id int) (*Train, error) {
	query := `select id, nama_diklat from riwayat_diklat where id=?`
	list, err := tr.fetch(query, id)
	if err != nil {
		return nil, err
	}
	train := &Train{}
	if len(list) > 0 {
		train = list[0]
	} else {
		return nil, err
	}
	return train, nil
}
