package views

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	page := strings.TrimPrefix(r.URL.Path, "/")
	page = strings.TrimSuffix(page, "/")

	var context map[string]interface{}

	switch page {
	case "index":
		context = IndexHandler(w, r)
	default:
		page = "index"
		context = IndexHandler(w, r)
	}
	context["Page"] = page
	Render(w, r, page, context)
}

func Render(w http.ResponseWriter, r *http.Request, page string, context map[string]interface{}) {
	uadmin.RenderMultiHTML(w, r, []string{fmt.Sprintf("./templates/%s.html", page)}, context)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")

	return map[string]interface{}{
		"Title": "HOME",
	}
}
