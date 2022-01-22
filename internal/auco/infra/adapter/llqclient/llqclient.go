package llqclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/thanhpp/zola/pkg/logger"
)

var (
	ErrNotStatusOK = errors.New("not status ok")
)

type LacLongQuanClient struct {
	host       string
	httpClient *http.Client
}

func NewLLQClient(host string) *LacLongQuanClient {
	return &LacLongQuanClient{
		host: host,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *LacLongQuanClient) doRequest(method, url, authHeader string, req, expected interface{}) error {
	var (
		reqB []byte
		err  error
	)
	if req != nil {
		reqB, err = json.Marshal(req)
		if err != nil {
			return err
		}
	}

	httpReq, err := http.NewRequest(method, url, bytes.NewBuffer(reqB))
	if err != nil {
		return err
	}
	if len(authHeader) != 0 {
		httpReq.Header.Set("Authorization", authHeader)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		logger.Errorf("LLQClient - status code: %d, body: %s", resp.StatusCode, string(body))
		return ErrNotStatusOK
	}

	if expected != nil {
		if err := json.NewDecoder(resp.Body).Decode(expected); err != nil {
			return err
		}
	}

	return nil
}

func (c *LacLongQuanClient) ValidateToken(token string) (*ValidateTokenResp, error) {
	// form url
	var (
		url  = c.host + "/internal/validatetoken"
		resp = new(ValidateTokenResp)
	)

	// add Bearer
	authHeader := "Bearer " + token

	if err := c.doRequest(http.MethodGet, url, authHeader, nil, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *LacLongQuanClient) GetUserInfo(userID string) (*GetUserInfoResp, error) {
	// form url
	var (
		url  = c.host + "/internal/user/" + userID
		resp = new(GetUserInfoResp)
	)

	if err := c.doRequest(http.MethodGet, url, "", nil, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *LacLongQuanClient) IsBlock(userAID, userBID string) (*IsBlockResp, error) {
	// form url
	var (
		url  = c.host + "/internal/isblock?usera=" + userAID + "&userb=" + userBID
		resp = new(IsBlockResp)
	)

	if err := c.doRequest(http.MethodGet, url, "", nil, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
