package requests

import (
    "errors"
    "strconv"
)

type ForwardMessage struct {
    ChatID              interface{}
    FromChatID          interface{}
    DisableNotification bool
    MessageID           int
}

func (r *ForwardMessage) IsMultipart() bool {
    return false
}

func (r *ForwardMessage) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch r.FromChatID.(type) {
    case int64:
        values["from_chat_id"] = []interface{}{strconv.FormatInt(r.FromChatID.(int64), 10)}
    case string:
        values["from_chat_id"] = []interface{}{r.FromChatID.(string)}
    default:
        return nil, errors.New("invalid from_chat_id")
    }

    if r.DisableNotification == true {
        values["disable_notification"] = []interface{}{"1"}
    }

    values["message_id"] = []interface{}{strconv.Itoa(r.MessageID)}

    return
}