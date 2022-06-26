// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: savourrpc/market.proto

package market

import (
	common "git.savour.io/savour/savourrpc/go-savourrpc/common"
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

type Symbol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SymbolId   string `protobuf:"bytes,1,opt,name=symbol_id,json=symbolId,proto3" json:"symbol_id,omitempty"`
	SymbolName string `protobuf:"bytes,2,opt,name=symbol_name,json=symbolName,proto3" json:"symbol_name,omitempty"`
}

func (x *Symbol) Reset() {
	*x = Symbol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_market_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Symbol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Symbol) ProtoMessage() {}

func (x *Symbol) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_market_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Symbol.ProtoReflect.Descriptor instead.
func (*Symbol) Descriptor() ([]byte, []int) {
	return file_savourrpc_market_proto_rawDescGZIP(), []int{0}
}

func (x *Symbol) GetSymbolId() string {
	if x != nil {
		return x.SymbolId
	}
	return ""
}

func (x *Symbol) GetSymbolName() string {
	if x != nil {
		return x.SymbolName
	}
	return ""
}

type SymbolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
}

func (x *SymbolRequest) Reset() {
	*x = SymbolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_market_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SymbolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymbolRequest) ProtoMessage() {}

func (x *SymbolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_market_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymbolRequest.ProtoReflect.Descriptor instead.
func (*SymbolRequest) Descriptor() ([]byte, []int) {
	return file_savourrpc_market_proto_rawDescGZIP(), []int{1}
}

func (x *SymbolRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

type SymbolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   *common.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Symbols []*Symbol     `protobuf:"bytes,2,rep,name=symbols,proto3" json:"symbols,omitempty"`
}

func (x *SymbolResponse) Reset() {
	*x = SymbolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_market_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SymbolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymbolResponse) ProtoMessage() {}

func (x *SymbolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_market_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymbolResponse.ProtoReflect.Descriptor instead.
func (*SymbolResponse) Descriptor() ([]byte, []int) {
	return file_savourrpc_market_proto_rawDescGZIP(), []int{2}
}

func (x *SymbolResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SymbolResponse) GetSymbols() []*Symbol {
	if x != nil {
		return x.Symbols
	}
	return nil
}

type SymbolPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SymbolName string `protobuf:"bytes,1,opt,name=symbol_name,json=symbolName,proto3" json:"symbol_name,omitempty"`
	BuyPrice   string `protobuf:"bytes,2,opt,name=buy_price,json=buyPrice,proto3" json:"buy_price,omitempty"`
	SellPrice  string `protobuf:"bytes,3,opt,name=sell_price,json=sellPrice,proto3" json:"sell_price,omitempty"`
	AvgPrice   string `protobuf:"bytes,4,opt,name=avg_price,json=avgPrice,proto3" json:"avg_price,omitempty"`
	UsdPrice   string `protobuf:"bytes,5,opt,name=usd_price,json=usdPrice,proto3" json:"usd_price,omitempty"`
	CnyPrice   string `protobuf:"bytes,6,opt,name=cny_price,json=cnyPrice,proto3" json:"cny_price,omitempty"`
	Margin     string `protobuf:"bytes,7,opt,name=margin,proto3" json:"margin,omitempty"`
}

func (x *SymbolPrice) Reset() {
	*x = SymbolPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_market_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SymbolPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymbolPrice) ProtoMessage() {}

func (x *SymbolPrice) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_market_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymbolPrice.ProtoReflect.Descriptor instead.
func (*SymbolPrice) Descriptor() ([]byte, []int) {
	return file_savourrpc_market_proto_rawDescGZIP(), []int{3}
}

func (x *SymbolPrice) GetSymbolName() string {
	if x != nil {
		return x.SymbolName
	}
	return ""
}

func (x *SymbolPrice) GetBuyPrice() string {
	if x != nil {
		return x.BuyPrice
	}
	return ""
}

func (x *SymbolPrice) GetSellPrice() string {
	if x != nil {
		return x.SellPrice
	}
	return ""
}

func (x *SymbolPrice) GetAvgPrice() string {
	if x != nil {
		return x.AvgPrice
	}
	return ""
}

func (x *SymbolPrice) GetUsdPrice() string {
	if x != nil {
		return x.UsdPrice
	}
	return ""
}

func (x *SymbolPrice) GetCnyPrice() string {
	if x != nil {
		return x.CnyPrice
	}
	return ""
}

func (x *SymbolPrice) GetMargin() string {
	if x != nil {
		return x.Margin
	}
	return ""
}

type SymbolPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	SymbolId      string `protobuf:"bytes,2,opt,name=symbol_id,json=symbolId,proto3" json:"symbol_id,omitempty"`
}

