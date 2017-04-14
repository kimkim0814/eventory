// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/eventory/design
// --out=$(GOPATH)
// --version=v1.1.0-dirty
//
// API "eventory": prefs Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// FollowPrefsPath computes a request path to the follow action of prefs.
func FollowPrefsPath(prefID int) string {
	param0 := strconv.Itoa(prefID)

	return fmt.Sprintf("/api/v2/prefs/%s/follow", param0)
}

// FollowPrefsPath2 computes a request path to the follow action of prefs.
func FollowPrefsPath2(prefID int) string {
	param0 := strconv.Itoa(prefID)

	return fmt.Sprintf("/api/v2/prefs/%s/follow", param0)
}

// ジャンルお気に入り操作
func (c *Client) FollowPrefs(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewFollowPrefsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFollowPrefsRequest create the request corresponding to the follow action endpoint of the prefs resource.
func (c *Client) NewFollowPrefsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.KeySigner != nil {
		c.KeySigner.Sign(req)
	}
	return req, nil
}
