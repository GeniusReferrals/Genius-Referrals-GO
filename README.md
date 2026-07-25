# Genius Referrals Go SDK

A Go SDK for integrating with the [Genius Referrals](https://www.geniusreferrals.com) RESTful API. This SDK provides convenient access to all API resources including accounts, advocates, bonuses, campaigns, products, referrals, redemption requests, reports, tags, vouchers, widget packages, and utilities.

## Requirements

- Go 1.21 or higher

## Installation

```bash
go get github.com/GeniusReferrals/Genius-Referrals-GO
```

## Authentication

Get your API token from the [Genius Referrals settings page](https://www.geniusreferrals.com/en/settings/api-access).

```go
import "github.com/GeniusReferrals/Genius-Referrals-GO/geniusreferrals"

client := geniusreferrals.NewClient("YOUR_API_TOKEN")
```

For a custom base URI:

```go
client := geniusreferrals.NewClientWithBaseURI("YOUR_API_TOKEN", "https://api.geniusreferrals.com")
```

## Usage

### Accounts

```go
accounts, err := client.GetAccounts(1, 10, "json", "")
account, err := client.GetAccount("your-account-slug", "")
newAccount, err := client.CreateAccount(geniusreferrals.AccountForm{Account: geniusreferrals.Account{Name: "My Account"}})
_, err := client.UpdateAccount("your-account-slug", geniusreferrals.AccountForm{Account: geniusreferrals.Account{Name: "Updated"}})
err := client.DeleteAccount("your-account-slug")
```

### Advocates

```go
advocates, err := client.GetAdvocates("your-account-slug", geniusreferrals.GetAdvocatesParam{Page: 1, Limit: 10})
advocate, err := client.GetAdvocate("your-account-slug", "advocate-token")
newAdvocate, err := client.PostAdvocate("your-account-slug", geniusreferrals.AdvocateForm{Advocate: geniusreferrals.Advocate{Name: "John", Email: "john@example.com"}})
_, err := client.PatchAdvocate("your-account-slug", "advocate-token", geniusreferrals.AdvocatePatchForm{Name: strPtr("Updated")})
err := client.DeleteAdvocate("your-account-slug", "advocate-token")
```

### Bonuses

```go
bonuses, err := client.GetBonuses("your-account-slug", 1, 10, "", "", "", "", "", "")
_, err := client.PostBonus("your-account-slug", geniusreferrals.BonusesForm{Bonus: geniusreferrals.Bonuses{AdvocateToken: "token", Reference: "ref"}})
_, err := client.PatchBonus("your-account-slug", "bonus-id", geniusreferrals.BonusPatchForm{Status: strPtr("approved")})
tags, err := client.GetBonusTags("your-account-slug", "bonus-id", 1, 10)
```

### Products

```go
products, err := client.GetProducts("your-account-slug", 1, 10)
product, err := client.GetProduct("your-account-slug", "product-slug")
newProduct, err := client.CreateProduct("your-account-slug", geniusreferrals.ProductForm{Product: geniusreferrals.Product{Title: "Widget"}})
```

### Tags

```go
tags, err := client.GetTags("your-account-slug", 1, 10)
tag, err := client.GetTag("your-account-slug", "tag-slug")
newTag, err := client.CreateTag("your-account-slug", geniusreferrals.TagForm{Tag: geniusreferrals.Tag{Name: "VIP"}})
```

### Vouchers

```go
vouchers, err := client.GetVouchers("your-account-slug", 1, 10)
voucher, err := client.GetVoucher("your-account-slug", "voucher-id")
newVoucher, err := client.CreateVoucher("your-account-slug", geniusreferrals.VoucherForm{Voucher: geniusreferrals.Voucher{Code: "CODE1", Value: 10.0, CurrencyCode: "USD", Status: "available"}})
```

### Widgets Packages

```go
packages, err := client.GetWidgetsPackages("your-account-slug", 1, 10)
widgets, err := client.GetWidgets("your-account-slug", "package-slug", 1, 10)
```

### Utilities

```go
methods, err := client.GetPaymentMethodsList(1, 10)
currencies, err := client.GetCurrencies(1, 10)
```

## Docker Development

```bash
docker compose build
docker compose run --rm sdk go test ./... -v
```

## Error Handling

All API errors are returned as `*geniusreferrals.APIError` which includes the HTTP status code and message.

```go
result, err := client.GetAccount("bad-slug", "")
if err != nil {
	if apiErr, ok := err.(*geniusreferrals.APIError); ok {
		fmt.Printf("API error %d: %s\n", apiErr.StatusCode, apiErr.Message)
	}
}
```

## License

MIT License. See LICENSE for details.
