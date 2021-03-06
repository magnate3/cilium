// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/protocol.proto

package envoy_api_v2_core

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// [#not-implemented-hide:]
type TcpProtocolOptions struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpProtocolOptions) Reset()         { *m = TcpProtocolOptions{} }
func (m *TcpProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*TcpProtocolOptions) ProtoMessage()    {}
func (*TcpProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d86476e078060b60, []int{0}
}

func (m *TcpProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProtocolOptions.Unmarshal(m, b)
}
func (m *TcpProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProtocolOptions.Marshal(b, m, deterministic)
}
func (m *TcpProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProtocolOptions.Merge(m, src)
}
func (m *TcpProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_TcpProtocolOptions.Size(m)
}
func (m *TcpProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProtocolOptions proto.InternalMessageInfo

type HttpProtocolOptions struct {
	// The idle timeout for connections. The idle timeout is defined as the
	// period in which there are no active requests. If not set, there is no idle timeout. When the
	// idle timeout is reached the connection will be closed. If the connection is an HTTP/2
	// downstream connection a drain sequence will occur prior to closing the connection, see
	// :ref:`drain_timeout
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.HttpConnectionManager.drain_timeout>`.
	// Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.
	// If not specified, this defaults to 1 hour. To disable idle timeouts explicitly set this to 0.
	//
	// .. warning::
	//   Disabling this timeout has a highly likelihood of yielding connection leaks due to lost TCP
	//   FIN packets, etc.
	IdleTimeout *duration.Duration `protobuf:"bytes,1,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// The maximum duration of a connection. The duration is defined as a period since a connection
	// was established. If not set, there is no max duration. When max_connection_duration is reached
	// the connection will be closed. Drain sequence will occur prior to closing the connection if
	// if's applicable. See :ref:`drain_timeout
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.HttpConnectionManager.drain_timeout>`.
	// Note: not implemented for upstream connections.
	MaxConnectionDuration *duration.Duration `protobuf:"bytes,3,opt,name=max_connection_duration,json=maxConnectionDuration,proto3" json:"max_connection_duration,omitempty"`
	// The maximum number of headers. If unconfigured, the default
	// maximum number of request headers allowed is 100. Requests that exceed this limit will receive
	// a 431 response for HTTP/1.x and cause a stream reset for HTTP/2.
	MaxHeadersCount      *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_headers_count,json=maxHeadersCount,proto3" json:"max_headers_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HttpProtocolOptions) Reset()         { *m = HttpProtocolOptions{} }
func (m *HttpProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*HttpProtocolOptions) ProtoMessage()    {}
func (*HttpProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d86476e078060b60, []int{1}
}

func (m *HttpProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpProtocolOptions.Unmarshal(m, b)
}
func (m *HttpProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpProtocolOptions.Marshal(b, m, deterministic)
}
func (m *HttpProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpProtocolOptions.Merge(m, src)
}
func (m *HttpProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_HttpProtocolOptions.Size(m)
}
func (m *HttpProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_HttpProtocolOptions proto.InternalMessageInfo

func (m *HttpProtocolOptions) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *HttpProtocolOptions) GetMaxConnectionDuration() *duration.Duration {
	if m != nil {
		return m.MaxConnectionDuration
	}
	return nil
}

func (m *HttpProtocolOptions) GetMaxHeadersCount() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxHeadersCount
	}
	return nil
}

type Http1ProtocolOptions struct {
	// Handle HTTP requests with absolute URLs in the requests. These requests
	// are generally sent by clients to forward/explicit proxies. This allows clients to configure
	// envoy as their HTTP proxy. In Unix, for example, this is typically done by setting the
	// *http_proxy* environment variable.
	AllowAbsoluteUrl *wrappers.BoolValue `protobuf:"bytes,1,opt,name=allow_absolute_url,json=allowAbsoluteUrl,proto3" json:"allow_absolute_url,omitempty"`
	// Handle incoming HTTP/1.0 and HTTP 0.9 requests.
	// This is off by default, and not fully standards compliant. There is support for pre-HTTP/1.1
	// style connect logic, dechunking, and handling lack of client host iff
	// *default_host_for_http_10* is configured.
	AcceptHttp_10 bool `protobuf:"varint,2,opt,name=accept_http_10,json=acceptHttp10,proto3" json:"accept_http_10,omitempty"`
	// A default host for HTTP/1.0 requests. This is highly suggested if *accept_http_10* is true as
	// Envoy does not otherwise support HTTP/1.0 without a Host header.
	// This is a no-op if *accept_http_10* is not true.
	DefaultHostForHttp_10 string `protobuf:"bytes,3,opt,name=default_host_for_http_10,json=defaultHostForHttp10,proto3" json:"default_host_for_http_10,omitempty"`
	// Describes how the keys for response headers should be formatted. By default, all header keys
	// are lower cased.
	HeaderKeyFormat      *Http1ProtocolOptions_HeaderKeyFormat `protobuf:"bytes,4,opt,name=header_key_format,json=headerKeyFormat,proto3" json:"header_key_format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *Http1ProtocolOptions) Reset()         { *m = Http1ProtocolOptions{} }
func (m *Http1ProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*Http1ProtocolOptions) ProtoMessage()    {}
func (*Http1ProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d86476e078060b60, []int{2}
}

func (m *Http1ProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Http1ProtocolOptions.Unmarshal(m, b)
}
func (m *Http1ProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Http1ProtocolOptions.Marshal(b, m, deterministic)
}
func (m *Http1ProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Http1ProtocolOptions.Merge(m, src)
}
func (m *Http1ProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_Http1ProtocolOptions.Size(m)
}
func (m *Http1ProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Http1ProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_Http1ProtocolOptions proto.InternalMessageInfo

func (m *Http1ProtocolOptions) GetAllowAbsoluteUrl() *wrappers.BoolValue {
	if m != nil {
		return m.AllowAbsoluteUrl
	}
	return nil
}

func (m *Http1ProtocolOptions) GetAcceptHttp_10() bool {
	if m != nil {
		return m.AcceptHttp_10
	}
	return false
}

func (m *Http1ProtocolOptions) GetDefaultHostForHttp_10() string {
	if m != nil {
		return m.DefaultHostForHttp_10
	}
	return ""
}

func (m *Http1ProtocolOptions) GetHeaderKeyFormat() *Http1ProtocolOptions_HeaderKeyFormat {
	if m != nil {
		return m.HeaderKeyFormat
	}
	return nil
}

type Http1ProtocolOptions_HeaderKeyFormat struct {
	// Types that are valid to be assigned to HeaderFormat:
	//	*Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords_
	HeaderFormat         isHttp1ProtocolOptions_HeaderKeyFormat_HeaderFormat `protobuf_oneof:"header_format"`
	XXX_NoUnkeyedLiteral struct{}                                            `json:"-"`
	XXX_unrecognized     []byte                                              `json:"-"`
	XXX_sizecache        int32                                               `json:"-"`
}

func (m *Http1ProtocolOptions_HeaderKeyFormat) Reset()         { *m = Http1ProtocolOptions_HeaderKeyFormat{} }
func (m *Http1ProtocolOptions_HeaderKeyFormat) String() string { return proto.CompactTextString(m) }
func (*Http1ProtocolOptions_HeaderKeyFormat) ProtoMessage()    {}
func (*Http1ProtocolOptions_HeaderKeyFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_d86476e078060b60, []int{2, 0}
}

func (m *Http1ProtocolOptions_HeaderKeyFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Http1ProtocolOptions_HeaderKeyFormat.Unmarshal(m, b)
}
func (m *Http1ProtocolOptions_HeaderKeyFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Http1ProtocolOptions_HeaderKeyFormat.Marshal(b, m, deterministic)
}
func (m *Http1ProtocolOptions_HeaderKeyFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Http1ProtocolOptions_HeaderKeyFormat.Merge(m, src)
}
func (m *Http1ProtocolOptions_HeaderKeyFormat) XXX_Size() int {
	return xxx_messageInfo_Http1ProtocolOptions_HeaderKeyFormat.Size(m)
}
func (m *Http1ProtocolOptions_HeaderKeyFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_Http1ProtocolOptions_HeaderKeyFormat.DiscardUnknown(m)
}

var xxx_messageInfo_Http1ProtocolOptions_HeaderKeyFormat proto.InternalMessageInfo

type isHttp1ProtocolOptions_HeaderKeyFormat_HeaderFormat interface {
	isHttp1ProtocolOptions_HeaderKeyFormat_HeaderFormat()
}

type Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords_ struct {
	ProperCaseWords *Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords `protobuf:"bytes,1,opt,name=proper_case_words,json=properCaseWords,proto3,oneof"`
}

func (*Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords_) isHttp1ProtocolOptions_HeaderKeyFormat_HeaderFormat() {
}

func (m *Http1ProtocolOptions_HeaderKeyFormat) GetHeaderFormat() isHttp1ProtocolOptions_HeaderKeyFormat_HeaderFormat {
	if m != nil {
		return m.HeaderFormat
	}
	return nil
}

func (m *Http1ProtocolOptions_HeaderKeyFormat) GetProperCaseWords() *Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords {
	if x, ok := m.GetHeaderFormat().(*Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords_); ok {
		return x.ProperCaseWords
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Http1ProtocolOptions_HeaderKeyFormat) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords_)(nil),
	}
}

type Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords) Reset() {
	*m = Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords{}
}
func (m *Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords) String() string {
	return proto.CompactTextString(m)
}
func (*Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords) ProtoMessage() {}
func (*Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords) Descriptor() ([]byte, []int) {
	return fileDescriptor_d86476e078060b60, []int{2, 0, 0}
}

