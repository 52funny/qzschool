package qzschool

// {
//     "errorCode":0,
//     "message":"成功",
//     "data":{
//         "accountId":001,
//         "projectId":367,
//         "identifier":null,
//         "sex":null,
//         "name":null,
//         "telPhone":"xxx",
//         "projectName":"xxxx洗浴",
//         "statusId":0,
//         "accountMoney":5070,
//         "accountGivenMoney":0,
//         "alias":"18225925413",
//         "loginCode":"xxxxxxxxxxxx",
//         "tags":"alLyrelease,367",
//         "isCard":0,
//         "cardStatusId":0,
//         "schoolDevice":null,
//         "isWxPrecharge":0,
//         "contractId":null,
//         "studentNo":null
//     }
// }
type loginResponse struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		AccountID    int    `json:"accountId"`
		ProjectID    int    `json:"projectId"`
		TelPhone     string `json:"telPhone"`
		ProjectName  string `json:"projectName"`
		StatusID     int    `json:"statusId"`
		AccountMoney int    `json:"accountMoney"`
		LoginCode    string `json:"loginCode"`
	} `json:"data"`
}

type versionResponse struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		Version    string      `json:"version"`
		UpdateCode int         `json:"updateCode"`
		Message    interface{} `json:"message"`
	} `json:"data"`
}

type accountMoneyResponse struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		Money string `json:"money"`
	} `json:"data"`
}

type billResponse struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Data      []Bill `json:"data"`
}

type Bill struct {
	DealDate      string  `json:"dealDate"`
	DealMark      string  `json:"dealMark"`
	UseCount      int     `json:"useCount"`
	AfterMoney    float64 `json:"afterMoney"`
	DealMoney     float64 `json:"dealMoney"`
	XfMoney       float64 `json:"xfMoney"`
	UpLeadMoney   float64 `json:"upLeadMoney"`
	UpMoney       float64 `json:"upMoney"`
	PerMoney      float64 `json:"perMoney"`
	TelPhone      string  `json:"telPhone"`
	AreaName      string  `json:"areaName"`
	AccountID     int     `json:"accountId"`
	GivenMark     string  `json:"givenMark"`
	CreditMark    int     `json:"creditMark"`
	OpName        string  `json:"opName"`
	ConsumeType   string  `json:"consumeType"`
	UpState       int     `json:"upState"`
	Description   string  `json:"description"`
	ComsumexfMode int     `json:"comsumexfMode"`
	ProjectID     int     `json:"projectId"`
}
