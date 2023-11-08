package api

import (
	"context"
	"net/http"
	"regexp"
	"text/template"

	"github.com/sirupsen/logrus"
	"github.com/yherasymets/go-shorter/proto"
	"google.golang.org/grpc"
)

// App struct
type ClientApp struct {
	Conn *grpc.ClientConn
}

func (app *ClientApp) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.get)
	mux.HandleFunc("/go-shorter", app.create)
	mux.HandleFunc("/go-shorter/result", app.result)
	return mux
}

// TODO: error handling, logging

func (app *ClientApp) create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/index.html")
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
		http.NotFound(w, r)
	}
}

func (app *ClientApp) get(w http.ResponseWriter, r *http.Request) {
	service := proto.NewShorterClient(app.Conn)
	// Extract the part of the URL after the slash ("/")
	path := r.URL.Path[1:]
	// Define a regular expression pattern to match alphanumeric characters
	pattern := `^[a-zA-Z0-9]+$`
	// Compile the regular expression
	re, err := regexp.Compile(pattern)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	// Check if the path matches the pattern
	if re.MatchString(path) {
		resp, err := service.Get(context.Background(), &proto.UrlRequest{
			FullURL: path,
		})
		if err != nil {
			logrus.Info(err)
		}
		http.Redirect(w, r, resp.Result, http.StatusMovedPermanently)
		return
	}
	http.NotFound(w, r)
}

func (app *ClientApp) result(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/result.html")
	if err != nil {
		logrus.Errorf("failed to parse template: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	service := proto.NewShorterClient(app.Conn)
	res, err := service.Create(r.Context(), &proto.UrlRequest{
		FullURL: r.PostFormValue("original-link"),
	})
	if err != nil {
		logrus.Errorf("failed to create URL: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err := t.ExecuteTemplate(w, "result.html", res); err != nil {
		logrus.Errorf("failed to execute template: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
