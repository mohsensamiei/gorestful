package restful

import (
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type OData struct {
	Top  int
	Skip int
}

func ParseOData(req *http.Request) (*OData, error) {
	var err error
	var odata OData
	if top := req.URL.Query().Get("$top"); top != "" {
		odata.Top, err = strconv.Atoi(top)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid odata query param $top=%s", top)
		}
	}
	if skip := req.URL.Query().Get("$skip"); skip != "" {
		odata.Skip, err = strconv.Atoi(skip)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid odata query param $skip=%s", skip)
		}
	}
	return &odata, nil
}
