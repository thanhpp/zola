package responsevalue

import "strconv"

type ResponseValue struct {
	Code    string
	Message string
}

var (
	ValueOK = ResponseValue{
		Code:    strconv.Itoa(CodeOK),
		Message: "OK",
	}
	ValueDBConnectErr = ResponseValue{
		Code:    strconv.Itoa(CodeDBConnectErr),
		Message: "DBConnectErr",
	}
	ValueNotEnoughParameter = ResponseValue{
		Code:    strconv.Itoa(CodeNotEnoughParameter),
		Message: "NotEnoughParameter",
	}
	ValueInvalidParameterType = ResponseValue{
		Code:    strconv.Itoa(CodeInvalidParameterType),
		Message: "InvalidParameterType",
	}
	ValueInvalidParameterValue = ResponseValue{
		Code:    strconv.Itoa(CodeInvalidParameterValue),
		Message: "InvalidParameterValue",
	}
	ValueUnknownError = ResponseValue{
		Code:    strconv.Itoa(CodeUnknownError),
		Message: "UnknownError",
	}
	ValueFileTooBig = ResponseValue{
		Code:    strconv.Itoa(CodeFileTooBig),
		Message: "FileTooBig",
	}
	ValueUploadFailed = ResponseValue{
		Code:    strconv.Itoa(CodeUploadFailed),
		Message: "UploadFailed",
	}
	ValueMaxImagesReached = ResponseValue{
		Code:    strconv.Itoa(CodeMaxImagesReached),
		Message: "MaxImagesReached",
	}
	ValueInvalidAccess = ResponseValue{
		Code:    strconv.Itoa(CodeInvalidAccess),
		Message: "InvalidAccess",
	}
	ValueActionHasBeenDone = ResponseValue{
		Code:    strconv.Itoa(CodeActionHasBeenDone),
		Message: "ActionHasBeenDone",
	}
)

var (
	ValuePostNotExist = ResponseValue{
		Code:    strconv.Itoa(CodePostNotExist),
		Message: "PostNotExist",
	}
	ValueIncorrectVerifyCode = ResponseValue{
		Code:    strconv.Itoa(CodeIncorrectVerifyCode),
		Message: "IncorrectVerifyCode",
	}
	ValueNoData = ResponseValue{
		Code:    strconv.Itoa(CodeNoData),
		Message: "NoData",
	}
	ValueInvalidateUser = ResponseValue{
		Code:    strconv.Itoa(CodeInvalidateUser),
		Message: "InvalidateUser",
	}
	ValueUserExisted = ResponseValue{
		Code:    strconv.Itoa(CodeUserExisted),
		Message: "UserExisted",
	}
	ValueInvalidMethod = ResponseValue{
		Code:    strconv.Itoa(CodeInvalidMethod),
		Message: "InvalidMethod",
	}
	ValueInvalidToken = ResponseValue{
		Code:    strconv.Itoa(CodeInvalidToken),
		Message: "InvalidToken",
	}
	ValueExceptionError = ResponseValue{
		Code:    strconv.Itoa(CodeExceptionError),
		Message: "ExceptionError",
	}
)
