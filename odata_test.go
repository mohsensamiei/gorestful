package restful

import (
	"fmt"
	"net/http"
	"testing"
)

func TestParseOData(t *testing.T) {
	t.Run("most return error on non int param $top", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/?$top=abc", nil)
		if _, err := ParseOData(request); err == nil {
			t.Errorf("unhandled error: want error got nil")
		}
	})
	t.Run("most return error on non int param $skip", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/?$skip=abc", nil)
		if _, err := ParseOData(request); err == nil {
			t.Errorf("unhandled error: want error got nil")
		}
	})
	t.Run("most return valid odata object", func(t *testing.T) {
		top, skip := 1, 2
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/?$top=%d&$skip=%d", top, skip), nil)

		odata, err := ParseOData(request)
		if err != nil {
			t.Errorf("unexpected error: want nil got %T(%s)", err, err)
		} else {
			if odata.Top != top {
				t.Errorf("unexpected top: want %d got %d", top, odata.Top)
			}
			if odata.Skip != skip {
				t.Errorf("unexpected skip: want %d got %d", skip, odata.Skip)
			}
		}
	})
}
