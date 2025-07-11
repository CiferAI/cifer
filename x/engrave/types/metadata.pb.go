// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cifer/engrave/metadata.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Metadata struct {
	CreatorWallet     string `protobuf:"bytes,1,opt,name=creatorWallet,proto3" json:"creatorWallet,omitempty"`
	GeneratedVia      string `protobuf:"bytes,2,opt,name=generatedVia,proto3" json:"generatedVia,omitempty"`
	ModelUsed         string `protobuf:"bytes,3,opt,name=modelUsed,proto3" json:"modelUsed,omitempty"`
	AssetId           string `protobuf:"bytes,4,opt,name=assetId,proto3" json:"assetId,omitempty"`
	AssetType         string `protobuf:"bytes,5,opt,name=assetType,proto3" json:"assetType,omitempty"`
	PromptAttribution string `protobuf:"bytes,6,opt,name=promptAttribution,proto3" json:"promptAttribution,omitempty"`
	IpOwnerName       string `protobuf:"bytes,7,opt,name=ipOwnerName,proto3" json:"ipOwnerName,omitempty"`
	IpOwnerWallet     string `protobuf:"bytes,8,opt,name=ipOwnerWallet,proto3" json:"ipOwnerWallet,omitempty"`
	UsageRights       string `protobuf:"bytes,9,opt,name=usageRights,proto3" json:"usageRights,omitempty"`
	LicenseExpiryDate string `protobuf:"bytes,10,opt,name=licenseExpiryDate,proto3" json:"licenseExpiryDate,omitempty"`
	CreatedOn         string `protobuf:"bytes,11,opt,name=createdOn,proto3" json:"createdOn,omitempty"`
	OnChainRecord     string `protobuf:"bytes,12,opt,name=onChainRecord,proto3" json:"onChainRecord,omitempty"`
	BlockHeight       int32  `protobuf:"varint,13,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaafcf95dd4ae959, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetCreatorWallet() string {
	if m != nil {
		return m.CreatorWallet
	}
	return ""
}

func (m *Metadata) GetGeneratedVia() string {
	if m != nil {
		return m.GeneratedVia
	}
	return ""
}

func (m *Metadata) GetModelUsed() string {
	if m != nil {
		return m.ModelUsed
	}
	return ""
}

func (m *Metadata) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *Metadata) GetAssetType() string {
	if m != nil {
		return m.AssetType
	}
	return ""
}

func (m *Metadata) GetPromptAttribution() string {
	if m != nil {
		return m.PromptAttribution
	}
	return ""
}

func (m *Metadata) GetIpOwnerName() string {
	if m != nil {
		return m.IpOwnerName
	}
	return ""
}

func (m *Metadata) GetIpOwnerWallet() string {
	if m != nil {
		return m.IpOwnerWallet
	}
	return ""
}

func (m *Metadata) GetUsageRights() string {
	if m != nil {
		return m.UsageRights
	}
	return ""
}

func (m *Metadata) GetLicenseExpiryDate() string {
	if m != nil {
		return m.LicenseExpiryDate
	}
	return ""
}

func (m *Metadata) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *Metadata) GetOnChainRecord() string {
	if m != nil {
		return m.OnChainRecord
	}
	return ""
}

func (m *Metadata) GetBlockHeight() int32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Metadata)(nil), "cifer.engrave.Metadata")
}

func init() { proto.RegisterFile("cifer/engrave/metadata.proto", fileDescriptor_eaafcf95dd4ae959) }

var fileDescriptor_eaafcf95dd4ae959 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcd, 0x4e, 0xfa, 0x40,
	0x14, 0xc5, 0xe9, 0x9f, 0x3f, 0x5f, 0x03, 0x2c, 0x6c, 0x62, 0x32, 0x0b, 0xd2, 0x10, 0xe2, 0x82,
	0x85, 0x81, 0x85, 0x4f, 0xe0, 0x57, 0xa2, 0x0b, 0x25, 0x69, 0xfc, 0x48, 0xdc, 0x0d, 0xed, 0xb5,
	0x4c, 0x6c, 0x67, 0x9a, 0x99, 0x8b, 0xc2, 0x5b, 0xf8, 0x58, 0x2e, 0x59, 0x19, 0x97, 0x06, 0x5e,
	0xc4, 0xcc, 0x4c, 0x45, 0x1a, 0x76, 0x9d, 0xdf, 0x39, 0xf7, 0xe6, 0x9c, 0xe6, 0x92, 0x5e, 0xc4,
	0x9f, 0x41, 0x8d, 0x41, 0x24, 0x8a, 0xbd, 0xc2, 0x38, 0x03, 0x64, 0x31, 0x43, 0x36, 0xca, 0x95,
	0x44, 0xe9, 0x77, 0xad, 0x3a, 0x2a, 0xd4, 0xc1, 0x67, 0x95, 0x34, 0x6f, 0x0a, 0x87, 0x7f, 0x44,
	0xba, 0x91, 0x02, 0x86, 0x52, 0x3d, 0xb2, 0x34, 0x05, 0xa4, 0x5e, 0xdf, 0x1b, 0xb6, 0xc2, 0x32,
	0xf4, 0x07, 0xa4, 0x93, 0x80, 0x00, 0xc5, 0x10, 0xe2, 0x07, 0xce, 0xe8, 0x3f, 0x6b, 0x2a, 0x31,
	0xbf, 0x47, 0x5a, 0x99, 0x8c, 0x21, 0xbd, 0xd7, 0x10, 0xd3, 0xaa, 0x35, 0xfc, 0x01, 0x9f, 0x92,
	0x06, 0xd3, 0x1a, 0xf0, 0x3a, 0xa6, 0xff, 0xad, 0xf6, 0xfb, 0x34, 0x73, 0xf6, 0xf3, 0x6e, 0x99,
	0x03, 0xad, 0xb9, 0xb9, 0x2d, 0xf0, 0x8f, 0xc9, 0x41, 0xae, 0x64, 0x96, 0xe3, 0x29, 0xa2, 0xe2,
	0xd3, 0x39, 0x72, 0x29, 0x68, 0xdd, 0xba, 0xf6, 0x05, 0xbf, 0x4f, 0xda, 0x3c, 0x9f, 0xbc, 0x09,
	0x50, 0xb7, 0x2c, 0x03, 0xda, 0xb0, 0xbe, 0x5d, 0x64, 0xfa, 0x16, 0xcf, 0xa2, 0x6f, 0xd3, 0xf5,
	0x2d, 0x41, 0xb3, 0x67, 0xae, 0x59, 0x02, 0x21, 0x4f, 0x66, 0xa8, 0x69, 0xcb, 0xed, 0xd9, 0x41,
	0x26, 0x57, 0xca, 0x23, 0x10, 0x1a, 0x2e, 0x17, 0x39, 0x57, 0xcb, 0x0b, 0x86, 0x40, 0x89, 0xcb,
	0xb5, 0x27, 0x98, 0x8e, 0xf6, 0x87, 0x42, 0x3c, 0x11, 0xb4, 0xed, 0x3a, 0x6e, 0x81, 0xc9, 0x24,
	0xc5, 0xf9, 0x8c, 0x71, 0x11, 0x42, 0x24, 0x55, 0x4c, 0x3b, 0x2e, 0x53, 0x09, 0x9a, 0x4c, 0xd3,
	0x54, 0x46, 0x2f, 0x57, 0x60, 0x12, 0xd0, 0x6e, 0xdf, 0x1b, 0xd6, 0xc2, 0x5d, 0x74, 0x36, 0xfe,
	0x58, 0x07, 0xde, 0x6a, 0x1d, 0x78, 0xdf, 0xeb, 0xc0, 0x7b, 0xdf, 0x04, 0x95, 0xd5, 0x26, 0xa8,
	0x7c, 0x6d, 0x82, 0xca, 0xd3, 0xa1, 0xbb, 0x8f, 0xc5, 0xf6, 0x42, 0x70, 0x99, 0x83, 0x9e, 0xd6,
	0xed, 0x7d, 0x9c, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x6f, 0xb6, 0xd8, 0x3f, 0x02, 0x00,
	0x00,
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintMetadata(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x68
	}
	if len(m.OnChainRecord) > 0 {
		i -= len(m.OnChainRecord)
		copy(dAtA[i:], m.OnChainRecord)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.OnChainRecord)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.CreatedOn) > 0 {
		i -= len(m.CreatedOn)
		copy(dAtA[i:], m.CreatedOn)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.CreatedOn)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.LicenseExpiryDate) > 0 {
		i -= len(m.LicenseExpiryDate)
		copy(dAtA[i:], m.LicenseExpiryDate)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.LicenseExpiryDate)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.UsageRights) > 0 {
		i -= len(m.UsageRights)
		copy(dAtA[i:], m.UsageRights)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.UsageRights)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.IpOwnerWallet) > 0 {
		i -= len(m.IpOwnerWallet)
		copy(dAtA[i:], m.IpOwnerWallet)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.IpOwnerWallet)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.IpOwnerName) > 0 {
		i -= len(m.IpOwnerName)
		copy(dAtA[i:], m.IpOwnerName)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.IpOwnerName)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.PromptAttribution) > 0 {
		i -= len(m.PromptAttribution)
		copy(dAtA[i:], m.PromptAttribution)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.PromptAttribution)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.AssetType) > 0 {
		i -= len(m.AssetType)
		copy(dAtA[i:], m.AssetType)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.AssetType)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AssetId) > 0 {
		i -= len(m.AssetId)
		copy(dAtA[i:], m.AssetId)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.AssetId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ModelUsed) > 0 {
		i -= len(m.ModelUsed)
		copy(dAtA[i:], m.ModelUsed)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.ModelUsed)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.GeneratedVia) > 0 {
		i -= len(m.GeneratedVia)
		copy(dAtA[i:], m.GeneratedVia)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.GeneratedVia)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CreatorWallet) > 0 {
		i -= len(m.CreatorWallet)
		copy(dAtA[i:], m.CreatorWallet)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.CreatorWallet)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CreatorWallet)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.GeneratedVia)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.ModelUsed)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.AssetId)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.AssetType)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.PromptAttribution)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.IpOwnerName)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.IpOwnerWallet)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.UsageRights)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.LicenseExpiryDate)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.CreatedOn)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.OnChainRecord)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovMetadata(uint64(m.BlockHeight))
	}
	return n
}

func sovMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetadata(x uint64) (n int) {
	return sovMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatorWallet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatorWallet = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GeneratedVia", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GeneratedVia = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelUsed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModelUsed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PromptAttribution", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PromptAttribution = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpOwnerName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpOwnerName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpOwnerWallet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpOwnerWallet = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsageRights", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UsageRights = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LicenseExpiryDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LicenseExpiryDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedOn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedOn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnChainRecord", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OnChainRecord = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadata
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetadata = fmt.Errorf("proto: unexpected end of group")
)
