package graphql

import (
	"github.com/go-ggz/ggz/pkg/config"
	"github.com/go-ggz/ggz/pkg/schema"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

// Handler initializes the prometheus middleware.
func Handler() gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   config.Server.GraphiQL,
		GraphiQL: config.Server.GraphiQL,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
