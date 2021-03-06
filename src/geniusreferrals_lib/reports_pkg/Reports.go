/*
 * geniusreferrals_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 02/18/2017
 */

package reports_pkg

import "time"

/*
 * Interface for the REPORTS_IMPL
 */
type REPORTS interface {
    GetReferralsSummaryPerOrigin (string) (interface{}, error)

    GetBonusesSummaryPerOrigin (string) (interface{}, error)

    GetTopAdvocates (*string, *string, *int64, *time.Time, *time.Time) (interface{}, error)

    GetShareDailyParticipation (*string, *string, *string, *time.Time, *time.Time) (interface{}, error)

    GetReferralDailyParticipation (*string, *string, *string, *time.Time, *time.Time) (interface{}, error)

    GetClickDailyParticipation (*string, *string, *string, *time.Time, *time.Time) (interface{}, error)

    GetBonusesDailyGiven (*string, *string, *string, *time.Time, *time.Time) (interface{}, error)
}

/*
 * Factory for the REPORTS interaface returning REPORTS_IMPL
 */
func NewREPORTS() REPORTS {
    return &REPORTS_IMPL{}
}