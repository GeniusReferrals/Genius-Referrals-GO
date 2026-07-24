# Migration Guide: v1.x to v2.0

## Overview

Version 2.0 is a complete rewrite. The old GOPATH-style, APIMATIC-generated SDK with the deprecated Unirest HTTP library has been replaced by a modern Go module using only the standard library.

## Key Changes

### 1. Import Path

**Before:**
```go
import "geniusreferrals_lib/accounts_pkg"
import "geniusreferrals_lib"
```

**After:**
```go
import "github.com/GeniusReferrals/Genius-Referrals-GO/geniusreferrals"
```

### 2. Authentication

**Before:**
```go
geniusreferrals_lib.Config.XAuthToken = "YOUR_TOKEN"
```

**After:**
```go
client := geniusreferrals.NewClient("YOUR_TOKEN")
```

### 3. Controller Pattern

**Before:**
```go
ctrl := accounts_pkg.NewACCOUNTS()
result, err := ctrl.GetAccount("slug")
```

**After:**
```go
result, err := client.GetAccount("slug", "")
```

### 4. HTTP Library

Removed dependency on `github.com/apimatic/unirest-go`. Now uses `net/http` from the Go standard library.

### 5. Response Types

All methods now return `map[string]interface{}, error` instead of `interface{}, error`.

## New Resources

v2.0 adds full support for:
- Products
- Tags
- Vouchers
- Widgets Packages
- Bonus Tags
- Utilities (payment-methods)
- Reports (revenue, 1099-tax-report)

## Upgrade Steps

1. Update your `go.mod` to require Go 1.21+
2. Replace all imports with the new package path
3. Instantiate a `*geniusreferrals.Client` with `NewClient(token)`
4. Replace controller calls with client method calls
5. Handle responses as `map[string]interface{}` instead of `interface{}`
