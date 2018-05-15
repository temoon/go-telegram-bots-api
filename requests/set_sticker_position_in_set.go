package requests

import (
    "strconv"
)

type SetStickerPositionInSet struct {
    Sticker  string
    Position int
}

func (r *SetStickerPositionInSet) IsMultipart() bool {
    return false
}

func (r *SetStickerPositionInSet) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    values["sticker"] = []interface{}{r.Sticker}
    values["position"] = []interface{}{strconv.Itoa(r.Position)}

    return
}