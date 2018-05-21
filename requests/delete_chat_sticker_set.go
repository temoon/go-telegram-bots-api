package requests

import (
    "errors"
    "strconv"
)

type DeleteChatStickerSet struct {
    ChatID interface{}
}

func (r *DeleteChatStickerSet) IsMultipart() bool {
    return false
}

func (r *DeleteChatStickerSet) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(r.ChatID.(uint64), 10)
    case string:
        values["chat_id"] = r.ChatID.(string)
    default:
        return nil, errors.New("invalid chat_id")
    }

    return
}
