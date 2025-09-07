package admingin

import (
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/go-advanced-admin/admin"
)

// Integrator adapts the admin panel to a gin.RouterGroup.
type Integrator struct {
	group *gin.RouterGroup
}

// NewIntegrator creates a new Integrator for the given gin.RouterGroup.
func NewIntegrator(g *gin.RouterGroup) *Integrator {
	return &Integrator{group: g}
}

// HandleRoute registers a route with the given method, path, and handler function.
// It adapts the admin.HandlerFunc to a gin.HandlerFunc.
func (i *Integrator) HandleRoute(method, path string, handler admin.HandlerFunc) {
	i.group.Handle(method, path, func(c *gin.Context) {
		redirectCodes := []int{http.StatusFound, http.StatusMovedPermanently, http.StatusSeeOther}
		code, body := handler(c)
		if slices.Contains(redirectCodes, int(code)) {
			c.Redirect(int(code), body)
			return
		}

		c.Data(int(code), "text/html; charset=utf-8", []byte(body))
	})
}

// ServeAssets serves static assets under the specified prefix using the provided renderer.
func (i *Integrator) ServeAssets(prefix string, renderer admin.TemplateRenderer) {
	i.group.GET(fmt.Sprintf("%s/*filepath", prefix), func(c *gin.Context) {
		fileName := c.Param("filepath")
		fileData, err := renderer.GetAsset(fileName)
		if err != nil {
			c.String(http.StatusNotFound, "File not found")
			return
		}
		contentType := mime.TypeByExtension(filepath.Ext(fileName))

		if contentType == "" {
			contentType = "application/octet-stream"
		}

		c.Data(http.StatusOK, contentType, fileData)
	})
}

// GetQueryParam retrieves the value of a query parameter from the context.
func (i *Integrator) GetQueryParam(ctx interface{}, name string) string {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return ""
	}
	return c.Query(name)
}

// GetPathParam retrieves the value of a path parameter from the context.
func (i *Integrator) GetPathParam(ctx interface{}, name string) string {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return ""
	}
	return c.Param(name)
}

// GetRequestMethod retrieves the HTTP method of the request from the context.
func (i *Integrator) GetRequestMethod(ctx interface{}) string {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return ""
	}
	return c.Request.Method
}

// GetFormData retrieves form data from the context.
func (i *Integrator) GetFormData(ctx interface{}) map[string][]string {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return nil
	}
	if err := c.Request.ParseForm(); err != nil {
		return nil
	}
	return c.Request.Form
}

// HandleJSONRoute registers a route that returns JSON responses.
func (i *Integrator) HandleJSONRoute(method, path string, handler admin.JSONHandlerFunc) {
	i.group.Handle(method, path, func(c *gin.Context) {
		err := handler(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, admin.NewErrorResponse([]string{err.Error()}))
		}
	})
}

// SetJSONResponse sets a JSON response with the given status code and data.
func (i *Integrator) SetJSONResponse(ctx interface{}, statusCode int, data interface{}) error {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return fmt.Errorf("invalid context type")
	}
	c.JSON(statusCode, data)
	return nil
}

// GetJSONBody retrieves JSON body data from the request context.
func (i *Integrator) GetJSONBody(ctx interface{}) (map[string]interface{}, error) {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return nil, fmt.Errorf("invalid context type")
	}

	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		return nil, err
	}
	return body, nil
}
