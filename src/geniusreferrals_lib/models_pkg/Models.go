/*
 * geniusreferrals_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 02/18/2017
 */

package models_pkg


/*
 * Structure for the custom type Advocate
 */
type Advocate struct {
    Name             string          `json:"name" form:"name"` //The advocate's name
    Lastname         string          `json:"lastname" form:"lastname"` //The advocate's last name
    Email            string          `json:"email" form:"email"` //The advocate's email
    PayoutThreshold  int64           `json:"payout_threshold" form:"payout_threshold"` //The total amount of bonuses that the advocate must generate before being able to create a bonus redemption request.
    AvatarUrl        *string         `json:"avatar_url,omitempty" form:"avatar_url,omitempty"` //The advocate's avatar URL
    Metadata         *string         `json:"metadata,omitempty" form:"metadata,omitempty"` //Useful to store extra information about the advocate. e.g, the phone number, address, etc.
    CanRefer         *bool           `json:"can_refer,omitempty" form:"can_refer,omitempty"` //Whether or not the advocate is allowed to refer services/products (Useful when using the Full Widget template).
}

/*
 * Structure for the custom type AdvocateForm
 */
type AdvocateForm struct {
    Advocate        Advocate        `json:"advocate" form:"advocate"` //The advocate's wrapper
}

/*
 * Structure for the custom type AdvocatePatchForm
 */
type AdvocatePatchForm struct {
    Name              *string         `json:"name,omitempty" form:"name,omitempty"` //The advocate's name
    Lastname          *string         `json:"lastname,omitempty" form:"lastname,omitempty"` //The advocate's last name
    Email             *string         `json:"email,omitempty" form:"email,omitempty"` //The advocate's email
    PayoutThreshold   *int64          `json:"payout_threshold,omitempty" form:"payout_threshold,omitempty"` //TODO: Write general description for this field
    ClaimedBalance    *int64          `json:"claimed_balance,omitempty" form:"claimed_balance,omitempty"` //The total amount of bonuses that the advocate must generate before being able to create a bonus redemption request.
    UnclaimedBalance  *int64          `json:"unclaimed_balance,omitempty" form:"unclaimed_balance,omitempty"` //The unclaimed balance
    CurrencyCode      *string         `json:"currency_code,omitempty" form:"currency_code,omitempty"` //The currency code
    AvatarUrl         *string         `json:"avatar_url,omitempty" form:"avatar_url,omitempty"` //The advocate's avatar URL
    Metadata          *string         `json:"metadata,omitempty" form:"metadata,omitempty"` //Useful to store extra information about the advocate. e.g, the phone number, address, etc.
    CanRefer          *bool           `json:"can_refer,omitempty" form:"can_refer,omitempty"` //Whether or not the advocate is allowed to refer services/products (Useful when using the Full Widget template).
    Token             *string         `json:"token,omitempty" form:"token,omitempty"` //The advocate's token
}

/*
 * Structure for the custom type Bonuses
 */
type Bonuses struct {
    AdvocateToken   string          `json:"advocate_token" form:"advocate_token"` //The referral's token. Usually the one that completed the purchase, or trigger an event.
    Reference       string          `json:"reference" form:"reference"` //The reference number for this request. Usually the order_id, payment_id, or timestamp.
    PaymentAmount   *float64        `json:"payment_amount,omitempty" form:"payment_amount,omitempty"` //The payment amount the referrals has made. Required for a percentage based campaign.
}

/*
 * Structure for the custom type BonusesForm
 */
type BonusesForm struct {
    Bonus           Bonuses         `json:"bonus" form:"bonus"` //The bonuses' wrapper
}

/*
 * Structure for the custom type ForceBonuses
 */
type ForceBonuses struct {
    AdvocateToken   string          `json:"advocate_token" form:"advocate_token"` //The referral's token.
    Reference       string          `json:"reference" form:"reference"` //The reference number for this request. Usually the order_id, payment_id, or timestamp.
    BonusAmount     int64           `json:"bonus_amount" form:"bonus_amount"` //The bonus amount to give to the advocate.
}

/*
 * Structure for the custom type ForceBonusesForm
 */
type ForceBonusesForm struct {
    Bonus           ForceBonuses    `json:"bonus" form:"bonus"` //The bonuses' wrapper
}

/*
 * Structure for the custom type RedemptionRequest
 */
type RedemptionRequest struct {
    AdvocateToken             string          `json:"advocate_token" form:"advocate_token"` //The advocate's token
    RequestStatusSlug         string          `json:"request_status_slug" form:"request_status_slug"` //The request status identifier
    RequestActionSlug         string          `json:"request_action_slug" form:"request_action_slug"` //The request action identifier
    CurrencyCode              string          `json:"currency_code" form:"currency_code"` //The currency code
    Amount                    int64           `json:"amount" form:"amount"` //The amount to be redeemed
    Description               *string         `json:"description,omitempty" form:"description,omitempty"` //The description of the requests, useful to store extra information about the request.
    AdvocatesPaypalUsername   *string         `json:"advocates_paypal_username,omitempty" form:"advocates_paypal_username,omitempty"` //The advocate's PayPal username
}

/*
 * Structure for the custom type RedemptionRequestForm
 */
type RedemptionRequestForm struct {
    RedemptionRequest  RedemptionRequest `json:"redemption_request" form:"redemption_request"` //The redemption request's wrapper
}

/*
 * Structure for the custom type Referral
 */
type Referral struct {
    ReferredAdvocateToken   string          `json:"referred_advocate_token" form:"referred_advocate_token"` //The referrals token
    ReferralOriginSlug      string          `json:"referral_origin_slug" form:"referral_origin_slug"` //The referral origin identifier
    CampaignSlug            string          `json:"campaign_slug" form:"campaign_slug"` //The campaign identifier
    HttpReferer             string          `json:"http_referer" form:"http_referer"` //The http_referrer URL
}

/*
 * Structure for the custom type ReferralForm
 */
type ReferralForm struct {
    Referral        Referral        `json:"referral" form:"referral"` //The referral's wrapper
}

/*
 * Structure for the custom type PaymentMethod
 */
type PaymentMethod struct {
    Username        string          `json:"username" form:"username"` //The username
    Description     string          `json:"description" form:"description"` //The description
    IsActive        []byte          `json:"is_active" form:"is_active"` //Is the payment method active? (true, false)
}

/*
 * Structure for the custom type PaymentMethodForm
 */
type PaymentMethodForm struct {
    AdvocatePaymentMethod   PaymentMethod   `json:"advocate_payment_method" form:"advocate_payment_method"` //The payment methods wrapper
}
