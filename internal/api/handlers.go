package api

import (
	"context"
	"net/http"
	"text/template"

	"github.com/yherasymets/go-shorter/pkg/errors"
	"github.com/yherasymets/go-shorter/pkg/utils"
	"github.com/yherasymets/go-shorter/proto"
	"google.golang.org/grpc"
)

type App struct {
	client *grpc.ClientConn
}

func NewApp(client *grpc.ClientConn) *App {
	return &App{client: client}
}

// Handler returns an HTTP handler that routes requests based on the path
func (app *App) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.get)
	mux.HandleFunc("/go-shorter", app.create)
	mux.HandleFunc("/go-shorter/result", app.result)
	return mux
}

// create is an HTTP handler that create short URL
func (app *App) create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/index.html")
	if err != nil {
		errors.TemplateError(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case http.MethodGet:
		if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
			errors.TemplateError(w, err)
		}
	case http.MethodPost:
		http.Redirect(w, r, "/go-shorter/result", http.StatusPermanentRedirect)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// get is an HTTP handler that processes requests to retrieve and redirect to the original URL
func (app *App) get(w http.ResponseWriter, r *http.Request) {
	service := proto.NewShorterClient(app.client)
	// Extract the part of the URL after the slash ("/")
	path := r.URL.Path[1:]
	if path == "" {
		http.NotFound(w, r)
		return
	}
	// Check if the path matches the pattern
	if utils.IsValidPath(path) {
		resp, err := service.Get(context.Background(), &proto.GetRequest{ShortURL: path})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, resp.Result, http.StatusFound)
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
	service := proto.NewShorterClient(app.client)
	res, err := service.Create(r.Context(), &proto.CreateRequest{
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
