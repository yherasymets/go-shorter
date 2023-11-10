package api

import (
	"context"
	"net/http"
	"regexp"
	"text/template"

	"github.com/sirupsen/logrus"
	"github.com/yherasymets/go-shorter/pkg/errors"
	"github.com/yherasymets/go-shorter/proto"
	"google.golang.org/grpc"
)

// App struct
type App struct {
	Conn *grpc.ClientConn
}

func (app *App) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.get)
	mux.HandleFunc("/go-shorter", app.create)
	mux.HandleFunc("/go-shorter/result", app.result)
	return mux
}

// TODO: error handling, logging

func (app *App) create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/index.html")
	if err != nil {
		logrus.Errorf("failed to parse template: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case http.MethodGet:
		if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
			logrus.Errorf("failed to execute template: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	case http.MethodPost:
		http.Redirect(w, r, "/go-shorter/result", http.StatusPermanentRedirect)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// get is an HTTP handler that processes requests to retrieve and redirect to the original URL
func (app *App) get(w http.ResponseWriter, r *http.Request) {
	service := proto.NewShorterClient(app.Conn)
	// Extract the part of the URL after the slash ("/")
	path := r.URL.Path[1:]
	if path == "" {
		http.NotFound(w, r)
		return
	}
	// Define a regular expression pattern to match alphanumeric characters
	pattern := `^[a-zA-Z0-9]+$`
	re, err := regexp.Compile(pattern)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Check if the path matches the pattern
	if re.MatchString(path) {
		resp, err := service.Get(context.Background(), &proto.UrlRequest{FullURL: path})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, resp.Result, http.StatusMovedPermanently)
		return
	}
	http.NotFound(w, r)
}

// result is an HTTP handler that processes a form submission to create a shortened URL
func (app *App) result(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("web/*.html")
	if err != nil {
		errors.TemplateError(w, err)
		return
	}
	service := proto.NewShorterClient(app.Conn)
	res, err := service.Create(r.Context(), &proto.UrlRequest{
		FullURL: r.PostFormValue("original-link"),
	})
	if err != nil {
		errors.RPCError(w, t, err)
		return
	}
	if err := t.ExecuteTemplate(w, "result.html", res); err != nil {
		errors.TemplateError(w, err)
	}
}
