package pprint

import (
	"fmt"
	"net/http"
)

func Print(w http.ResponseWriter, str string) {
	fmt.Fprintf(w, str)
}
