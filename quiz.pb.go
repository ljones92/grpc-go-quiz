// Code generated by protoc-gen-go. DO NOT EDIT.
// source: quiz.proto

package quiz

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type QuestionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuestionRequest) Reset()         { *m = QuestionRequest{} }
func (m *QuestionRequest) String() string { return proto.CompactTextString(m) }
func (*QuestionRequest) ProtoMessage()    {}
func (*QuestionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_quiz_00eb7f74ca50e653, []int{0}
}
func (m *QuestionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuestionRequest.Unmarshal(m, b)
}
func (m *QuestionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuestionRequest.Marshal(b, m, deterministic)
}
func (dst *QuestionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestionRequest.Merge(dst, src)
}
func (m *QuestionRequest) XXX_Size() int {
	return xxx_messageInfo_QuestionRequest.Size(m)
}
func (m *QuestionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuestionRequest proto.InternalMessageInfo

type QuestionResponse struct {
	Question             string   `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	QuestionId           int32    `protobuf:"varint,2,opt,name=questionId,proto3" json:"questionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuestionResponse) Reset()         { *m = QuestionResponse{} }
func (m *QuestionResponse) String() string { return proto.CompactTextString(m) }
func (*QuestionResponse) ProtoMessage()    {}
func (*QuestionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_quiz_00eb7f74ca50e653, []int{1}
}
func (m *QuestionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuestionResponse.Unmarshal(m, b)
}
func (m *QuestionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuestionResponse.Marshal(b, m, deterministic)
}
func (dst *QuestionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestionResponse.Merge(dst, src)
}
func (m *QuestionResponse) XXX_Size() int {
	return xxx_messageInfo_QuestionResponse.Size(m)
}
func (m *QuestionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuestionResponse proto.InternalMessageInfo

func (m *QuestionResponse) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *QuestionResponse) GetQuestionId() int32 {
	if m != nil {
		return m.QuestionId
	}
	return 0
}

type AnswerRequest struct {
	Answer               string   `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
	QuestionId           int32    `protobuf:"varint,2,opt,name=questionId,proto3" json:"questionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnswerRequest) Reset()         { *m = AnswerRequest{} }
func (m *AnswerRequest) String() string { return proto.CompactTextString(m) }
func (*AnswerRequest) ProtoMessage()    {}
func (*AnswerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_quiz_00eb7f74ca50e653, []int{2}
}
func (m *AnswerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnswerRequest.Unmarshal(m, b)
}
func (m *AnswerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnswerRequest.Marshal(b, m, deterministic)
}
func (dst *AnswerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnswerRequest.Merge(dst, src)
}
func (m *AnswerRequest) XXX_Size() int {
	return xxx_messageInfo_AnswerRequest.Size(m)
}
func (m *AnswerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnswerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnswerRequest proto.InternalMessageInfo

func (m *AnswerRequest) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func (m *AnswerRequest) GetQuestionId() int32 {
	if m != nil {
		return m.QuestionId
	}
	return 0
}

type AnswerResponse struct {
	Correct              bool     `protobuf:"varint,1,opt,name=correct,proto3" json:"correct,omitempty"`
	Answer               string   `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnswerResponse) Reset()         { *m = AnswerResponse{} }
func (m *AnswerResponse) String() string { return proto.CompactTextString(m) }
func (*AnswerResponse) ProtoMessage()    {}
func (*AnswerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_quiz_00eb7f74ca50e653, []int{3}
}
func (m *AnswerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnswerResponse.Unmarshal(m, b)
}
func (m *AnswerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnswerResponse.Marshal(b, m, deterministic)
}
func (dst *AnswerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnswerResponse.Merge(dst, src)
}
func (m *AnswerResponse) XXX_Size() int {
	return xxx_messageInfo_AnswerResponse.Size(m)
}
func (m *AnswerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnswerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnswerResponse proto.InternalMessageInfo

func (m *AnswerResponse) GetCorrect() bool {
	if m != nil {
		return m.Correct
	}
	return false
}

func (m *AnswerResponse) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func init() {
	proto.RegisterType((*QuestionRequest)(nil), "quiz.QuestionRequest")
	proto.RegisterType((*QuestionResponse)(nil), "quiz.QuestionResponse")
	proto.RegisterType((*AnswerRequest)(nil), "quiz.AnswerRequest")
	proto.RegisterType((*AnswerResponse)(nil), "quiz.AnswerResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QuizClient is the client API for Quiz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuizClient interface {
	GetQuestion(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*QuestionResponse, error)
	CheckAnswer(ctx context.Context, in *AnswerRequest, opts ...grpc.CallOption) (*AnswerResponse, error)
}

type quizClient struct {
	cc *grpc.ClientConn
}

func NewQuizClient(cc *grpc.ClientConn) QuizClient {
	return &quizClient{cc}
}

func (c *quizClient) GetQuestion(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*QuestionResponse, error) {
	out := new(QuestionResponse)
	err := c.cc.Invoke(ctx, "/quiz.Quiz/GetQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizClient) CheckAnswer(ctx context.Context, in *AnswerRequest, opts ...grpc.CallOption) (*AnswerResponse, error) {
	out := new(AnswerResponse)
	err := c.cc.Invoke(ctx, "/quiz.Quiz/CheckAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuizServer is the server API for Quiz service.
type QuizServer interface {
	GetQuestion(context.Context, *QuestionRequest) (*QuestionResponse, error)
	CheckAnswer(context.Context, *AnswerRequest) (*AnswerResponse, error)
}

func RegisterQuizServer(s *grpc.Server, srv QuizServer) {
	s.RegisterService(&_Quiz_serviceDesc, srv)
}

func _Quiz_GetQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServer).GetQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quiz.Quiz/GetQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServer).GetQuestion(ctx, req.(*QuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quiz_CheckAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServer).CheckAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quiz.Quiz/CheckAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServer).CheckAnswer(ctx, req.(*AnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Quiz_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quiz.Quiz",
	HandlerType: (*QuizServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuestion",
			Handler:    _Quiz_GetQuestion_Handler,
		},
		{
			MethodName: "CheckAnswer",
			Handler:    _Quiz_CheckAnswer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quiz.proto",
}

func init() { proto.RegisterFile("quiz.proto", fileDescriptor_quiz_00eb7f74ca50e653) }

var fileDescriptor_quiz_00eb7f74ca50e653 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2c, 0xcd, 0xac,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x04, 0xb9, 0xf8, 0x03, 0x4b,
	0x53, 0x8b, 0x4b, 0x32, 0xf3, 0xf3, 0x82, 0x52, 0x0b, 0x41, 0x2c, 0x25, 0x3f, 0x2e, 0x01, 0x84,
	0x50, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x14, 0x17, 0x47, 0x21, 0x54, 0x4c, 0x82, 0x51,
	0x81, 0x51, 0x83, 0x33, 0x08, 0xce, 0x17, 0x92, 0x03, 0x19, 0x0b, 0x61, 0x7b, 0xa6, 0x48, 0x30,
	0x29, 0x30, 0x6a, 0xb0, 0x06, 0x21, 0x89, 0x28, 0xb9, 0x73, 0xf1, 0x3a, 0xe6, 0x15, 0x97, 0xa7,
	0x16, 0x41, 0x2d, 0x10, 0x12, 0xe3, 0x62, 0x4b, 0x04, 0x0b, 0x40, 0x8d, 0x82, 0xf2, 0x08, 0x1a,
	0xe4, 0xc4, 0xc5, 0x07, 0x33, 0x08, 0xea, 0x2c, 0x09, 0x2e, 0xf6, 0xe4, 0xfc, 0xa2, 0xa2, 0xd4,
	0xe4, 0x12, 0xb0, 0x51, 0x1c, 0x41, 0x30, 0x2e, 0x92, 0x1d, 0x4c, 0xc8, 0x76, 0x18, 0x35, 0x31,
	0x72, 0xb1, 0x04, 0x96, 0x66, 0x56, 0x09, 0xd9, 0x71, 0x71, 0xbb, 0xa7, 0x96, 0xc0, 0x3c, 0x2a,
	0x24, 0xaa, 0x07, 0x0e, 0x1a, 0xb4, 0xb0, 0x90, 0x12, 0x43, 0x17, 0x86, 0x58, 0xac, 0xc4, 0x20,
	0x64, 0xc5, 0xc5, 0xed, 0x9c, 0x91, 0x9a, 0x9c, 0x0d, 0x71, 0x91, 0x90, 0x30, 0x44, 0x21, 0x8a,
	0x47, 0xa5, 0x44, 0x50, 0x05, 0x61, 0x7a, 0x93, 0xd8, 0xc0, 0x31, 0x60, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x75, 0x1f, 0x21, 0x6f, 0x8f, 0x01, 0x00, 0x00,
}
