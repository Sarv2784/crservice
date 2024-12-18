// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.20.3
// source: weatherForecast.proto

package __

import (
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

// Request message
type MarineWaveHeightForecastReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  string `protobuf:"bytes,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude string `protobuf:"bytes,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *MarineWaveHeightForecastReq) Reset() {
	*x = MarineWaveHeightForecastReq{}
	mi := &file_weatherForecast_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarineWaveHeightForecastReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarineWaveHeightForecastReq) ProtoMessage() {}

func (x *MarineWaveHeightForecastReq) ProtoReflect() protoreflect.Message {
	mi := &file_weatherForecast_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarineWaveHeightForecastReq.ProtoReflect.Descriptor instead.
func (*MarineWaveHeightForecastReq) Descriptor() ([]byte, []int) {
	return file_weatherForecast_proto_rawDescGZIP(), []int{0}
}

func (x *MarineWaveHeightForecastReq) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *MarineWaveHeightForecastReq) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

// HourlyUnits defines the units for hourly data
type HourlyUnits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time       string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	WaveHeight string `protobuf:"bytes,2,opt,name=wave_height,json=waveHeight,proto3" json:"wave_height,omitempty"`
}

func (x *HourlyUnits) Reset() {
	*x = HourlyUnits{}
	mi := &file_weatherForecast_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HourlyUnits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HourlyUnits) ProtoMessage() {}

func (x *HourlyUnits) ProtoReflect() protoreflect.Message {
	mi := &file_weatherForecast_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HourlyUnits.ProtoReflect.Descriptor instead.
func (*HourlyUnits) Descriptor() ([]byte, []int) {
	return file_weatherForecast_proto_rawDescGZIP(), []int{1}
}

func (x *HourlyUnits) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *HourlyUnits) GetWaveHeight() string {
	if x != nil {
		return x.WaveHeight
	}
	return ""
}

// Hourly defines the hourly data
type Hourly struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time       []string  `protobuf:"bytes,1,rep,name=time,proto3" json:"time,omitempty"`
	WaveHeight []float64 `protobuf:"fixed64,2,rep,packed,name=wave_height,json=waveHeight,proto3" json:"wave_height,omitempty"`
}

func (x *Hourly) Reset() {
	*x = Hourly{}
	mi := &file_weatherForecast_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Hourly) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hourly) ProtoMessage() {}

func (x *Hourly) ProtoReflect() protoreflect.Message {
	mi := &file_weatherForecast_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hourly.ProtoReflect.Descriptor instead.
func (*Hourly) Descriptor() ([]byte, []int) {
	return file_weatherForecast_proto_rawDescGZIP(), []int{2}
}

func (x *Hourly) GetTime() []string {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Hourly) GetWaveHeight() []float64 {
	if x != nil {
		return x.WaveHeight
	}
	return nil
}

// MarineForecast defines the response message
type MarineForecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude             float64      `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64      `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	GenerationTimeMs     float64      `protobuf:"fixed64,3,opt,name=generation_time_ms,json=generationTimeMs,proto3" json:"generation_time_ms,omitempty"`
	UtcOffsetSeconds     int32        `protobuf:"varint,4,opt,name=utc_offset_seconds,json=utcOffsetSeconds,proto3" json:"utc_offset_seconds,omitempty"`
	Timezone             string       `protobuf:"bytes,5,opt,name=timezone,proto3" json:"timezone,omitempty"`
	TimezoneAbbreviation string       `protobuf:"bytes,6,opt,name=timezone_abbreviation,json=timezoneAbbreviation,proto3" json:"timezone_abbreviation,omitempty"`
	Elevation            float64      `protobuf:"fixed64,7,opt,name=elevation,proto3" json:"elevation,omitempty"`
	HourlyUnits          *HourlyUnits `protobuf:"bytes,8,opt,name=hourly_units,json=hourlyUnits,proto3" json:"hourly_units,omitempty"`
	Hourly               *Hourly      `protobuf:"bytes,9,opt,name=hourly,proto3" json:"hourly,omitempty"`
}

func (x *MarineForecast) Reset() {
	*x = MarineForecast{}
	mi := &file_weatherForecast_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarineForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarineForecast) ProtoMessage() {}

