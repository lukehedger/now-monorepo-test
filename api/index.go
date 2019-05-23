package root

import (
	"net/http"

	"github.com/lukehedger/now-monorepo-test/api/pprint"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	pprint.Print(w, "<h1>Hello from private Go on Now!</h1>")
}