func (m *Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords.Unmarshal(m, b)
}
func (m *Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords.Marshal(b, m, deterministic)
}
func (m *Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords.Merge(m, src)
}
func (m *Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords) XXX_Size() int {
	return xxx_messageInfo_Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords.Size(m)
}
func (m *Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords) XXX_DiscardUnknown() {
	xxx_messageInfo_Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords.DiscardUnknown(m)
}

var xxx_messageInfo_Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords proto.InternalMessageInfo

// [#next-free-field: 13]
type Http2ProtocolOptions struct {
	// `Maximum table size <https://httpwg.org/specs/rfc7541.html#rfc.section.4.2>`_
	// (in octets) that the encoder is permitted to use for the dynamic HPACK table. Valid values
	// range from 0 to 4294967295 (2^32 - 1) and defaults to 4096. 0 effectively disables header
	// compression.
	HpackTableSize *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=hpack_table_size,json=hpackTableSize,proto3" json:"hpack_table_size,omitempty"`
	// `Maximum concurrent streams <https://httpwg.org/specs/rfc7540.html#rfc.section.5.1.2>`_
	// allowed for peer on one HTTP/2 connection. Valid values range from 1 to 2147483647 (2^31 - 1)
	// and defaults to 2147483647.
	MaxConcurrentStreams *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_concurrent_streams,json=maxConcurrentStreams,proto3" json:"max_concurrent_streams,omitempty"`
	// `Initial stream-level flow-control window
	// <https://httpwg.org/specs/rfc7540.html#rfc.section.6.9.2>`_ size. Valid values range from 65535
	// (2^16 - 1, HTTP/2 default) to 2147483647 (2^31 - 1, HTTP/2 maximum) and defaults to 268435456
	// (256 * 1024 * 1024).
	//
	// NOTE: 65535 is the initial window size from HTTP/2 spec. We only support increasing the default
	// window size now, so it's also the minimum.
	//
	// This field also acts as a soft limit on the number of bytes Envoy will buffer per-stream in the
	// HTTP/2 codec buffers. Once the buffer reaches this pointer, watermark callbacks will fire to
	// stop the flow of data to the codec buffers.
	InitialStreamWindowSize *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=initial_stream_window_size,json=initialStreamWindowSize,proto3" json:"initial_stream_window_size,omitempty"`
	// Similar to *initial_stream_window_size*, but for connection-level flow-control
	// window. Currently, this has the same minimum/maximum/default as *initial_stream_window_size*.
	InitialConnectionWindowSize *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=initial_connection_window_size,json=initialConnectionWindowSize,proto3" json:"initial_connection_window_size,omitempty"`
	// Allows proxying Websocket and other upgrades over H2 connect.
	AllowConnect bool `protobuf:"varint,5,opt,name=allow_connect,json=allowConnect,proto3" json:"allow_connect,omitempty"`
	// [#not-implemented-hide:] Hiding until envoy has full metadata support.
	// Still under implementation. DO NOT USE.
	//
	// Allows metadata. See [metadata
	// docs](https://github.com/envoyproxy/envoy/blob/master/source/docs/h2_metadata.md) for more
	// information.
	AllowMetadata bool `protobuf:"varint,6,opt,name=allow_metadata,json=allowMetadata,proto3" json:"allow_metadata,omitempty"`
	// Limit the number of pending outbound downstream frames of all types (frames that are waiting to
	// be written into the socket). Exceeding this limit triggers flood mitigation and connection is
	// terminated. The ``http2.outbound_flood`` stat tracks the number of terminated connections due
	// to flood mitigation. The default limit is 10000.
	// [#comment:TODO: implement same limits for upstream outbound frames as well.]
	MaxOutboundFrames *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=max_outbound_frames,json=maxOutboundFrames,proto3" json:"max_outbound_frames,omitempty"`
	// Limit the number of pending outbound downstream frames of types PING, SETTINGS and RST_STREAM,
	// preventing high memory utilization when receiving continuous stream of these frames. Exceeding
	// this limit triggers flood mitigation and connection is terminated. The
	// ``http2.outbound_control_flood`` stat tracks the number of terminated connections due to flood
	// mitigation. The default limit is 1000.
	// [#comment:TODO: implement same limits for upstream outbound frames as well.]
	MaxOutboundControlFrames *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=max_outbound_control_frames,json=maxOutboundControlFrames,proto3" json:"max_outbound_control_frames,omitempty"`
	// Limit the number of consecutive inbound frames of types HEADERS, CONTINUATION and DATA with an
	// empty payload and no end stream flag. Those frames have no legitimate use and are abusive, but
	// might be a result of a broken HTTP/2 implementation. The `http2.inbound_empty_frames_flood``
	// stat tracks the number of connections terminated due to flood mitigation.
	// Setting this to 0 will terminate connection upon receiving first frame with an empty payload
	// and no end stream flag. The default limit is 1.
	// [#comment:TODO: implement same limits for upstream inbound frames as well.]
	MaxConsecutiveInboundFramesWithEmptyPayload *wrappers.UInt32Value `protobuf:"bytes,9,opt,name=max_consecutive_inbound_frames_with_empty_payload,json=maxConsecutiveInboundFramesWithEmptyPayload,proto3" json:"max_consecutive_inbound_frames_with_empty_payload,omitempty"`
	// Limit the number of inbound PRIORITY frames allowed per each opened stream. If the number
	// of PRIORITY frames received over the lifetime of connection exceeds the value calculated
	// using this formula::
	//
	//     max_inbound_priority_frames_per_stream * (1 + inbound_streams)
	//
	// the connection is terminated. The ``http2.inbound_priority_frames_flood`` stat tracks
	// the number of connections terminated due to flood mitigation. The default limit is 100.
	// [#comment:TODO: implement same limits for upstream inbound frames as well.]
	MaxInboundPriorityFramesPerStream *wrappers.UInt32Value `protobuf:"bytes,10,opt,name=max_inbound_priority_frames_per_stream,json=maxInboundPriorityFramesPerStream,proto3" json:"max_inbound_priority_frames_per_stream,omitempty"`
	// Limit the number of inbound WINDOW_UPDATE frames allowed per DATA frame sent. If the number
	// of WINDOW_UPDATE frames received over the lifetime of connection exceeds the value calculated
	// using this formula::
	//
	//     1 + 2 * (inbound_streams +
	//              max_inbound_window_update_frames_per_data_frame_sent * outbound_data_frames)
	//
	// the connection is terminated. The ``http2.inbound_priority_frames_flood`` stat tracks
	// the number of connections terminated due to flood mitigation. The default limit is 10.
	// Setting this to 1 should be enough to support HTTP/2 implementations with basic flow control,
	// but more complex implementations that try to estimate available bandwidth require at least 2.
	// [#comment:TODO: implement same limits for upstream inbound frames as well.]
	MaxInboundWindowUpdateFramesPerDataFrameSent *wrappers.UInt32Value `protobuf:"bytes,11,opt,name=max_inbound_window_update_frames_per_data_frame_sent,json=maxInboundWindowUpdateFramesPerDataFrameSent,proto3" json:"max_inbound_window_update_frames_per_data_frame_sent,omitempty"`
	// Allows invalid HTTP messaging and headers. When this option is disabled (default), then
	// the whole HTTP/2 connection is terminated upon receiving invalid HEADERS frame. However,
	// when this option is enabled, only the offending stream is terminated.
	//
	// See `RFC7540, sec. 8.1 <https://tools.ietf.org/html/rfc7540#section-8.1>`_ for details.
	StreamErrorOnInvalidHttpMessaging bool     `protobuf:"varint,12,opt,name=stream_error_on_invalid_http_messaging,json=streamErrorOnInvalidHttpMessaging,proto3" json:"stream_error_on_invalid_http_messaging,omitempty"`
	XXX_NoUnkeyedLiteral              struct{} `json:"-"`
	XXX_unrecognized                  []byte   `json:"-"`
	XXX_sizecache                     int32    `json:"-"`
}

