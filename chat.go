package drbot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cellularhacker/logger-go"
)

const (
	ValueChannelAdmin      = "admin"
	ValueIntendId          = ""
	ValueIsStartFalse      = 0
	ValueIsStartTrue       = 1
	ValueLangEn            = "en"
	ValueLangJa            = "ja"
	ValueLangKo            = "ko"
	ValueModelMessage      = "message"
	ValueRequestTypeButton = "button"
	ValueRequestTypeText   = "text"
	ValueTimeZoneAsiaSeoul = "Asia/Seoul"
	ValueUtteranceStart    = "[start]"
	ValueVisitCountStart   = int64(1)
)

type ChatReq struct {
	UserRequest *UserRequest `json:"userRequest"`
	VisitCount  int64        `json:"visit_count"`
	Context     *ChatContext `json:"context"`
}

func (scr *ChatReq) Bytes() []byte {
	bs, _ := json.Marshal(scr)
	return bs
}
func (scr *ChatReq) String() string {
	return string(scr.Bytes())
}

type SendChatResp struct {
	UserRequest  UserRequest `json:"userRequest"`
	Animation    string      `json:"animation"`
	Input        ChatInput   `json:"input"`
	Context      ChatContext `json:"context"`
	Intent       ChatIntent  `json:"intent"`
	Response     ChatResp    `json:"response"`
	VisitCount   int         `json:"visit_count"`
	CreatedAt    string      `json:"created_at"`
	ResponseTime int         `json:"response_time"`
}

type UserRequest struct {
	Model       string   `json:"model"`
	IntentId    string   `json:"intentId"`
	IsStart     int      `json:"isStart"`
	Timezone    string   `json:"timezone"`
	Utterance   string   `json:"utterance"`
	Lang        string   `json:"lang"`
	Channel     string   `json:"channel"`
	TopFolderId string   `json:"top_folder_id,omitempty"`
	CbcType     string   `json:"cbc_type,omitempty"`
	CbcRound    int64    `json:"cbc_round,omitempty"`
	CbcTerm     int64    `json:"cbc_term,omitempty"`
	RequestType string   `json:"requestType"`
	User        *UserReq `json:"user"`
}

type UserReq struct {
	Id        string `json:"id"`
	SessionId string `json:"sessionId"`
}
type ChatInput struct {
	Text string `json:"text"`
}
type ChatContext struct {
}
type ChatIntent struct {
	Module      string  `json:"module"`
	IntentId    string  `json:"intent_id"`
	IntentTitle string  `json:"intent_title"`
	Confidence  float64 `json:"confidence"`
}
type ChatResp struct {
	Outputs []ChatOutput `json:"outputs"`
}
type ChatOutput struct {
	Type  string     `json:"type"`
	Items []ChatItem `json:"items"`
}
type ChatItem []struct {
	Title       string          `json:"title"`
	Subtitle    string          `json:"subtitle,omitempty"`
	Description string          `json:"description,omitempty"`
	Similarity  string          `json:"similarity,omitempty"`
	Source      string          `json:"source,omitempty"`
	Thumbnail   json.RawMessage `json:"thumbnail,omitempty"` // MARK: Guess -> *string
	Buttons     json.RawMessage `json:"buttons,omitempty"`
}

func NewInitChatReq() *ChatReq {
	return &ChatReq{
		UserRequest: &UserRequest{
			Model:       "",
			IntentId:    "",
			IsStart:     0,
			Timezone:    "",
			Utterance:   "",
			Lang:        "",
			Channel:     "",
			RequestType: "",
			User: &UserReq{
				Id:        "",
				SessionId: "",
			},
		},
		VisitCount: 0,
		Context:    &ChatContext{},
	}
}

type ChatData struct {
	UserID     string `json:"user_id"`
	VisitCount int64  `json:"visit_count"`
	SessionId  string `json:"session_id"`
	Message    string `json:"message"`
	Language   string `json:"language,omitempty"`
}

func NewChatData() *ChatData {
	return &ChatData{}
}

//type SendChatResp struct {
//	SessionId      string `json:"session_id"`
//	NextVisitCount int64  `json:"next_visit_count"`
//}

func NewSendChatResp() *SendChatResp {
	return &SendChatResp{}
}

func SendChat(isStart bool, initData ...*ChatData) (*SendChatResp, error) {
	req := NewInitChatReq()

	if isStart {
		req.VisitCount = ValueVisitCountStart
		req.UserRequest.User.SessionId = GetNewUUID()
		req.UserRequest.RequestType = ValueRequestTypeButton
		req.UserRequest.Utterance = ValueUtteranceStart
	} else {
		if len(initData) < 1 {
			return nil, ErrNeedChatDataExceptStartingTheChat
		}

		chatData := initData[0]

		req.VisitCount = chatData.VisitCount
		req.UserRequest.User.Id = chatData.UserID
		req.UserRequest.User.SessionId = chatData.SessionId
		req.UserRequest.RequestType = ValueRequestTypeText
		req.UserRequest.Utterance = chatData.Message
		if len(chatData.Language) <= 0 {
			req.UserRequest.Lang = ValueLangKo
		} else {
			req.UserRequest.Lang = chatData.Language
		}
	}

	// MARK: Setting common values...
	req.UserRequest.Timezone = ValueTimeZoneAsiaSeoul
	req.UserRequest.Channel = ValueChannelAdmin
	req.UserRequest.IsStart = ValueIsStartFalse

	resp, err := MakeRequestChat(http.MethodPost, PathAPI, nil, req.Bytes())
	if err != nil {
		err = fmt.Errorf("MakeRequestChat(http.MethodPost, PathAPI, nil, req.Bytes()): %s", err)
		logger.L.Errorln("\t", err)
		logger.L.Errorln("\tRequest Body")
		logger.L.Errorln("\t", req.String())
		logger.L.Errorln("\tResponse Body")
		logger.L.Errorln("\t", string(resp))
		return nil, err
	}

	respBody := NewSendChatResp()
	if err = json.Unmarshal(resp, respBody); err != nil {
		logger.L.Errorln(ErrInvalidMakeRequestChatResponseData)
		logger.L.Errorln("\t", err)
		logger.L.Errorln("\tRequest Body")
		logger.L.Errorln("\tResponse Body")
		logger.L.Errorln("\t", string(resp))
		return nil, fmt.Errorf("ErrInvalidMakeRequestChatResponseData: %s", err)
	}

	return respBody, nil
}
