package web

import (
	"embed"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/SUSE/console-for-sap-applications/internal/consul"
)

//go:embed frontend/assets
var assetsFS embed.FS

//go:embed templates
var templatesFS embed.FS

type App struct {
	host string
	port int
	Dependencies
}

type Dependencies struct {
	consul consul.Client
	engine *gin.Engine
}

func DefaultDependencies() Dependencies {
	consulClient, _ := consul.DefaultClient()
	engine := gin.Default()

	return Dependencies{consulClient, engine}
}

// shortcut to use default dependencies
func NewApp(host string, port int) (*App, error) {
	return NewAppWithDeps(host, port, DefaultDependencies())
}

func NewAppWithDeps(host string, port int, deps Dependencies) (*App, error) {
	app := &App{
		Dependencies: deps,
		host:         host,
		port:         port,
	}

	engine := deps.engine
	engine.HTMLRender = NewLayoutRender(templatesFS, "templates/*.tmpl")
	engine.Use(ErrorHandler)
	engine.StaticFS("/static", http.FS(assetsFS))
	engine.GET("/", HomeHandler)
	engine.GET("/environments", NewEnvironmentsListHandler(deps.consul))
	apiGroup := engine.Group("/api")
	{
		apiGroup.GET("/ping", ApiPingHandler)
	}

	return app, nil
}

func (a *App) Start() error {
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", a.host, a.port),
		Handler:        a,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()
}

func (a *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	a.engine.ServeHTTP(w, req)
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	c.Negotiate(500, gin.Negotiate{
		Offered:  []string{gin.MIMEJSON, gin.MIMEHTML, gin.MIMEPlain},
		HTMLName: "error.html.tmpl",
		Data:     c.Errors,
	})

	c.Abort()
}
