package handlers_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/catninpo/gophi"
	"github.com/catninpo/gophi/handlers"
	"github.com/catninpo/gophi/mock"
)

const (
	Success = "\033[0;32m\u2713\033[0m"
	Fail    = "\033[0;31m\u2717\033[0m"
)

func TestGetUser(t *testing.T) {
	tt := map[string]struct {
		ByIDFn func(id int) (*gophi.User, error)

		RequestMethod string
		RequestURL    string

		ExpectedStatusCode   int
		ExpectedInvokeStatus bool
		ExpectedUser         *gophi.User
	}{
		"existing user is returned correctly": {
			ByIDFn: func(id int) (*gophi.User, error) {
				return &gophi.User{ID: 1, Name: "Gophi"}, nil
			},

			RequestMethod: http.MethodGet,
			RequestURL:    "/user/1",

			ExpectedStatusCode:   http.StatusOK,
			ExpectedInvokeStatus: true,
			ExpectedUser: &gophi.User{
				ID:   1,
				Name: "Gophi",
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			us := mock.UserService{
				ByIDFn: tc.ByIDFn,
			}

			h := handlers.UserHandler{
				UserService: &us,
			}

			r, err := http.NewRequest(tc.RequestMethod, tc.RequestURL, nil)
			if err != nil {
				t.Fatalf("[%s] unable to create new http request: err=%v", Fail, err)
			}
			r.SetPathValue("id", path.Base(tc.RequestURL))

			w := httptest.NewRecorder()

			h.GetUser(w, r)

			resp := w.Result()
			if resp.StatusCode != tc.ExpectedStatusCode {
				t.Fatalf("[%s] status codes do not match: got=%d want=%d", Fail,
					resp.StatusCode, tc.ExpectedStatusCode)
			}
			t.Logf("[%s] correct status code returned: code=%d", Success,
				tc.ExpectedStatusCode)

			if us.ByIDInvoked != tc.ExpectedInvokeStatus {
				t.Fatalf("[%s] UserByID invoked status was not correct: got=%v want=%v",
					Fail, us.ByIDInvoked, tc.ExpectedInvokeStatus)
			}
			t.Logf("[%s] UserByID invoke status was correct: invoked_status=%v",
				Success, tc.ExpectedInvokeStatus)

			b, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("[%s] unable to read response body from request: err=%v",
					Fail, err)
			}
			defer resp.Body.Close()

			var u gophi.User
			if err := json.Unmarshal(b, &u); err != nil {
				t.Fatalf("[%s] unable to deocde response body to json: err=%v",
					Fail, err)
			}
			t.Logf("[%s] response body decoded to json correctly", Success)

			if u.ID != tc.ExpectedUser.ID {
				t.Fatalf("[%s] id of response user did not match expected: got=%d want=%d",
					Fail, u.ID, tc.ExpectedUser.ID)
			}
			t.Logf("[%s] id of response user correctly matched expected", Success)

			if u.Name != tc.ExpectedUser.Name {
				t.Fatalf("[%s] name of response user did not match expected: got=%s want=%s",
					Fail, u.Name, tc.ExpectedUser.Name)
			}
			t.Logf("[%s] name of response user correctly matched expected", Success)
		})
	}
}
