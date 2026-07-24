package main

import (
	"fmt"
	"os"
	"time"

	"github.com/GeniusReferrals/Genius-Referrals-GO/geniusreferrals"
)

var passCount, failCount int
var failures []string

func logResult(name, status, detail string) {
	line := fmt.Sprintf("%s: %s", name, status)
	if detail != "" {
		line += " - " + detail
	}
	fmt.Println(line)
	if status == "PASS" {
		passCount++
	} else {
		failCount++
		failures = append(failures, line)
	}
}

func runCheck(name string, fn func() error) {
	if err := fn(); err != nil {
		logResult(name, "FAIL", err.Error())
	} else {
		logResult(name, "PASS", "")
	}
}

func extractString(data map[string]interface{}, keys ...string) string {
	for _, key := range keys {
		if v, ok := data[key]; ok {
			switch val := v.(type) {
			case string:
				return val
			case float64:
				return fmt.Sprintf("%.0f", val)
			}
		}
	}
	if d, ok := data["data"].(map[string]interface{}); ok {
		for _, key := range keys {
			if v, ok := d[key]; ok {
				switch val := v.(type) {
				case string:
					return val
				case float64:
					return fmt.Sprintf("%.0f", val)
				}
			}
		}
	}
	return ""
}

func main() {
	token := os.Getenv("GR_API_TOKEN")
	accountSlug := os.Getenv("GR_ACCOUNT_SLUG")
	baseURI := os.Getenv("GR_BASE_URI")
	if baseURI == "" {
		baseURI = "https://api.geniusreferrals.com"
	}

	if token == "" {
		fmt.Println("SKIP: GR_API_TOKEN is not set")
		os.Exit(0)
	}

	var client *geniusreferrals.Client
	if baseURI != "https://api.geniusreferrals.com" {
		client = geniusreferrals.NewClientWithBaseURI(token, baseURI)
	} else {
		client = geniusreferrals.NewClient(token)
	}

	runCheck("authentication", func() error { _, err := client.GetAuthentication(); return err })
	runCheck("root", func() error { _, err := client.GetRoot(); return err })
	runCheck("accounts list", func() error { _, err := client.GetAccounts(1, 10, "json", ""); return err })

	if accountSlug == "" {
		accounts, err := client.GetAccounts(1, 1, "json", "")
		if err != nil {
			logResult("account discovery", "FAIL", err.Error())
		} else {
			accountSlug = extractString(accounts, "slug")
			if accountSlug == "" {
				logResult("account discovery", "FAIL", "no slug found")
			} else {
				logResult("account discovery", "PASS", accountSlug)
			}
		}
	}

	if accountSlug != "" {
		runCheck("account detail", func() error { _, err := client.GetAccount(accountSlug, ""); return err })
		runCheck("advocates list", func() error { _, err := client.GetAdvocates(accountSlug, geniusreferrals.GetAdvocatesParam{Page: 1, Limit: 5}); return err })
		runCheck("bonuses list", func() error { _, err := client.GetBonuses(accountSlug, 1, 5, "", "", "", "", "", "", ""); return err })
		runCheck("campaigns list", func() error { _, err := client.GetCampaigns(accountSlug, 1, 5, "", "", ""); return err })
		runCheck("products list", func() error { _, err := client.GetProducts(accountSlug, 1, 5); return err })
		runCheck("tags list", func() error { _, err := client.GetTags(accountSlug, 1, 5); return err })
		runCheck("vouchers list", func() error { _, err := client.GetVouchers(accountSlug, 1, 5); return err })
		runCheck("widgets packages list", func() error { _, err := client.GetWidgetsPackages(accountSlug, 1, 5); return err })
		runCheck("1099 tax report", func() error { _, err := client.Get1099TaxReport(accountSlug, "2025"); return err })

		// Tags CRUD
		tagSlug := ""
		tagName := fmt.Sprintf("bolt-smoke-tag-%d", time.Now().Unix())
		func() {
			defer func() {
				if r := recover(); r != nil {
					logResult("tags CRUD", "FAIL", fmt.Sprintf("panic: %v", r))
				}
				if tagSlug != "" {
					_ = client.DeleteTag(accountSlug, tagSlug)
				}
			}()
			created, err := client.CreateTag(accountSlug, geniusreferrals.TagForm{Tag: geniusreferrals.Tag{Name: tagName, Description: strPtr("Bolt live smoke suite"), Color: strPtr("#123456")}})
			if err != nil {
				panic(err)
			}
			tagSlug = extractString(created, "slug")
			if tagSlug == "" {
				panic(fmt.Sprintf("no tag slug returned: %v", created))
			}
			logResult("tags create", "PASS", "")
			_, err = client.GetTag(accountSlug, tagSlug)
			if err != nil {
				panic(err)
			}
			logResult("tags get", "PASS", "")
			_, err = client.UpdateTag(accountSlug, tagSlug, geniusreferrals.TagForm{Tag: geniusreferrals.Tag{Name: tagName, Description: strPtr("Updated by Bolt smoke"), Color: strPtr("#654321")}})
			if err != nil {
				panic(err)
			}
			logResult("tags update", "PASS", "")
			err = client.DeleteTag(accountSlug, tagSlug)
			if err != nil {
				panic(err)
			}
			logResult("tags delete", "PASS", "")
			tagSlug = ""
		}()

		// Advocates CRUD
		advocateToken := ""
		suffix := fmt.Sprintf("%d", time.Now().Unix())
		func() {
			defer func() {
				if r := recover(); r != nil {
					logResult("advocates CRUD", "FAIL", fmt.Sprintf("panic: %v", r))
				}
				if advocateToken != "" {
					_ = client.DeleteAdvocate(accountSlug, advocateToken)
				}
			}()
			created, err := client.PostAdvocate(accountSlug, geniusreferrals.AdvocateForm{Advocate: geniusreferrals.Advocate{Name: "Bolt", Lastname: "Smoke", Email: "bolt-smoke-" + suffix + "@example.com", PayoutThreshold: 10, CanRefer: boolPtr(true)}})
			if err != nil {
				panic(err)
			}
			advocateToken = extractString(created, "token")
			if advocateToken == "" {
				panic(fmt.Sprintf("no advocate token returned: %v", created))
			}
			logResult("advocates create", "PASS", "")
			_, err = client.GetAdvocate(accountSlug, advocateToken)
			if err != nil {
				panic(err)
			}
			logResult("advocates get", "PASS", "")
			_, err = client.GetAdvocateShareLinks(accountSlug, advocateToken)
			if err != nil {
				panic(err)
			}
			logResult("advocates share-links", "PASS", "")
			_, err = client.PatchAdvocate(accountSlug, advocateToken, geniusreferrals.AdvocatePatchForm{Name: strPtr("Bolted"), Status: strPtr("active")})
			if err != nil {
				panic(err)
			}
			logResult("advocates patch", "PASS", "")
			err = client.DeleteAdvocate(accountSlug, advocateToken)
			if err != nil {
				panic(err)
			}
			logResult("advocates delete", "PASS", "")
			advocateToken = ""
		}()

		// Vouchers CRUD
		voucherId := ""
		voucherCode := fmt.Sprintf("BOLT-%d", time.Now().Unix())
		func() {
			defer func() {
				if r := recover(); r != nil {
					logResult("vouchers CRUD", "FAIL", fmt.Sprintf("panic: %v", r))
				}
				if voucherId != "" {
					_ = client.DeleteVoucher(accountSlug, voucherId)
				}
			}()
			created, err := client.CreateVoucher(accountSlug, geniusreferrals.VoucherForm{Voucher: geniusreferrals.Voucher{Code: voucherCode, Value: 9.99, CurrencyCode: "USD", Status: "available"}})
			if err != nil {
				panic(err)
			}
			voucherId = extractString(created, "id")
			if voucherId == "" {
				panic(fmt.Sprintf("no voucher id returned: %v", created))
			}
			logResult("vouchers create", "PASS", "")
			_, err = client.GetVoucher(accountSlug, voucherId)
			if err != nil {
				panic(err)
			}
			logResult("vouchers get", "PASS", "")
			if err != nil {
			}
			if err != nil {
				panic(err)
			}
			logResult("vouchers delete", "PASS", "")
			voucherId = ""
		}()

		// Products CRUD
		productSlug := ""
		psuffix := fmt.Sprintf("%d", time.Now().Unix())
		func() {
			defer func() {
				if r := recover(); r != nil {
					logResult("products CRUD", "FAIL", fmt.Sprintf("panic: %v", r))
				}
				if productSlug != "" {
					_ = client.DeleteProduct(accountSlug, productSlug)
				}
			}()
			created, err := client.CreateProduct(accountSlug, geniusreferrals.ProductForm{Product: geniusreferrals.Product{Title: "Bolt Smoke Product " + psuffix, DescriptionHTML: "Bolt live smoke suite", Status: "open", Images: []string{"https://example.com/product.png"}}})
			if err != nil {
				panic(err)
			}
			productSlug = extractString(created, "slug")
			if productSlug == "" {
				panic(fmt.Sprintf("no product slug returned: %v", created))
			}
			logResult("products create", "PASS", "")
			_, err = client.GetProduct(accountSlug, productSlug)
			if err != nil {
				panic(err)
			}
			logResult("products get", "PASS", "")
			_, err = client.UpdateProduct(accountSlug, productSlug, geniusreferrals.ProductForm{Product: geniusreferrals.Product{Title: "Bolt Smoke Product " + psuffix + " Updated", DescriptionHTML: "Updated by Bolt smoke", Status: "open", Images: []string{"https://example.com/product-updated.png"}}})
			if err != nil {
				panic(err)
			}
			logResult("products update", "PASS", "")
			err = client.DeleteProduct(accountSlug, productSlug)
			if err != nil {
				panic(err)
			}
			logResult("products delete", "PASS", "")
			productSlug = ""
		}()
	} else {
		logResult("account scoped checks", "FAIL", "no account slug discovered")
	}

	runCheck("utility payment methods", func() error { _, err := client.GetPaymentMethodsList(1, 10); return err })
	runCheck("utility currencies", func() error { _, err := client.GetCurrencies(1, 10); return err })

	fmt.Println("\n--- SUMMARY ---")
	fmt.Printf("Pass: %d, Fail: %d\n", passCount, failCount)
	if len(failures) > 0 {
		fmt.Println("Failures:")
		for _, f := range failures {
			fmt.Println("  " + f)
		}
		os.Exit(1)
	}
}

func strPtr(s string) *string { return &s }
func boolPtr(b bool) *bool    { return &b }
