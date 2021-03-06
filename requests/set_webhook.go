package requests

import (
    "encoding/json"
    "io"
    "strconv"
)

type SetWebhook struct {
    URL            string
    Certificate    io.Reader
    MaxConnections uint32
    AllowedUpdates []string
}

func (r *SetWebhook) IsMultipart() bool {
    return r.Certificate != nil
}

func (r *SetWebhook) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["url"] = r.URL

    if r.Certificate != nil {
        values["certificate"] = r.Certificate
    }

    if r.MaxConnections != 0 {
        values["max_connections"] = strconv.FormatUint(uint64(r.MaxConnections), 10)
    }

    if r.AllowedUpdates != nil {
        var data []byte
        if data, err = json.Marshal(r.AllowedUpdates); err != nil {
            return
        }

        values["allowed_updates"] = string(data)
    }

    return
}
