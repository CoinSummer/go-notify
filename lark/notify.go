package lark

import (
	"encoding/json"
	"errors"
	"github.com/imroc/req"
)

type Options struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
}

type client struct {
	opt Options
}

func New(opt Options) *client {
	return &client{opt: opt}
}

type Resp struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

type Webhook struct {
	Content string `json:"content"`
}

type WebhookRequestData struct {
	MsgType string `json:"msg_type"`
	Content string `json:"content"`
}

type Text struct {
	Text string `json:"text"`
}

func (c *client) Send(message string) error {

	if "" == message {
		return errors.New("missing message")
	}

	t := Text{Text: message}
	tj, _ := json.Marshal(t)
	rd := WebhookRequestData{
		MsgType: "text",
		Content: string(tj),
	}
	rj, _ := json.Marshal(rd)

	webhook := c.opt.Token
	resp, err := req.Post(webhook, string(rj))
	if err != nil {
		return err
	}

	r := &Resp{}
	return resp.ToJSON(r)
}
