package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	api "github.com/relationskat/helpdesk-go/src/internal/api/http"
	"github.com/relationskat/helpdesk-go/src/internal/api/http/dto"
	"github.com/relationskat/helpdesk-go/src/internal/config"
	"github.com/relationskat/helpdesk-go/src/internal/service"
	"github.com/relationskat/helpdesk-go/src/internal/store"
)

const expectedCreatedID = 1

func newHandler() http.Handler {
	return api.NewServer(api.Deps{
		Config:  &config.Config{},
		Tickets: service.NewTickets(store.NewMemory()),
	}).Handler()
}

func createTicket(h http.Handler, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/tickets", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	h.ServeHTTP(response, req)
	return response
}

// T-1: FR-2 / US-2, successful example from docs/interface-contract.md.
func TestCreateTicketSuccess(t *testing.T) {
	h := newHandler()
	response := createTicket(h, `{"title":"  Не работает VPN  "}`)
	if response.Code != http.StatusCreated {
		t.Fatalf("status = %d, want 201; body = %q", response.Code, response.Body.String())
	}
	if got := response.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("Content-Type = %q, want application/json", got)
	}
	var got dto.TicketResponse
	if err := json.Unmarshal(response.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	want := dto.TicketResponse{ID: expectedCreatedID, Title: "Не работает VPN", Status: "open"}
	if got != want {
		t.Fatalf("ticket = %+v, want %+v", got, want)
	}

	list := httptest.NewRecorder()
	h.ServeHTTP(list, httptest.NewRequest(http.MethodGet, "/tickets", nil))
	if list.Code != http.StatusOK {
		t.Fatalf("list status = %d, want 200", list.Code)
	}
	var items []dto.TicketResponse
	if err := json.Unmarshal(list.Body.Bytes(), &items); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(items, []dto.TicketResponse{want}) {
		t.Fatalf("stored tickets = %+v, want [%+v]", items, want)
	}
}

// T-2: FR-2 / US-2, blank title returns the contracted error without mutation.
func TestCreateTicketBlankTitle(t *testing.T) {
	h := newHandler()
	if response := createTicket(h, `{"title":"Не работает VPN"}`); response.Code != http.StatusCreated {
		t.Fatalf("seed status = %d, want 201", response.Code)
	}
	response := createTicket(h, `{"title":"   "}`)
	if response.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", response.Code)
	}
	if got := response.Header().Get("Content-Type"); got != "text/plain; charset=utf-8" {
		t.Fatalf("Content-Type = %q, want text/plain; charset=utf-8", got)
	}
	const wantError = "expected JSON with non-empty title (max 4096 bytes)\n"
	if got := response.Body.String(); got != wantError {
		t.Fatalf("body = %q, want %q", got, wantError)
	}

	list := httptest.NewRecorder()
	h.ServeHTTP(list, httptest.NewRequest(http.MethodGet, "/tickets", nil))
	const wantList = "[{\"id\":1,\"title\":\"Не работает VPN\",\"status\":\"open\"}]\n"
	if list.Code != http.StatusOK || list.Body.String() != wantList {
		t.Fatalf("after rejection: status=%d body=%q, want 200 and %q", list.Code, list.Body.String(), wantList)
	}
	next := createTicket(h, `{"title":"Вторая заявка"}`)
	var item dto.TicketResponse
	if err := json.Unmarshal(next.Body.Bytes(), &item); err != nil {
		t.Fatal(err)
	}
	if next.Code != http.StatusCreated || item.ID != 2 {
		t.Fatalf("next status=%d id=%d, want 201 and 2 (rejected request must not consume ID)", next.Code, item.ID)
	}
}
