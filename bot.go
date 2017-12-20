// https://core.telegram.org/bots/api
// Bot API 3.5
package tgbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const urlPattern = "https://api.telegram.org/bot%s/%s"

type Bot struct {
	token string
	Httpc *http.Client
}

func NewBot(token string) *Bot {
	return &Bot{
		token: token,
		Httpc: http.DefaultClient,
	}
}

func (b *Bot) request(u string, params interface{}) (json.RawMessage, error) {
	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)
	if err := enc.Encode(params); err != nil {
		return nil, fmt.Errorf("could not encode params: %s", err)
	}

	req, err := http.NewRequest("GET", u, &buf)
	if err != nil {
		return nil, fmt.Errorf("could not new request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := b.Httpc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not do request: %s", err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var rb ResponseBody
	if err := dec.Decode(&rb); err != nil {
		return nil, fmt.Errorf("could not decode response: %s", err)
	}

	if !rb.Ok {
		return nil, fmt.Errorf("could not get result: %s", rb.Description)
	}

	return rb.Result, nil
}

func (b *Bot) GetUpdates(params *GetUpdatesParams) ([]*Update, error) {
	u := fmt.Sprintf(urlPattern, b.token, "getUpdates")

	result, err := b.request(u, params)
	if err != nil {
		return nil, err
	}

	var updates []*Update
	if err := json.Unmarshal(result, &updates); err != nil {
		return nil, fmt.Errorf("could not unmarshal updates: %s", err)
	}

	return updates, nil
}

func (b *Bot) SendMessage(params *SendMessageParams) (*Message, error) {
	u := fmt.Sprintf(urlPattern, b.token, "sendMessage")

	result, err := b.request(u, params)
	if err != nil {
		return nil, err
	}

	var msg Message
	if err := json.Unmarshal(result, &msg); err != nil {
		return nil, fmt.Errorf("could not unmarshal message: %s", err)
	}

	return &msg, nil
}

func (b *Bot) EditMessageText(params *EditMessageTextParams) (*Message, error) {
	u := fmt.Sprintf(urlPattern, b.token, "editMessageText")

	result, err := b.request(u, params)
	if err != nil {
		return nil, err
	}

	var msg Message
	if err := json.Unmarshal(result, &msg); err != nil {
		return nil, fmt.Errorf("could not unmarshal message: %s", err)
	}

	return &msg, nil
}
