package violetlog

import (
	"encoding/json"
	"net/http"

	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

func (v VioletLog) Custom(level, message, stackTrace string) (Log, error) {
	requestData := LogRequest{
		ErrorLevel: level,
		Message:    message,
		StackTrace: stackTrace,
	}

	request := v.client.Request()
	request.Method(http.MethodPost)
	request.AddPath("/errors")
	request.Use(body.JSON(requestData))

	response, err := request.Send()
	if err != nil {
		return Log{}, err
	}

	log := Log{}

	err = json.Unmarshal(response.Bytes(), &log)
	if err != nil {
		return Log{}, err
	}

	return log, nil
}

func (v VioletLog) Several(message, stackTrace string) (Log, error) {
	return v.Custom("several", message, stackTrace)
}

func (v VioletLog) Error(message, stackTrace string) (Log, error) {
	return v.Custom("error", message, stackTrace)
}

func (v VioletLog) Warning(message, stackTrace string) (Log, error) {
	return v.Custom("warning", message, stackTrace)
}

func (v VioletLog) Info(message, stackTrace string) (Log, error) {
	return v.Custom("info", message, stackTrace)
}

func (v VioletLog) Debug(message, stackTrace string) (Log, error) {
	return v.Custom("debug", message, stackTrace)
}

func (v VioletLog) Verbose(message, stackTrace string) (Log, error) {
	return v.Custom("verbose", message, stackTrace)
}
