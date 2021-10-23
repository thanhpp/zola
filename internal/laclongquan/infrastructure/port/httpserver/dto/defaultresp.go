package dto

type DefaultResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewDefaultResp(code int, message string, data interface{}) *DefaultResp {
	return &DefaultResp{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (resp *DefaultResp) SetCode(code int) {
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

func (resp *DefaultResp) SetData(data interface{}) {
	if resp == nil || data == nil {
		return
	}

	resp.Data = data
}