func (m *Http2ProtocolOptions) Reset()         { *m = Http2ProtocolOptions{} }
func (m *Http2ProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*Http2ProtocolOptions) ProtoMessage()    {}
func (*Http2ProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d86476e078060b60, []int{3}
}

func (m *Http2ProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Http2ProtocolOptions.Unmarshal(m, b)
}
func (m *Http2ProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Http2ProtocolOptions.Marshal(b, m, deterministic)
}
func (m *Http2ProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Http2ProtocolOptions.Merge(m, src)
}
func (m *Http2ProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_Http2ProtocolOptions.Size(m)
}
func (m *Http2ProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Http2ProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_Http2ProtocolOptions proto.InternalMessageInfo

func (m *Http2ProtocolOptions) GetHpackTableSize() *wrappers.UInt32Value {
	if m != nil {
		return m.HpackTableSize
	}
	return nil
}

func (m *Http2ProtocolOptions) GetMaxConcurrentStreams() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConcurrentStreams
	}
	return nil
}

func (m *Http2ProtocolOptions) GetInitialStreamWindowSize() *wrappers.UInt32Value {
	if m != nil {
		return m.InitialStreamWindowSize
	}
	return nil
}

func (m *Http2ProtocolOptions) GetInitialConnectionWindowSize() *wrappers.UInt32Value {
	if m != nil {
		return m.InitialConnectionWindowSize
	}
	return nil
}

