package api

import (
	_ "net/url"
)

type CreateLinodeRequest struct {
	DataCenterId int
	PlanId       int
	PaymentTerm  int
}

// SetDataCenter ... Use DataCenter object to set DataCenter.
/*func (c *CreateLinodeRequest) SetDataCenter(dc DataCenter) *CreateLinodeRequest {
	c.DataCenterId = dc.DataCenterId
	return c
}*/

func (c *CreateLinodeRequest) SetPlan() *CreateLinodeRequest {
	return c
}

func (c *LinodeClient) CreateLinode(request *CreateLinodeRequest) {

}
