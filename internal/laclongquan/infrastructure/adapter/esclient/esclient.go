package esclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

type EsClient struct {
	host       string
	httpClient *http.Client
	syncLock   sync.Mutex
}

func NewEsClient(host string) *EsClient {
	return &EsClient{
		host: host,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

var (
	ErrNotStatusOK = errors.New("not status ok")
)

func (es *EsClient) CreateOrUpdateUser(user *entity.User) error {
	// form url
	url := es.host + userPostfix

	// form request
	body, err := es.formUserReq(user)
	if err != nil {
		return err
	}

	// send request
	if err := es.makeReq(http.MethodPost, url, body, nil); err != nil {
		return err
	}

	return nil
}

func (es *EsClient) DeleteUser(userID string) error {
	// form url
	url := es.host + userPostfix + "/" + userID

	if err := es.makeReq(http.MethodDelete, url, nil, nil); err != nil {
		return err
	}

	return nil
}

func (es *EsClient) makeReq(method, url string, body []byte, expected interface{}) error {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := es.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// log.Printf("[DEBUG] status code: %d\n", resp.StatusCode)
		return ErrNotStatusOK
	}

	if expected != nil {
		if err := json.NewDecoder(resp.Body).Decode(expected); err != nil {
			return err
		}
	}

	return nil
}

func (es *EsClient) formUserReq(user *entity.User) ([]byte, error) {
	var req = &UserDataReq{
		UserID:   user.ID().String(),
		Phone:    user.Account().Phone,
		Username: user.GetUsername(),
		Name:     user.Name(),
		State:    user.State().String(),
	}

	return json.Marshal(req)
}

func (es *EsClient) SearchUser(keyword string) ([]string, error) {
	// form url
	url := es.host + "/search" + userPostfix
	// logger.Debugf("search user url: ", url)

	var req = &SearchReq{
		Keyword: keyword,
	}
	reqB, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// log.Printf("[DEBUG] req: %s\n", string(reqB))

	var resp = new(SearchResp)
	if err := es.makeReq(http.MethodPost, url, reqB, resp); err != nil {
		return nil, err
	}

	var ids []string
	for _, r := range *resp {
		ids = append(ids, r.ID)
	}

	return ids, nil
}
