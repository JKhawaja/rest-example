package ssot

import (
	// . "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("GitHub SSH Keys", func() {
	Title("GitHub SSH Keys")
	Description(`This API is for retrieving lists of public SSH keys based on github usernames.`)
	Version("0.0.1")
	Scheme("http")
	Consumes("application/json", func() {
		Package("github.com/goadesign/goa/encoding/json")
	})
	Produces("application/json", func() {
		Package("github.com/goadesign/goa/encoding/json")
	})
})
