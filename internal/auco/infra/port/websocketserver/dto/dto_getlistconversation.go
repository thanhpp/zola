package dto

import (
	"strconv"

	"github.com/thanhpp/zola/internal/auco/app"
)

type GetListConversationResp struct {
	DefaultResp
	Data []GetListConversationRespElem `json:"data"`
}

type GetListConversationRespElem struct {
	ID          string          `json:"id"`
	Partner     PartnerData     `json:"partner"`
	LastMessage LastMessageData `json:"lastmessage"`
}

func (resp *GetListConversationResp) SetData(data *app.GetListConversationRes, requestorID string) {
	if resp == nil || data == nil {
		return
	}

	resp.Data = make([]GetListConversationRespElem, 0, len(data.Data))
	for i := range data.Data {
		var partnerID string
		if data.Data[i].Room.UserA == requestorID {
			partnerID = data.Data[i].Room.UserB
		} else {
			partnerID = data.Data[i].Room.UserA
		}

		resp.Data = append(resp.Data, GetListConversationRespElem{
			ID: data.Data[i].Room.ID,
			Partner: PartnerData{
				ID: partnerID,
			},
			LastMessage: LastMessageData{
				Messsage: data.Data[i].LastMessage.Content,
				Created:  data.Data[i].LastMessage.Created,
				Unread:   boolTranslate(data.Data[i].LastMessage.IsSeen()),
			},
		})
	}
}

type PartnerData struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}

func (pd *PartnerData) SetData(id, username, name, avatar string) {
	if pd == nil {
		return
	}
	pd.ID = id
	pd.Username = username
	pd.Name = name
	pd.Avatar = avatar
}

type LastMessageData struct {
	Messsage string `json:"message"`
	Created  string `json:"created"`
	Unread   string `json:"unread"`
}

func (ld *LastMessageData) SetData(msg string, created int64, unread bool) {
	if ld == nil {
		return
	}

	ld.Messsage = msg
	ld.Created = strconv.FormatInt(created, 10)
	ld.Unread = boolTranslate(unread)
}
