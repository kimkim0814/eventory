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

// <b>ログイン認証</b><br>
// 正規ユーザーのメールアドレスとパスワードのハッシュを送ることで、ユーザー認証を行う<br>
// 正しくユーザー認証が完了した場合、正規ユーザーのIDを仮ユーザーIDに紐付けを行い。<br>
// ユーザーの行動を別端末で引き継ぐことが出来る。<br>
func (c *Client) LoginUsers(ctx context.Context, path string, email string, passwordHash string) (*http.Response, error) {
	req, err := c.NewLoginUsersRequest(ctx, path, email, passwordHash)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLoginUsersRequest create the request corresponding to the login action endpoint of the users resource.
func (c *Client) NewLoginUsersRequest(ctx context.Context, path string, email string, passwordHash string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("email", email)
	values.Set("password_hash", passwordHash)
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

// RegularCreateUsersPath computes a request path to the regular create action of users.
func RegularCreateUsersPath() string {
	return fmt.Sprintf("/api/v2/users/new")
}

// <b>正規ユーザーの作成</b><br>
// メールアドレスとパスワードハッシュを使って、正規ユーザーの作成を行う。<br>
// もし、既に存在するアカウントだった場合は、"alreadyExists"を返す。<br>
// 正しく実行された場合は、"ok"を返す。
func (c *Client) RegularCreateUsers(ctx context.Context, path string, email string, passwordHash string) (*http.Response, error) {
	req, err := c.NewRegularCreateUsersRequest(ctx, path, email, passwordHash)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRegularCreateUsersRequest create the request corresponding to the regular create action endpoint of the users resource.
func (c *Client) NewRegularCreateUsersRequest(ctx context.Context, path string, email string, passwordHash string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("email", email)
	values.Set("password_hash", passwordHash)
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

// StatusUsersPath computes a request path to the status action of users.
func StatusUsersPath() string {
	return fmt.Sprintf("/api/v2/users/status")
}

// <b>ユーザーの端末情報更新</b><br>
// 利用者のバージョンや端末情報を更新する。この更新処理は起動時に行われるものとする。
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
	if c.UserTokenSigner != nil {
		c.UserTokenSigner.Sign(req)
	}
	return req, nil
}

// TmpCreateUsersPath computes a request path to the tmp create action of users.
func TmpCreateUsersPath() string {
	return fmt.Sprintf("/api/v2/users/tmp")
}

// <b>一時ユーザーの作成</b><br>
// 初回起動時に仮ユーザーを作成する。ここで与えられるユーザーIDは、メールアドレスなどとひも付きがないため、<br>
// 端末が変わるとtokenが変わるので、別端末で共有するには、正規ユーザーの登録が必要になる。
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
