package dto

// Client
const (
	MsgEventJoin          = "joinchat"
	MsgEventReconnect     = "reconnecting"
	MsgEventAvaliable     = "avaliable"
	MsgEventDisconnect    = "disconnect"
	MsgEventDeleteMessage = "deletemessage"
	MsgEventSend          = "send"
)

// Server
const (
	MsgEventConnectionTimeout = "connection_timeout"
	MsgEventConnectionError   = "connection_error"
	MsgEventReconnectAttemp   = "reconnect_attemp"
	MsgEventOnMessage         = "onmessage"
)
