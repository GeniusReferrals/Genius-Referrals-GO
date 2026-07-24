package geniusreferrals

// Advocate represents an advocate.
type Advocate struct {
	Name            string  `json:"name,omitempty"`
	Lastname        string  `json:"lastname,omitempty"`
	Email           string  `json:"email,omitempty"`
	PayoutThreshold int64   `json:"payout_threshold,omitempty"`
	AvatarURL       *string `json:"avatar_url,omitempty"`
	Metadata        *string `json:"metadata,omitempty"`
	CanRefer        *bool   `json:"can_refer,omitempty"`
}

// AdvocateForm wraps an Advocate for create/update.
type AdvocateForm struct {
	Advocate Advocate `json:"advocate"`
}

// AdvocatePatchForm for partial advocate updates.
type AdvocatePatchForm struct {
	Name            *string `json:"name,omitempty"`
	Lastname        *string `json:"lastname,omitempty"`
	Email           *string `json:"email,omitempty"`
	PayoutThreshold *int64  `json:"payout_threshold,omitempty"`
	ClaimedBalance  *int64  `json:"claimed_balance,omitempty"`
	UnclaimedBalance *int64 `json:"unclaimed_balance,omitempty"`
	CurrencyCode    *string `json:"currency_code,omitempty"`
	AvatarURL       *string `json:"avatar_url,omitempty"`
	Metadata        *string `json:"metadata,omitempty"`
	CanRefer        *bool   `json:"can_refer,omitempty"`
	Token           *string `json:"token,omitempty"`
	Status          *string `json:"status,omitempty"`
}

// Bonuses represents a bonus.
type Bonuses struct {
	AdvocateToken string   `json:"advocate_token"`
	Reference     string   `json:"reference"`
	PaymentAmount *float64 `json:"payment_amount,omitempty"`
}

// BonusesForm wraps Bonuses for create.
type BonusesForm struct {
	Bonus Bonuses `json:"bonus"`
}

// ForceBonuses represents a forced bonus.
type ForceBonuses struct {
	AdvocateToken string `json:"advocate_token"`
	Reference     string `json:"reference"`
	BonusAmount   int64  `json:"bonus_amount"`
}

// ForceBonusesForm wraps ForceBonuses.
type ForceBonusesForm struct {
	Bonus ForceBonuses `json:"bonus"`
}

// BonusPatchForm for partial bonus updates.
type BonusPatchForm struct {
	Status *string `json:"status,omitempty"`
}

// RedemptionRequest represents a redemption request.
type RedemptionRequest struct {
	AdvocateToken           string  `json:"advocate_token"`
	RequestStatusSlug       string  `json:"request_status_slug"`
	RequestActionSlug       string  `json:"request_action_slug"`
	CurrencyCode            string  `json:"currency_code"`
	Amount                  int64   `json:"amount"`
	Description             *string `json:"description,omitempty"`
	AdvocatesPaypalUsername *string `json:"advocates_paypal_username,omitempty"`
}

// RedemptionRequestForm wraps RedemptionRequest.
type RedemptionRequestForm struct {
	RedemptionRequest RedemptionRequest `json:"redemption_request"`
}

// Referral represents a referral.
type Referral struct {
	ReferredAdvocateToken string `json:"referred_advocate_token"`
	ReferralOriginSlug    string `json:"referral_origin_slug"`
	CampaignSlug          string `json:"campaign_slug"`
	HTTPReferer           string `json:"http_referer"`
}

// ReferralForm wraps Referral.
type ReferralForm struct {
	Referral Referral `json:"referral"`
}

// PaymentMethod represents an advocate payment method.
type PaymentMethod struct {
	Username    string `json:"username"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

// PaymentMethodForm wraps PaymentMethod.
type PaymentMethodForm struct {
	AdvocatePaymentMethod PaymentMethod `json:"advocate_payment_method"`
}

// Account represents an account.
type Account struct {
	Name     string `json:"name,omitempty"`
	URL      string `json:"url,omitempty"`
	Slug     string `json:"slug,omitempty"`
	IsActive *bool  `json:"is_active,omitempty"`
	IsLive   *bool  `json:"is_live,omitempty"`
}

// AccountForm wraps Account for create/update.
type AccountForm struct {
	Account Account `json:"account"`
}

// Tag represents a tag.
type Tag struct {
	Name        string  `json:"name,omitempty"`
	Slug        string  `json:"slug,omitempty"`
	Description *string `json:"description,omitempty"`
	Color       *string `json:"color,omitempty"`
}

// TagForm wraps Tag.
type TagForm struct {
	Tag Tag `json:"tag"`
}

// Voucher represents a voucher.
type Voucher struct {
	Code         string  `json:"code,omitempty"`
	Value        float64 `json:"value,omitempty"`
	CurrencyCode string  `json:"currency_code,omitempty"`
	Status       string  `json:"status,omitempty"`
}

// VoucherForm wraps Voucher.
type VoucherForm struct {
	Voucher Voucher `json:"voucher"`
}

// Product represents a product.
type Product struct {
	Title           string   `json:"title,omitempty"`
	Slug            string   `json:"slug,omitempty"`
	DescriptionHTML string   `json:"description_html,omitempty"`
	Status          string   `json:"status,omitempty"`
	Images          []string `json:"images,omitempty"`
}

// ProductForm wraps Product.
type ProductForm struct {
	Product Product `json:"product"`
}

// ProductVariant represents a product variant.
type ProductVariant struct {
	Title             string `json:"title,omitempty"`
	SKU               string `json:"sku,omitempty"`
	Barcode           string `json:"barcode,omitempty"`
	InventoryQuantity int64  `json:"inventory_quantity,omitempty"`
	Features          string `json:"features,omitempty"`
}

// ProductVariantForm wraps ProductVariant.
type ProductVariantForm struct {
	Variant ProductVariant `json:"variant"`
}

// ProductVariantPriceForm for variant price updates.
type ProductVariantPriceForm struct {
	Price struct {
		CurrencyCode string `json:"currency_code,omitempty"`
		Amount       string `json:"amount,omitempty"`
	} `json:"price,omitempty"`
}

// WidgetsPackage represents a widgets package.
type WidgetsPackage struct {
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// WidgetsPackageForm wraps WidgetsPackage.
type WidgetsPackageForm struct {
	WidgetsPackage WidgetsPackage `json:"widgets_package"`
}

// Widget represents a widget.
type Widget struct {
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// WidgetForm wraps Widget.
type WidgetForm struct {
	Widget Widget `json:"widget"`
}

// WidgetTranslation represents a widget translation.
type WidgetTranslation struct {
	Locale      string `json:"locale,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// WidgetTranslationForm wraps WidgetTranslation.
type WidgetTranslationForm struct {
	WidgetTranslation WidgetTranslation `json:"widget_translation"`
}

// PaginatedResponse is a generic paginated API response wrapper.
type PaginatedResponse struct {
	Page    int64       `json:"page,omitempty"`
	Limit   int64       `json:"limit,omitempty"`
	Total   int64       `json:"total,omitempty"`
	Results interface{} `json:"results,omitempty"`
}
