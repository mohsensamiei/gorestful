package restful

import (
	"encoding/json"
	"encoding/xml"
	"github.com/pkg/errors"
	"net/http"
)

func SendStruct(response http.ResponseWriter, request *http.Request, model interface{}, code int) {
	mime, err := GetAccept(request)
	if err != nil {
		switch errors.Cause(err) {
		case ErrNotAcceptable:
			http.Error(response, err.Error(), http.StatusNotAcceptable)
		default:
			http.Error(response, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var bytes []byte
	switch mime {
	case MimeXml:
		bytes, err = xml.Marshal(model)
	case MimeJson:
		bytes, err = json.Marshal(model)
	default:
		http.Error(response, ErrNotAcceptable.Error(), http.StatusNotAcceptable)
		return
	}
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(code)
	response.Header().Set(HeadContentType, mime)
	if _, err := response.Write(bytes); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}