func (m *Http2ProtocolOptions) GetAllowConnect() bool {
	if m != nil {
		return m.AllowConnect
	}
	return false
}

func (m *Http2ProtocolOptions) GetAllowMetadata() bool {
	if m != nil {
		return m.AllowMetadata
	}
	return false
}

func (m *Http2ProtocolOptions) GetMaxOutboundFrames() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxOutboundFrames
	}
	return nil
}

func (m *Http2ProtocolOptions) GetMaxOutboundControlFrames() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxOutboundControlFrames
	}
	return nil
}

func (m *Http2ProtocolOptions) GetMaxConsecutiveInboundFramesWithEmptyPayload() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConsecutiveInboundFramesWithEmptyPayload
	}
	return nil
}

func (m *Http2ProtocolOptions) GetMaxInboundPriorityFramesPerStream() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxInboundPriorityFramesPerStream
	}
	return nil
}

func (m *Http2ProtocolOptions) GetMaxInboundWindowUpdateFramesPerDataFrameSent() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxInboundWindowUpdateFramesPerDataFrameSent
	}
	return nil
}

func (m *Http2ProtocolOptions) GetStreamErrorOnInvalidHttpMessaging() bool {
	if m != nil {
		return m.StreamErrorOnInvalidHttpMessaging
	}
	return false
}

