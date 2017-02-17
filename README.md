# Getting started

## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=Genius%20Referrals-GoLang&projectName=geniusreferrals_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=geniusreferrals_lib)

## How to Use

The following section explains how to use the GeniusReferrals library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=geniusreferrals_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=geniusreferrals_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=geniusreferrals_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=geniusreferrals_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=geniusreferrals_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=Genius%20Referrals-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=geniusreferrals_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=geniusreferrals_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| contentType | The content type |
| xAuthToken | Your API Token, you can get your token here https://www.geniusreferrals.com/en/settings/api-access |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


## Class Reference

### <a name="list_of_controllers"></a>List of Controllers

* [roots_pkg](#roots_pkg)
* [authentications_pkg](#authentications_pkg)
* [advocates_pkg](#advocates_pkg)
* [accounts_pkg](#accounts_pkg)
* [reports_pkg](#reports_pkg)
* [referrals_pkg](#referrals_pkg)
* [redemptionrequests_pkg](#redemptionrequests_pkg)
* [bonuses_pkg](#bonuses_pkg)
* [campaigns_pkg](#campaigns_pkg)

### <a name="roots_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".roots_pkg") roots_pkg

#### Get instance

Factory for the ``` ROOTS ``` interface can be accessed from the package roots_pkg.

```go
roots := roots_pkg.NewROOTS()
```

#### <a name="get_root"></a>![Method: ](https://apidocs.io/img/method.png ".roots_pkg.GetRoot") GetRoot

> The root of the API


```go
func (me *ROOTS_IMPL) GetRoot()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = roots.GetRoot()

```


[Back to List of Controllers](#list_of_controllers)

### <a name="authentications_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".authentications_pkg") authentications_pkg

#### Get instance

Factory for the ``` AUTHENTICATIONS ``` interface can be accessed from the package authentications_pkg.

```go
authentications := authentications_pkg.NewAUTHENTICATIONS()
```

#### <a name="get_authentications"></a>![Method: ](https://apidocs.io/img/method.png ".authentications_pkg.GetAuthentications") GetAuthentications

> Allow clients to test authentication on Genius Referrals platform.


```go
func (me *AUTHENTICATIONS_IMPL) GetAuthentications()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = authentications.GetAuthentications()

```


[Back to List of Controllers](#list_of_controllers)

### <a name="advocates_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".advocates_pkg") advocates_pkg

#### Get instance

Factory for the ``` ADVOCATES ``` interface can be accessed from the package advocates_pkg.

```go
advocates := advocates_pkg.NewADVOCATES()
```

#### <a name="delete_advocate"></a>![Method: ](https://apidocs.io/img/method.png ".advocates_pkg.DeleteAdvocate") DeleteAdvocate

> Delete an advocate


```go
func (me *ADVOCATES_IMPL) DeleteAdvocate(
            accountSlug string,
            advocateToken string)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| advocateToken |  ``` Required ```  | The advocate's token |


#### Example Usage

```go
accountSlug := "account_slug"
advocateToken := "advocate_token"

var result 
result,_ = advocates.DeleteAdvocate(accountSlug, advocateToken)

```


#### <a name="put_advocate"></a>![Method: ](https://apidocs.io/img/method.png ".advocates_pkg.PutAdvocate") PutAdvocate

> Update an advocate.


```go
func (me *ADVOCATES_IMPL) PutAdvocate(
            accountSlug string,
            advocateToken string,
            advocateForm *models_pkg.AdvocateForm)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| advocateToken |  ``` Required ```  | The advocate's token |
| advocateForm |  ``` Required ```  | The body of the request |


#### Example Usage

```go
accountSlug := "account_slug"
advocateToken := "advocate_token"
var advocateForm *models_pkg.AdvocateForm

var result 
result,_ = advocates.PutAdvocate(accountSlug, advocateToken, advocateForm)

```


#### <a name="post_advocate"></a>![Method: ](https://apidocs.io/img/method.png ".advocates_pkg.PostAdvocate") PostAdvocate

> Create a new advocate.


```go
func (me *ADVOCATES_IMPL) PostAdvocate(
            accountSlug string,
            advocateForm *models_pkg.AdvocateForm)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| advocateForm |  ``` Required ```  | The body of the request |


#### Example Usage

```go
accountSlug := "account_slug"
var advocateForm *models_pkg.AdvocateForm

var result interface{}
result,_ = advocates.PostAdvocate(accountSlug, advocateForm)

```


#### <a name="get_advocate"></a>![Method: ](https://apidocs.io/img/method.png ".advocates_pkg.GetAdvocate") GetAdvocate

> Get an advocate by a given token.


```go
func (me *ADVOCATES_IMPL) GetAdvocate(
            accountSlug string,
            advocateToken string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| advocateToken |  ``` Required ```  | The advocate's token |


#### Example Usage

```go
accountSlug := "account_slug"
advocateToken := "advocate_token"

var result interface{}
result,_ = advocates.GetAdvocate(accountSlug, advocateToken)

```


#### <a name="delete_advocates"></a>![Method: ](https://apidocs.io/img/method.png ".advocates_pkg.DeleteAdvocates") DeleteAdvocates

> Delete all advocates


```go
func (me *ADVOCATES_IMPL) DeleteAdvocates(accountSlug string)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |


#### Example Usage

```go
accountSlug := "account_slug"

var result 
result,_ = advocates.DeleteAdvocates(accountSlug)

```


#### <a name="get_advocates"></a>![Method: ](https://apidocs.io/img/method.png ".advocates_pkg.GetAdvocates") GetAdvocates

> Get the list of advocates


```go
func (me *ADVOCATES_IMPL) GetAdvocates(
            accountSlug string,
            page *int64,
            limit *int64,
            filter *string,
            sort *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| page |  ``` Optional ```  ``` DefaultValue ```  | Page number, e.g. 1 would start at the first result, and 10 would start at the tenth result. |
| limit |  ``` Optional ```  ``` DefaultValue ```  | Maximum number of results to return in the response. Default (10), threshold (100) |
| filter |  ``` Optional ```  | Allowed fields: name, lastname, email, advocate_token, bonus_exchange_method_slug, campaign_slug, can_refer, is_referral, from, to, created. Use the following delimiters to build your filters params. The vertical bar ('\|') to separate individual filter phrases and a double colon ('::') to separate the names and values. The delimiter of the double colon (':') separates the property name from the comparison value, enabling the comparison value to contain spaces. Example: www.example.com/users?filter='name::todd\|city::denver\|title::grand poobah' |
| sort |  ``` Optional ```  | Allowed fields: name, lastname, email, created. Use sort query-string parameter that contains a delimited set of property names. For each property name, sort in ascending order, and for each property prefixed with a dash ('-') sort in descending order. Separate each property name with a vertical bar ('\|'), which is consistent with the separation of the name\|value pairs in filtering, above. For example, if we want to retrieve users in order of their last name (ascending), first name (ascending) and hire date (descending), the request might look like this www.example.com/users?sort='last_name\|first_name\|-hire_date' |


#### Example Usage

```go
accountSlug := "account_slug"
page,_ := strconv.ParseInt("1", 10, 8)
limit,_ := strconv.ParseInt("10", 10, 8)
filter := "filter"
sort := "sort"

var result interface{}
result,_ = advocates.GetAdvocates(accountSlug, page, limit, filter, sort)

```


#### <a name="patch_advocate"></a>![Method: ](https://apidocs.io/img/method.png ".advocates_pkg.PatchAdvocate") PatchAdvocate

> Update partial elements of an advocate.


```go
func (me *ADVOCATES_IMPL) PatchAdvocate(
            accountSlug string,
            advocateToken string,
            advocatePatchForm []*models_pkg.AdvocatePatchForm)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| advocateToken |  ``` Required ```  | The advocate's token |
| advocatePatchForm |  ``` Required ```  ``` Collection ```  | The body of the request |


#### Example Usage

```go
accountSlug := "account_slug"
advocateToken := "advocate_token"
var advocatePatchForm []*models_pkg.AdvocatePatchForm

var result interface{}
result,_ = advocates.PatchAdvocate(accountSlug, advocateToken, advocatePatchForm)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="accounts_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".accounts_pkg") accounts_pkg

#### Get instance

Factory for the ``` ACCOUNTS ``` interface can be accessed from the package accounts_pkg.

```go
accounts := accounts_pkg.NewACCOUNTS()
```

#### <a name="get_account"></a>![Method: ](https://apidocs.io/img/method.png ".accounts_pkg.GetAccount") GetAccount

> Get an account by a given slug.


```go
func (me *ACCOUNTS_IMPL) GetAccount(accountSlug string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |


#### Example Usage

```go
accountSlug := "account_slug"

var result interface{}
result,_ = accounts.GetAccount(accountSlug)

```


#### <a name="get_accounts"></a>![Method: ](https://apidocs.io/img/method.png ".accounts_pkg.GetAccounts") GetAccounts

> Get the list of accounts.


```go
func (me *ACCOUNTS_IMPL) GetAccounts(
            page *int64,
            limit *int64,
            filter *string,
            sort *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page number, e.g. 1 would start at the first result, and 10 would start at the tenth result. |
| limit |  ``` Optional ```  | Maximum number of results to return in the response. Default (10), threshold (100) |
| filter |  ``` Optional ```  | Allowed fields: name. Use the following delimiters to build your filters params. The vertical bar ('\|') to separate individual filter phrases and a double colon ('::') to separate the names and values. The delimiter of the double colon (':') separates the property name from the comparison value, enabling the comparison value to contain spaces. Example: www.example.com/users?filter='name::todd\|city::denver\|title::grand poobah' |
| sort |  ``` Optional ```  | Allowed fields: name, created. Use sort query-string parameter that contains a delimited set of property names. For each property name, sort in ascending order, and for each property prefixed with a dash ('-') sort in descending order. Separate each property name with a vertical bar ('\|'), which is consistent with the separation of the name\|value pairs in filtering, above. For example, if we want to retrieve users in order of their last name (ascending), first name (ascending) and hire date (descending), the request might look like this www.example.com/users?sort=last_name\|first_name\|-hire_date |


#### Example Usage

```go
page,_ := strconv.ParseInt("200", 10, 8)
limit,_ := strconv.ParseInt("200", 10, 8)
filter := "filter"
sort := "sort"

var result interface{}
result,_ = accounts.GetAccounts(page, limit, filter, sort)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="reports_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".reports_pkg") reports_pkg

#### Get instance

Factory for the ``` REPORTS ``` interface can be accessed from the package reports_pkg.

```go
reports := reports_pkg.NewREPORTS()
```

#### <a name="get_referrals_summary_per_origin"></a>![Method: ](https://apidocs.io/img/method.png ".reports_pkg.GetReferralsSummaryPerOrigin") GetReferralsSummaryPerOrigin

> Get referrals summary per referral origin.


```go
func (me *REPORTS_IMPL) GetReferralsSummaryPerOrigin(advocateToken string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| advocateToken |  ``` Required ```  | The advocate's token |


#### Example Usage

```go
advocateToken := "advocate_token"

var result interface{}
result,_ = reports.GetReferralsSummaryPerOrigin(advocateToken)

```


#### <a name="get_bonuses_summary_per_origin"></a>![Method: ](https://apidocs.io/img/method.png ".reports_pkg.GetBonusesSummaryPerOrigin") GetBonusesSummaryPerOrigin

> Get bonuses summary per referral origin.


```go
func (me *REPORTS_IMPL) GetBonusesSummaryPerOrigin(advocateToken string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| advocateToken |  ``` Required ```  | The advocate's token |


#### Example Usage

```go
advocateToken := "advocate_token"

var result interface{}
result,_ = reports.GetBonusesSummaryPerOrigin(advocateToken)

```


#### <a name="get_top_advocates"></a>![Method: ](https://apidocs.io/img/method.png ".reports_pkg.GetTopAdvocates") GetTopAdvocates

> Get top 10 advocates.


```go
func (me *REPORTS_IMPL) GetTopAdvocates(
            accountSlug *string,
            campaignSlug *string,
            limit *int64,
            from *time.Time,
            to *time.Time)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Optional ```  | The account identifier |
| campaignSlug |  ``` Optional ```  | The campaign identifier |
| limit |  ``` Optional ```  | Maximum number of results to return in the response. Default (10) |
| from |  ``` Optional ```  | The datetime were the range of the search starts |
| to |  ``` Optional ```  | The datetime were the range of the search stops |


#### Example Usage

```go
accountSlug := "account_slug"
campaignSlug := "campaign_slug"
limit,_ := strconv.ParseInt("200", 10, 8)
from := time.Now()
to := time.Now()

var result interface{}
result,_ = reports.GetTopAdvocates(accountSlug, campaignSlug, limit, from, to)

```


#### <a name="get_share_daily_participation"></a>![Method: ](https://apidocs.io/img/method.png ".reports_pkg.GetShareDailyParticipation") GetShareDailyParticipation

> Get share daily participation.


```go
func (me *REPORTS_IMPL) GetShareDailyParticipation(
            accountSlug *string,
            campaignSlug *string,
            advocateToken *string,
            from *time.Time,
            to *time.Time)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Optional ```  | The account identifier |
| campaignSlug |  ``` Optional ```  | The campaign identifier |
| advocateToken |  ``` Optional ```  | The advocate's token |
| from |  ``` Optional ```  | The datetime were the range of the search starts |
| to |  ``` Optional ```  | The datetime were the range of the search stops |


#### Example Usage

```go
accountSlug := "account_slug"
campaignSlug := "campaign_slug"
advocateToken := "advocate_token"
from := time.Now()
to := time.Now()

var result interface{}
result,_ = reports.GetShareDailyParticipation(accountSlug, campaignSlug, advocateToken, from, to)

```


#### <a name="get_referral_daily_participation"></a>![Method: ](https://apidocs.io/img/method.png ".reports_pkg.GetReferralDailyParticipation") GetReferralDailyParticipation

> Get referral daily participation.


```go
func (me *REPORTS_IMPL) GetReferralDailyParticipation(
            accountSlug *string,
            campaignSlug *string,
            advocateToken *string,
            from *time.Time,
            to *time.Time)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Optional ```  | The account identifier |
| campaignSlug |  ``` Optional ```  | The campaign identifier |
| advocateToken |  ``` Optional ```  | The advocate's token |
| from |  ``` Optional ```  | The datetime were the range of the search starts |
| to |  ``` Optional ```  | The datetime were the range of the search stops |


#### Example Usage

```go
accountSlug := "account_slug"
campaignSlug := "campaign_slug"
advocateToken := "advocate_token"
from := time.Now()
to := time.Now()

var result interface{}
result,_ = reports.GetReferralDailyParticipation(accountSlug, campaignSlug, advocateToken, from, to)

```


#### <a name="get_click_daily_participation"></a>![Method: ](https://apidocs.io/img/method.png ".reports_pkg.GetClickDailyParticipation") GetClickDailyParticipation

> Get click daily participation.


```go
func (me *REPORTS_IMPL) GetClickDailyParticipation(
            accountSlug *string,
            campaignSlug *string,
            advocateToken *string,
            from *time.Time,
            to *time.Time)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Optional ```  | The account identifier |
| campaignSlug |  ``` Optional ```  | The campaign identifier |
| advocateToken |  ``` Optional ```  | The advocate's token |
| from |  ``` Optional ```  | The datetime were the range of the search starts |
| to |  ``` Optional ```  | The datetime were the range of the search stops |


#### Example Usage

```go
accountSlug := "account_slug"
campaignSlug := "campaign_slug"
advocateToken := "advocate_token"
from := time.Now()
to := time.Now()

var result interface{}
result,_ = reports.GetClickDailyParticipation(accountSlug, campaignSlug, advocateToken, from, to)

```


#### <a name="get_bonuses_daily_given"></a>![Method: ](https://apidocs.io/img/method.png ".reports_pkg.GetBonusesDailyGiven") GetBonusesDailyGiven

> Get bonuses daily given.


```go
func (me *REPORTS_IMPL) GetBonusesDailyGiven(
            accountSlug *string,
            campaignSlug *string,
            advocateToken *string,
            from *time.Time,
            to *time.Time)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Optional ```  | The account identifier |
| campaignSlug |  ``` Optional ```  | The campaign identifier |
| advocateToken |  ``` Optional ```  | The advocate identifier |
| from |  ``` Optional ```  | The datetime were the range of the search starts |
| to |  ``` Optional ```  | The datetime were the range of the search stops |


#### Example Usage

```go
accountSlug := "account_slug"
campaignSlug := "campaign_slug"
advocateToken := "advocate_token"
from := time.Now()
to := time.Now()

var result interface{}
result,_ = reports.GetBonusesDailyGiven(accountSlug, campaignSlug, advocateToken, from, to)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="referrals_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".referrals_pkg") referrals_pkg

#### Get instance

Factory for the ``` REFERRALS ``` interface can be accessed from the package referrals_pkg.

```go
referrals := referrals_pkg.NewREFERRALS()
```

#### <a name="get_referral_origin"></a>![Method: ](https://apidocs.io/img/method.png ".referrals_pkg.GetReferralOrigin") GetReferralOrigin

> Get a referral origin by a given slug.


```go
func (me *REFERRALS_IMPL) GetReferralOrigin(referralOriginSlug string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| referralOriginSlug |  ``` Required ```  | The referral origin identifier |


#### Example Usage

```go
referralOriginSlug := "referral_origin_slug"

var result interface{}
result,_ = referrals.GetReferralOrigin(referralOriginSlug)

```


#### <a name="get_referral_origins"></a>![Method: ](https://apidocs.io/img/method.png ".referrals_pkg.GetReferralOrigins") GetReferralOrigins

> Get referral origins. This is needed when creating (POST) a new referral, as referral_origin_slug refers to one of this origins.


```go
func (me *REFERRALS_IMPL) GetReferralOrigins()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = referrals.GetReferralOrigins()

```


#### <a name="get_referral"></a>![Method: ](https://apidocs.io/img/method.png ".referrals_pkg.GetReferral") GetReferral

> Get a referral by a given id.


```go
func (me *REFERRALS_IMPL) GetReferral(
            accountSlug string,
            advocateToken string,
            referralId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| advocateToken |  ``` Required ```  | The advocate's token |
| referralId |  ``` Required ```  | The referral id |


#### Example Usage

```go
accountSlug := "account_slug"
advocateToken := "advocate_token"
referralId := "referral_id"

var result interface{}
result,_ = referrals.GetReferral(accountSlug, advocateToken, referralId)

```


#### <a name="delete_referral"></a>![Method: ](https://apidocs.io/img/method.png ".referrals_pkg.DeleteReferral") DeleteReferral

> Delete a referral.


```go
func (me *REFERRALS_IMPL) DeleteReferral(
            accountSlug string,
            advocateToken string,
            referralId string)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| advocateToken |  ``` Required ```  | The advocate's token |
| referralId |  ``` Required ```  | The referral identifier |


#### Example Usage

```go
accountSlug := "account_slug"
advocateToken := "advocate_token"
referralId := "referral_id"

var result 
result,_ = referrals.DeleteReferral(accountSlug, advocateToken, referralId)

```


#### <a name="post_referrals"></a>![Method: ](https://apidocs.io/img/method.png ".referrals_pkg.PostReferrals") PostReferrals

> Create a new referral.


```go
func (me *REFERRALS_IMPL) PostReferrals(
            accountSlug string,
            advocateToken string,
            referralForm *models_pkg.ReferralForm)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| advocateToken |  ``` Required ```  | The advocate's token |
| referralForm |  ``` Required ```  | The body of the request |


#### Example Usage

```go
accountSlug := "account_slug"
advocateToken := "advocate_token"
var referralForm *models_pkg.ReferralForm

var result interface{}
result,_ = referrals.PostReferrals(accountSlug, advocateToken, referralForm)

```


#### <a name="put_referral"></a>![Method: ](https://apidocs.io/img/method.png ".referrals_pkg.PutReferral") PutReferral

> Update a referral.


```go
func (me *REFERRALS_IMPL) PutReferral(
            accountSlug string,
            advocateToken string,
            referralId string,
            referralForm *models_pkg.ReferralForm)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| advocateToken |  ``` Required ```  | The advocate's token |
| referralId |  ``` Required ```  | The referral id |
| referralForm |  ``` Required ```  | The body of the request |


#### Example Usage

```go
accountSlug := "account_slug"
advocateToken := "advocate_token"
referralId := "referral_id"
var referralForm *models_pkg.ReferralForm

var result 
result,_ = referrals.PutReferral(accountSlug, advocateToken, referralId, referralForm)

```


#### <a name="get_referrals"></a>![Method: ](https://apidocs.io/img/method.png ".referrals_pkg.GetReferrals") GetReferrals

> Get the list of referrals for a given advocate.


```go
func (me *REFERRALS_IMPL) GetReferrals(
            accountSlug string,
            advocateToken string,
            page *int64,
            limit *int64,
            filter *string,
            sort *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| advocateToken |  ``` Required ```  | The advocate's token |
| page |  ``` Optional ```  | Page number, e.g. 1 would start at the first result, and 10 would start at the tenth result. |
| limit |  ``` Optional ```  | Maximum number of results to return in the response. Default (10), threshold (100) |
| filter |  ``` Optional ```  | Allowed fields: url, referral_origin_slug, created. Use the following delimiters to build your filters params. Use the following delimiters to build your filters params. The vertical bar ('\|') to separate individual filter phrases and a double colon ('::') to separate the names and values. The delimiter of the double colon (':') separates the property name from the comparison value, enabling the comparison value to contain spaces. Example: www.example.com/users?filter='name::todd\|city::denver\|title::grand poobah' |
| sort |  ``` Optional ```  | Allowed fields: created. Use sort query-string parameter that contains a delimited set of property names. For each property name, sort in ascending order, and for each property prefixed with a dash ('-') sort in descending order. Separate each property name with a vertical bar ('\|'), which is consistent with the separation of the name\|value pairs in filtering, above. For example, if we want to retrieve users in order of their last name (ascending), first name (ascending) and hire date (descending), the request might look like this www.example.com/users?sort='last_name\|first_name\|-hire_date' |


#### Example Usage

```go
accountSlug := "account_slug"
advocateToken := "advocate_token"
page,_ := strconv.ParseInt("200", 10, 8)
limit,_ := strconv.ParseInt("200", 10, 8)
filter := "filter"
sort := "sort"

var result interface{}
result,_ = referrals.GetReferrals(accountSlug, advocateToken, page, limit, filter, sort)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="redemptionrequests_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".redemptionrequests_pkg") redemptionrequests_pkg

#### Get instance

Factory for the ``` REDEMPTIONREQUESTS ``` interface can be accessed from the package redemptionrequests_pkg.

```go
redemptionRequests := redemptionrequests_pkg.NewREDEMPTIONREQUESTS()
```

#### <a name="get_redemption_request_status"></a>![Method: ](https://apidocs.io/img/method.png ".redemptionrequests_pkg.GetRedemptionRequestStatus") GetRedemptionRequestStatus

> Get a redemption request status.


```go
func (me *REDEMPTIONREQUESTS_IMPL) GetRedemptionRequestStatus(redemptionRequestStatusSlug string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| redemptionRequestStatusSlug |  ``` Required ```  | The redemption request status identifier |


#### Example Usage

```go
redemptionRequestStatusSlug := "redemption_request_status_slug"

var result interface{}
result,_ = redemptionRequests.GetRedemptionRequestStatus(redemptionRequestStatusSlug)

```


#### <a name="get_redemption_request_statuses"></a>![Method: ](https://apidocs.io/img/method.png ".redemptionrequests_pkg.GetRedemptionRequestStatuses") GetRedemptionRequestStatuses

> Get redemption request statuses.


```go
func (me *REDEMPTIONREQUESTS_IMPL) GetRedemptionRequestStatuses()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = redemptionRequests.GetRedemptionRequestStatuses()

```


#### <a name="get_redemption_request_action"></a>![Method: ](https://apidocs.io/img/method.png ".redemptionrequests_pkg.GetRedemptionRequestAction") GetRedemptionRequestAction

> Get a redemption request action.


```go
func (me *REDEMPTIONREQUESTS_IMPL) GetRedemptionRequestAction(redemptionRequestActionSlug string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| redemptionRequestActionSlug |  ``` Required ```  | The redemption request action identifier |


#### Example Usage

```go
redemptionRequestActionSlug := "redemption_request_action_slug"

var result interface{}
result,_ = redemptionRequests.GetRedemptionRequestAction(redemptionRequestActionSlug)

```


#### <a name="get_redemption_request_actions"></a>![Method: ](https://apidocs.io/img/method.png ".redemptionrequests_pkg.GetRedemptionRequestActions") GetRedemptionRequestActions

> Get redemption request actions.


```go
func (me *REDEMPTIONREQUESTS_IMPL) GetRedemptionRequestActions()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = redemptionRequests.GetRedemptionRequestActions()

```


#### <a name="patch_redemption_request"></a>![Method: ](https://apidocs.io/img/method.png ".redemptionrequests_pkg.PatchRedemptionRequest") PatchRedemptionRequest

> Redeem a redemption request. After the redemption request is created it needs to be redeemed. This will deduct the amount of the advocate unclaimed balance and move the request to the completed state.


```go
func (me *REDEMPTIONREQUESTS_IMPL) PatchRedemptionRequest(
            accountSlug string,
            redemptionRequestId int64)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| redemptionRequestId |  ``` Required ```  | The redemption request id |


#### Example Usage

```go
accountSlug := "account_slug"
redemptionRequestId,_ := strconv.ParseInt("200", 10, 8)

var result 
result,_ = redemptionRequests.PatchRedemptionRequest(accountSlug, redemptionRequestId)

```


#### <a name="post_redemption_request"></a>![Method: ](https://apidocs.io/img/method.png ".redemptionrequests_pkg.PostRedemptionRequest") PostRedemptionRequest

> Create a redemption request.


```go
func (me *REDEMPTIONREQUESTS_IMPL) PostRedemptionRequest(
            accountSlug string,
            redemptionRequestForm *models_pkg.RedemptionRequestForm)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| redemptionRequestForm |  ``` Required ```  | The body of the request |


#### Example Usage

```go
accountSlug := "account_slug"
var redemptionRequestForm *models_pkg.RedemptionRequestForm

var result interface{}
result,_ = redemptionRequests.PostRedemptionRequest(accountSlug, redemptionRequestForm)

```


#### <a name="get_redemption_request"></a>![Method: ](https://apidocs.io/img/method.png ".redemptionrequests_pkg.GetRedemptionRequest") GetRedemptionRequest

> Get a redemption request by a given id.


```go
func (me *REDEMPTIONREQUESTS_IMPL) GetRedemptionRequest(
            accountSlug string,
            redemptionRequestId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| redemptionRequestId |  ``` Required ```  | The redemption request identifier |


#### Example Usage

```go
accountSlug := "account_slug"
redemptionRequestId := "redemption_request_id"

var result interface{}
result,_ = redemptionRequests.GetRedemptionRequest(accountSlug, redemptionRequestId)

```


#### <a name="get_redemption_requests"></a>![Method: ](https://apidocs.io/img/method.png ".redemptionrequests_pkg.GetRedemptionRequests") GetRedemptionRequests

> Get the list of redemption requests.


```go
func (me *REDEMPTIONREQUESTS_IMPL) GetRedemptionRequests(
            accountSlug string,
            page *string,
            limit *string,
            filter *string,
            sort *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| page |  ``` Optional ```  | Page number, e.g. 1 would start at the first result, and 10 would start at the tenth result. |
| limit |  ``` Optional ```  | Maximum number of results to return in the response. Default (10), threshold (100) |
| filter |  ``` Optional ```  | Allowed fields: redemption_request_id, name, lastname, email, request_status_slug, request_action_slug, from, to, created. Use the following delimiters to build your filters params. The vertical bar ('\|') to separate individual filter phrases and a double colon ('::') to separate the names and values. The delimiter of the double colon (':') separates the property name from the comparison value, enabling the comparison value to contain spaces. Example: www.example.com/users?filter='name::todd\|city::denver\|title::grand poobah' |
| sort |  ``` Optional ```  | Allowed fields: name, lastname, email, created. Use sort query-string parameter that contains a delimited set of property names. For each property name, sort in ascending order, and for each property prefixed with a dash ('-') sort in descending order. Separate each property name with a vertical bar ('\|'), which is consistent with the separation of the name\|value pairs in filtering, above. For example, if we want to retrieve users in order of their last name (ascending), first name (ascending) and hire date (descending), the request might look like this www.example.com/users?sort='last_name\|first_name\|-hire_date' |


#### Example Usage

```go
accountSlug := "account_slug"
page := "page"
limit := "limit"
filter := "filter"
sort := "sort"

var result interface{}
result,_ = redemptionRequests.GetRedemptionRequests(accountSlug, page, limit, filter, sort)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="bonuses_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".bonuses_pkg") bonuses_pkg

#### Get instance

Factory for the ``` BONUSES ``` interface can be accessed from the package bonuses_pkg.

```go
bonuses := bonuses_pkg.NewBONUSES()
```

#### <a name="get_bonuses"></a>![Method: ](https://apidocs.io/img/method.png ".bonuses_pkg.GetBonuses") GetBonuses

> Get the list of bonuses for a given account.


```go
func (me *BONUSES_IMPL) GetBonuses(
            accountSlug string,
            page *int64,
            limit *int64,
            filter *string,
            sort *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| page |  ``` Optional ```  | Page number, e.g. 1 would start at the first result, and 10 would start at the tenth result. |
| limit |  ``` Optional ```  | Maximum number of results to return in the response. Default (10), threshold (100) |
| filter |  ``` Optional ```  | Allowed fields: name, lastname, email, campaign_slug, from, to, created. Use the following delimiters to build your filters params. The vertical bar ('\|') to separate individual filter phrases and a double colon ('::') to separate the names and values. The delimiter of the double colon (':') separates the property name from the comparison value, enabling the comparison value to contain spaces. Example: www.example.com/users?filter='name::todd\|city::denver\|title::grand poobah' |
| sort |  ``` Optional ```  | Allowed fields: name, lastname, email, created. Use sort query-string parameter that contains a delimited set of property names. For each property name, sort in ascending order, and for each property prefixed with a dash ('-') sort in descending order. Separate each property name with a vertical bar ('\|'), which is consistent with the separation of the name\|value pairs in filtering, above. For example, if we want to retrieve users in order of their last name (ascending), first name (ascending) and hire date (descending), the request might look like this www.example.com/users?sort='last_name\|first_name\|-hire_date' |


#### Example Usage

```go
accountSlug := "account_slug"
page,_ := strconv.ParseInt("36", 10, 8)
limit,_ := strconv.ParseInt("36", 10, 8)
filter := "filter"
sort := "sort"

var result interface{}
result,_ = bonuses.GetBonuses(accountSlug, page, limit, filter, sort)

```


#### <a name="post_bonuses"></a>![Method: ](https://apidocs.io/img/method.png ".bonuses_pkg.PostBonuses") PostBonuses

> Make an attempt to give a bonus for to the advocate's referrer. The system processes the given advocate (referral) and creates a bonus for the advocate's referrer if is needed. All restrictions set on the campaigns for this account will be check out before giving the bonus to the advocate's referrer.


```go
func (me *BONUSES_IMPL) PostBonuses(
            accountSlug string,
            bonusesForm *models_pkg.BonusesForm)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| bonusesForm |  ``` Required ```  | The body of the request |


#### Example Usage

```go
accountSlug := "account_slug"
var bonusesForm *models_pkg.BonusesForm

var result interface{}
result,_ = bonuses.PostBonuses(accountSlug, bonusesForm)

```


#### <a name="get_bonuses_checkup"></a>![Method: ](https://apidocs.io/img/method.png ".bonuses_pkg.GetBonusesCheckup") GetBonusesCheckup

> Check if there is a bonus to be given to the advocate. Allows the clients to check if there is a bonus to be given, it simulates the behaivor of a POST request to /accounts/{account_slug}/bonuses resource. This resource is idempotent.


```go
func (me *BONUSES_IMPL) GetBonusesCheckup(
            accountSlug string,
            advocateToken string,
            reference string,
            paymentAmount float64)([]byte,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| advocateToken |  ``` Required ```  | The referral's token. Usually the one that completed the purchase, or trigger an event. |
| reference |  ``` Required ```  | The reference number for this request. Usually the order_id, payment_id, or timestamp. |
| paymentAmount |  ``` Required ```  | The payment amount the referrals has made. Required for a percentage based campaign. |


#### Example Usage

```go
accountSlug := "account_slug"
advocateToken := "advocate_token"
reference := "reference"
paymentAmount := 36.6559414736256

var result []byte
result,_ = bonuses.GetBonusesCheckup(accountSlug, advocateToken, reference, paymentAmount)

```


#### <a name="post_bonuses_force"></a>![Method: ](https://apidocs.io/img/method.png ".bonuses_pkg.PostBonusesForce") PostBonusesForce

> Force the system to give a bonus to an advocate. The system will not take into account the restriccions specified on the campaigns.


```go
func (me *BONUSES_IMPL) PostBonusesForce(
            accountSlug string,
            bonusForm *models_pkg.BonusesForm1)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| bonusForm |  ``` Required ```  | The body of the request |


#### Example Usage

```go
accountSlug := "account_slug"
var bonusForm *models_pkg.BonusesForm1

var result interface{}
result,_ = bonuses.PostBonusesForce(accountSlug, bonusForm)

```


#### <a name="get_bonuses_trace"></a>![Method: ](https://apidocs.io/img/method.png ".bonuses_pkg.GetBonusesTrace") GetBonusesTrace

> Get a bonus request trace.


```go
func (me *BONUSES_IMPL) GetBonusesTrace(
            accountSlug string,
            traceId int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| traceId |  ``` Required ```  | The trace id |


#### Example Usage

```go
accountSlug := "account_slug"
traceId,_ := strconv.ParseInt("36", 10, 8)

var result interface{}
result,_ = bonuses.GetBonusesTrace(accountSlug, traceId)

```


#### <a name="delete_bonus"></a>![Method: ](https://apidocs.io/img/method.png ".bonuses_pkg.DeleteBonus") DeleteBonus

> Delete a bonus


```go
func (me *BONUSES_IMPL) DeleteBonus(
            accountSlug string,
            bonusId int64)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| bonusId |  ``` Required ```  | The bonus id |


#### Example Usage

```go
accountSlug := "account_slug"
bonusId,_ := strconv.ParseInt("36", 10, 8)

var result 
result,_ = bonuses.DeleteBonus(accountSlug, bonusId)

```


#### <a name="get_bonus"></a>![Method: ](https://apidocs.io/img/method.png ".bonuses_pkg.GetBonus") GetBonus

> Get a bonus by a given id.


```go
func (me *BONUSES_IMPL) GetBonus(
            accountSlug string,
            bonusId int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| bonusId |  ``` Required ```  | The bonus id |


#### Example Usage

```go
accountSlug := "account_slug"
bonusId,_ := strconv.ParseInt("36", 10, 8)

var result interface{}
result,_ = bonuses.GetBonus(accountSlug, bonusId)

```


#### <a name="get_bonuses_traces"></a>![Method: ](https://apidocs.io/img/method.png ".bonuses_pkg.GetBonusesTraces") GetBonusesTraces

> Get the list of bonuses traces (audit trail). Every time the system tries to give a bonus the an advocate a new trace is created.


```go
func (me *BONUSES_IMPL) GetBonusesTraces(
            accountSlug string,
            page *int64,
            limit *int64,
            filter *string,
            sort *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| page |  ``` Optional ```  | Page number, e.g. 1 would start at the first result, and 10 would start at the tenth result. |
| limit |  ``` Optional ```  | Maximum number of results to return in the response. Default (10), threshold (100) |
| filter |  ``` Optional ```  | Allowed fields: reference, result, bonus_amount, advocate_token, advocate_referrer_token, campaign_slug, from, to, created. Use the following delimiters to build your filters params. The vertical bar ('\|') to separate individual filter phrases and a double colon ('::') to separate the names and values. The delimiter of the double colon (':') separates the property name from the comparison value, enabling the comparison value to contain spaces. Example: www.example.com/users?filter='name::todd\|city::denver\|title::grand poobah' |
| sort |  ``` Optional ```  | Allowed fields: created. Use sort query-string parameter that contains a delimited set of property names. For each property name, sort in ascending order, and for each property prefixed with a dash ('-') sort in descending order. Separate each property name with a vertical bar ('\|'), which is consistent with the separation of the name\|value pairs in filtering, above. For example, if we want to retrieve users in order of their last name (ascending), first name (ascending) and hire date (descending), the request might look like this www.example.com/users?sort='last_name\|first_name\|-hire_date' |


#### Example Usage

```go
accountSlug := "account_slug"
page,_ := strconv.ParseInt("36", 10, 8)
limit,_ := strconv.ParseInt("36", 10, 8)
filter := "filter"
sort := "sort"

var result interface{}
result,_ = bonuses.GetBonusesTraces(accountSlug, page, limit, filter, sort)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="campaigns_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".campaigns_pkg") campaigns_pkg

#### Get instance

Factory for the ``` CAMPAIGNS ``` interface can be accessed from the package campaigns_pkg.

```go
campaigns := campaigns_pkg.NewCAMPAIGNS()
```

#### <a name="get_campaign"></a>![Method: ](https://apidocs.io/img/method.png ".campaigns_pkg.GetCampaign") GetCampaign

> Get a campaign by a given slug.


```go
func (me *CAMPAIGNS_IMPL) GetCampaign(
            accountSlug string,
            campaignSlug string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| campaignSlug |  ``` Required ```  | The campaign identifier |


#### Example Usage

```go
accountSlug := "account_slug"
campaignSlug := "campaign_slug"

var result interface{}
result,_ = campaigns.GetCampaign(accountSlug, campaignSlug)

```


#### <a name="get_campaigns"></a>![Method: ](https://apidocs.io/img/method.png ".campaigns_pkg.GetCampaigns") GetCampaigns

> Get the list of campaings for a given account.


```go
func (me *CAMPAIGNS_IMPL) GetCampaigns(
            accountSlug string,
            page *int64,
            limit *int64,
            filter *string,
            sort *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSlug |  ``` Required ```  | The account identifier |
| page |  ``` Optional ```  | Page number, e.g. 1 would start at the first result, and 10 would start at the tenth result. |
| limit |  ``` Optional ```  | Maximum number of results to return in the response. Default (10), threshold (100) |
| filter |  ``` Optional ```  | Allowed fields: name, description, start_date, end_date, is_active (true\|false), created. Use the following delimiters to build your filters params. The vertical bar ('\|') to separate individual filter phrases and a double colon ('::') to separate the names and values. The delimiter of the double colon (':') separates the property name from the comparison value, enabling the comparison value to contain spaces. Example: www.example.com/users?filter='name::todd\|city::denver\|title::grand poobah' |
| sort |  ``` Optional ```  | Allowed fields: campaign_slug, created, start_date, end_date, is_active. Use sort query-string parameter that contains a delimited set of property names. For each property name, sort in ascending order, and for each property prefixed with a dash ('-') sort in descending order. Separate each property name with a vertical bar ('\|'), which is consistent with the separation of the name\|value pairs in filtering, above. For example, if we want to retrieve users in order of their last name (ascending), first name (ascending) and hire date (descending), the request might look like this www.example.com/users?sort='last_name\|first_name\|-hire_date' |


#### Example Usage

```go
accountSlug := "account_slug"
page,_ := strconv.ParseInt("36", 10, 8)
limit,_ := strconv.ParseInt("36", 10, 8)
filter := "filter"
sort := "sort"

var result interface{}
result,_ = campaigns.GetCampaigns(accountSlug, page, limit, filter, sort)

```


[Back to List of Controllers](#list_of_controllers)



