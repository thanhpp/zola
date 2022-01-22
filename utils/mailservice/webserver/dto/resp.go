package dto

type RespErr struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (rE *RespErr) SetCode(code int) {
	rE.Error.Code = code
}

func (rE *RespErr) SetMessage(message string) {
	rE.Error.Message = message
}

func (rE *RespErr) SetCodeMessage(code int, msg string) {
	rE.Error.Code = code
	rE.Error.Message = msg
}

type Resp struct {
	RespErr
	Data interface{} `json:"data,omitempty"`
}

func (r *Resp) SetData(data interface{}) {
	r.Data = data
}
