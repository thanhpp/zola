package esclient

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

func (es *EsClient) formPostReq(post *entity.Post, authorName string) ([]byte, error) {
	var req = PostDataReq{
		ID:        post.ID(),
		Described: post.Content(),
		Author: PostAuthorData{
			ID:   post.Creator(),
			Name: authorName,
		},
		Created:  strconv.FormatInt(post.CreatedAt(), 10),
		Modified: strconv.FormatInt(post.UpdatedAt(), 10),
	}

	return json.Marshal(req)
}

func (es *EsClient) CreateUpdatePost(post *entity.Post, authorName string) error {
	// url
	url := es.host + postPostfix

	body, err := es.formPostReq(post, authorName)
	if err != nil {
		return err
	}

	// send request
	if err := es.makeReq(http.MethodPost, url, body, nil); err != nil {
		return err
	}

	return nil
}

func (es *EsClient) DeletePost(postID string) error {
	// url
	url := es.host + postPostfix + "/" + postID

	if err := es.makeReq(http.MethodDelete, url, nil, nil); err != nil {
		return err
	}

	return nil
}

func (es *EsClient) SearchPost(keyword string, index, count int) ([]string, error) {
	// form url
	url := es.host + "/search" + postPostfix

	var req = &SearchReq{
		Keyword: keyword,
		Index:   strconv.Itoa(index),
		Count:   strconv.Itoa(count),
	}
	reqB, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var resp = new(SearchPostResp)
	if err := es.makeReq(http.MethodPost, url, reqB, resp); err != nil {
		return nil, err
	}

	var ids = make([]string, 0, len(*resp))
	for _, item := range *resp {
		ids = append(ids, item.ID)
	}

	return ids, nil
}

func (es *EsClient) GetAllPost() ([]string, error) {
	// form url
	url := es.host + "/posts"

	var resp = new(SearchPostResp)
	if err := es.makeReq(http.MethodGet, url, nil, resp); err != nil {
		return nil, err
	}

	var ids = make([]string, 0, len(*resp))
	for _, item := range *resp {
		ids = append(ids, item.ID)
	}

	return ids, nil
}
