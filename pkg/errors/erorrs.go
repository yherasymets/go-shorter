package errors

import (
	"net/http"
	"text/template"

	"github.com/sirupsen/logrus"
)

func TemplateError(w http.ResponseWriter, err error) {
	logrus.Errorf("failed to parse template: %v", err)
	http.Error(w, "internal server error", http.StatusInternalServerError)
}

func RPCError(w http.ResponseWriter, t *template.Template, err error) {
	logrus.Errorf("failed to create URL: %v", err)
	if err := t.ExecuteTemplate(w, "errmsg.html", err); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
