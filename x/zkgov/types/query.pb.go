// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sdk/zkgov/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryCommitmentMerkleProofRequest
type QueryCommitmentMerkleProofRequest struct {
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Commitment string `protobuf:"bytes,2,opt,name=commitment,proto3" json:"commitment,omitempty"`
}

func (m *QueryCommitmentMerkleProofRequest) Reset()         { *m = QueryCommitmentMerkleProofRequest{} }
func (m *QueryCommitmentMerkleProofRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCommitmentMerkleProofRequest) ProtoMessage()    {}
func (*QueryCommitmentMerkleProofRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac2a1f040192f382, []int{0}
}
func (m *QueryCommitmentMerkleProofRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCommitmentMerkleProofRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCommitmentMerkleProofRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCommitmentMerkleProofRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCommitmentMerkleProofRequest.Merge(m, src)
}
func (m *QueryCommitmentMerkleProofRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCommitmentMerkleProofRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCommitmentMerkleProofRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCommitmentMerkleProofRequest proto.InternalMessageInfo

// QueryCommitmentMerkleProofResponse
type QueryCommitmentMerkleProofResponse struct {
	MerkleProof     [][]byte `protobuf:"bytes,1,rep,name=merkle_proof,json=merkleProof,proto3" json:"merkle_proof,omitempty"`
	Root            []byte   `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	CommitmentIndex uint64   `protobuf:"varint,3,opt,name=commitment_index,json=commitmentIndex,proto3" json:"commitment_index,omitempty"`
}

func (m *QueryCommitmentMerkleProofResponse) Reset()         { *m = QueryCommitmentMerkleProofResponse{} }
func (m *QueryCommitmentMerkleProofResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCommitmentMerkleProofResponse) ProtoMessage()    {}
func (*QueryCommitmentMerkleProofResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac2a1f040192f382, []int{1}
}
func (m *QueryCommitmentMerkleProofResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCommitmentMerkleProofResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCommitmentMerkleProofResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCommitmentMerkleProofResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCommitmentMerkleProofResponse.Merge(m, src)
}
func (m *QueryCommitmentMerkleProofResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCommitmentMerkleProofResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCommitmentMerkleProofResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCommitmentMerkleProofResponse proto.InternalMessageInfo

func (m *QueryCommitmentMerkleProofResponse) GetMerkleProof() [][]byte {
	if m != nil {
		return m.MerkleProof
	}
	return nil
}

func (m *QueryCommitmentMerkleProofResponse) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *QueryCommitmentMerkleProofResponse) GetCommitmentIndex() uint64 {
	if m != nil {
		return m.CommitmentIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryCommitmentMerkleProofRequest)(nil), "sdk.zkgov.v1beta1.QueryCommitmentMerkleProofRequest")
	proto.RegisterType((*QueryCommitmentMerkleProofResponse)(nil), "sdk.zkgov.v1beta1.QueryCommitmentMerkleProofResponse")
}

func init() { proto.RegisterFile("sdk/zkgov/v1beta1/query.proto", fileDescriptor_ac2a1f040192f382) }

var fileDescriptor_ac2a1f040192f382 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x4e, 0xc9, 0xd6,
	0xaf, 0xca, 0x4e, 0xcf, 0x2f, 0xd3, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x2c,
	0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2c, 0x4e, 0xc9, 0xd6, 0x03,
	0x4b, 0xeb, 0x41, 0xa5, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xb2, 0xfa, 0x20, 0x16, 0x44,
	0xa1, 0x94, 0x64, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0x71, 0x3c, 0x44, 0x02, 0xc2, 0x81, 0x4a, 0xc9,
	0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97,
	0x24, 0x96, 0x64, 0xe6, 0xe7, 0x41, 0x65, 0x95, 0xf2, 0xb8, 0x14, 0x03, 0x41, 0x16, 0x3a, 0xe7,
	0xe7, 0xe6, 0x66, 0x96, 0xe4, 0xa6, 0xe6, 0x95, 0xf8, 0xa6, 0x16, 0x65, 0xe7, 0xa4, 0x06, 0x14,
	0xe5, 0xe7, 0xa7, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xc9, 0x73, 0x71, 0x17, 0x14,
	0xe5, 0x17, 0xe4, 0x17, 0x27, 0xe6, 0xc4, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04,
	0x71, 0xc1, 0x84, 0x3c, 0x53, 0x84, 0xe4, 0xb8, 0xb8, 0x92, 0xe1, 0x06, 0x48, 0x30, 0x29, 0x30,
	0x6a, 0x70, 0x06, 0x21, 0x89, 0x58, 0x71, 0x74, 0x2c, 0x90, 0x67, 0x78, 0xb1, 0x40, 0x9e, 0x41,
	0xa9, 0x8d, 0x91, 0x4b, 0x09, 0x9f, 0x85, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x8a, 0x5c,
	0x3c, 0xb9, 0x60, 0x61, 0x90, 0x8f, 0xf2, 0xd3, 0x24, 0x18, 0x15, 0x98, 0x35, 0x78, 0x82, 0xb8,
	0x73, 0x11, 0x4a, 0x85, 0x84, 0xb8, 0x58, 0x8a, 0xf2, 0xf3, 0x21, 0xb6, 0xf1, 0x04, 0x81, 0xd9,
	0x42, 0x9a, 0x5c, 0x02, 0x08, 0x5b, 0xe3, 0x33, 0xf3, 0x52, 0x52, 0x2b, 0x24, 0x98, 0xc1, 0xae,
	0xe5, 0x47, 0x88, 0x7b, 0x82, 0x84, 0x8d, 0x76, 0x31, 0x72, 0xb1, 0x82, 0x1d, 0x22, 0xb4, 0x81,
	0x91, 0x4b, 0x14, 0xab, 0x6b, 0x84, 0x4c, 0xf4, 0x30, 0xc2, 0x5f, 0x8f, 0x60, 0x68, 0x49, 0x99,
	0x92, 0xa8, 0x0b, 0xe2, 0x65, 0x25, 0xad, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0xa9, 0x08, 0x29, 0xe9,
	0x57, 0x65, 0xeb, 0x82, 0x12, 0x44, 0x35, 0x52, 0x98, 0xd7, 0xea, 0x57, 0x23, 0x3c, 0x50, 0xeb,
	0xe4, 0x79, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xfa, 0xe9, 0x99, 0x25,
	0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x65, 0x99, 0xc5, 0x19, 0x89, 0x39, 0xba, 0xd9,
	0x89, 0x79, 0x79, 0x89, 0xfa, 0x55, 0xd9, 0x30, 0x73, 0x2b, 0xa0, 0x09, 0xae, 0xa4, 0xb2, 0x20,
	0xb5, 0x38, 0x89, 0x0d, 0x9c, 0x0e, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x88, 0xa2,
	0x6b, 0x8a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// CommitmentMerkleProof
	CommitmentMerkleProof(ctx context.Context, in *QueryCommitmentMerkleProofRequest, opts ...grpc.CallOption) (*QueryCommitmentMerkleProofResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CommitmentMerkleProof(ctx context.Context, in *QueryCommitmentMerkleProofRequest, opts ...grpc.CallOption) (*QueryCommitmentMerkleProofResponse, error) {
	out := new(QueryCommitmentMerkleProofResponse)
	err := c.cc.Invoke(ctx, "/sdk.zkgov.v1beta1.Query/CommitmentMerkleProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// CommitmentMerkleProof
	CommitmentMerkleProof(context.Context, *QueryCommitmentMerkleProofRequest) (*QueryCommitmentMerkleProofResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CommitmentMerkleProof(ctx context.Context, req *QueryCommitmentMerkleProofRequest) (*QueryCommitmentMerkleProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitmentMerkleProof not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CommitmentMerkleProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCommitmentMerkleProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CommitmentMerkleProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk.zkgov.v1beta1.Query/CommitmentMerkleProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CommitmentMerkleProof(ctx, req.(*QueryCommitmentMerkleProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sdk.zkgov.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommitmentMerkleProof",
			Handler:    _Query_CommitmentMerkleProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sdk/zkgov/v1beta1/query.proto",
}

func (m *QueryCommitmentMerkleProofRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCommitmentMerkleProofRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCommitmentMerkleProofRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Commitment) > 0 {
		i -= len(m.Commitment)
		copy(dAtA[i:], m.Commitment)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Commitment)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryCommitmentMerkleProofResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCommitmentMerkleProofResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCommitmentMerkleProofResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CommitmentIndex != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.CommitmentIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Root) > 0 {
		i -= len(m.Root)
		copy(dAtA[i:], m.Root)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Root)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MerkleProof) > 0 {
		for iNdEx := len(m.MerkleProof) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.MerkleProof[iNdEx])
			copy(dAtA[i:], m.MerkleProof[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.MerkleProof[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryCommitmentMerkleProofRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovQuery(uint64(m.ProposalId))
	}
	l = len(m.Commitment)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCommitmentMerkleProofResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MerkleProof) > 0 {
		for _, b := range m.MerkleProof {
			l = len(b)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.CommitmentIndex != 0 {
		n += 1 + sovQuery(uint64(m.CommitmentIndex))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCommitmentMerkleProofRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCommitmentMerkleProofRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCommitmentMerkleProofRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commitment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commitment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCommitmentMerkleProofResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCommitmentMerkleProofResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCommitmentMerkleProofResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MerkleProof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MerkleProof = append(m.MerkleProof, make([]byte, postIndex-iNdEx))
			copy(m.MerkleProof[len(m.MerkleProof)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = append(m.Root[:0], dAtA[iNdEx:postIndex]...)
			if m.Root == nil {
				m.Root = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitmentIndex", wireType)
			}
			m.CommitmentIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommitmentIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
