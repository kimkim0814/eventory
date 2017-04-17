package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// CreateGenresPath computes a request path to the create action of genres.
func CreateGenresPath() string {
	return fmt.Sprintf("/api/v2/genres/new")
}

// <b>ジャンルの新規作成</b><br>
// 新しく作成するジャンル名を送信して、新規作成を行う。追加処理が完了とするとジャンルIDが返ってくるので、それを自動でフォローするようにする。<br>
// 但し、ジャンルを新規作成する前に、ジャンル名を検索するフローを挟み、検索結果に出てこなかった場合に追加できるようにする。
func (c *Client) CreateGenres(ctx context.Context, path string, name string) (*http.Response, error) {
	req, err := c.NewCreateGenresRequest(ctx, path, name)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateGenresRequest create the request corresponding to the create action endpoint of the genres resource.
func (c *Client) NewCreateGenresRequest(ctx context.Context, path string, name string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("name", name)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.UserTokenSigner != nil {
		c.UserTokenSigner.Sign(req)
	}
	return req, nil
}

// FollowGenresPath computes a request path to the follow action of genres.
func FollowGenresPath(genreID int) string {
	return fmt.Sprintf("/api/v2/genres/%v/follow", genreID)
}

// FollowGenresPath2 computes a request path to the follow action of genres.
func FollowGenresPath2(genreID int) string {
	return fmt.Sprintf("/api/v2/genres/%v/follow", genreID)
}

// <b>ジャンルフォロー操作</b><br>
// PUTでフォロー、DELETEでアンフォローをする。<br>
// HTTPメソッド意外は同じパラメーターで動作する。<br>
// 存在しない都道府県へのリクエストは404エラーを返す。
func (c *Client) FollowGenres(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewFollowGenresRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFollowGenresRequest create the request corresponding to the follow action endpoint of the genres resource.
func (c *Client) NewFollowGenresRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.UserTokenSigner != nil {
		c.UserTokenSigner.Sign(req)
	}
	return req, nil
}

// ListGenresPath computes a request path to the list action of genres.
func ListGenresPath() string {
	return fmt.Sprintf("/api/v2/genres")
}

// <b>ジャンル検索</b><br>
// ジャンル名で検索し、当てはまるジャンルを返す。その際に対象となるジャンルがなかった場合、<br>
// ジャンル追加ボタンを表示し、追加出来るようにする。
func (c *Client) ListGenres(ctx context.Context, path string, page *int, q *string, sort *string) (*http.Response, error) {
	req, err := c.NewListGenresRequest(ctx, path, page, q, sort)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListGenresRequest create the request corresponding to the list action endpoint of the genres resource.
func (c *Client) NewListGenresRequest(ctx context.Context, path string, page *int, q *string, sort *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if page != nil {
		tmp17 := strconv.Itoa(*page)
		values.Set("page", tmp17)
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
