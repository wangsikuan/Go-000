// solmyr

package mysql

import (
	"Week02/models"
)

import (
	"github.com/pkg/errors"
)
type BookMysql struct {

}

func (bookDao BookMysql) GetByName(name string) (models.Book, error) {
	query := "SELECT id, name FROM books where name = " + name
	db := get()
	defer db.Close()

	var book models.Book
	stmt, err := db.Prepare(query)
	if err != nil {
		return book, errors.Wrap(err, "db prepare error")
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return book, errors.Wrap(err, "db query error")
	}

	for rows.Next() {
		var row models.Book
		err := rows.Scan(&row.Name, &row.Author, &row.Price)
		return row, errors.Wrap(err, "db scann error")
	}
	return book, err
}