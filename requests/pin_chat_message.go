package requests

import (
    "errors"
    "strconv"
)

type PinChatMessage struct {
    ChatID              interface{}
    MessageID           uint32
    DisableNotification bool
}

func (r *PinChatMessage) IsMultipart() bool {
    return false
}

func (r *PinChatMessage) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["message_id"] = strconv.FormatUint(uint64(r.MessageID), 10)

    if r.DisableNotification {
        values["disable_notification"] = "1"
    }

    return
}
