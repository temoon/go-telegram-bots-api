package requests

import (
    "errors"
    "strconv"
)

type UnpinChatMessage struct {
    ChatID interface{}
}

func (r *UnpinChatMessage) IsMultipart() bool {
    return false
}

func (r *UnpinChatMessage) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    return
}