func (x *SymbolPriceRequest) Reset() {
	*x = SymbolPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_market_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SymbolPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymbolPriceRequest) ProtoMessage() {}

func (x *SymbolPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_market_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymbolPriceRequest.ProtoReflect.Descriptor instead.
func (*SymbolPriceRequest) Descriptor() ([]byte, []int) {
	return file_savourrpc_market_proto_rawDescGZIP(), []int{4}
}

func (x *SymbolPriceRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *SymbolPriceRequest) GetSymbolId() string {
	if x != nil {
		return x.SymbolId
	}
	return ""
}

type SymbolPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error        *common.Error  `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	SymbolPrices []*SymbolPrice `protobuf:"bytes,2,rep,name=symbol_prices,json=symbolPrices,proto3" json:"symbol_prices,omitempty"`
}

func (x *SymbolPriceResponse) Reset() {
	*x = SymbolPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_market_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SymbolPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymbolPriceResponse) ProtoMessage() {}

func (x *SymbolPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_market_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymbolPriceResponse.ProtoReflect.Descriptor instead.
func (*SymbolPriceResponse) Descriptor() ([]byte, []int) {
	return file_savourrpc_market_proto_rawDescGZIP(), []int{5}
}

func (x *SymbolPriceResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SymbolPriceResponse) GetSymbolPrices() []*SymbolPrice {
	if x != nil {
		return x.SymbolPrices
	}
	return nil
}

type StableCoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinId   string `protobuf:"bytes,1,opt,name=coin_id,json=coinId,proto3" json:"coin_id,omitempty"`
	CoinName string `protobuf:"bytes,2,opt,name=coin_name,json=coinName,proto3" json:"coin_name,omitempty"`
}

func (x *StableCoin) Reset() {
	*x = StableCoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_market_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StableCoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StableCoin) ProtoMessage() {}

func (x *StableCoin) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_market_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StableCoin.ProtoReflect.Descriptor instead.
func (*StableCoin) Descriptor() ([]byte, []int) {
	return file_savourrpc_market_proto_rawDescGZIP(), []int{6}
}

func (x *StableCoin) GetCoinId() string {
	if x != nil {
		return x.CoinId
	}
	return ""
}

func (x *StableCoin) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

type StableCoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
}

func (x *StableCoinRequest) Reset() {
	*x = StableCoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_market_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StableCoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StableCoinRequest) ProtoMessage() {}

func (x *StableCoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_market_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StableCoinRequest.ProtoReflect.Descriptor instead.
func (*StableCoinRequest) Descriptor() ([]byte, []int) {
	return file_savourrpc_market_proto_rawDescGZIP(), []int{7}
}

func (x *StableCoinRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

type StableCoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error       *common.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	StableCoins []*StableCoin `protobuf:"bytes,2,rep,name=stable_coins,json=stableCoins,proto3" json:"stable_coins,omitempty"`
}

func (x *StableCoinResponse) Reset() {
	*x = StableCoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_market_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StableCoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StableCoinResponse) ProtoMessage() {}

func (x *StableCoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_market_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StableCoinResponse.ProtoReflect.Descriptor instead.
func (*StableCoinResponse) Descriptor() ([]byte, []int) {
	return file_savourrpc_market_proto_rawDescGZIP(), []int{8}
}

func (x *StableCoinResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *StableCoinResponse) GetStableCoins() []*StableCoin {
	if x != nil {
		return x.StableCoins
	}
	return nil
}

type StableCoinPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UsdPrice string `protobuf:"bytes,3,opt,name=usd_price,json=usdPrice,proto3" json:"usd_price,omitempty"`
	CnyPrice string `protobuf:"bytes,4,opt,name=cny_price,json=cnyPrice,proto3" json:"cny_price,omitempty"`
}

func (x *StableCoinPrice) Reset() {
	*x = StableCoinPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_market_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StableCoinPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StableCoinPrice) ProtoMessage() {}

func (x *StableCoinPrice) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_market_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StableCoinPrice.ProtoReflect.Descriptor instead.
func (*StableCoinPrice) Descriptor() ([]byte, []int) {
	return file_savourrpc_market_proto_rawDescGZIP(), []int{9}
}

func (x *StableCoinPrice) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StableCoinPrice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StableCoinPrice) GetUsdPrice() string {
	if x != nil {
		return x.UsdPrice
	}
	return ""
}

func (x *StableCoinPrice) GetCnyPrice() string {
	if x != nil {
		return x.CnyPrice
	}
	return ""
}

type StableCoinPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	CoinId        string `protobuf:"bytes,2,opt,name=coin_id,json=coinId,proto3" json:"coin_id,omitempty"`
}

func (x *StableCoinPriceRequest) Reset() {
	*x = StableCoinPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_market_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StableCoinPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StableCoinPriceRequest) ProtoMessage() {}

func (x *StableCoinPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_market_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StableCoinPriceRequest.ProtoReflect.Descriptor instead.
func (*StableCoinPriceRequest) Descriptor() ([]byte, []int) {
	return file_savourrpc_market_proto_rawDescGZIP(), []int{10}
}

func (x *StableCoinPriceRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *StableCoinPriceRequest) GetCoinId() string {
	if x != nil {
		return x.CoinId
	}
	return ""
}

type StableCoinPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error      *common.Error      `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	CoinPrices []*StableCoinPrice `protobuf:"bytes,2,rep,name=coin_prices,json=coinPrices,proto3" json:"coin_prices,omitempty"`
}

