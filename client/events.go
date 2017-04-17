package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// KeepEventsPath computes a request path to the keep action of events.
func KeepEventsPath(eventID int) string {
	return fmt.Sprintf("/api/v2/events/%v/keep", eventID)
}

// <b>イベントお気に入り操作</b><br>
// isKeepがtrueだった場合はフォロー、falseの場合はアンフォローとする。<br>
// 存在しないイベントへのリクエストは404エラーを返す。
func (c *Client) KeepEvents(ctx context.Context, path string, isKeep bool) (*http.Response, error) {
	req, err := c.NewKeepEventsRequest(ctx, path, isKeep)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewKeepEventsRequest create the request corresponding to the keep action endpoint of the events resource.
func (c *Client) NewKeepEventsRequest(ctx context.Context, path string, isKeep bool) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	tmp15 := strconv.FormatBool(isKeep)
	values.Set("isKeep", tmp15)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.UserTokenSigner != nil {
		c.UserTokenSigner.Sign(req)
	}
	return req, nil
}

// ListEventsPath computes a request path to the list action of events.
func ListEventsPath(id string) string {
	return fmt.Sprintf("/api/v2/events/genre/%v", id)
}

// ListEventsPath2 computes a request path to the list action of events.
func ListEventsPath2() string {
	return fmt.Sprintf("/api/v2/events/new")
}

// ListEventsPath3 computes a request path to the list action of events.
func ListEventsPath3() string {
	return fmt.Sprintf("/api/v2/events/keep")
}

// ListEventsPath4 computes a request path to the list action of events.
func ListEventsPath4() string {
	return fmt.Sprintf("/api/v2/events/nokeep")
}

// ListEventsPath5 computes a request path to the list action of events.
func ListEventsPath5() string {
	return fmt.Sprintf("/api/v2/events/popular")
}

// ListEventsPath6 computes a request path to the list action of events.
func ListEventsPath6() string {
	return fmt.Sprintf("/api/v2/events/recommend")
}

// <b>イベント情報取得</b><br>
// <ul>
// <li>/genre/:id -> ジャンル別新着情報</li>
// <li>/new -> ユーザー別新着情報</li>
// <li>/keep -> ユーザーがキープしているイベント</li>
// <li>/nokeep -> ユーザーが興味なしにしたイベント</li>
// <li>/popular -> キープ数が多い。注目されているイベント</li>
// <li>/recommend -> ユーザー属性に合わせたおすすめイベント</li>
// </ul>
// イベントの情報は区切って送信され、スクロールイベントで次のページのイベント情報を取得することを想定している。<br>
// また、キープや興味なしの操作は１日に１回行われるバッチ処理時に確定されるまでは、分類されずに表示される。
func (c *Client) ListEvents(ctx context.Context, path string, page *int, q *string, sort *string) (*http.Response, error) {
	req, err := c.NewListEventsRequest(ctx, path, page, q, sort)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListEventsRequest create the request corresponding to the list action endpoint of the events resource.
func (c *Client) NewListEventsRequest(ctx context.Context, path string, page *int, q *string, sort *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if page != nil {
		tmp16 := strconv.Itoa(*page)
		values.Set("page", tmp16)
	}
	if q != nil {
		values.Set("q", *q)
	}
	if sort != nil {
		values.Set("sort", *sort)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.UserTokenSigner != nil {
		c.UserTokenSigner.Sign(req)
	}
	return req, nil
}
