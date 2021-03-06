package rs

import (
	"time"
	"strings"
	"strconv"
	"encoding/json"
	"github.com/qiniu/api/url"
	"github.com/qiniu/api/auth/digest"
)

// ----------------------------------------------------------

type GetPolicy struct {
	Expires	uint32
}

func (r GetPolicy) MakeRequest(baseUrl string, mac *digest.Mac) (privateUrl string) {

	if r.Expires == 0 {
		r.Expires = 3600
	}
	deadline := time.Now().Unix() + int64(r.Expires)

	if strings.Contains(baseUrl, "?") {
		baseUrl += "&e="
	} else {
		baseUrl += "?e="
	}
	baseUrl += strconv.FormatInt(deadline, 10)

	token := digest.Sign(mac, []byte(baseUrl))
	return baseUrl + "&token=" + token
}

func MakeBaseUrl(domain, key string) (baseUrl string) {

	return "http://" + domain + "/" + url.Escape(key)
}

// --------------------------------------------------------------------------------

type PutPolicy struct {
	Scope            string `json:"scope,omitempty"`
	CallbackUrl      string `json:"callbackUrl,omitempty"`
	CallbackBody     string `json:"callbackBody,omitempty"`
	ReturnUrl        string `json:"returnUrl,omitempty"`
	ReturnBody       string `json:"returnBody,omitempty"`
	AsyncOps         string `json:"asyncOps,omitempty"`
	EndUser          string `json:"endUser,omitempty"`
	Expires          uint32 `json:"deadline"` 			// 截止时间（以秒为单位）
}

func (r *PutPolicy) Token(mac *digest.Mac) string {
	if r.Expires == 0 {
		r.Expires = 3600
	}
	r.Expires += uint32(time.Now().Unix())
	b, _ := json.Marshal(&r)
	return digest.SignWithData(mac, b)
}

// ----------------------------------------------------------

