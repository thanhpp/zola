package responsevalue

const (
	CodeOK int = 1000 + iota
	CodeDBConnectErr
	CodeNotEnoughParameter
	CodeInvalidParameterType
	CodeInvalidParameterValue
	CodeUnknownError
	CodeFileTooBig
	CodeUploadFailed
	CodeMaxImagesReached
	CodeInvalidAccess
	CodeActionHasBeenDone
)

const (
	CodePostNotExist int = 9992 + iota
	CodeIncorrectVerifyCode
	CodeNoData
	CodeInvalidateUser
	CodeUserExisted
	CodeInvalidMethod
	CodeInvalidToken
	CodeExceptionError
)
