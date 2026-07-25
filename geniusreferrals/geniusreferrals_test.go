package geniusreferrals

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func newTestClient(serverURL string) *Client {
	c := NewClientWithBaseURI("test-token", serverURL)
	return c
}

func newMockServer(handler http.HandlerFunc) (*httptest.Server, *Client) {
	srv := httptest.NewServer(handler)
	return srv, newTestClient(srv.URL)
}

// === Accounts ===
func TestGetAccounts(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts" {
			http.Error(w, "bad request", 400)
			return
		}
		if r.Header.Get("X-Auth-Token") != "test-token" {
			http.Error(w, "unauthorized", 401)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetAccounts(1, 10, "json", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetAccount(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"slug": "test1000"}})
	})
	defer srv.Close()

	res, err := c.GetAccount("test1000", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestCreateAccount(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"slug": "new"}})
	})
	defer srv.Close()

	res, err := c.CreateAccount(AccountForm{Account: Account{Name: "New"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestUpdateAccount(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" || r.URL.Path != "/accounts/test1000" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.UpdateAccount("test1000", AccountForm{Account: Account{Name: "Updated"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil result for 204")
	}
}

func TestDeleteAccount(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/accounts/test1000" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	if err := c.DeleteAccount("test1000"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAccountAuthFailure(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]interface{}{"code": 401, "message": "Unauthorized"})
	})
	defer srv.Close()

	_, err := c.GetAccount("test1000", "")
	if err == nil {
		t.Fatal("expected error for 401")
	}
	if apiErr, ok := err.(*APIError); !ok || apiErr.StatusCode != 401 {
		t.Fatalf("expected APIError 401, got %T %v", err, err)
	}
}

// === Advocates ===
func TestGetAdvocates(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/advocates" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetAdvocates("test1000", GetAdvocatesParam{Page: 1, Limit: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetAdvocate(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/advocates/abc" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"token": "abc"}})
	})
	defer srv.Close()

	res, err := c.GetAdvocate("test1000", "abc")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestPostAdvocate(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts/test1000/advocates" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"token": "new"}})
	})
	defer srv.Close()

	res, err := c.PostAdvocate("test1000", AdvocateForm{Advocate: Advocate{Name: "John", Email: "john@example.com"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestPatchAdvocate(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" || r.URL.Path != "/accounts/test1000/advocates/abc" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.PatchAdvocate("test1000", "abc", AdvocatePatchForm{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil for 204")
	}
}

func TestDeleteAdvocate(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/accounts/test1000/advocates/abc" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	if err := c.DeleteAdvocate("test1000", "abc"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDeleteAdvocates(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/accounts/test1000/advocates" {
			http.Error(w, "bad request", 400)
			return
		}
		advocates := r.URL.Query().Get("advocates")
		if !strings.Contains(advocates, "abc") {
			http.Error(w, "missing advocates param", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	if err := c.DeleteAdvocates("test1000", []string{"abc", "def"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetAdvocateShareLinks(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/advocates/abc/share-links" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": []string{}})
	})
	defer srv.Close()

	res, err := c.GetAdvocateShareLinks("test1000", "abc")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetAdvocatePaymentMethods(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/advocates/abc/payment-methods" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetAdvocatePaymentMethods("test1000", "abc", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestPostAdvocatePaymentMethod(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts/test1000/advocates/abc/payment-methods" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"id": "pm1"}})
	})
	defer srv.Close()

	res, err := c.PostAdvocatePaymentMethod("test1000", "abc", PaymentMethodForm{AdvocatePaymentMethod: PaymentMethod{Username: "user", Description: "desc", IsActive: true}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

// === Authentications ===
func TestGetAuthentication(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/test-authentication" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"code": 200, "message": "Hello!"})
	})
	defer srv.Close()

	res, err := c.GetAuthentication()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

// === Roots ===
func TestGetRoot(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"_links": map[string]interface{}{"self": map[string]string{"href": "/"}}})
	})
	defer srv.Close()

	res, err := c.GetRoot()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

// === Bonuses ===
func TestGetBonuses(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/bonuses" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetBonuses("test1000", 1, 10, "", "", "", "", "", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestPostBonus(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts/test1000/bonuses" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"id": 1}})
	})
	defer srv.Close()

	res, err := c.PostBonus("test1000", BonusesForm{Bonus: Bonuses{AdvocateToken: "abc", Reference: "ref1"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestForceBonus(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts/test1000/bonuses/force" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"id": 2}})
	})
	defer srv.Close()

	res, err := c.ForceBonus("test1000", ForceBonusesForm{Bonus: ForceBonuses{AdvocateToken: "abc", Reference: "ref2", BonusAmount: 100}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestDeleteBonus(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/accounts/test1000/bonuses/1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	if err := c.DeleteBonus("test1000", "1"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheckupBonuses(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/bonuses/checkup" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": []string{}})
	})
	defer srv.Close()

	res, err := c.CheckupBonuses("test1000")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetBonusTraces(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/bonuses/traces" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetBonusTraces("test1000", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetBonusTrace(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/bonuses/traces/t1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"id": "t1"}})
	})
	defer srv.Close()

	res, err := c.GetBonusTrace("test1000", "t1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestPatchBonus(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" || r.URL.Path != "/accounts/test1000/bonuses/1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.PatchBonus("test1000", "1", BonusPatchForm{Status: strPtr("approved")})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil for 204")
	}
}

