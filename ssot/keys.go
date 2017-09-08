package ssot

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("keys", func() {

	BasePath("/keys")
	Response(BadRequest, ErrorMedia)   // 400
	Response(Unauthorized, ErrorMedia) // 401
	Response(NotFound, ErrorMedia)     // 404
	Response(InternalServerError)      // 500

	Action("list", func() {
		Description("Given a list of GitHub usernames, responds with list of public SSH keys for each User (associated to their GitHub account).")
		Routing(POST(""))
		Payload(ArrayOf(String))
		Response(OK, CollectionOf(User))
		Metadata("swagger:summary", "List public SSH Keys based on username (or set of usernames)")
	})
})

// UserKey is the type that represents a Github user's public SSH Key
var UserKey = Type("UserKey", func() {
	Description("Type for a GitHub user's public SSH Key")

	Attribute("id", Integer, "The ID of the public SSH key on GitHub.", func() {
		Example(12345)
	})
	Attribute("key", String, "The public SSH key", func() {
		Example("ssh-rsa ABC123 ...")
	})
	Required("id", "key")
})

// User ...
var User = MediaType("application/vnd.User+json", func() {
	Description("Response Type for a GitHub User's list of public SSH Keys")

	Attribute("username", String, "The username of the GitHub user.", func() {
		Example("myname")
	})
	Attribute("keys", ArrayOf(UserKey), "The list of the Github user's public SSH keys.", func() {
		NoExample()
	})

	Required("username", "keys")

	View("default", func() {
		Attribute("username")
		Attribute("keys")
	})
})
