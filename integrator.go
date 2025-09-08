package admingin

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    admin "github.com/ovnicraft/go-advanced-admin"
)

// Integrator adapts the admin panel to a gin.RouterGroup.
type Integrator struct {
    group *gin.RouterGroup
}

// NewIntegrator creates a new Integrator for the given gin.RouterGroup.
func NewIntegrator(group *gin.RouterGroup) *Integrator {
    return &Integrator{group: group}
}

// HandleRoute registers a route with the given method, path, and handler function.
// It adapts the admin.HandlerFunc to a gin.HandlerFunc.
func (i *Integrator) HandleRoute(method, path string, handler admin.HandlerFunc) {
    i.group.Handle(method, path, func(c *gin.Context) {
        code, body := handler(c)
        switch int(code) {
        case http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther,
            http.StatusTemporaryRedirect, http.StatusPermanentRedirect:
            c.Redirect(int(code), body)
            return
        }
        c.Data(int(code), "text/html; charset=utf-8", []byte(body))
    })
}

// HandleJSONRoute registers a route that returns JSON responses.
func (i *Integrator) HandleJSONRoute(method, path string, handler admin.JSONHandlerFunc) {
    i.group.Handle(method, path, func(c *gin.Context) {
        if err := handler(c); err != nil {
            _ = i.SetJSONResponse(c, http.StatusInternalServerError, admin.NewErrorResponse([]string{err.Error()}))
        }
    })
}

// ServeAssets serves static assets under the specified prefix using the provided renderer.
func (i *Integrator) ServeAssets(prefix string, renderer admin.TemplateRenderer) {
    p := "/" + strings.TrimPrefix(prefix, "/")
    grp := i.group.Group(p)

    grp.HEAD("/*filepath", func(c *gin.Context) {
        path := strings.TrimPrefix(c.Param("filepath"), "/")
        if _, err := renderer.GetAsset(path); err != nil {
            c.Status(http.StatusNotFound)
            return
        }
        c.Status(http.StatusOK)
    })

    grp.GET("/*filepath", func(c *gin.Context) {
        path := strings.TrimPrefix(c.Param("filepath"), "/")
        data, err := renderer.GetAsset(path)
        if err != nil {
            c.Status(http.StatusNotFound)
            return
        }
        c.Data(http.StatusOK, detectContentType(path), data)
    })
}

// GetQueryParam retrieves the value of a query parameter from the context.
func (i *Integrator) GetQueryParam(ctx interface{}, name string) string {
    if c, ok := ctx.(*gin.Context); ok {
        return c.Query(name)
    }
    return ""
}

// GetPathParam retrieves the value of a path parameter from the context.
func (i *Integrator) GetPathParam(ctx interface{}, name string) string {
    if c, ok := ctx.(*gin.Context); ok {
        return c.Param(name)
    }
    return ""
}

// GetRequestMethod retrieves the HTTP method of the request from the context.
func (i *Integrator) GetRequestMethod(ctx interface{}) string {
    if c, ok := ctx.(*gin.Context); ok {
        return c.Request.Method
    }
    return ""
}

// GetFormData retrieves form data from the context.
func (i *Integrator) GetFormData(ctx interface{}) map[string][]string {
    if c, ok := ctx.(*gin.Context); ok {
        if err := c.Request.ParseForm(); err != nil {
            return nil
        }
        return c.Request.Form
    }
    return nil
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

func detectContentType(path string) string {
    switch {
    case strings.HasSuffix(path, ".css"):
        return "text/css; charset=utf-8"
    case strings.HasSuffix(path, ".js"):
        return "application/javascript"
    case strings.HasSuffix(path, ".png"):
        return "image/png"
    case strings.HasSuffix(path, ".jpg"), strings.HasSuffix(path, ".jpeg"):
        return "image/jpeg"
    case strings.HasSuffix(path, ".svg"):
        return "image/svg+xml"
    case strings.HasSuffix(path, ".woff2"):
        return "font/woff2"
    case strings.HasSuffix(path, ".woff"):
        return "font/woff"
    case strings.HasSuffix(path, ".ttf"):
        return "font/ttf"
    case strings.HasSuffix(path, ".map"):
        return "application/json"
    default:
        return "application/octet-stream"
    }
}
