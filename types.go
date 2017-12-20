package tgbot

import "encoding/json"

type GetUpdatesParams struct {
	Offset         int64    `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type SendMessageParams struct {
	ChatId      int64                 `json:"chat_id"`
	Text        string                `json:"text"`
	ParseMode   string                `json:"parse_mode,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageTextParams struct {
	ChatId      int64                 `json:"chat_id"`
	MessageId   int64                 `json:"message_id	"`
	Text        string                `json:"text"`
	ParseMode   string                `json:"parse_mode,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type Update struct {
	Id            int64          `json:"update_id"`
	Message       *Message       `json:"message"`
	CallbackQuery *CallbackQuery `json:"callback_query"`
}

type Message struct {
	Id       int64             `json:"message_id"`
	Chat     *Chat             `json:"chat"`
	Text     string            `json:"text"`
	Entities []*MessageEnitity `json:"entities"`
}

func (m *Message) Command() (string, string) {
	for _, entry := range m.Entities {
		if entry.Type == "bot_command" {
			cmd := m.Text[entry.Offset : entry.Offset+entry.Length]

			argsOffset := entry.Offset + entry.Length + 1
			if argsOffset >= len(m.Text) {
				return cmd, ""
			}

			return cmd, m.Text[argsOffset:len(m.Text)]
		}
	}
	return "", ""
}

type Chat struct {
	Id   int64  `json:"id"`
	Type string `json:"type"`
}

type MessageEnitity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

type CallbackQuery struct {
	Id      string   `json:"id"`
	Message *Message `json:"message"`
	Data    string   `json:"data"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data,omitempty"`
}

type ResponseBody struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
}
