// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: tv_service/user_info.proto

package tv_service

import (
	authentication_service_v2 "gitlab.sweet.tv/proto/authentication_service_v2"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetUserInfoResponse_Result int32

const (
	GetUserInfoResponse_OK     GetUserInfoResponse_Result = 0
	GetUserInfoResponse_NoAuth GetUserInfoResponse_Result = 1
)

// Enum value maps for GetUserInfoResponse_Result.
var (
	GetUserInfoResponse_Result_name = map[int32]string{
		0: "OK",
		1: "NoAuth",
	}
	GetUserInfoResponse_Result_value = map[string]int32{
		"OK":     0,
		"NoAuth": 1,
	}
)

func (x GetUserInfoResponse_Result) Enum() *GetUserInfoResponse_Result {
	p := new(GetUserInfoResponse_Result)
	*p = x
	return p
}

func (x GetUserInfoResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetUserInfoResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_tv_service_user_info_proto_enumTypes[0].Descriptor()
}

func (GetUserInfoResponse_Result) Type() protoreflect.EnumType {
	return &file_tv_service_user_info_proto_enumTypes[0]
}

func (x GetUserInfoResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *GetUserInfoResponse_Result) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = GetUserInfoResponse_Result(num)
	return nil
}

// Deprecated: Use GetUserInfoResponse_Result.Descriptor instead.
func (GetUserInfoResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_tv_service_user_info_proto_rawDescGZIP(), []int{2, 0}
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId                *int64     `protobuf:"varint,1,req,name=account_id,json=accountId" json:"account_id,omitempty"`
	Balance                  *float32   `protobuf:"fixed32,2,req,name=balance" json:"balance,omitempty"`
	Cost                     *float32   `protobuf:"fixed32,3,req,name=cost" json:"cost,omitempty"`
	Fullname                 *string    `protobuf:"bytes,4,req,name=fullname" json:"fullname,omitempty"`
	Tariff                   *string    `protobuf:"bytes,5,req,name=tariff" json:"tariff,omitempty"`
	TvPacks                  *string    `protobuf:"bytes,6,req,name=tv_packs,json=tvPacks" json:"tv_packs,omitempty"`
	IsVod                    *bool      `protobuf:"varint,7,req,name=is_vod,json=isVod" json:"is_vod,omitempty"`
	IsBlocked                *bool      `protobuf:"varint,8,req,name=is_blocked,json=isBlocked" json:"is_blocked,omitempty"`
	CompanyId                *int32     `protobuf:"varint,9,opt,name=company_id,json=companyId" json:"company_id,omitempty"`
	Login                    *string    `protobuf:"bytes,10,opt,name=login" json:"login,omitempty"`
	Password                 *string    `protobuf:"bytes,11,opt,name=password" json:"password,omitempty"`
	TariffId                 *int32     `protobuf:"varint,12,opt,name=tariff_id,json=tariffId" json:"tariff_id,omitempty"`
	PartnerId                *int32     `protobuf:"varint,13,opt,name=partner_id,json=partnerId" json:"partner_id,omitempty"`
	OnTest                   *bool      `protobuf:"varint,14,opt,name=on_test,json=onTest" json:"on_test,omitempty"`
	ToPay                    *int32     `protobuf:"varint,15,opt,name=to_pay,json=toPay" json:"to_pay,omitempty"`
	RealTariffId             *int32     `protobuf:"varint,16,opt,name=real_tariff_id,json=realTariffId" json:"real_tariff_id,omitempty"`
	SubscriptionId           *int32     `protobuf:"varint,17,opt,name=subscription_id,json=subscriptionId" json:"subscription_id,omitempty"`
	SubscriptionEndTime      *int32     `protobuf:"varint,18,opt,name=subscription_end_time,json=subscriptionEndTime" json:"subscription_end_time,omitempty"`
	Locale                   *string    `protobuf:"bytes,19,opt,name=locale" json:"locale,omitempty"`
	TariffPaidFor            *int32     `protobuf:"varint,20,opt,name=tariff_paid_for,json=tariffPaidFor" json:"tariff_paid_for,omitempty"`
	AutoUser                 *bool      `protobuf:"varint,21,opt,name=auto_user,json=autoUser" json:"auto_user,omitempty"`
	PromoCode                *string    `protobuf:"bytes,22,opt,name=promo_code,json=promoCode" json:"promo_code,omitempty"`
	Email                    *string    `protobuf:"bytes,23,opt,name=email" json:"email,omitempty"`
	Services                 []*Service `protobuf:"bytes,24,rep,name=services" json:"services,omitempty"`
	DateOfBirth              *int64     `protobuf:"zigzag64,25,opt,name=date_of_birth,json=dateOfBirth" json:"date_of_birth,omitempty"`
	CurrencyToPay            *int32     `protobuf:"varint,26,opt,name=currency_to_pay,json=currencyToPay" json:"currency_to_pay,omitempty"`
	CurrencyBalance          *float32   `protobuf:"fixed32,27,opt,name=currency_balance,json=currencyBalance" json:"currency_balance,omitempty"`
	CurrencyCost             *float32   `protobuf:"fixed32,28,opt,name=currency_cost,json=currencyCost" json:"currency_cost,omitempty"`
	AutopaymentEnabled       *bool      `protobuf:"varint,29,opt,name=autopayment_enabled,json=autopaymentEnabled" json:"autopayment_enabled,omitempty"`
	PhoneNumber              *string    `protobuf:"bytes,30,opt,name=phone_number,json=phoneNumber" json:"phone_number,omitempty"`
	NotificationDay          *int32     `protobuf:"varint,31,opt,name=notification_day,json=notificationDay" json:"notification_day,omitempty"`
	DateContract             *int32     `protobuf:"varint,32,opt,name=date_contract,json=dateContract" json:"date_contract,omitempty"`
	DateTariffBinding        *int32     `protobuf:"varint,33,opt,name=date_tariff_binding,json=dateTariffBinding" json:"date_tariff_binding,omitempty"`
	BoundTariffId            *int32     `protobuf:"varint,34,opt,name=bound_tariff_id,json=boundTariffId" json:"bound_tariff_id,omitempty"`
	SubscriptionStore        *int32     `protobuf:"varint,35,opt,name=subscription_store,json=subscriptionStore" json:"subscription_store,omitempty"`
	BloggerPromo             *string    `protobuf:"bytes,36,opt,name=blogger_promo,json=bloggerPromo" json:"blogger_promo,omitempty"`
	SubscriptionStoreEndTime *int32     `protobuf:"varint,37,opt,name=subscription_store_end_time,json=subscriptionStoreEndTime" json:"subscription_store_end_time,omitempty"`
	// Deprecated: Do not use.
	BloggerAbonCount *int32 `protobuf:"varint,38,opt,name=blogger_abon_count,json=bloggerAbonCount" json:"blogger_abon_count,omitempty"`
	// Deprecated: Do not use.
	SubscriptionStoreProductId *string                                  `protobuf:"bytes,39,opt,name=subscription_store_product_id,json=subscriptionStoreProductId" json:"subscription_store_product_id,omitempty"`
	ParentalControlEnabled     *bool                                    `protobuf:"varint,40,opt,name=parental_control_enabled,json=parentalControlEnabled" json:"parental_control_enabled,omitempty"`
	Referrer                   *int32                                   `protobuf:"varint,41,opt,name=referrer" json:"referrer,omitempty"`
	GeoZoneId                  *int32                                   `protobuf:"varint,42,opt,name=geo_zone_id,json=geoZoneId" json:"geo_zone_id,omitempty"`
	ContractId                 *int64                                   `protobuf:"varint,43,req,name=contract_id,json=contractId" json:"contract_id,omitempty"`
	Subscriptions              []*Subscription                          `protobuf:"bytes,44,rep,name=subscriptions" json:"subscriptions,omitempty"`
	SignupFullname             *string                                  `protobuf:"bytes,45,opt,name=signup_fullname,json=signupFullname" json:"signup_fullname,omitempty"`
	SignupNote                 *string                                  `protobuf:"bytes,46,opt,name=signup_note,json=signupNote" json:"signup_note,omitempty"`
	SignupPhoneNumber          *string                                  `protobuf:"bytes,47,opt,name=signup_phone_number,json=signupPhoneNumber" json:"signup_phone_number,omitempty"`
	SignupEmail                *string                                  `protobuf:"bytes,48,opt,name=signup_email,json=signupEmail" json:"signup_email,omitempty"`
	SignupOAuthProvider        *authentication_service_v2.OAuthProvider `protobuf:"varint,49,opt,name=signup_oAuthProvider,json=signupOAuthProvider,enum=authentication_service_v2.OAuthProvider" json:"signup_oAuthProvider,omitempty"`
	MonoSmartphone             *bool                                    `protobuf:"varint,100,opt,name=mono_smartphone,json=monoSmartphone" json:"mono_smartphone,omitempty"`
	NotificationsEnabled       *bool                                    `protobuf:"varint,200,opt,name=notifications_enabled,json=notificationsEnabled" json:"notifications_enabled,omitempty"`
	ReportRequested            *bool                                    `protobuf:"varint,201,opt,name=report_requested,json=reportRequested" json:"report_requested,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tv_service_user_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tv_service_user_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_tv_service_user_info_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetAccountId() int64 {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return 0
}

func (x *UserInfo) GetBalance() float32 {
	if x != nil && x.Balance != nil {
		return *x.Balance
	}
	return 0
}

func (x *UserInfo) GetCost() float32 {
	if x != nil && x.Cost != nil {
		return *x.Cost
	}
	return 0
}

func (x *UserInfo) GetFullname() string {
	if x != nil && x.Fullname != nil {
		return *x.Fullname
	}
	return ""
}

func (x *UserInfo) GetTariff() string {
	if x != nil && x.Tariff != nil {
		return *x.Tariff
	}
	return ""
}

func (x *UserInfo) GetTvPacks() string {
	if x != nil && x.TvPacks != nil {
		return *x.TvPacks
	}
	return ""
}

func (x *UserInfo) GetIsVod() bool {
	if x != nil && x.IsVod != nil {
		return *x.IsVod
	}
	return false
}

func (x *UserInfo) GetIsBlocked() bool {
	if x != nil && x.IsBlocked != nil {
		return *x.IsBlocked
	}
	return false
}

func (x *UserInfo) GetCompanyId() int32 {
	if x != nil && x.CompanyId != nil {
		return *x.CompanyId
	}
	return 0
}

func (x *UserInfo) GetLogin() string {
	if x != nil && x.Login != nil {
		return *x.Login
	}
	return ""
}

func (x *UserInfo) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *UserInfo) GetTariffId() int32 {
	if x != nil && x.TariffId != nil {
		return *x.TariffId
	}
	return 0
}

func (x *UserInfo) GetPartnerId() int32 {
	if x != nil && x.PartnerId != nil {
		return *x.PartnerId
	}
	return 0
}

func (x *UserInfo) GetOnTest() bool {
	if x != nil && x.OnTest != nil {
		return *x.OnTest
	}
	return false
}

func (x *UserInfo) GetToPay() int32 {
	if x != nil && x.ToPay != nil {
		return *x.ToPay
	}
	return 0
}

func (x *UserInfo) GetRealTariffId() int32 {
	if x != nil && x.RealTariffId != nil {
		return *x.RealTariffId
	}
	return 0
}

func (x *UserInfo) GetSubscriptionId() int32 {
	if x != nil && x.SubscriptionId != nil {
		return *x.SubscriptionId
	}
	return 0
}

func (x *UserInfo) GetSubscriptionEndTime() int32 {
	if x != nil && x.SubscriptionEndTime != nil {
		return *x.SubscriptionEndTime
	}
	return 0
}

func (x *UserInfo) GetLocale() string {
	if x != nil && x.Locale != nil {
		return *x.Locale
	}
	return ""
}

func (x *UserInfo) GetTariffPaidFor() int32 {
	if x != nil && x.TariffPaidFor != nil {
		return *x.TariffPaidFor
	}
	return 0
}

func (x *UserInfo) GetAutoUser() bool {
	if x != nil && x.AutoUser != nil {
		return *x.AutoUser
	}
	return false
}

func (x *UserInfo) GetPromoCode() string {
	if x != nil && x.PromoCode != nil {
		return *x.PromoCode
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UserInfo) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *UserInfo) GetDateOfBirth() int64 {
	if x != nil && x.DateOfBirth != nil {
		return *x.DateOfBirth
	}
	return 0
}

func (x *UserInfo) GetCurrencyToPay() int32 {
	if x != nil && x.CurrencyToPay != nil {
		return *x.CurrencyToPay
	}
	return 0
}

func (x *UserInfo) GetCurrencyBalance() float32 {
	if x != nil && x.CurrencyBalance != nil {
		return *x.CurrencyBalance
	}
	return 0
}

func (x *UserInfo) GetCurrencyCost() float32 {
	if x != nil && x.CurrencyCost != nil {
		return *x.CurrencyCost
	}
	return 0
}

func (x *UserInfo) GetAutopaymentEnabled() bool {
	if x != nil && x.AutopaymentEnabled != nil {
		return *x.AutopaymentEnabled
	}
	return false
}

func (x *UserInfo) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *UserInfo) GetNotificationDay() int32 {
	if x != nil && x.NotificationDay != nil {
		return *x.NotificationDay
	}
	return 0
}

func (x *UserInfo) GetDateContract() int32 {
	if x != nil && x.DateContract != nil {
		return *x.DateContract
	}
	return 0
}

func (x *UserInfo) GetDateTariffBinding() int32 {
	if x != nil && x.DateTariffBinding != nil {
		return *x.DateTariffBinding
	}
	return 0
}

func (x *UserInfo) GetBoundTariffId() int32 {
	if x != nil && x.BoundTariffId != nil {
		return *x.BoundTariffId
	}
	return 0
}

func (x *UserInfo) GetSubscriptionStore() int32 {
	if x != nil && x.SubscriptionStore != nil {
		return *x.SubscriptionStore
	}
	return 0
}

func (x *UserInfo) GetBloggerPromo() string {
	if x != nil && x.BloggerPromo != nil {
		return *x.BloggerPromo
	}
	return ""
}

func (x *UserInfo) GetSubscriptionStoreEndTime() int32 {
	if x != nil && x.SubscriptionStoreEndTime != nil {
		return *x.SubscriptionStoreEndTime
	}
	return 0
}

// Deprecated: Do not use.
func (x *UserInfo) GetBloggerAbonCount() int32 {
	if x != nil && x.BloggerAbonCount != nil {
		return *x.BloggerAbonCount
	}
	return 0
}

// Deprecated: Do not use.
func (x *UserInfo) GetSubscriptionStoreProductId() string {
	if x != nil && x.SubscriptionStoreProductId != nil {
		return *x.SubscriptionStoreProductId
	}
	return ""
}

func (x *UserInfo) GetParentalControlEnabled() bool {
	if x != nil && x.ParentalControlEnabled != nil {
		return *x.ParentalControlEnabled
	}
	return false
}

func (x *UserInfo) GetReferrer() int32 {
	if x != nil && x.Referrer != nil {
		return *x.Referrer
	}
	return 0
}

func (x *UserInfo) GetGeoZoneId() int32 {
	if x != nil && x.GeoZoneId != nil {
		return *x.GeoZoneId
	}
	return 0
}

func (x *UserInfo) GetContractId() int64 {
	if x != nil && x.ContractId != nil {
		return *x.ContractId
	}
	return 0
}

func (x *UserInfo) GetSubscriptions() []*Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

func (x *UserInfo) GetSignupFullname() string {
	if x != nil && x.SignupFullname != nil {
		return *x.SignupFullname
	}
	return ""
}

func (x *UserInfo) GetSignupNote() string {
	if x != nil && x.SignupNote != nil {
		return *x.SignupNote
	}
	return ""
}

func (x *UserInfo) GetSignupPhoneNumber() string {
	if x != nil && x.SignupPhoneNumber != nil {
		return *x.SignupPhoneNumber
	}
	return ""
}

func (x *UserInfo) GetSignupEmail() string {
	if x != nil && x.SignupEmail != nil {
		return *x.SignupEmail
	}
	return ""
}

func (x *UserInfo) GetSignupOAuthProvider() authentication_service_v2.OAuthProvider {
	if x != nil && x.SignupOAuthProvider != nil {
		return *x.SignupOAuthProvider
	}
	return authentication_service_v2.OAuthProvider(0)
}

func (x *UserInfo) GetMonoSmartphone() bool {
	if x != nil && x.MonoSmartphone != nil {
		return *x.MonoSmartphone
	}
	return false
}

func (x *UserInfo) GetNotificationsEnabled() bool {
	if x != nil && x.NotificationsEnabled != nil {
		return *x.NotificationsEnabled
	}
	return false
}

func (x *UserInfo) GetReportRequested() bool {
	if x != nil && x.ReportRequested != nil {
		return *x.ReportRequested
	}
	return false
}

type GetUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth *string `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
}

