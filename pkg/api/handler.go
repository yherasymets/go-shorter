package api

import (
	"context"
	"html/template"
	"net/http"
	"regexp"

	"github.com/sirupsen/logrus"
	"github.com/yherasymets/go-shorter/proto"
	"google.golang.org/grpc"
)

// App struct
type ClientApp struct {
	HTMLDir   string
	StaticDir string
	Conn      *grpc.ClientConn
}

func (app *ClientApp) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Get)
	mux.HandleFunc("/go-shorter", app.Create)
	return mux

}

func (app *ClientApp) Create(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseGlob("frontend/index.html")
	service := proto.NewShorterClient(app.Conn)
	if r.Method == http.MethodPost {
		res, err := service.Create(r.Context(), &proto.UrlRequest{
			FullURL: r.PostFormValue("original-link"),
		})
		if err != nil {
			logrus.Infof("%v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		t.ExecuteTemplate(w, "index.html", res)
		return
	}
	t.ExecuteTemplate(w, "index.html", nil)
}

func (app *ClientApp) Get(w http.ResponseWriter, r *http.Request) {
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
			logrus.Infof("%v", err)
		}
		http.Redirect(w, r, resp.Result, http.StatusMovedPermanently)
		return
	}
	http.NotFound(w, r)
}
