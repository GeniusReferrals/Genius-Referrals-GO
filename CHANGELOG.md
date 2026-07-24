# Changelog

## 2.0.0 - 2026-07-24

### Breaking Changes
- Migrated from GOPATH-style to Go modules (requires Go 1.21+)
- Replaced deprecated `github.com/apimatic/unirest-go` with standard `net/http`
- Removed global `Config` struct; authentication now passed via `NewClient(apiToken)`
- All methods are now receiver methods on `*Client` instead of separate controller packages
- Response types changed from `interface{}` to `map[string]interface{}` for consistency

### Added
- Full API coverage (~80 endpoints across 13 resources + utilities)
- New resources: Products, Tags, Vouchers, Widgets Packages
- New endpoints: Account CRUD (POST/PUT/DELETE), Bonus PATCH, Bonus Tags, Reports (revenue, 1099), Utilities payment-methods, Voucher denominations, Widget translations
- Table-driven unit tests with `httptest` mock server (90+ tests)
- Docker and compose.yaml for development
- Live smoke test suite

### Removed
- Dependency on `github.com/apimatic/unirest-go`
- Legacy APIMATIC v2.0 generated code and package structure
