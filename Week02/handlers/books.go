// solmyr

package handlers

import (
	"fmt"
	"net/http"
)

import (
	"Week02/dao/mysql"
)

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	bookDao := mysql.BookMysql{}
	book, err := bookDao.GetByName("foo")

	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Fprintf(w, book.Name)
}