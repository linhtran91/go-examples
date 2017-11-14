// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "My API": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/linh-snoopy/go-examples/goatest/design
// --out=c:\Users\LENOVO\go\src\github.com\linh-snoopy\go-examples\goatest
// --version=v1.3.0

package app

// MyUser media type (default view)
//
// Identifier: vnd.my.user; view=default
type MyUser struct {
	// User's email address
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// User's name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Phone's number
	Phone *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
}

// MyUserCollection is the media type for an array of MyUser (default view)
//
// Identifier: vnd.my.user; type=collection; view=default
type MyUserCollection []*MyUser