func TestGetBonusTags(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/bonuses/1/tags" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetBonusTags("test1000", "1", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestAddBonusTag(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts/test1000/bonuses/1/tags" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"id": 1}})
	})
	defer srv.Close()

	res, err := c.AddBonusTag("test1000", "1", TagForm{Tag: Tag{Name: "vip"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestRemoveBonusTag(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/accounts/test1000/bonuses/1/tags/vip" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	if err := c.RemoveBonusTag("test1000", "1", "vip"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// === Campaigns ===
func TestGetCampaigns(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/campaigns" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetCampaigns("test1000", 1, 10, "", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetCampaign(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/campaigns/default" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"slug": "default"}})
	})
	defer srv.Close()

	res, err := c.GetCampaign("test1000", "default")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

// === Products ===
func TestGetProducts(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/products" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetProducts("test1000", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetProduct(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/products/widget" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"slug": "widget"}})
	})
	defer srv.Close()

	res, err := c.GetProduct("test1000", "widget")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestCreateProduct(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts/test1000/products" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"slug": "new"}})
	})
	defer srv.Close()

	res, err := c.CreateProduct("test1000", ProductForm{Product: Product{Title: "Widget"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestUpdateProduct(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" || r.URL.Path != "/accounts/test1000/products/widget" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.UpdateProduct("test1000", "widget", ProductForm{Product: Product{Title: "Updated"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil for 204")
	}
}

func TestDeleteProduct(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/accounts/test1000/products/widget" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	if err := c.DeleteProduct("test1000", "widget"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetProductVariants(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/products/widget/variants" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetProductVariants("test1000", "widget", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestCreateProductVariant(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts/test1000/products/widget/variants" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"id": "v1"}})
	})
	defer srv.Close()

	res, err := c.CreateProductVariant("test1000", "widget", ProductVariantForm{Variant: ProductVariant{Title: "V1"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestPatchProductVariant(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" || r.URL.Path != "/accounts/test1000/products/widget/variants/v1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.PatchProductVariant("test1000", "widget", "v1", ProductVariantForm{Variant: ProductVariant{Title: "V1U"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil for 204")
	}
}

func TestDeleteProductVariant(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/accounts/test1000/products/widget/variants/v1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	if err := c.DeleteProductVariant("test1000", "widget", "v1"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPatchProductVariantPrice(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" || r.URL.Path != "/accounts/test1000/products/widget/variants/v1/prices/USD" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.PatchProductVariantPrice("test1000", "widget", "v1", "USD", ProductVariantPriceForm{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil for 204")
	}
}

// === Referrals ===
func TestGetReferrals(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/advocates/abc/referrals" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetReferrals("test1000", "abc", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetReferral(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/advocates/abc/referrals/r1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"id": "r1"}})
	})
	defer srv.Close()

	res, err := c.GetReferral("test1000", "abc", "r1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestPostReferral(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts/test1000/advocates/abc/referrals" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"id": "new"}})
	})
	defer srv.Close()

	res, err := c.PostReferral("test1000", "abc", ReferralForm{Referral: Referral{ReferredAdvocateToken: "def", ReferralOriginSlug: "web", CampaignSlug: "default", HTTPReferer: "https://example.com"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestPutReferral(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" || r.URL.Path != "/accounts/test1000/advocates/abc/referrals/r1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.PutReferral("test1000", "abc", "r1", ReferralForm{Referral: Referral{ReferredAdvocateToken: "def", ReferralOriginSlug: "web", CampaignSlug: "default", HTTPReferer: "https://example.com"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil for 204")
	}
}

func TestDeleteReferral(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/accounts/test1000/advocates/abc/referrals/r1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	if err := c.DeleteReferral("test1000", "abc", "r1"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// === Redemption Requests ===
func TestGetRedemptionRequests(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/redemption-requests" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetRedemptionRequests("test1000", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetRedemptionRequest(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/redemption-requests/rr1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"id": "rr1"}})
	})
	defer srv.Close()

	res, err := c.GetRedemptionRequest("test1000", "rr1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestPostRedemptionRequest(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts/test1000/redemption-requests" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"id": "new"}})
	})
	defer srv.Close()

	res, err := c.PostRedemptionRequest("test1000", RedemptionRequestForm{RedemptionRequest: RedemptionRequest{AdvocateToken: "abc", RequestStatusSlug: "pending", RequestActionSlug: "paypal", CurrencyCode: "USD", Amount: 100}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestPatchRedemptionRequest(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" || r.URL.Path != "/accounts/test1000/redemption-requests/rr1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.PatchRedemptionRequest("test1000", "rr1", RedemptionRequestForm{RedemptionRequest: RedemptionRequest{AdvocateToken: "abc", RequestStatusSlug: "approved", RequestActionSlug: "paypal", CurrencyCode: "USD", Amount: 100}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil for 204")
	}
}

func TestRedeemRedemptionRequest(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" || r.URL.Path != "/accounts/test1000/redemption-requests/rr1/redemption" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.RedeemRedemptionRequest("test1000", "rr1", map[string]interface{}{"status": "redeemed"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil for 204")
	}
}

// === Reports ===
func TestGetBonusesDailyGiven(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/reports/bonuses-daily-given" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": []string{}})
	})
	defer srv.Close()

	res, err := c.GetBonusesDailyGiven("test1000", "2025-01-01", "2025-01-31", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGet1099TaxReport(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/reports/1099-tax-report" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": []string{}})
	})
	defer srv.Close()

	res, err := c.Get1099TaxReport("test1000", "2025")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetRevenue(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/reports/revenue" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(map[string]interface{}{"code": 404, "message": "Not found"})
	})
	defer srv.Close()

	_, err := c.GetRevenue("test1000", "2025-01-01", "2025-01-31")
	if err == nil {
		t.Fatal("expected error for 404")
	}
}

