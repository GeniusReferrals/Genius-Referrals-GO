/*
 * geniusreferrals_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 02/18/2017
 */
package roots_pkg


import(
	"github.com/apimatic/unirest-go"
	"geniusreferrals_lib"
	"geniusreferrals_lib/apihelper_pkg"
)
/*
 * Client structure as interface implementation
 */
type ROOTS_IMPL struct { }

/**
 * The root of the API
 * @return	Returns the interface{} response from the API call
 */
func (me *ROOTS_IMPL) GetRoot () (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := geniusreferrals_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/"

    //variable to hold errors
    var err error = nil
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

