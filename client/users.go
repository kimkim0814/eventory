// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/eventory/design
// --out=$(GOPATH)
// --version=v1.1.0-dirty
//
// API "eventory": users Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// LoginUsersPath computes a request path to the login action of users.
func LoginUsersPath() string {

	return fmt.Sprintf("/api/v2/users/login")
}

// ログイン
func (c *Client) LoginUsers(ctx context.Context, path string, email string, password string) (*http.Response, error) {
	req, err := c.NewLoginUsersRequest(ctx, path, email, password)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLoginUsersRequest create the request corresponding to the login action endpoint of the users resource.
func (c *Client) NewLoginUsersRequest(ctx context.Context, path string, email string, password string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("email", email)
	values.Set("password", password)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.KeySigner != nil {
		c.KeySigner.Sign(req)
	}
	return req, nil
}

// RegularCreateUsersPath computes a request path to the regular create action of users.
func RegularCreateUsersPath() string {

	return fmt.Sprintf("/api/v2/users/new")
}

// 正規ユーザーの作成
func (c *Client) RegularCreateUsers(ctx context.Context, path string, email string, identifier string) (*http.Response, error) {
	req, err := c.NewRegularCreateUsersRequest(ctx, path, email, identifier)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRegularCreateUsersRequest create the request corresponding to the regular create action endpoint of the users resource.
func (c *Client) NewRegularCreateUsersRequest(ctx context.Context, path string, email string, identifier string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("email", email)
	values.Set("identifier", identifier)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.KeySigner != nil {
		c.KeySigner.Sign(req)
	}
	return req, nil
}

// StatusUsersPath computes a request path to the status action of users.
func StatusUsersPath() string {

	return fmt.Sprintf("/api/v2/users/status")
}

// 一時ユーザーの作成
func (c *Client) StatusUsers(ctx context.Context, path string, clientVersion string, platform string) (*http.Response, error) {
	req, err := c.NewStatusUsersRequest(ctx, path, clientVersion, platform)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewStatusUsersRequest create the request corresponding to the status action endpoint of the users resource.
func (c *Client) NewStatusUsersRequest(ctx context.Context, path string, clientVersion string, platform string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("client_version", clientVersion)
	values.Set("platform", platform)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.KeySigner != nil {
		c.KeySigner.Sign(req)
	}
	return req, nil
}

// TmpCreateUsersPath computes a request path to the tmp create action of users.
func TmpCreateUsersPath() string {

	return fmt.Sprintf("/api/v2/users/tmp")
}

// 一時ユーザーの作成
func (c *Client) TmpCreateUsers(ctx context.Context, path string, clientVersion string, identifier string, platform string) (*http.Response, error) {
	req, err := c.NewTmpCreateUsersRequest(ctx, path, clientVersion, identifier, platform)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewTmpCreateUsersRequest create the request corresponding to the tmp create action endpoint of the users resource.
func (c *Client) NewTmpCreateUsersRequest(ctx context.Context, path string, clientVersion string, identifier string, platform string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("client_version", clientVersion)
	values.Set("identifier", identifier)
	values.Set("platform", platform)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