// === Tags ===
func TestGetTags(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/tags" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetTags("test1000", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetTag(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/tags/vip" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"slug": "vip"}})
	})
	defer srv.Close()

	res, err := c.GetTag("test1000", "vip")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestCreateTag(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts/test1000/tags" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"slug": "new"}})
	})
	defer srv.Close()

	res, err := c.CreateTag("test1000", TagForm{Tag: Tag{Name: "New"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestUpdateTag(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" || r.URL.Path != "/accounts/test1000/tags/vip" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.UpdateTag("test1000", "vip", TagForm{Tag: Tag{Name: "VIP"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil for 204")
	}
}

func TestDeleteTag(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/accounts/test1000/tags/vip" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	if err := c.DeleteTag("test1000", "vip"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// === Vouchers ===
func TestGetVouchers(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/vouchers" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetVouchers("test1000", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetVoucher(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/vouchers/v1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"id": "v1"}})
	})
	defer srv.Close()

	res, err := c.GetVoucher("test1000", "v1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestCreateVoucher(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts/test1000/vouchers" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"id": "new"}})
	})
	defer srv.Close()

	res, err := c.CreateVoucher("test1000", VoucherForm{Voucher: Voucher{Code: "CODE1", Value: 10.0, CurrencyCode: "USD", Status: "available"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestPatchVoucher(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" || r.URL.Path != "/accounts/test1000/vouchers/v1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.PatchVoucher("test1000", "v1", VoucherForm{Voucher: Voucher{Status: "redeemed"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil for 204")
	}
}

func TestDeleteVoucher(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/accounts/test1000/vouchers/v1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	if err := c.DeleteVoucher("test1000", "v1"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetVoucherDenominations(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/vouchers-denominations/USD" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": []string{}})
	})
	defer srv.Close()

	res, err := c.GetVoucherDenominations("test1000", "USD")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

// === Widgets Packages ===
func TestGetWidgetsPackages(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/widgets-packages" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetWidgetsPackages("test1000", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetWidgetsPackage(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/widgets-packages/wp1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"slug": "wp1"}})
	})
	defer srv.Close()

	res, err := c.GetWidgetsPackage("test1000", "wp1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestCreateWidgetsPackage(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts/test1000/widgets-packages" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"slug": "new"}})
	})
	defer srv.Close()

	res, err := c.CreateWidgetsPackage("test1000", WidgetsPackageForm{WidgetsPackage: WidgetsPackage{Name: "Pack"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestUpdateWidgetsPackage(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" || r.URL.Path != "/accounts/test1000/widgets-packages/wp1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.UpdateWidgetsPackage("test1000", "wp1", WidgetsPackageForm{WidgetsPackage: WidgetsPackage{Name: "Updated"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil for 204")
	}
}

func TestDeleteWidgetsPackage(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" || r.URL.Path != "/accounts/test1000/widgets-packages/wp1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	if err := c.DeleteWidgetsPackage("test1000", "wp1"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetWidgets(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/widgets-packages/wp1/widgets" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetWidgets("test1000", "wp1", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetWidget(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/widgets-packages/wp1/widgets/w1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"id": "w1"}})
	})
	defer srv.Close()

	res, err := c.GetWidget("test1000", "wp1", "w1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestCreateWidget(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts/test1000/widgets-packages/wp1/widgets" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"id": "new"}})
	})
	defer srv.Close()

	res, err := c.CreateWidget("test1000", "wp1", WidgetForm{Widget: Widget{Name: "W1"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestPatchWidget(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" || r.URL.Path != "/accounts/test1000/widgets-packages/wp1/widgets/w1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.PatchWidget("test1000", "wp1", "w1", WidgetForm{Widget: Widget{Name: "W1U"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil for 204")
	}
}

func TestPutWidget(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" || r.URL.Path != "/accounts/test1000/widgets-packages/wp1/widgets/w1" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.PutWidget("test1000", "wp1", "w1", WidgetForm{Widget: Widget{Name: "W1R"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil for 204")
	}
}

func TestGetWidgetTranslations(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/widgets-packages/wp1/widgets/w1/translations" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetWidgetTranslations("test1000", "wp1", "w1", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetWidgetTranslation(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/accounts/test1000/widgets-packages/wp1/widgets/w1/translations/en" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"locale": "en"}})
	})
	defer srv.Close()

	res, err := c.GetWidgetTranslation("test1000", "wp1", "w1", "en")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestCreateWidgetTranslation(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/accounts/test1000/widgets-packages/wp1/widgets/w1/translations" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"locale": "en"}})
	})
	defer srv.Close()

	res, err := c.CreateWidgetTranslation("test1000", "wp1", "w1", WidgetTranslationForm{WidgetTranslation: WidgetTranslation{Locale: "en", Name: "Name"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestPatchWidgetTranslation(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" || r.URL.Path != "/accounts/test1000/widgets-packages/wp1/widgets/w1/translations/en" {
			http.Error(w, "bad request", 400)
			return
		}
		w.WriteHeader(204)
	})
	defer srv.Close()

	res, err := c.PatchWidgetTranslation("test1000", "wp1", "w1", "en", WidgetTranslationForm{WidgetTranslation: WidgetTranslation{Name: "Updated"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatal("expected nil for 204")
	}
}

// === Utilities ===
func TestGetBonusesRedemptionMethods(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/utilities/bonuses-redemption-methods" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetBonusesRedemptionMethods(1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetBonusesRedemptionMethod(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/utilities/bonuses-redemption-methods/paypal" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"slug": "paypal"}})
	})
	defer srv.Close()

	res, err := c.GetBonusesRedemptionMethod("paypal")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetCurrencies(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/utilities/currencies" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetCurrencies(1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetCurrency(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/utilities/currencies/USD" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"code": "USD"}})
	})
	defer srv.Close()

	res, err := c.GetCurrency("USD")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetPaymentMethodsList(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/utilities/payment-methods" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetPaymentMethodsList(1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetRedemptionRequestActions(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/utilities/redemption-request-actions" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetRedemptionRequestActions(1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetRedemptionRequestAction(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/utilities/redemption-request-actions/paypal" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"slug": "paypal"}})
	})
	defer srv.Close()

	res, err := c.GetRedemptionRequestAction("paypal")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetRedemptionRequestStatuses(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/utilities/redemption-request-statuses" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetRedemptionRequestStatuses(1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetRedemptionRequestStatus(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/utilities/redemption-request-statuses/pending" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"slug": "pending"}})
	})
	defer srv.Close()

	res, err := c.GetRedemptionRequestStatus("pending")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetReferralOrigins(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/utilities/referral-origins" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"results": []string{}, "page": 1, "limit": 10, "total": 0}})
	})
	defer srv.Close()

	res, err := c.GetReferralOrigins(1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestGetReferralOrigin(t *testing.T) {
	srv, c := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/utilities/referral-origins/web" {
			http.Error(w, "bad request", 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{"slug": "web"}})
	})
	defer srv.Close()

	res, err := c.GetReferralOrigin("web")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected non-nil result")
	}
}

// === Query Param Helpers ===
func TestSetOpt(t *testing.T) {
	q := url.Values{}
	setOpt(&q, "key", "value")
	if q.Get("key") != "value" {
		t.Fatalf("expected value, got %s", q.Get("key"))
	}
	setOpt(&q, "empty", "")
	if q.Get("empty") != "" {
		t.Fatalf("expected empty, got %s", q.Get("empty"))
	}
}

func TestSetBoolOpt(t *testing.T) {
	q := url.Values{}
	v := true
	setBoolOpt(&q, "active", &v)
	if q.Get("active") != "true" {
		t.Fatalf("expected true, got %s", q.Get("active"))
	}
}

func strPtr(s string) *string {
	return &s
}
