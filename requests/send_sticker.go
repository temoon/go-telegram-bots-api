package requests

import (
    "encoding/json"
    "errors"
    "os"
    "strconv"
)

type SendSticker struct {
    ChatID              interface{}
    Sticker             interface{}
    DisableNotification bool
    ReplyToMessageID    int
    ReplyMarkup         interface{}
}

func (r *SendSticker) IsMultipart() bool {
    _, ok := r.Sticker.(*os.File)

    return ok
}

func (r *SendSticker) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["sticker"] = []interface{}{r.Sticker}

    if r.DisableNotification {
        values["disable_notification"] = []interface{}{"1"}
    }

    if r.ReplyToMessageID != 0 {
        values["reply_to_message_id"] = []interface{}{strconv.Itoa(r.ReplyToMessageID)}
    }

    if r.ReplyMarkup != nil {
        var data []byte
        if data, err = json.Marshal(r.ReplyMarkup); err != nil {
            return
        }

        values["reply_markup"] = []interface{}{string(data)}
    }

    return
}