package middleware

type ResponseMsgLang struct {
	Data     map[ResponseCode]string
	Lang     string
	Defacode string
}

func createResponseMsgLang(lang string) *ResponseMsgLang {
	return &ResponseMsgLang{Lang: lang, Data: make(map[ResponseCode]string)}
}

func (r *ResponseMsgLang) setLang(lang string) {
	r.Lang = lang
}

func (r *ResponseMsgLang) setDefacode(defacode string) {
	r.Defacode = defacode
}

func (r *ResponseMsgLang) getMessage(code ResponseCode) string {
	if msg, ok := r.Data[code]; ok {
		return msg
	}
	return r.Defacode
}

func (r *ResponseMsgLang) setMessage(data map[ResponseCode]string) {
	if len(r.Data) == 0 {
		r.Data = data
		return
	}
	for code, msg := range data {
		r.Data[code] = msg
	}
	return
}

func (r *ResponseMsgLang) editMessage(code ResponseCode, msg string) {
	r.Data[code] = msg
}
