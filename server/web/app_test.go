package web

import (
	"DGraph-Example/model"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockDb struct {
	user []*model.User
	err  error
}

func (m *MockDb) GetUsers() ([]*model.User, error) {
	return m.user, m.err
}

func TestApp_GetUsers(t *testing.T) {
	app := App{d: &MockDb{
		user: []*model.User{
			// {"Tech1", "Details1"},
			// {"Tech2", "Details2"},
		},
	}}

	r, _ := http.NewRequest("GET", "/api/users", nil)
	w := httptest.NewRecorder()

	app.GetUsers(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}

	want := `[{"name":"Tech1","details":"Details1"},{"name":"Tech2","details":"Details2"}]` + "\n"
	if got := w.Body.String(); got != want {
		t.Errorf("handler returned unexpected body: got %v want %v", got, want)
	}
}

func TestApp_GetTechnologies_WithDBError(t *testing.T) {
	app := App{d: &MockDb{
		user: nil,
		err:  errors.New("unknown error"),
	}}

	r, _ := http.NewRequest("GET", "/api/technologies", nil)
	w := httptest.NewRecorder()

	app.GetUsers(w, r)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}
}
