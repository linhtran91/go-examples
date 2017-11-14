// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "My API": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/linh-snoopy/go-examples/goatest/design
// --out=c:\Users\LENOVO\go\src\github.com\linh-snoopy\go-examples\goatest
// --version=v1.3.0

package client

import (
	"net/http"
)

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

// DecodeMyUser decodes the MyUser instance encoded in resp body.
func (c *Client) DecodeMyUser(resp *http.Response) (*MyUser, error) {
	var decoded MyUser
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// MyUserCollection is the media type for an array of MyUser (default view)
//
// Identifier: vnd.my.user; type=collection; view=default
type MyUserCollection []*MyUser

// DecodeMyUserCollection decodes the MyUserCollection instance encoded in resp body.
func (c *Client) DecodeMyUserCollection(resp *http.Response) (MyUserCollection, error) {
	var decoded MyUserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}