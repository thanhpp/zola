package dto

type DefaultResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (resp *DefaultResp) SetCode(code string) {
	if resp == nil {
		return
	}
	resp.Code = code
}

func (resp *DefaultResp) SetMsg(msg string) {
	if resp == nil {
		return
	}
	resp.Message = msg
}

func boolTranslate(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