func (x *StableCoinPriceResponse) Reset() {
	*x = StableCoinPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_market_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StableCoinPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StableCoinPriceResponse) ProtoMessage() {}

func (x *StableCoinPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_market_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StableCoinPriceResponse.ProtoReflect.Descriptor instead.
func (*StableCoinPriceResponse) Descriptor() ([]byte, []int) {
	return file_savourrpc_market_proto_rawDescGZIP(), []int{11}
}

func (x *StableCoinPriceResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *StableCoinPriceResponse) GetCoinPrices() []*StableCoinPrice {
	if x != nil {
		return x.CoinPrices
	}
	return nil
}

var File_savourrpc_market_proto protoreflect.FileDescriptor

var file_savourrpc_market_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72,
	0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x1a, 0x16, 0x73, 0x61, 0x76, 0x6f,
	0x75, 0x72, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x46, 0x0a, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x0d, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x6c, 0x0a, 0x0e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x07,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x07, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73,
	0x22, 0xd9, 0x01, 0x0a, 0x0b, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x76, 0x67, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x76, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73,
	0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6e, 0x79, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6e, 0x79, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x22, 0x58, 0x0a, 0x12,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x49, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x42, 0x0a, 0x0d, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x0c, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x0a, 0x53, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x69, 0x6e, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3a,
	0x0a, 0x11, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7d, 0x0a, 0x12, 0x53, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3f, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x0b, 0x73, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x22, 0x6f, 0x0a, 0x0f, 0x53, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6e, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6e, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x58, 0x0a, 0x16, 0x53, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f,
	0x69, 0x6e, 0x49, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x17, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43,
	0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x42, 0x0a, 0x0b, 0x63, 0x6f, 0x69, 0x6e,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2e, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x0a, 0x63, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x32, 0x8f, 0x03, 0x0a,
	0x0c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a,
	0x0a, 0x67, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x61,
	0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x53,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73,
	0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x60, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x61, 0x76, 0x6f,
	0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43,
	0x6f, 0x69, 0x6e, 0x73, 0x12, 0x23, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63,
	0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x61, 0x76, 0x6f,
	0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x6b, 0x0a, 0x12, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f,
	0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72,
	0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x43, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4a,
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x2e, 0x73, 0x61, 0x76, 0x6f,
	0x75, 0x72, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x2f, 0x73, 0x61, 0x76,
	0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72,
	0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_savourrpc_market_proto_rawDescOnce sync.Once
	file_savourrpc_market_proto_rawDescData = file_savourrpc_market_proto_rawDesc
)

func file_savourrpc_market_proto_rawDescGZIP() []byte {
	file_savourrpc_market_proto_rawDescOnce.Do(func() {
		file_savourrpc_market_proto_rawDescData = protoimpl.X.CompressGZIP(file_savourrpc_market_proto_rawDescData)
	})
	return file_savourrpc_market_proto_rawDescData
}

