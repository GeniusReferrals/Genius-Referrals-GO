/*
 * geniusreferrals_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 02/17/2017
 */

package accounts_pkg


/*
 * Interface for the ACCOUNTS_IMPL
 */
type ACCOUNTS interface {
    GetAccount (string) (interface{}, error)

    GetAccounts (*int64, *int64, *string, *string) (interface{}, error)
}

/*
 * Factory for the ACCOUNTS interaface returning ACCOUNTS_IMPL
 */
func NewACCOUNTS() ACCOUNTS {
    return &ACCOUNTS_IMPL{}
}