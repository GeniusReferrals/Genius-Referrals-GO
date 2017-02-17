/*
 * geniusreferrals_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 02/17/2017
 */
package campaigns_pkg


import(
	"github.com/apimatic/unirest-go"
	"geniusreferrals_lib"
	"geniusreferrals_lib/apihelper_pkg"
)
/*
 * Client structure as interface implementation
 */
type CAMPAIGNS_IMPL struct { }

/**
 * Get a campaign by a given slug.
 * @param    string        accountSlug       parameter: Required
 * @param    string        campaignSlug      parameter: Required
 * @return	Returns the interface{} response from the API call
 */
func (me *CAMPAIGNS_IMPL) GetCampaign (
            accountSlug string,
            campaignSlug string) (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := geniusreferrals_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/accounts/{account_slug}/campaigns/{campaign_slug}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "account_slug" : accountSlug,
        "campaign_slug" : campaignSlug,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "X-Auth-Token" : geniusreferrals_lib.Config.XAuthToken,
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("You are not authenticated", _response.Code, _response.RawBody)
    } else if (_response.Code == 403) {
        err = apihelper_pkg.NewAPIError("User not authorized to perform the operation", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("Resource not found", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal interface{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Get the list of campaings for a given account.
 * @param    string         accountSlug      parameter: Required
 * @param    *int64         page             parameter: Optional
 * @param    *int64         limit            parameter: Optional
 * @param    *string        filter           parameter: Optional
 * @param    *string        sort             parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *CAMPAIGNS_IMPL) GetCampaigns (
            accountSlug string,
            page *int64,
            limit *int64,
            filter *string,
            sort *string) (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := geniusreferrals_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/accounts/{account_slug}/campaigns"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "account_slug" : accountSlug,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "page" : page,
        "limit" : limit,
        "filter" : filter,
        "sort" : sort,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "X-Auth-Token" : geniusreferrals_lib.Config.XAuthToken,
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("You are not authenticated", _response.Code, _response.RawBody)
    } else if (_response.Code == 403) {
        err = apihelper_pkg.NewAPIError("User not authorized to perform the operation", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("Resource not found", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal interface{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

