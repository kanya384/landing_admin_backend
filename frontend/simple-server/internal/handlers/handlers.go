package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Registry()
}

type handlers struct {
	content map[string]interface{}
	appHost string
	mu      sync.RWMutex
	router  *gin.Engine
}

func NewHandlers(router *gin.Engine, appHost string) Handlers {
	return &handlers{
		content: map[string]interface{}{},
		appHost: appHost,
		router:  router,
		mu:      sync.RWMutex{},
	}
}

func (h *handlers) Registry() {
	h.router.LoadHTMLFiles("./build/index.html")
	h.router.Use(gzip.Gzip(gzip.DefaultCompression))

	h.router.GET("/", h.HandleSite)
	h.router.GET("/update/", h.UpdateContent)
	h.router.Use(static.Serve("/", static.LocalFile("./build", false)))
	h.router.NoRoute(h.HandleSite)
}

func (h *handlers) HandleSite(c *gin.Context) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.content) == 0 {
		err := h.getContent()
		fmt.Println(err)
	}
	jsonContent, _ := json.Marshal(h.content)
	context := map[string]interface{}{
		"content": string(jsonContent),
	}

	c.HTML(http.StatusOK, "index.html", context)
}

func (h *handlers) UpdateContent(c *gin.Context) {
	h.mu.Lock()
	h.getContent()
	h.mu.Unlock()
}

func (h *handlers) getContent() (err error) {
	content := map[string]interface{}{}
	r, err := http.Get("http://" + h.appHost + "/api/content")
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &content)
	if err != nil {
		return
	}

	h.content = content

	return
}
