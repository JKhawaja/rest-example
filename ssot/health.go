package ssot

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("health", func() {

	BasePath("/health")

	Action("healthcheck", func() {
		Description("Returns a 200 if service is available.")
		Routing(GET(""))
		Response(OK)
		Metadata("swagger:summary", "Returns a 200 if service is available.")
	})
})