var file_savourrpc_market_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_savourrpc_market_proto_goTypes = []interface{}{
	(*Symbol)(nil),                  // 0: savourrpc.market.Symbol
	(*SymbolRequest)(nil),           // 1: savourrpc.market.SymbolRequest
	(*SymbolResponse)(nil),          // 2: savourrpc.market.SymbolResponse
	(*SymbolPrice)(nil),             // 3: savourrpc.market.SymbolPrice
	(*SymbolPriceRequest)(nil),      // 4: savourrpc.market.SymbolPriceRequest
	(*SymbolPriceResponse)(nil),     // 5: savourrpc.market.SymbolPriceResponse
	(*StableCoin)(nil),              // 6: savourrpc.market.StableCoin
	(*StableCoinRequest)(nil),       // 7: savourrpc.market.StableCoinRequest
	(*StableCoinResponse)(nil),      // 8: savourrpc.market.StableCoinResponse
	(*StableCoinPrice)(nil),         // 9: savourrpc.market.StableCoinPrice
	(*StableCoinPriceRequest)(nil),  // 10: savourrpc.market.StableCoinPriceRequest
	(*StableCoinPriceResponse)(nil), // 11: savourrpc.market.StableCoinPriceResponse
	(*common.Error)(nil),            // 12: savourrpc.Error
}
var file_savourrpc_market_proto_depIdxs = []int32{
	12, // 0: savourrpc.market.SymbolResponse.error:type_name -> savourrpc.Error
	0,  // 1: savourrpc.market.SymbolResponse.symbols:type_name -> savourrpc.market.Symbol
	12, // 2: savourrpc.market.SymbolPriceResponse.error:type_name -> savourrpc.Error
	3,  // 3: savourrpc.market.SymbolPriceResponse.symbol_prices:type_name -> savourrpc.market.SymbolPrice
	12, // 4: savourrpc.market.StableCoinResponse.error:type_name -> savourrpc.Error
	6,  // 5: savourrpc.market.StableCoinResponse.stable_coins:type_name -> savourrpc.market.StableCoin
	12, // 6: savourrpc.market.StableCoinPriceResponse.error:type_name -> savourrpc.Error
	9,  // 7: savourrpc.market.StableCoinPriceResponse.coin_prices:type_name -> savourrpc.market.StableCoinPrice
	1,  // 8: savourrpc.market.PriceService.getSymbols:input_type -> savourrpc.market.SymbolRequest
	4,  // 9: savourrpc.market.PriceService.getSymbolPrices:input_type -> savourrpc.market.SymbolPriceRequest
	7,  // 10: savourrpc.market.PriceService.getStableCoins:input_type -> savourrpc.market.StableCoinRequest
	10, // 11: savourrpc.market.PriceService.getStableCoinPrice:input_type -> savourrpc.market.StableCoinPriceRequest
	2,  // 12: savourrpc.market.PriceService.getSymbols:output_type -> savourrpc.market.SymbolResponse
	5,  // 13: savourrpc.market.PriceService.getSymbolPrices:output_type -> savourrpc.market.SymbolPriceResponse
	8,  // 14: savourrpc.market.PriceService.getStableCoins:output_type -> savourrpc.market.StableCoinResponse
	11, // 15: savourrpc.market.PriceService.getStableCoinPrice:output_type -> savourrpc.market.StableCoinPriceResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_savourrpc_market_proto_init() }
func file_savourrpc_market_proto_init() {
	if File_savourrpc_market_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_savourrpc_market_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Symbol); i {
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
		file_savourrpc_market_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SymbolRequest); i {
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
		file_savourrpc_market_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SymbolResponse); i {
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
		file_savourrpc_market_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SymbolPrice); i {
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
		file_savourrpc_market_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SymbolPriceRequest); i {
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
		file_savourrpc_market_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SymbolPriceResponse); i {
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
		file_savourrpc_market_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StableCoin); i {
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
		file_savourrpc_market_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StableCoinRequest); i {
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
		file_savourrpc_market_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StableCoinResponse); i {
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
		file_savourrpc_market_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StableCoinPrice); i {
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
		file_savourrpc_market_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StableCoinPriceRequest); i {
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
		file_savourrpc_market_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StableCoinPriceResponse); i {
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
			RawDescriptor: file_savourrpc_market_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_savourrpc_market_proto_goTypes,
		DependencyIndexes: file_savourrpc_market_proto_depIdxs,
		MessageInfos:      file_savourrpc_market_proto_msgTypes,
	}.Build()
	File_savourrpc_market_proto = out.File
	file_savourrpc_market_proto_rawDesc = nil
	file_savourrpc_market_proto_goTypes = nil
	file_savourrpc_market_proto_depIdxs = nil
}
