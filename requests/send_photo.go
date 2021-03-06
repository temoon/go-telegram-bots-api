package requests

import (
    "encoding/json"
    "errors"
    "io"
    "strconv"
)

type SendPhoto struct {
    ChatID              interface{}
    Photo               interface{}
    Caption             string
    ParseMode           string
    DisableNotification bool
    ReplyToMessageID    uint32
    ReplyMarkup         interface{}
}

func (r *SendPhoto) IsMultipart() bool {
    _, ok := r.Photo.(io.Reader)

    return ok
}

func (r *SendPhoto) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch photo := r.Photo.(type) {
    case string:
        values["photo"] = photo
    case io.Reader:
        values["photo"] = photo
    default:
        return nil, errors.New("invalid photo")
    }

    if r.Caption != "" {
        values["caption"] = r.Caption
    }

    if r.ParseMode != "" {
        values["parse_mode"] = r.ParseMode
    }

    if r.DisableNotification {
        values["disable_notification"] = "1"
    }

    if r.ReplyToMessageID != 0 {
        values["reply_to_message_id"] = strconv.FormatUint(uint64(r.ReplyToMessageID), 10)
    }

    if r.ReplyMarkup != nil {
        var data []byte
        if data, err = json.Marshal(r.ReplyMarkup); err != nil {
            return
        }

        values["reply_markup"] = string(data)
    }

    return
}
