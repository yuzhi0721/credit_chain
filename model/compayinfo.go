package model

type CompanyEssinfos struct {
	Id           int    `json:"id" form:"id"`
	Company      string `json:"company" form:"company"`
	Wallet_addr  string `json:"wallet_addr" form:"wallet_addr"`
	Uscc         string `json:"uscc" form:"uscc"`
	Logo         string `json:"logo" form:"logo"`
	Business_lic string `json:"business_lic" form:"business_lic"`
	Create_time  int    `json:"create_time" form:"create_time"`
	Update_time  int    `json:"update_time" form:"update_time"`
	Userid       int    `json:"userid" form:"userid"`
}
