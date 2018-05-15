package requests

import (
    "errors"
    "strconv"
)

type RestrictChatMember struct {
    ChatID                interface{}
    UserID                int
    UntilDate             int
    CanSendMessages       bool
    CanSendMediaMessages  bool
    CanSendOtherMessages  bool
    CanAddWebPagePreviews bool
}

func (r *RestrictChatMember) IsMultipart() bool {
    return false
}

func (r *RestrictChatMember) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = strconv.FormatInt(r.ChatID.(int64), 10)
    case string:
        values["chat_id"] = r.ChatID.(string)
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["user_id"] = strconv.Itoa(r.UserID)

    if r.UntilDate != 0 {
        values["until_date"] = strconv.Itoa(r.UntilDate)
    }

    if r.CanSendMessages {
        values["can_send_messages"] = "1"
    }

    if r.CanSendMediaMessages {
        values["can_send_media_messages"] = "1"
    }

    if r.CanSendOtherMessages {
        values["can_send_other_messages"] = "1"
    }

    if r.CanAddWebPagePreviews {
        values["can_add_web_page_previews"] = "1"
    }

    return
}
