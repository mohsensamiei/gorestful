package restful

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
	"strings"
)

const (
	HeadSchemaBearer  = "Bearer"
	HeadAuthorization = "Authorization"
	HeadContentType   = "Content-Type"
	HeadAccept        = "Accept"
	MimeXml           = "application/xml"
	MimeJson          = "application/json"
)

func ParseBody(req *http.Request, model interface{}) error {
	switch value := req.Header.Get(HeadContentType); value {
	case MimeXml:
		if err := xml.NewDecoder(req.Body).Decode(model); err != nil {
			return errors.Wrapf(err, "cannot parse body from %s to %T", MimeXml, model)
		}
	case MimeJson:
		if err := json.NewDecoder(req.Body).Decode(model); err != nil {
			return errors.Wrapf(err, "cannot parse body from %s to %T", MimeJson, model)
		}
	default:
		return ErrUnsupportedMediaType
	}
	return nil
}

func SetContext(req *http.Request, key string, model interface{}) {
	*req = *req.WithContext(context.WithValue(req.Context(), key, model))
}
func GetContext(req *http.Request, key string, model interface{}) error {
	reflect.ValueOf(model).Elem().Set(reflect.ValueOf(req.Context().Value(key)))
	if model == nil {
		return ErrValueDoesNotExist
	}
	return nil
}

func SetClaims(req *http.Request, model interface{}) {
	SetContext(req, "claims", model)
}
func GetClaims(req *http.Request, model interface{}) error {
	if err := GetContext(req, "claims", model); err != nil {
		return errors.Wrap(err, "cannot get claims from context")
	}
	return nil
}

func GetBearerToken(req *http.Request) (token string, err error) {
	dump := strings.Split(req.Header.Get(HeadAuthorization), HeadSchemaBearer)
	if len(dump) != 2 {
		err = ErrUnauthorized
	} else if token = strings.TrimSpace(dump[1]); len(token) < 1 {
		err = ErrUnauthorized
	}
	return
}

func GetAccept(request *http.Request) (string, error) {
	for _, item := range strings.FieldsFunc(request.Header.Get(HeadAccept), func(r rune) bool {
		return r == ',' || r == ';'
	}) {
		switch mime := strings.TrimSpace(item); mime {
		case MimeJson:
			return mime, nil
		case MimeXml:
			return mime, nil
		}
	}
	return "", ErrNotAcceptable
}
