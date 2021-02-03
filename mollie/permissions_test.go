package mollie

import (
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v2/testdata"
)

func TestPermissionsService_Get(t *testing.T) {
	setup()
	defer teardown()
	id := string(PaymentsRead)
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/permissions/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetPermissionsResponse))
	})

	p, err := tClient.Permissions.Get(id)
	if err != nil {
		t.Error(err)
	}

	if p.ID != PaymentsRead {
		t.Errorf("the response content doesn't match expectations")
	}
}

func TestPermissionsService_List(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/permissions", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListPermissionsResponse))
	})

	pl, err := tClient.Permissions.List()
	if err != nil {
		t.Error(err)
	}

	if pl.Count != 15 {
		t.Errorf("the response content doesn't match expectations")
	}
}
