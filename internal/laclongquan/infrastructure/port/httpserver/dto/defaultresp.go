package dto

type DefaultRespWithoutData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (resp *DefaultRespWithoutData) SetCode(code string) {
	if resp == nil {
		return
	}

	resp.Code = code
}

func (resp *DefaultRespWithoutData) SetMsg(msg string) {
	if resp == nil {
		return
	}

	resp.Message = msg
}

type DefaultResp struct {
	DefaultRespWithoutData
	Data interface{} `json:"data"`
}

func (resp *DefaultResp) SetData(data interface{}) {
	if resp == nil || data == nil {
		return
	}

	resp.Data = data
}

func NewDefaultResp(code, message string, data interface{}) *DefaultResp {
	return &DefaultResp{
		DefaultRespWithoutData: DefaultRespWithoutData{
			Code:    code,
			Message: message,
		},
		Data: data,
	}
}