// [#not-implemented-hide:]
type GrpcProtocolOptions struct {
	Http2ProtocolOptions *Http2ProtocolOptions `protobuf:"bytes,1,opt,name=http2_protocol_options,json=http2ProtocolOptions,proto3" json:"http2_protocol_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GrpcProtocolOptions) Reset()         { *m = GrpcProtocolOptions{} }
func (m *GrpcProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*GrpcProtocolOptions) ProtoMessage()    {}
func (*GrpcProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d86476e078060b60, []int{4}
}

func (m *GrpcProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcProtocolOptions.Unmarshal(m, b)
}
func (m *GrpcProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcProtocolOptions.Marshal(b, m, deterministic)
}
func (m *GrpcProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcProtocolOptions.Merge(m, src)
}
func (m *GrpcProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_GrpcProtocolOptions.Size(m)
}
func (m *GrpcProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcProtocolOptions proto.InternalMessageInfo

func (m *GrpcProtocolOptions) GetHttp2ProtocolOptions() *Http2ProtocolOptions {
	if m != nil {
		return m.Http2ProtocolOptions
	}
	return nil
}

func init() {
	proto.RegisterType((*TcpProtocolOptions)(nil), "envoy.api.v2.core.TcpProtocolOptions")
	proto.RegisterType((*HttpProtocolOptions)(nil), "envoy.api.v2.core.HttpProtocolOptions")
	proto.RegisterType((*Http1ProtocolOptions)(nil), "envoy.api.v2.core.Http1ProtocolOptions")
	proto.RegisterType((*Http1ProtocolOptions_HeaderKeyFormat)(nil), "envoy.api.v2.core.Http1ProtocolOptions.HeaderKeyFormat")
	proto.RegisterType((*Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords)(nil), "envoy.api.v2.core.Http1ProtocolOptions.HeaderKeyFormat.ProperCaseWords")
	proto.RegisterType((*Http2ProtocolOptions)(nil), "envoy.api.v2.core.Http2ProtocolOptions")
	proto.RegisterType((*GrpcProtocolOptions)(nil), "envoy.api.v2.core.GrpcProtocolOptions")
}

func init() { proto.RegisterFile("envoy/api/v2/core/protocol.proto", fileDescriptor_d86476e078060b60) }

var fileDescriptor_d86476e078060b60 = []byte{
	// 935 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x73, 0x1b, 0x35,
	0x14, 0x67, 0xeb, 0xd0, 0x24, 0xca, 0x87, 0xeb, 0x8d, 0x69, 0x16, 0x97, 0x29, 0xad, 0x81, 0x92,
	0x09, 0xcc, 0xba, 0x71, 0x19, 0xb8, 0x70, 0x61, 0xd3, 0xa6, 0xce, 0x30, 0x9d, 0xb8, 0x9b, 0x84,
	0x9c, 0x18, 0x8d, 0xbc, 0xab, 0xc4, 0xa2, 0xbb, 0x92, 0x46, 0xd2, 0xfa, 0xa3, 0x67, 0x2e, 0x1c,
	0xf9, 0x53, 0xf8, 0x6b, 0xf8, 0x57, 0x20, 0x17, 0x33, 0xfa, 0xd8, 0xd4, 0xb1, 0x0b, 0x38, 0x39,
	0x79, 0xf7, 0xe9, 0xf7, 0xf1, 0x9e, 0xf4, 0xf4, 0xd6, 0xe0, 0x11, 0xa6, 0x03, 0x36, 0x6e, 0x21,
	0x4e, 0x5a, 0x83, 0x76, 0x2b, 0x61, 0x02, 0xb7, 0xb8, 0x60, 0x8a, 0x25, 0x2c, 0x0b, 0xcd, 0x83,
	0x5f, 0x33, 0x88, 0x10, 0x71, 0x12, 0x0e, 0xda, 0xa1, 0x46, 0x34, 0x1e, 0x5e, 0x30, 0x76, 0x91,
	0x39, 0x64, 0xaf, 0x38, 0x6f, 0xa5, 0x85, 0x40, 0x8a, 0x30, 0x6a, 0x29, 0xf3, 0xeb, 0x43, 0x81,
	0x38, 0xc7, 0x42, 0xba, 0xf5, 0xed, 0x01, 0xca, 0x48, 0x8a, 0x14, 0x6e, 0x95, 0x0f, 0x76, 0xa1,
	0x59, 0x07, 0xfe, 0x49, 0xc2, 0xbb, 0x2e, 0x81, 0x23, 0xae, 0x35, 0x65, 0xf3, 0x6f, 0x0f, 0x6c,
	0x75, 0x94, 0x9a, 0x8d, 0xfb, 0xdf, 0x83, 0x75, 0x92, 0x66, 0x18, 0x2a, 0x92, 0x63, 0x56, 0xa8,
	0xc0, 0x7b, 0xe4, 0xed, 0xac, 0xb5, 0x3f, 0x0e, 0xad, 0x7b, 0x58, 0xba, 0x87, 0xcf, 0x5d, 0x76,
	0xf1, 0x9a, 0x86, 0x9f, 0x58, 0xb4, 0xff, 0x1a, 0x6c, 0xe7, 0x68, 0x04, 0x13, 0x46, 0x29, 0x4e,
	0xf4, 0x32, 0x2c, 0xab, 0x08, 0x2a, 0xff, 0x27, 0xf4, 0x51, 0x8e, 0x46, 0xfb, 0x57, 0xc4, 0x32,
	0xec, 0xbf, 0x06, 0x35, 0x2d, 0xd9, 0xc7, 0x28, 0xc5, 0x42, 0xc2, 0x84, 0x15, 0x54, 0x05, 0x77,
	0x8c, 0xd8, 0x27, 0x73, 0x62, 0xa7, 0x87, 0x54, 0x3d, 0x6b, 0xff, 0x84, 0xb2, 0x02, 0x47, 0xcb,
	0x97, 0xd1, 0xd2, 0xee, 0x9d, 0x1d, 0x2f, 0xae, 0xe6, 0x68, 0xd4, 0xb1, 0xf4, 0x7d, 0xcd, 0x6e,
	0xfe, 0x59, 0x01, 0x75, 0x5d, 0xfb, 0xde, 0x6c, 0xf1, 0x1d, 0xe0, 0xa3, 0x2c, 0x63, 0x43, 0x88,
	0x7a, 0x92, 0x65, 0x85, 0xc2, 0xb0, 0x10, 0x99, 0xdb, 0x82, 0xc6, 0x9c, 0x59, 0xc4, 0x58, 0x66,
	0xac, 0xe2, 0x7b, 0x86, 0xf5, 0x83, 0x23, 0x9d, 0x8a, 0xcc, 0xff, 0x1c, 0x6c, 0xa2, 0x24, 0xc1,
	0x5c, 0xc1, 0xbe, 0x52, 0x1c, 0xee, 0x3d, 0x35, 0x29, 0xaf, 0xc4, 0xeb, 0x36, 0x6a, 0xdc, 0x9f,
	0xfa, 0xdf, 0x82, 0x20, 0xc5, 0xe7, 0xa8, 0xc8, 0x14, 0xec, 0x33, 0xa9, 0xe0, 0x39, 0x13, 0x57,
	0x78, 0xbd, 0x5f, 0xab, 0x71, 0xdd, 0xad, 0x77, 0x98, 0x54, 0x07, 0x4c, 0x38, 0x5e, 0x02, 0x6a,
	0x76, 0x3f, 0xe0, 0x1b, 0x3c, 0xd6, 0xac, 0x1c, 0xa9, 0x60, 0xc9, 0xa4, 0xf9, 0x5d, 0x38, 0xd7,
	0x5a, 0xe1, 0xfb, 0x6a, 0x0d, 0xed, 0x8e, 0xfc, 0x88, 0xc7, 0x07, 0x86, 0x1e, 0x57, 0xfb, 0xd7,
	0x03, 0x8d, 0x3f, 0x3c, 0x50, 0x9d, 0x01, 0xf9, 0x05, 0xa8, 0x71, 0xc1, 0x38, 0x16, 0x30, 0x41,
	0x12, 0xc3, 0x21, 0x13, 0xa9, 0x74, 0xfb, 0xf3, 0xf2, 0x96, 0xc6, 0x61, 0xd7, 0x08, 0xee, 0x23,
	0x89, 0xcf, 0xb4, 0x5c, 0xe7, 0x83, 0xb8, 0xca, 0xaf, 0x87, 0x1a, 0x35, 0x50, 0x9d, 0x41, 0x45,
	0x75, 0xb0, 0xe1, 0xb6, 0xc0, 0x96, 0xef, 0x57, 0xfe, 0x8a, 0xbc, 0xe6, 0x6f, 0xab, 0xf6, 0x64,
	0xdb, 0xb3, 0x27, 0x7b, 0x00, 0xee, 0xf5, 0x39, 0x4a, 0xde, 0x40, 0x85, 0x7a, 0x19, 0x86, 0x92,
	0xbc, 0xc5, 0x2e, 0xef, 0xff, 0x6c, 0xa2, 0x78, 0xd3, 0xb0, 0x4e, 0x34, 0xe9, 0x98, 0xbc, 0xc5,
	0x3e, 0x02, 0xf7, 0x5d, 0x83, 0x27, 0x85, 0x10, 0x98, 0x2a, 0x28, 0x95, 0xc0, 0x28, 0x97, 0x0b,
	0xb5, 0xe4, 0xc6, 0x65, 0x04, 0x76, 0x57, 0x82, 0xc9, 0x64, 0x32, 0x59, 0xde, 0xf1, 0xe2, 0xba,
	0xed, 0x78, 0xa7, 0x74, 0x6c, 0x85, 0xfc, 0x5f, 0x40, 0x83, 0x50, 0xa2, 0x08, 0xca, 0x9c, 0x36,
	0x1c, 0x12, 0x9a, 0xb2, 0xa1, 0x4d, 0xba, 0xb2, 0x80, 0x4d, 0xf5, 0x32, 0x5a, 0xdf, 0x05, 0xce,
	0x66, 0x32, 0xa9, 0xc4, 0xdb, 0x4e, 0xd0, 0x5a, 0x9c, 0x19, 0x39, 0x53, 0x8e, 0x00, 0x0f, 0x4b,
	0xaf, 0xa9, 0x3b, 0x3b, 0xed, 0xb7, 0x74, 0x1b, 0xbf, 0x07, 0x4e, 0xf4, 0xdd, 0x75, 0x9e, 0xf2,
	0xfc, 0x0c, 0x6c, 0xd8, 0x4b, 0xe6, 0x1c, 0x83, 0x0f, 0xdd, 0xcd, 0xd0, 0x41, 0xc7, 0xf0, 0xbf,
	0x00, 0x9b, 0x16, 0x94, 0x63, 0x85, 0x52, 0xa4, 0x50, 0x70, 0xd7, 0xa0, 0x2c, 0xf5, 0x95, 0x0b,
	0xfa, 0xa7, 0x60, 0x4b, 0x1f, 0x07, 0x2b, 0x54, 0x8f, 0x15, 0x34, 0x85, 0xe7, 0x02, 0xe5, 0x58,
	0x06, 0xcb, 0x37, 0x19, 0x0f, 0x7a, 0xbc, 0x1c, 0x39, 0x81, 0x03, 0xc3, 0xf7, 0x53, 0xf0, 0xe0,
	0x9a, 0x6c, 0xc2, 0xa8, 0x12, 0x2c, 0x2b, 0xe5, 0x57, 0x6e, 0x22, 0x1f, 0x4c, 0xc9, 0xef, 0x5b,
	0x1d, 0xe7, 0xf2, 0xab, 0x07, 0xf6, 0x5c, 0x33, 0x49, 0x9c, 0x14, 0x8a, 0x0c, 0x30, 0x24, 0x74,
	0xba, 0x10, 0x38, 0x24, 0xaa, 0x0f, 0x71, 0xce, 0xd5, 0x18, 0x72, 0x34, 0xce, 0x18, 0x4a, 0x83,
	0xd5, 0x05, 0xba, 0xf6, 0x2b, 0xdb, 0x58, 0xa5, 0xea, 0x21, 0x9d, 0x2a, 0xee, 0x8c, 0xa8, 0xfe,
	0x0b, 0xad, 0xd8, 0xb5, 0x82, 0x3e, 0x03, 0x4f, 0x74, 0x16, 0xa5, 0x33, 0x17, 0x84, 0x09, 0xa2,
	0xc6, 0x65, 0x0a, 0xfa, 0xb2, 0xdb, 0x3e, 0x0c, 0xc0, 0x02, 0xd6, 0x8f, 0x73, 0x34, 0x72, 0x7e,
	0x5d, 0xa7, 0x64, 0x7d, 0xbb, 0x58, 0xd8, 0xf6, 0xf3, 0x7f, 0xf7, 0xc0, 0x37, 0xd3, 0x8e, 0xae,
	0xdd, 0x0a, 0xae, 0x3f, 0x5a, 0xd3, 0xb6, 0xfa, 0x88, 0xed, 0x3b, 0x94, 0x98, 0xaa, 0x60, 0xed,
	0x26, 0xfb, 0xfe, 0xf5, 0xbb, 0x44, 0x6c, 0xf3, 0x9d, 0x1a, 0xfd, 0xab, 0x64, 0x9e, 0x23, 0x85,
	0xcc, 0xcb, 0x31, 0xa6, 0xfa, 0xc3, 0xf5, 0xc4, 0x5d, 0x36, 0x2c, 0x04, 0x13, 0x90, 0x51, 0x48,
	0xa8, 0xf9, 0x90, 0xda, 0x81, 0x9c, 0x63, 0x29, 0xd1, 0x05, 0xa1, 0x17, 0xc1, 0xba, 0xe9, 0xc3,
	0xc7, 0x16, 0xfd, 0x42, 0x83, 0x8f, 0xe8, 0xa1, 0x85, 0xea, 0xc9, 0xf3, 0xaa, 0x04, 0x36, 0x15,
	0xd8, 0x7a, 0x29, 0x78, 0x32, 0x3b, 0x89, 0x7e, 0x06, 0xf7, 0xb5, 0x62, 0x1b, 0x96, 0x7f, 0x09,
	0x20, 0xb3, 0x2b, 0x6e, 0x1e, 0x7d, 0xf9, 0x2f, 0x73, 0x74, 0x76, 0xa4, 0xc5, 0xf5, 0xfe, 0x7b,
	0xa2, 0x51, 0x1b, 0x7c, 0x4a, 0x98, 0x95, 0xe0, 0x82, 0x8d, 0xc6, 0xf3, 0x6a, 0xd1, 0x46, 0xc9,
	0x31, 0xbf, 0x5d, 0xaf, 0x77, 0xd7, 0xa4, 0xf2, 0xec, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb4,
	0x90, 0xc4, 0x2e, 0xb8, 0x08, 0x00, 0x00,
}
