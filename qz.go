package qzschool

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// default client url
const defaultBackend = "http://wfwapi.china-qzxy.cn"

// default client url host
const host = "wfwapi.china-qzxy.cn"

// default user agent
const userAgent = "okhttp/4.2.2"

const (
	// default version
	// 	will get version from server
	DefaultVersion uint = 0
	// custom version
	// 	will set custom version to client
	CustomVersion = 1
)

type QzSchool struct {
	// login token
	LoginCode string
	// username
	Username string
	// password
	Password string
	// project id
	projectId string
	// version of client
	version string
	// http client
	client *http.Client
}

// init normal info qzschool struct
func NewQzSchool(username, password string) *QzSchool {
	client := http.DefaultClient
	return &QzSchool{
		Username: username,
		Password: password,
		client:   client,
	}
}

// init more info qzschool struct
func NewDetailQzSchool(username, password, loginCode, projectId, version string) *QzSchool {
	client := http.DefaultClient
	return &QzSchool{
		Username:  username,
		Password:  password,
		LoginCode: loginCode,
		projectId: projectId,
		version:   version,
		client:    client,
	}
}

// get login code from server
//	if versionType is defaultVersion this will get version from server
// 	if versionType is customVersion this will set custom version to client do not forgot set second param of version
func (q *QzSchool) Login(versionType uint, version ...string) error {
	// get server public version
	if versionType == DefaultVersion {
		v, err := q.getServerVersion()
		if err != nil {
			return err
		}
		version = append(version, v)
	}
	// if versionType is customVersion and version is large than one param
	if versionType == CustomVersion && len(version) != 1 {
		return errors.New("version param length is not correct")
	}
	// set client version
	q.version = version[0]

	formData := url.Values{
		"password":    {encodePass(q.Password)},
		"telPhone":    {q.Username},
		"openId":      {""},
		"typeId":      {"0"},
		"phoneSystem": {"android"},
		"version":     {q.version},
	}
	req, err := http.NewRequest("POST", defaultBackend+"/user/login?"+formData.Encode(), nil)
	if err != nil {
		return err
	}
	// set header
	setHeader(req)
	// send request
	resp, err := q.client.Do(req)
	if err != nil {
		return err
	}
	// close the body
	defer resp.Body.Close()
	// read the body
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(bs))
	loginRes := loginResponse{}
	err = json.Unmarshal(bs, &loginRes)
	if err != nil {
		return err
	}
	// check error
	if loginRes.ErrorCode != 0 {
		return errors.New(loginRes.Message)
	}

	// transfer int to string
	pId := strconv.Itoa(loginRes.Data.ProjectID)

	// set project id
	q.projectId = pId

	// set login code
	q.LoginCode = loginRes.Data.LoginCode

	return nil
}

// get server public version
func (q *QzSchool) getServerVersion() (version string, err error) {
	formData := url.Values{
		"phoneSystem": {"android"},
		"version":     {"6.2.8"},
	}
	req, err := http.NewRequest("GET", defaultBackend+"/user/version?"+formData.Encode(), nil)
	if err != nil {
		return "", err
	}
	setHeader(req)
	resp, err := q.client.Do(req)
	if err != nil {
		return "", err
	}
	// close the body
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("get server version failed")
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	v := &versionResponse{}
	err = json.Unmarshal(bs, v)
	if err != nil {
		return "", err
	}
	return v.Data.Version, nil
}

// set request header
func setHeader(req *http.Request) {
	req.Header.Set("Host", host)
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Connection", "close")
}

// encrypt password
func encodePass(password string) string {
	return fmt.Sprintf("%X", md5.Sum([]byte(password)))[22:]
}

// get account money
//	you should login first before you call this function
//	if has error will return -1 and error
func (q *QzSchool) GetAccountMoney() (float64, error) {
	if q.LoginCode == "" {
		return -1, errors.New("please login first")
	}
	formData := url.Values{
		"telPhone":    {q.Username},
		"loginCode":   {q.LoginCode},
		"projectId":   {q.projectId},
		"phoneSystem": {"android"},
		"version":     {q.version},
	}
	req, err := http.NewRequest("GET", defaultBackend+"/wallet/money?"+formData.Encode(), nil)
	if err != nil {
		return -1, err
	}
	resp, err := q.client.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return -1, errors.New("get account money failed")
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}
	accountMoneyRes := &accountMoneyResponse{}
	err = json.Unmarshal(bs, accountMoneyRes)
	if err != nil {
		return -1, err
	}
	if accountMoneyRes.ErrorCode != 0 {
		return -1, errors.New(accountMoneyRes.Message)
	}
	money, err := strconv.ParseFloat(accountMoneyRes.Data.Money, 64)
	if err != nil {
		return -1, err
	}
	return money, nil
}
