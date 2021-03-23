package web

import (
	"embed"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/SUSE/console-for-sap-applications/web/api"
	"github.com/SUSE/console-for-sap-applications/web/envronments"
)

//go:embed frontend/assets
var assetsFS embed.FS

//go:embed templates
var templatesFS embed.FS

var layoutData = gin.H{
	"title":     "SUSE Console for SAP Applications",
	"copyright": "© 2019-2020 SUSE, all rights reserved.",
}

type App struct {
	engine *gin.Engine
	host   string
	port   int
}

func NewApp(host string, port int) *App {
	engine := gin.Default()
	engine.HTMLRender = NewLayoutRender(templatesFS, layoutData, "templates/*.tmpl")

	engine.StaticFS("/static", http.FS(assetsFS))
	engine.GET("/", homeHandler)
	engine.GET("/environments", envronments.ListHandler)

	apiGroup := engine.Group("/api")
	{
		apiGroup.GET("/ping", api.PingHandler)
	}

	return &App{engine, host, port}
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
