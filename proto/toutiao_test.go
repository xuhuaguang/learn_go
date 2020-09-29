package proto

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"testing"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 设备类型
type BidRequestDeviceDeviceType int32

const (
	BidRequestDeviceDeviceUnknown BidRequestDeviceDeviceType = 0
	BidRequestDevicePhone         BidRequestDeviceDeviceType = 1 // 手机。
	BidRequestDeviceTablet        BidRequestDeviceDeviceType = 2 // 平板。
	BidRequestDeviceTv            BidRequestDeviceDeviceType = 3 // 智能电视。
)

var BidRequestDeviceDeviceTypeName = map[int32]string{
	0: "DEVICE_UNKNOWN",
	1: "PHONE",
	2: "TABLET",
	3: "TV",
}

var BidRequestDeviceDeviceTypeValue = map[string]int32{
	"DEVICE_UNKNOWN": 0,
	"PHONE":          1,
	"TABLET":         2,
	"TV":             3,
}

func (x BidRequestDeviceDeviceType) Enum() *BidRequestDeviceDeviceType {
	p := new(BidRequestDeviceDeviceType)
	*p = x
	return p
}

func (x BidRequestDeviceDeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

func TestDeviceType(t *testing.T) {
	enum := BidRequestDeviceDeviceType.Enum(1)
	fmt.Println(enum.Number())
}

type BidRequest struct {
	Version          *uint32  `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	RequestId        *string  `protobuf:"bytes,2,req,name=request_id" json:"request_id,omitempty"`
	Test             *bool    `protobuf:"varint,3,opt,name=test,def=0" json:"test,omitempty"`
	SspToken         *string  `protobuf:"bytes,4,req,name=ssp_token" json:"ssp_token,omitempty"`
	Ip               *string  `protobuf:"bytes,5,opt,name=ip" json:"ip,omitempty"`
	UserAgent        *string  `protobuf:"bytes,6,opt,name=user_agent" json:"user_agent,omitempty"`
	Language         *string  `protobuf:"bytes,7,opt,name=language" json:"language,omitempty"`
	Category         []uint32 `protobuf:"varint,9,rep,name=category" json:"category,omitempty"`
	ExcludedCategory []uint32 `protobuf:"varint,10,rep,name=excluded_category" json:"excluded_category,omitempty"`
	ExcludeDomain    []string `protobuf:"bytes,11,rep,name=exclude_domain" json:"exclude_domain,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BidRequest) Reset() {
	*m = BidRequest{}
}

func (m *BidRequest) String() string {
	return proto.CompactTextString(m)
}

func (*BidRequest) ProtoMessage() {

}

const Default_BidRequest_Test bool = false

func (m *BidRequest) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *BidRequest) GetRequestId() string {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return ""
}

func (m *BidRequest) GetTest() bool {
	if m != nil && m.Test != nil {
		return *m.Test
	}
	return Default_BidRequest_Test
}

type PublisherName string

const (
	PublisherZhanyueSDK PublisherName = "zy"
	PublisherApiApp     PublisherName = "client"
	PublisherApiOPPO    PublisherName = "oppo"
)

var PublisherMap = map[string]PublisherName{
	"zy":     PublisherZhanyueSDK,
	"client": PublisherApiApp,
	"oppo":   PublisherApiOPPO,
}

func TestPublisher(t *testing.T) {
	name := PublisherMap["zy"]
	fmt.Println(name)
}
