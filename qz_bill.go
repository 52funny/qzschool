package qzschool

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	// all bill
	BillAllType = "0"
	// pay bill
	BillPayType = "1"
	// consume bill
	BillConsumeType = "2"
	// refund bill
	BillRefundType = "3"
	// other bill
	BillOtherType = "4"
)

// get bill info
//	typeId param is the type of bill
//	monthStr param is start date of bill just like "2020-01"
// 	page param is page number of bill. 20 items per page
func (q *QzSchool) GetPageMonthBill(typeId, monthStr string, page int) ([]Bill, error) {
	formData := url.Values{
		"telPhone":    {q.Username},
		"curNum":      {strconv.Itoa(page)},
		"loginCode":   {q.LoginCode},
		"typeId":      {typeId},
		"monthStr":    {monthStr},
		"projectId":   {q.projectId},
		"phoneSystem": {"android"},
		"version":     {q.version},
	}
	req, err := http.NewRequest("GET", defaultBackend+"/wallet/billList/month?"+formData.Encode(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := q.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("status code is not 200")
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bill := &billResponse{}
	err = json.Unmarshal(bs, bill)
	if err != nil {
		return nil, err
	}
	return bill.Data, nil
}

// get this month all bill info
//	typeId param is the type of bill
//	monthStr param is start date of bill just like "2020-01"
func (q *QzSchool) GetAllMonthBill(typeId, monthStr string) ([]Bill, error) {
	billSlice := make([]Bill, 0)
	i := 1
	for {
		bill, err := q.GetPageMonthBill(typeId, monthStr, i)
		if err != nil {
			return nil, err
		}
		billSlice = append(billSlice, bill...)
		if len(bill) == 0 {
			break
		}
		i++
	}
	return billSlice, nil
}
