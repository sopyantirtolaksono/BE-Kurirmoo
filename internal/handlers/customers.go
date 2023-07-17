package handlers

import (
	"context"
	"kurirmoo"
	"kurirmoo/gen/models"
	"kurirmoo/gen/restapi/operations/customers"
	"strings"
)

func (h *handler) Customers(ctx context.Context, rt *kurirmoo.Runtime, params customers.GetAllCustomerParams) ([]*customers.GetAllCustomerOKBodyItems0, error) {
	var customerList []*customers.GetAllCustomerOKBodyItems0

	rt.DB().WithContext(ctx).Model(&models.Customer{}).Select("id", "name", "ktp_number", "phone_number", "account_status", "is_corporate", "created_at").Find(&customerList)

	if params.Name != nil {
		rt.DB().WithContext(ctx).Model(&models.Customer{}).Where("LOWER(name) LIKE ?", "%"+strings.ToLower(*params.Name)+"%").Select("id", "name", "ktp_number", "phone_number", "account_status", "is_corporate", "created_at").Find(&customerList)
	}

	if params.Status != nil {
		if *params.Status == "both" {
			rt.DB().WithContext(ctx).Model(&models.Customer{}).Where("account_status IN ?", []string{"active", "non-active"}).Select("id", "name", "ktp_number", "phone_number", "account_status", "is_corporate", "created_at").Find(&customerList)
		} else {
			rt.DB().WithContext(ctx).Model(&models.Customer{}).Where("account_status = ?", params.Status).Select("id", "name", "ktp_number", "phone_number", "account_status", "is_corporate", "created_at").Find(&customerList)
		}
	}

	if params.Type != nil {
		if *params.Type == "corporate" {
			rt.DB().WithContext(ctx).Model(&models.Customer{}).Where("is_corporate = ?", true).Select("id", "name", "ktp_number", "phone_number", "account_status", "is_corporate", "created_at").Find(&customerList)
		} else if *params.Type == "personal" {
			rt.DB().WithContext(ctx).Model(&models.Customer{}).Where("is_corporate = ?", false).Select("id", "name", "ktp_number", "phone_number", "account_status", "is_corporate", "created_at").Find(&customerList)
		}
	}

	return customerList, nil
}
