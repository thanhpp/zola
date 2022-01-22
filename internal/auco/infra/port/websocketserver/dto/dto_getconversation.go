package dto

import (
	"github.com/thanhpp/zola/internal/auco/app"
)

type GetConversationResp struct {
	DefaultResp
	Data struct {
		Conversation []MessageData `json:"conversation"`
		IsBlocked    string        `json:"is_blocked"`
	} `json:"data"`
}

type MessageData struct {
	Message   string      `json:"message"`
	MessageID string      `json:"message_id"`
	Unread    string      `json:"unread"`
	Created   string      `json:"created"`
	Sender    PartnerData `json:"sender"`
}

func (resp *GetConversationResp) SetData(data *app.GetConversationRes) {
	if resp == nil {
		return
	}
	resp.Data.Conversation = make([]MessageData, 0, len(data.Data))
	for i := range data.Data {
		resp.Data.Conversation = append(resp.Data.Conversation, MessageData{
			Message:   data.Data[i].Message.Content,
			MessageID: data.Data[i].Message.MsgID,
			Unread:    boolTranslate(!data.Data[i].Message.IsSeen()),
			Created:   data.Data[i].Message.Created,
			Sender: PartnerData{
				ID:       data.Data[i].Sender.ID,
				Name:     data.Data[i].Sender.Name,
				Username: data.Data[i].Sender.Username,
				Avatar:   data.Data[i].Sender.Avatar,
			},
		})
	}
	resp.Data.IsBlocked = boolTranslate(false)
}

func (resp *GetConversationResp) SetIsBlocked() {
	if resp == nil {
		return
	}
	resp.Data.IsBlocked = boolTranslate(true)
}