func (x *GetUserInfoRequest) Reset() {
	*x = GetUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tv_service_user_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoRequest) ProtoMessage() {}

func (x *GetUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tv_service_user_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoRequest.ProtoReflect.Descriptor instead.
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_tv_service_user_info_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserInfoRequest) GetAuth() string {
	if x != nil && x.Auth != nil {
		return *x.Auth
	}
	return ""
}

type GetUserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *GetUserInfoResponse_Result `protobuf:"varint,1,req,name=status,enum=tv_service.GetUserInfoResponse_Result" json:"status,omitempty"`
	Info   *UserInfo                   `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (x *GetUserInfoResponse) Reset() {
	*x = GetUserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tv_service_user_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResponse) ProtoMessage() {}

func (x *GetUserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tv_service_user_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResponse.ProtoReflect.Descriptor instead.
func (*GetUserInfoResponse) Descriptor() ([]byte, []int) {
	return file_tv_service_user_info_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserInfoResponse) GetStatus() GetUserInfoResponse_Result {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return GetUserInfoResponse_OK
}

func (x *GetUserInfoResponse) GetInfo() *UserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_tv_service_user_info_proto protoreflect.FileDescriptor

var file_tv_service_user_info_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x76,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x18, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x36, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x32, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x0f, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28, 0x02, 0x52, 0x04,
	0x63, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x18, 0x05, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x76, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x74, 0x76, 0x50, 0x61,
	0x63, 0x6b, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x76, 0x6f, 0x64, 0x18, 0x07, 0x20,
	0x02, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x56, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x08, 0x20, 0x02, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6f, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x74, 0x6f, 0x5f, 0x70, 0x61, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x74,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x72, 0x65, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x65, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x70, 0x61, 0x69, 0x64,
	0x5f, 0x66, 0x6f, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x50, 0x61, 0x69, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74,
	0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x75,
	0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6d,
	0x6f, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2f, 0x0a, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0d,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x12, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68,
	0x12, 0x26, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x6f, 0x5f,
	0x70, 0x61, 0x79, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x54, 0x6f, 0x50, 0x61, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x1b, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f,
	0x63, 0x6f, 0x73, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x61, 0x75, 0x74, 0x6f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x61, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x79,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x2e, 0x0a, 0x13,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x62, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x21, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0f,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x18,
	0x22, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x23, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x12, 0x3d, 0x0a, 0x1b, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x25, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x12, 0x62, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x5f, 0x61, 0x62, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x26, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x02, 0x18, 0x01, 0x52, 0x10, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x41, 0x62, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x1d, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x27, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x1a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x38, 0x0a, 0x18, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x16, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x18, 0x29, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0b, 0x67, 0x65, 0x6f, 0x5f, 0x7a, 0x6f,
	0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x67, 0x65, 0x6f,
	0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x2b, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x2c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x46, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x18,
	0x2e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x4e, 0x6f, 0x74,
	0x65, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x5f, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x2f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x30, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x5b, 0x0a, 0x14, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x5f, 0x6f,
	0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x31, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x28, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x32, 0x2e, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x13, 0x73, 0x69,
	0x67, 0x6e, 0x75, 0x70, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x6f, 0x6e, 0x6f, 0x5f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6d, 0x6f, 0x6e, 0x6f,
	0x53, 0x6d, 0x61, 0x72, 0x74, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x22, 0x28, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x9d, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74,
	0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x1c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x6f,
	0x41, 0x75, 0x74, 0x68, 0x10, 0x01, 0x42, 0x53, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x61,
	0x2e, 0x6d, 0x79, 0x74, 0x72, 0x69, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x76, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0d, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x73, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x74, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
}

var (
	file_tv_service_user_info_proto_rawDescOnce sync.Once
	file_tv_service_user_info_proto_rawDescData = file_tv_service_user_info_proto_rawDesc
)

func file_tv_service_user_info_proto_rawDescGZIP() []byte {
	file_tv_service_user_info_proto_rawDescOnce.Do(func() {
		file_tv_service_user_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_tv_service_user_info_proto_rawDescData)
	})
	return file_tv_service_user_info_proto_rawDescData
}

var file_tv_service_user_info_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tv_service_user_info_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tv_service_user_info_proto_goTypes = []interface{}{
	(GetUserInfoResponse_Result)(0),              // 0: tv_service.GetUserInfoResponse.Result
	(*UserInfo)(nil),                             // 1: tv_service.UserInfo
	(*GetUserInfoRequest)(nil),                   // 2: tv_service.GetUserInfoRequest
	(*GetUserInfoResponse)(nil),                  // 3: tv_service.GetUserInfoResponse
	(*Service)(nil),                              // 4: tv_service.Service
	(*Subscription)(nil),                         // 5: tv_service.Subscription
	(authentication_service_v2.OAuthProvider)(0), // 6: authentication_service_v2.OAuthProvider
}
var file_tv_service_user_info_proto_depIdxs = []int32{
	4, // 0: tv_service.UserInfo.services:type_name -> tv_service.Service
	5, // 1: tv_service.UserInfo.subscriptions:type_name -> tv_service.Subscription
	6, // 2: tv_service.UserInfo.signup_oAuthProvider:type_name -> authentication_service_v2.OAuthProvider
	0, // 3: tv_service.GetUserInfoResponse.status:type_name -> tv_service.GetUserInfoResponse.Result
	1, // 4: tv_service.GetUserInfoResponse.info:type_name -> tv_service.UserInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tv_service_user_info_proto_init() }
func file_tv_service_user_info_proto_init() {
	if File_tv_service_user_info_proto != nil {
		return
	}
	file_tv_service_billing_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tv_service_user_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tv_service_user_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tv_service_user_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tv_service_user_info_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tv_service_user_info_proto_goTypes,
		DependencyIndexes: file_tv_service_user_info_proto_depIdxs,
		EnumInfos:         file_tv_service_user_info_proto_enumTypes,
		MessageInfos:      file_tv_service_user_info_proto_msgTypes,
	}.Build()
	File_tv_service_user_info_proto = out.File
	file_tv_service_user_info_proto_rawDesc = nil
	file_tv_service_user_info_proto_goTypes = nil
	file_tv_service_user_info_proto_depIdxs = nil
}