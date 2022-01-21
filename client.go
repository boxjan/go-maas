package maas

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	HttpMethodGet    = "GET"
	HttpMethodPost   = "POST"
	HttpMethodPut    = "PUT"
	HttpMethodDelete = "DELETE"
)

type Client struct {
	apiBaseUrl *url.URL
	client     http.Client
	withHeader map[string]string
	oauth      OAuthSigner
}

// NewMaasClient http://localhost/MAAS/, 2.0, string
func NewMaasClient(baseUrl, apiVersion, apiKey string, defaultTimeout time.Duration,
	proxyUrl string, extraHeader map[string]string) (*Client, error) {

	elements := strings.Split(apiKey, ":")
	if len(elements) != 3 {
		return nil, fmt.Errorf("invalid API key %q; expected \"<consumer secret>:<token key>:<token secret>\"", apiKey)
	}

	token := &OAuthToken{
		ConsumerKey: elements[0],
		// The consumer secret is the empty string in MAAS' authentication.
		ConsumerSecret: "",
		TokenKey:       elements[1],
		TokenSecret:    elements[2],
	}

	signer, err := NewPlainTestOAuthSigner(token, "MAAS API")
	if err != nil {
		return nil, err
	}

	midUrl, err := url.Parse(JoinURLs(baseUrl, "api/"))
	if err != nil {
		return nil, err
	}

	httpClient := http.Client{Timeout: defaultTimeout}

	if len(proxyUrl) > 0 {
		proxy, err := url.Parse(proxyUrl)
		if err != nil {
			return nil, errors.New("the proxy url:" + proxyUrl + "can not be parse")
		}
		httpClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	}

	req, err := http.NewRequest("GET", JoinURLs(midUrl.String(), "version/"), nil)
	if err != nil {
		return nil, fmt.Errorf("generate %s GET request failed with: %+v", midUrl.String()+"version/", err)
	}

	for k, v := range extraHeader {
		req.Header.Set(k, v)
	}

	if rsp, err := httpClient.Do(req); err != nil {
		return nil, fmt.Errorf("try to get maas api version failed with error: %+v", err)
	} else {
		responseHtmlByte, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			return nil, fmt.Errorf("try to purse maas api version failed with error: %+v", err)
		}
		_ = rsp.Body.Close()

		maasSupportApiVersion := string(responseHtmlByte)
		if maasSupportApiVersion != apiVersion {
			return nil, fmt.Errorf("maas only support: %s version, but be asked: %s",
				maasSupportApiVersion, apiVersion,
			)
		}
	}

	if parsedURL, err := url.Parse(midUrl.String() + apiVersion + "/"); err != nil {
		return nil, err
	} else {
		return &Client{
			client:     httpClient,
			apiBaseUrl: parsedURL,
			withHeader: extraHeader,
			oauth:      signer,
		}, nil
	}
}

func (c *Client) Do(r *http.Request) (*http.Response, error) {
	if c.withHeader != nil {
		for k, v := range c.withHeader {
			r.Header.Set(k, v)
		}
	}

	c.oauth.OAuthSign(r)

	return c.client.Do(r)
}

func (c *Client) buildRequestUrl(api string) (*url.URL, error) {

	if strings.HasPrefix(api, c.apiBaseUrl.Path) {
		api = strings.TrimPrefix(api, c.apiBaseUrl.Path)
	}
	uS := JoinURLs(c.apiBaseUrl.String(), api)

	uS = EnsureTrailingSlash(uS)

	finalUri, err := url.Parse(uS)
	if err != nil {
		return nil, fmt.Errorf("format url: %s failed with err: %+v", uS, err)
	}

	return finalUri, nil
}

func (c *Client) Get(resource, operation string, parameters url.Values) (*http.Response, error) {
	if parameters == nil {
		parameters = make(url.Values)
	}
	opParameter := parameters.Get("op")
	if opParameter != "" {
		return nil, fmt.Errorf("reserved parameter 'op' passed (with value '%s')", opParameter)
	}
	if operation != "" {
		parameters.Set("op", operation)
	}

	finalUri, err := c.buildRequestUrl(resource)
	if err != nil {
		return nil, err
	}

	finalUri.RawQuery = parameters.Encode()

	req, err := http.NewRequest(HttpMethodGet, finalUri.String(), nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) Post(resource, operation string, parameters url.Values, body map[string][]byte) (*http.Response, error) {
	finalUri, err := c.buildRequestUrl(resource)
	if err != nil {
		return nil, err
	}

	finalUri.RawQuery = url.Values{"op": []string{operation}}.Encode()
	if body == nil {
		return c.postSimply(finalUri, parameters)
	}
	return c.postWithBody(finalUri, parameters, body)
}

func (c *Client) postWithBody(u *url.URL, parameters url.Values, file map[string][]byte) (*http.Response, error) {

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	// add file body
	for k, v := range file {
		fw, err := writer.CreateFormFile(k, k)
		if err != nil {
			return nil, fmt.Errorf("multipart add `%s` failed with err: %+v", k, err)
		}
		_, err = io.Copy(fw, bytes.NewBuffer(v))
		if err != nil {
			return nil, fmt.Errorf("multipart add `%s` ctx(base64): `%s` failed with err: %+v",
				k, base64.StdEncoding.EncodeToString(v), err)
		}
	}

	// add parameters
	for k, vs := range parameters {
		for _, v := range vs {
			fw, err := writer.CreateFormField(k)
			if err != nil {
				return nil, fmt.Errorf("multipart add `%s` failed with err: %+v", k, err)
			}
			_, err = io.Copy(fw, strings.NewReader(v))
			if err != nil {
				return nil, fmt.Errorf("multipart add `%s` ctx: `%s` failed with err: %+v", k, v, err)
			}
		}
	}

	writer.Close()

	req, err := http.NewRequest(HttpMethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return c.Do(req)

}

func (c *Client) postSimply(u *url.URL, parameters url.Values) (*http.Response, error) {
	req, err := http.NewRequest(HttpMethodPost, u.String(), strings.NewReader(string(parameters.Encode())))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.Do(req)
}

func (c *Client) Put(resource string, parameters url.Values) (*http.Response, error) {
	finalUri, err := c.buildRequestUrl(resource)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(HttpMethodPut, finalUri.String(), strings.NewReader(string(parameters.Encode())))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.Do(req)
}

func (c *Client) Delete(resource string) (*http.Response, error) {
	finalUri, err := c.buildRequestUrl(resource)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(HttpMethodDelete, finalUri.String(), nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func responseReadAndClose(rc io.ReadCloser) ([]byte, error) {
	if rc == nil {
		return nil, nil
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}

func (c *Client) TurnResponse(response *http.Response, RErr error) ([]byte, error) {
	if RErr != nil {
		return nil, RErr
	}
	body, err := responseReadAndClose(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return body, fmt.Errorf("status: %s", response.Status)
	}

	return body, nil
}