func (x *MarineForecast) ProtoReflect() protoreflect.Message {
	mi := &file_weatherForecast_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarineForecast.ProtoReflect.Descriptor instead.
func (*MarineForecast) Descriptor() ([]byte, []int) {
	return file_weatherForecast_proto_rawDescGZIP(), []int{3}
}

func (x *MarineForecast) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *MarineForecast) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *MarineForecast) GetGenerationTimeMs() float64 {
	if x != nil {
		return x.GenerationTimeMs
	}
	return 0
}

func (x *MarineForecast) GetUtcOffsetSeconds() int32 {
	if x != nil {
		return x.UtcOffsetSeconds
	}
	return 0
}

func (x *MarineForecast) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *MarineForecast) GetTimezoneAbbreviation() string {
	if x != nil {
		return x.TimezoneAbbreviation
	}
	return ""
}

func (x *MarineForecast) GetElevation() float64 {
	if x != nil {
		return x.Elevation
	}
	return 0
}

func (x *MarineForecast) GetHourlyUnits() *HourlyUnits {
	if x != nil {
		return x.HourlyUnits
	}
	return nil
}

func (x *MarineForecast) GetHourly() *Hourly {
	if x != nil {
		return x.Hourly
	}
	return nil
}

var File_weatherForecast_proto protoreflect.FileDescriptor

var file_weatherForecast_proto_rawDesc = []byte{
	0x0a, 0x15, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x22, 0x57, 0x0a, 0x1b, 0x4d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x57, 0x61, 0x76, 0x65, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x42, 0x0a, 0x0b, 0x48, 0x6f, 0x75,
	0x72, 0x6c, 0x79, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x77, 0x61, 0x76, 0x65, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x77, 0x61, 0x76, 0x65, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x3d, 0x0a,
	0x06, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x77,
	0x61, 0x76, 0x65, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01,
	0x52, 0x0a, 0x77, 0x61, 0x76, 0x65, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xf7, 0x02, 0x0a,
	0x0e, 0x4d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x75, 0x74, 0x63, 0x5f, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x75, 0x74, 0x63, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x12, 0x33, 0x0a, 0x15, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x61, 0x62,
	0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x41, 0x62, 0x62, 0x72, 0x65, 0x76,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x65, 0x6c, 0x65, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0c, 0x68, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x5f, 0x75,
	0x6e, 0x69, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x55, 0x6e, 0x69, 0x74, 0x73,
	0x52, 0x0b, 0x68, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x27, 0x0a,
	0x06, 0x68, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x52, 0x06,
	0x68, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x32, 0x6e, 0x0a, 0x16, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x54, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x57, 0x61, 0x76,
	0x65, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x4d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x57, 0x61, 0x76, 0x65, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x72, 0x69, 0x6e, 0x65, 0x46, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_weatherForecast_proto_rawDescOnce sync.Once
	file_weatherForecast_proto_rawDescData = file_weatherForecast_proto_rawDesc
)

func file_weatherForecast_proto_rawDescGZIP() []byte {
	file_weatherForecast_proto_rawDescOnce.Do(func() {
		file_weatherForecast_proto_rawDescData = protoimpl.X.CompressGZIP(file_weatherForecast_proto_rawDescData)
	})
	return file_weatherForecast_proto_rawDescData
}

var file_weatherForecast_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_weatherForecast_proto_goTypes = []any{
	(*MarineWaveHeightForecastReq)(nil), // 0: example.MarineWaveHeightForecastReq
	(*HourlyUnits)(nil),                 // 1: example.HourlyUnits
	(*Hourly)(nil),                      // 2: example.Hourly
	(*MarineForecast)(nil),              // 3: example.MarineForecast
}
var file_weatherForecast_proto_depIdxs = []int32{
	1, // 0: example.MarineForecast.hourly_units:type_name -> example.HourlyUnits
	2, // 1: example.MarineForecast.hourly:type_name -> example.Hourly
	0, // 2: example.WeatherForecastService.GetMarineWaveHeight:input_type -> example.MarineWaveHeightForecastReq
	3, // 3: example.WeatherForecastService.GetMarineWaveHeight:output_type -> example.MarineForecast
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_weatherForecast_proto_init() }
func file_weatherForecast_proto_init() {
	if File_weatherForecast_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_weatherForecast_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weatherForecast_proto_goTypes,
		DependencyIndexes: file_weatherForecast_proto_depIdxs,
		MessageInfos:      file_weatherForecast_proto_msgTypes,
	}.Build()
	File_weatherForecast_proto = out.File
	file_weatherForecast_proto_rawDesc = nil
	file_weatherForecast_proto_goTypes = nil
	file_weatherForecast_proto_depIdxs = nil
}
