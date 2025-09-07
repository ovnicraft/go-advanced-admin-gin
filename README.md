# Go Advanced Admin - Gin Integration

[Gin](https://github.com/gin-gonic/gin) framework integration for the Go Advanced Admin Panel.

This package provides integration with the Gin web framework for the Go Advanced Admin Panel, enabling you to use Gin as your web framework.

## Installation

Add the module to your project by running:

```sh
go get github.com/ovnicraft/go-advanced-admin-gin
```

## Documentation

For detailed documentation on how to use the Gin integration, please visit the [official documentation website](https://goadmin.dev/Gin.html).

Note: [Gin Quickstart](https://github.com/go-advanced-admin/docs/pull/21) Gin documentation PR is opened.

## Quick Start

```go
import (
    "github.com/ovnicraft/admin"
    "github.com/ovnicraft/go-advanced-admin-gin"
    "github.com/gin-gonic/gi"
)

func main() {
    // Initialize Echo
    e := echo.New()

    // Initialize the web integrator
    webIntegrator := adminecho.NewIntegrator(e.Group("/admin"))

    // Use webIntegrator when initializing the admin panel
}
```

For more detailed examples and configuration options, please refer to the [Gin Integration Guide](https://goadmin.dev/Gin.html).

## Contributing

Contributions are always welcome! Please refer to the [Contributing Guidelines](https://github.com/go-advanced-admin/admin/blob/main/CONTRIBUTING.md) in the main repository.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.
