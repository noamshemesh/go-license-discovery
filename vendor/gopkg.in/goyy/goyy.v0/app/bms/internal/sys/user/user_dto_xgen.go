// generated by xgen -- DO NOT EDIT
package user

import (
	"gopkg.in/goyy/goyy.v0/data/dto"
)

type DTOs []DTO

type DTO struct {
	dto.Sys
	Name      string `json:"name"`
	Code      string `json:"code"`
	Passwd    string `json:"passwd"`
	Genre     string `json:"genre"`
	Email     string `json:"email"`
	Tel       string `json:"tel"`
	Mobile    string `json:"mobile"`
	AreaId    string `json:"areaId"`
	OrgId     string `json:"orgId"`
	LoginName string `json:"loginName"`
	LoginIp   string `json:"loginIp"`
	LoginTime int64  `json:"loginTime"`
}
