package api

import (
	"encoding/json"
	_ "net/url"
)

type CreateLinodeRequest struct {
	DataCenterId int
	PlanId       int
	PaymentTerm  int
}

// SetDataCenter ... Use DataCenter object to set DataCenter.
func (c *CreateLinodeRequest) SetDataCenter(dc DataCenter) *CreateLinodeRequest {
	c.DataCenterId = dc.Id
	return c
}

func (c *CreateLinodeRequest) SetPlan(plan Plan) *CreateLinodeRequest {
	c.PlanId = plan.Id
	return c
}

func (c *CreateLinodeRequest) SetPaymentTerm(term int) *CreateLinodeRequest {
	c.PaymentTerm = term
	return c
}

func (c *LinodeClient) CreateLinode(request *CreateLinodeRequest) *Linode {
	params := c.InitialValues("linode.create")
	params.Add("DataCenterId", string(request.DataCenterId))
	params.Add("PlanId", string(request.PlanId))
	if request.PaymentTerm != 0 {
		params.Add("PaymentTerm", string(request.PaymentTerm))
	}
	response, err := c.HandleCall(params.Encode())
	if err != nil {
		println(err.Error())
	}
	data, err := GetRawJson(response)
	if err != nil {
		println(err.Error())
	}
	info := struct {
		LinodeId int
	}{
		0,
	}
	json.Unmarshal(data, &info)
	return &Linode{
		Id: info.LinodeId,
	}
}
