/*
 * geniusreferrals_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 02/17/2017
 */

package roots_pkg


/*
 * Interface for the ROOTS_IMPL
 */
type ROOTS interface {
    GetRoot () (interface{}, error)
}

/*
 * Factory for the ROOTS interaface returning ROOTS_IMPL
 */
func NewROOTS() ROOTS {
    return &ROOTS_IMPL{}
}