// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/recommendationengine/v1beta1/catalog.proto

package recommendationengine

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Item stock state. If this field is unspecified, the item is
// assumed to be in stock.
type ProductCatalogItem_StockState int32

const (
	// Default item stock status. Should never be used.
	ProductCatalogItem_STOCK_STATE_UNSPECIFIED ProductCatalogItem_StockState = 0
	// Item in stock.
	ProductCatalogItem_IN_STOCK ProductCatalogItem_StockState = 0
	// Item out of stock.
	ProductCatalogItem_OUT_OF_STOCK ProductCatalogItem_StockState = 1
	// Item that is in pre-order state.
	ProductCatalogItem_PREORDER ProductCatalogItem_StockState = 2
	// Item that is back-ordered (i.e. temporarily out of stock).
	ProductCatalogItem_BACKORDER ProductCatalogItem_StockState = 3
)

var ProductCatalogItem_StockState_name = map[int32]string{
	0: "STOCK_STATE_UNSPECIFIED",
	// Duplicate value: 0: "IN_STOCK",
	1: "OUT_OF_STOCK",
	2: "PREORDER",
	3: "BACKORDER",
}

var ProductCatalogItem_StockState_value = map[string]int32{
	"STOCK_STATE_UNSPECIFIED": 0,
	"IN_STOCK":                0,
	"OUT_OF_STOCK":            1,
	"PREORDER":                2,
	"BACKORDER":               3,
}

func (x ProductCatalogItem_StockState) String() string {
	return proto.EnumName(ProductCatalogItem_StockState_name, int32(x))
}

func (ProductCatalogItem_StockState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_988daa8a4f3967d9, []int{1, 0}
}

// CatalogItem captures all metadata information of items to be recommended.
type CatalogItem struct {
	// Required. Catalog item identifier. UTF-8 encoded string with a length limit
	// of 128 bytes.
	//
	// This id must be unique among all catalog items within the same catalog. It
	// should also be used when logging user events in order for the user events
	// to be joined with the Catalog.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Required. Catalog item categories. This field is repeated for supporting
	// one catalog item belonging to several parallel category hierarchies.
	//
	// For example, if a shoes product belongs to both
	// ["Shoes & Accessories" -> "Shoes"] and
	// ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be
	// represented as:
	//
	//      "categoryHierarchies": [
	//        { "categories": ["Shoes & Accessories", "Shoes"]},
	//        { "categories": ["Sports & Fitness", "Athletic Clothing", "Shoes"] }
	//      ]
	CategoryHierarchies []*CatalogItem_CategoryHierarchy `protobuf:"bytes,2,rep,name=category_hierarchies,json=categoryHierarchies,proto3" json:"category_hierarchies,omitempty"`
	// Required. Catalog item title. UTF-8 encoded string with a length limit of 1
	// KiB.
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// Optional. Catalog item description. UTF-8 encoded string with a length
	// limit of 5 KiB.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Optional. Highly encouraged. Extra catalog item attributes to be
	// included in the recommendation model. For example, for retail products,
	// this could include the store name, vendor, style, color, etc. These are
	// very strong signals for recommendation model, thus we highly recommend
	// providing the item attributes here.
	ItemAttributes *FeatureMap `protobuf:"bytes,5,opt,name=item_attributes,json=itemAttributes,proto3" json:"item_attributes,omitempty"`
	// Optional. Language of the title/description/item_attributes. Use language
	// tags defined by BCP 47. https://www.rfc-editor.org/rfc/bcp/bcp47.txt. Our
	// supported language codes include 'en', 'es', 'fr', 'de', 'ar', 'fa', 'zh',
	// 'ja', 'ko', 'sv', 'ro', 'nl'. For other languages, contact
	// your Google account manager.
	LanguageCode string `protobuf:"bytes,6,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional. Filtering tags associated with the catalog item. Each tag should
	// be a UTF-8 encoded string with a length limit of 1 KiB.
	//
	// This tag can be used for filtering recommendation results by passing the
	// tag as part of the predict request filter. The tags have to satisfy the
	// following restrictions:
	//
	// * Only contain alphanumeric characters (`a-z`, `A-Z`, `0-9`), underscores
	//   (`_`) and dashes (`-`).
	Tags []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	// Optional. Variant group identifier for prediction results. UTF-8 encoded
	// string with a length limit of 128 bytes.
	//
	// This field must be enabled before it can be used. [Learn
	// more](/recommendations-ai/docs/catalog#item-group-id).
	ItemGroupId string `protobuf:"bytes,9,opt,name=item_group_id,json=itemGroupId,proto3" json:"item_group_id,omitempty"`
	// Extra catalog item metadata for different recommendation types.
	//
	// Types that are valid to be assigned to RecommendationType:
	//	*CatalogItem_ProductMetadata
	RecommendationType   isCatalogItem_RecommendationType `protobuf_oneof:"recommendation_type"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CatalogItem) Reset()         { *m = CatalogItem{} }
func (m *CatalogItem) String() string { return proto.CompactTextString(m) }
func (*CatalogItem) ProtoMessage()    {}
func (*CatalogItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_988daa8a4f3967d9, []int{0}
}

func (m *CatalogItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatalogItem.Unmarshal(m, b)
}
func (m *CatalogItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatalogItem.Marshal(b, m, deterministic)
}
func (m *CatalogItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatalogItem.Merge(m, src)
}
func (m *CatalogItem) XXX_Size() int {
	return xxx_messageInfo_CatalogItem.Size(m)
}
func (m *CatalogItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CatalogItem.DiscardUnknown(m)
}

var xxx_messageInfo_CatalogItem proto.InternalMessageInfo

func (m *CatalogItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CatalogItem) GetCategoryHierarchies() []*CatalogItem_CategoryHierarchy {
	if m != nil {
		return m.CategoryHierarchies
	}
	return nil
}

func (m *CatalogItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CatalogItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CatalogItem) GetItemAttributes() *FeatureMap {
	if m != nil {
		return m.ItemAttributes
	}
	return nil
}

func (m *CatalogItem) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *CatalogItem) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *CatalogItem) GetItemGroupId() string {
	if m != nil {
		return m.ItemGroupId
	}
	return ""
}

type isCatalogItem_RecommendationType interface {
	isCatalogItem_RecommendationType()
}

type CatalogItem_ProductMetadata struct {
	ProductMetadata *ProductCatalogItem `protobuf:"bytes,10,opt,name=product_metadata,json=productMetadata,proto3,oneof"`
}

func (*CatalogItem_ProductMetadata) isCatalogItem_RecommendationType() {}

func (m *CatalogItem) GetRecommendationType() isCatalogItem_RecommendationType {
	if m != nil {
		return m.RecommendationType
	}
	return nil
}

func (m *CatalogItem) GetProductMetadata() *ProductCatalogItem {
	if x, ok := m.GetRecommendationType().(*CatalogItem_ProductMetadata); ok {
		return x.ProductMetadata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CatalogItem) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CatalogItem_ProductMetadata)(nil),
	}
}

// Category represents catalog item category hierarchy.
type CatalogItem_CategoryHierarchy struct {
	// Required. Catalog item categories. Each category should be a UTF-8
	// encoded string with a length limit of 1 KiB.
	//
	// Note that the order in the list denotes the specificity (from least to
	// most specific).
	Categories           []string `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CatalogItem_CategoryHierarchy) Reset()         { *m = CatalogItem_CategoryHierarchy{} }
func (m *CatalogItem_CategoryHierarchy) String() string { return proto.CompactTextString(m) }
func (*CatalogItem_CategoryHierarchy) ProtoMessage()    {}
func (*CatalogItem_CategoryHierarchy) Descriptor() ([]byte, []int) {
	return fileDescriptor_988daa8a4f3967d9, []int{0, 0}
}

func (m *CatalogItem_CategoryHierarchy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatalogItem_CategoryHierarchy.Unmarshal(m, b)
}
func (m *CatalogItem_CategoryHierarchy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatalogItem_CategoryHierarchy.Marshal(b, m, deterministic)
}
func (m *CatalogItem_CategoryHierarchy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatalogItem_CategoryHierarchy.Merge(m, src)
}
func (m *CatalogItem_CategoryHierarchy) XXX_Size() int {
	return xxx_messageInfo_CatalogItem_CategoryHierarchy.Size(m)
}
func (m *CatalogItem_CategoryHierarchy) XXX_DiscardUnknown() {
	xxx_messageInfo_CatalogItem_CategoryHierarchy.DiscardUnknown(m)
}

var xxx_messageInfo_CatalogItem_CategoryHierarchy proto.InternalMessageInfo

func (m *CatalogItem_CategoryHierarchy) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

// ProductCatalogItem captures item metadata specific to retail products.
type ProductCatalogItem struct {
	// Product price. Only one of 'exactPrice'/'priceRange' can be provided.
	//
	// Types that are valid to be assigned to Price:
	//	*ProductCatalogItem_ExactPrice_
	//	*ProductCatalogItem_PriceRange_
	Price isProductCatalogItem_Price `protobuf_oneof:"price"`
	// Optional. A map to pass the costs associated with the product.
	//
	// For example:
	// {"manufacturing": 45.5} The profit of selling this item is computed like
	// so:
	//
	// * If 'exactPrice' is provided, profit = displayPrice - sum(costs)
	// * If 'priceRange' is provided, profit = minPrice - sum(costs)
	Costs map[string]float32 `protobuf:"bytes,3,rep,name=costs,proto3" json:"costs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	// Optional. Only required if the price is set. Currency code for price/costs. Use
	// three-character ISO-4217 code.
	CurrencyCode string `protobuf:"bytes,4,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// Optional. Online stock state of the catalog item. Default is `IN_STOCK`.
	StockState ProductCatalogItem_StockState `protobuf:"varint,5,opt,name=stock_state,json=stockState,proto3,enum=google.cloud.recommendationengine.v1beta1.ProductCatalogItem_StockState" json:"stock_state,omitempty"`
	// Optional. The available quantity of the item.
	AvailableQuantity int64 `protobuf:"varint,6,opt,name=available_quantity,json=availableQuantity,proto3" json:"available_quantity,omitempty"`
	// Optional. Canonical URL directly linking to the item detail page with a
	// length limit of 5 KiB..
	CanonicalProductUri string `protobuf:"bytes,7,opt,name=canonical_product_uri,json=canonicalProductUri,proto3" json:"canonical_product_uri,omitempty"`
	// Optional. Product images for the catalog item.
	Images               []*Image `protobuf:"bytes,8,rep,name=images,proto3" json:"images,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductCatalogItem) Reset()         { *m = ProductCatalogItem{} }
func (m *ProductCatalogItem) String() string { return proto.CompactTextString(m) }
func (*ProductCatalogItem) ProtoMessage()    {}
func (*ProductCatalogItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_988daa8a4f3967d9, []int{1}
}

func (m *ProductCatalogItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductCatalogItem.Unmarshal(m, b)
}
func (m *ProductCatalogItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductCatalogItem.Marshal(b, m, deterministic)
}
func (m *ProductCatalogItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductCatalogItem.Merge(m, src)
}
func (m *ProductCatalogItem) XXX_Size() int {
	return xxx_messageInfo_ProductCatalogItem.Size(m)
}
func (m *ProductCatalogItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductCatalogItem.DiscardUnknown(m)
}

var xxx_messageInfo_ProductCatalogItem proto.InternalMessageInfo

type isProductCatalogItem_Price interface {
	isProductCatalogItem_Price()
}

type ProductCatalogItem_ExactPrice_ struct {
	ExactPrice *ProductCatalogItem_ExactPrice `protobuf:"bytes,1,opt,name=exact_price,json=exactPrice,proto3,oneof"`
}

type ProductCatalogItem_PriceRange_ struct {
	PriceRange *ProductCatalogItem_PriceRange `protobuf:"bytes,2,opt,name=price_range,json=priceRange,proto3,oneof"`
}

func (*ProductCatalogItem_ExactPrice_) isProductCatalogItem_Price() {}

func (*ProductCatalogItem_PriceRange_) isProductCatalogItem_Price() {}

func (m *ProductCatalogItem) GetPrice() isProductCatalogItem_Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *ProductCatalogItem) GetExactPrice() *ProductCatalogItem_ExactPrice {
	if x, ok := m.GetPrice().(*ProductCatalogItem_ExactPrice_); ok {
		return x.ExactPrice
	}
	return nil
}

func (m *ProductCatalogItem) GetPriceRange() *ProductCatalogItem_PriceRange {
	if x, ok := m.GetPrice().(*ProductCatalogItem_PriceRange_); ok {
		return x.PriceRange
	}
	return nil
}

func (m *ProductCatalogItem) GetCosts() map[string]float32 {
	if m != nil {
		return m.Costs
	}
	return nil
}

func (m *ProductCatalogItem) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *ProductCatalogItem) GetStockState() ProductCatalogItem_StockState {
	if m != nil {
		return m.StockState
	}
	return ProductCatalogItem_STOCK_STATE_UNSPECIFIED
}

func (m *ProductCatalogItem) GetAvailableQuantity() int64 {
	if m != nil {
		return m.AvailableQuantity
	}
	return 0
}

func (m *ProductCatalogItem) GetCanonicalProductUri() string {
	if m != nil {
		return m.CanonicalProductUri
	}
	return ""
}

func (m *ProductCatalogItem) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ProductCatalogItem) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ProductCatalogItem_ExactPrice_)(nil),
		(*ProductCatalogItem_PriceRange_)(nil),
	}
}

// Exact product price.
type ProductCatalogItem_ExactPrice struct {
	// Optional. Display price of the product.
	DisplayPrice float32 `protobuf:"fixed32,1,opt,name=display_price,json=displayPrice,proto3" json:"display_price,omitempty"`
	// Optional. Price of the product without any discount. If zero, by default
	// set to be the 'displayPrice'.
	OriginalPrice        float32  `protobuf:"fixed32,2,opt,name=original_price,json=originalPrice,proto3" json:"original_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductCatalogItem_ExactPrice) Reset()         { *m = ProductCatalogItem_ExactPrice{} }
func (m *ProductCatalogItem_ExactPrice) String() string { return proto.CompactTextString(m) }
func (*ProductCatalogItem_ExactPrice) ProtoMessage()    {}
func (*ProductCatalogItem_ExactPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_988daa8a4f3967d9, []int{1, 0}
}

func (m *ProductCatalogItem_ExactPrice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductCatalogItem_ExactPrice.Unmarshal(m, b)
}
func (m *ProductCatalogItem_ExactPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductCatalogItem_ExactPrice.Marshal(b, m, deterministic)
}
func (m *ProductCatalogItem_ExactPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductCatalogItem_ExactPrice.Merge(m, src)
}
func (m *ProductCatalogItem_ExactPrice) XXX_Size() int {
	return xxx_messageInfo_ProductCatalogItem_ExactPrice.Size(m)
}
func (m *ProductCatalogItem_ExactPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductCatalogItem_ExactPrice.DiscardUnknown(m)
}

var xxx_messageInfo_ProductCatalogItem_ExactPrice proto.InternalMessageInfo

func (m *ProductCatalogItem_ExactPrice) GetDisplayPrice() float32 {
	if m != nil {
		return m.DisplayPrice
	}
	return 0
}

func (m *ProductCatalogItem_ExactPrice) GetOriginalPrice() float32 {
	if m != nil {
		return m.OriginalPrice
	}
	return 0
}

// Product price range when there are a range of prices for different
// variations of the same product.
type ProductCatalogItem_PriceRange struct {
	// Required. The minimum product price.
	Min float32 `protobuf:"fixed32,1,opt,name=min,proto3" json:"min,omitempty"`
	// Required. The maximum product price.
	Max                  float32  `protobuf:"fixed32,2,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductCatalogItem_PriceRange) Reset()         { *m = ProductCatalogItem_PriceRange{} }
func (m *ProductCatalogItem_PriceRange) String() string { return proto.CompactTextString(m) }
func (*ProductCatalogItem_PriceRange) ProtoMessage()    {}
func (*ProductCatalogItem_PriceRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_988daa8a4f3967d9, []int{1, 1}
}

func (m *ProductCatalogItem_PriceRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductCatalogItem_PriceRange.Unmarshal(m, b)
}
func (m *ProductCatalogItem_PriceRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductCatalogItem_PriceRange.Marshal(b, m, deterministic)
}
func (m *ProductCatalogItem_PriceRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductCatalogItem_PriceRange.Merge(m, src)
}
func (m *ProductCatalogItem_PriceRange) XXX_Size() int {
	return xxx_messageInfo_ProductCatalogItem_PriceRange.Size(m)
}
func (m *ProductCatalogItem_PriceRange) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductCatalogItem_PriceRange.DiscardUnknown(m)
}

var xxx_messageInfo_ProductCatalogItem_PriceRange proto.InternalMessageInfo

func (m *ProductCatalogItem_PriceRange) GetMin() float32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *ProductCatalogItem_PriceRange) GetMax() float32 {
	if m != nil {
		return m.Max
	}
	return 0
}

// Catalog item thumbnail/detail image.
type Image struct {
	// Required. URL of the image with a length limit of 5 KiB.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Optional. Height of the image in number of pixels.
	Height int32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	// Optional. Width of the image in number of pixels.
	Width                int32    `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_988daa8a4f3967d9, []int{2}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Image) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Image) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func init() {
	proto.RegisterEnum("google.cloud.recommendationengine.v1beta1.ProductCatalogItem_StockState", ProductCatalogItem_StockState_name, ProductCatalogItem_StockState_value)
	proto.RegisterType((*CatalogItem)(nil), "google.cloud.recommendationengine.v1beta1.CatalogItem")
	proto.RegisterType((*CatalogItem_CategoryHierarchy)(nil), "google.cloud.recommendationengine.v1beta1.CatalogItem.CategoryHierarchy")
	proto.RegisterType((*ProductCatalogItem)(nil), "google.cloud.recommendationengine.v1beta1.ProductCatalogItem")
	proto.RegisterMapType((map[string]float32)(nil), "google.cloud.recommendationengine.v1beta1.ProductCatalogItem.CostsEntry")
	proto.RegisterType((*ProductCatalogItem_ExactPrice)(nil), "google.cloud.recommendationengine.v1beta1.ProductCatalogItem.ExactPrice")
	proto.RegisterType((*ProductCatalogItem_PriceRange)(nil), "google.cloud.recommendationengine.v1beta1.ProductCatalogItem.PriceRange")
	proto.RegisterType((*Image)(nil), "google.cloud.recommendationengine.v1beta1.Image")
}

func init() {
	proto.RegisterFile("google/cloud/recommendationengine/v1beta1/catalog.proto", fileDescriptor_988daa8a4f3967d9)
}

var fileDescriptor_988daa8a4f3967d9 = []byte{
	// 936 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x8e, 0xe4, 0x38, 0x4d, 0x8e, 0x93, 0xd4, 0x65, 0x1a, 0x54, 0x73, 0x07, 0x2c, 0xc8, 0x30,
	0x2c, 0x1d, 0x30, 0x7b, 0xc9, 0xb0, 0x35, 0xe8, 0xb0, 0x0b, 0xdb, 0x55, 0x1a, 0xa3, 0x48, 0xe3,
	0xc9, 0x49, 0x2f, 0x06, 0x0c, 0x02, 0x4d, 0xb1, 0x32, 0x57, 0x49, 0xd4, 0x28, 0x2a, 0x8b, 0xef,
	0x86, 0x3d, 0xca, 0x2e, 0xf7, 0x1c, 0x7b, 0x90, 0x3d, 0xc7, 0xae, 0x06, 0xfe, 0xd8, 0x72, 0x93,
	0x5c, 0xc4, 0xd8, 0xee, 0xc4, 0xef, 0x3b, 0xfc, 0xbe, 0x73, 0x78, 0x0e, 0x45, 0x78, 0x1e, 0x73,
	0x1e, 0x27, 0xb4, 0x43, 0x12, 0x5e, 0x46, 0x1d, 0x41, 0x09, 0x4f, 0x53, 0x9a, 0x45, 0x58, 0x32,
	0x9e, 0xd1, 0x2c, 0x66, 0x19, 0xed, 0x5c, 0x1d, 0x8e, 0xa9, 0xc4, 0x87, 0x1d, 0x82, 0x25, 0x4e,
	0x78, 0xdc, 0xce, 0x05, 0x97, 0x1c, 0x3d, 0x33, 0x1b, 0xdb, 0x7a, 0x63, 0xfb, 0xae, 0x8d, 0x6d,
	0xbb, 0xb1, 0xf5, 0x89, 0xf5, 0xc0, 0x39, 0xeb, 0xbc, 0x63, 0x34, 0x89, 0xc2, 0x31, 0x9d, 0xe0,
	0x2b, 0xc6, 0x85, 0xd1, 0x6a, 0x7d, 0xbb, 0x44, 0x12, 0x3c, 0x4d, 0x79, 0x66, 0xf7, 0x7d, 0x6c,
	0xf7, 0xe9, 0xd5, 0xb8, 0x7c, 0xd7, 0x29, 0xa4, 0x28, 0x89, 0xbc, 0xc1, 0x2a, 0x5b, 0x9c, 0x65,
	0x5c, 0x6a, 0xbd, 0xc2, 0xb0, 0xfb, 0xff, 0xac, 0x42, 0xa3, 0x6f, 0x2a, 0x1a, 0x48, 0x9a, 0xa2,
	0x1d, 0x70, 0x59, 0xe4, 0x39, 0x7b, 0xce, 0xc1, 0x46, 0xaf, 0xf6, 0x77, 0xd7, 0x0d, 0x5c, 0x16,
	0xa1, 0xdf, 0x1c, 0x78, 0x4c, 0xb0, 0xa4, 0x31, 0x17, 0xd3, 0x70, 0xc2, 0xa8, 0xc0, 0x82, 0x4c,
	0x18, 0x2d, 0x3c, 0x77, 0xaf, 0x76, 0xd0, 0x38, 0x3a, 0x6d, 0xdf, 0xfb, 0x10, 0xda, 0x0b, 0x5e,
	0xea, 0x5b, 0x4b, 0x9e, 0x5a, 0xc5, 0xa9, 0x71, 0xdc, 0x21, 0x37, 0x70, 0x46, 0x0b, 0xf4, 0x11,
	0xd4, 0x25, 0x93, 0x09, 0xf5, 0x6a, 0x55, 0x6a, 0x06, 0x41, 0x9f, 0x41, 0x23, 0xa2, 0x05, 0x11,
	0x2c, 0x57, 0x7e, 0xde, 0xea, 0x2c, 0xc0, 0x09, 0x16, 0x71, 0x44, 0xe0, 0x21, 0x93, 0x34, 0x0d,
	0xb1, 0x94, 0x82, 0x8d, 0x4b, 0x49, 0x0b, 0xaf, 0xbe, 0xe7, 0x1c, 0x34, 0x8e, 0xbe, 0x59, 0x22,
	0xfd, 0x13, 0x8a, 0x65, 0x29, 0xe8, 0x19, 0xce, 0x8d, 0xc3, 0xb6, 0x92, 0xec, 0xce, 0x15, 0xd1,
	0x01, 0x6c, 0x25, 0x38, 0x8b, 0x4b, 0x1c, 0xd3, 0x90, 0xf0, 0x88, 0x7a, 0x6b, 0x55, 0x36, 0x9b,
	0x33, 0xa6, 0xcf, 0x23, 0x8a, 0x9e, 0xc0, 0xaa, 0xc4, 0x71, 0xe1, 0xad, 0xef, 0xd5, 0x66, 0x01,
	0x1a, 0x40, 0x9f, 0xc3, 0x96, 0xce, 0x33, 0x16, 0xbc, 0xcc, 0x43, 0x16, 0x79, 0x1b, 0x0b, 0x05,
	0x29, 0xe6, 0x95, 0x22, 0x06, 0x11, 0xe2, 0xd0, 0xcc, 0x05, 0x8f, 0x4a, 0x22, 0xc3, 0x94, 0x4a,
	0x1c, 0x61, 0x89, 0x3d, 0xd0, 0x15, 0x7d, 0xbf, 0x44, 0x45, 0x43, 0x23, 0xb1, 0xd0, 0x17, 0x6d,
	0x75, 0xba, 0x12, 0x3c, 0xb4, 0xea, 0x67, 0x56, 0xbc, 0x75, 0x0c, 0x8f, 0x6e, 0xb5, 0x0c, 0x7d,
	0x0a, 0x60, 0xfb, 0xa5, 0x06, 0xc2, 0x99, 0x55, 0xe3, 0x06, 0x0b, 0x70, 0x6f, 0x17, 0x76, 0x3e,
	0x4c, 0x22, 0x94, 0xd3, 0x9c, 0xee, 0xff, 0xbe, 0x0e, 0xe8, 0xb6, 0x3f, 0xca, 0xa1, 0x41, 0xaf,
	0x31, 0x91, 0x61, 0x2e, 0x18, 0xa1, 0x7a, 0x18, 0x97, 0x1b, 0xb2, 0xdb, 0x9a, 0x6d, 0x5f, 0x09,
	0x0e, 0x95, 0xde, 0xac, 0x3c, 0xa0, 0x73, 0x48, 0x39, 0x6a, 0xaf, 0x50, 0xe0, 0x2c, 0xa6, 0x9e,
	0xfb, 0x7f, 0x38, 0x6a, 0xe5, 0x40, 0xe9, 0xcd, 0x1d, 0xf3, 0x39, 0x84, 0x08, 0xd4, 0x09, 0x2f,
	0x64, 0xe1, 0xd5, 0x96, 0xbe, 0x42, 0x77, 0x78, 0xf5, 0x95, 0x94, 0x9f, 0x49, 0x31, 0x35, 0x73,
	0x62, 0xb4, 0xd5, 0x34, 0x92, 0x52, 0x08, 0x9a, 0x91, 0xa9, 0x99, 0xc6, 0x85, 0xbb, 0xb1, 0x39,
	0x63, 0xf4, 0x34, 0x66, 0xd0, 0x28, 0x24, 0x27, 0xef, 0xc3, 0x42, 0x62, 0x49, 0xf5, 0xc5, 0xd8,
	0xfe, 0xaf, 0x49, 0x8d, 0x94, 0xe0, 0x48, 0xe9, 0x19, 0x47, 0x28, 0xe6, 0x00, 0x3a, 0x02, 0x84,
	0xaf, 0x30, 0x4b, 0xf0, 0x38, 0xa1, 0xe1, 0x2f, 0x25, 0xce, 0x24, 0x93, 0x53, 0x7d, 0x59, 0x6a,
	0x26, 0xf8, 0xd1, 0x9c, 0xfe, 0xc1, 0xb2, 0xe8, 0x39, 0xec, 0x12, 0x9c, 0xf1, 0x8c, 0x11, 0x9c,
	0x84, 0xb3, 0xc9, 0x2f, 0x05, 0xf3, 0x1e, 0x54, 0x55, 0xed, 0xcc, 0x23, 0x6c, 0x42, 0x97, 0x82,
	0xa1, 0x33, 0x58, 0x63, 0x29, 0x8e, 0xa9, 0xb9, 0x6c, 0x8d, 0xa3, 0xaf, 0x96, 0xa8, 0x6b, 0xa0,
	0x36, 0x1a, 0x6d, 0x2b, 0xd2, 0x1a, 0x03, 0x54, 0xd3, 0xa4, 0xce, 0x38, 0x62, 0x45, 0x9e, 0xe0,
	0xe9, 0xc2, 0xb8, 0xba, 0xf6, 0x8c, 0x2d, 0x63, 0x22, 0xbf, 0x80, 0x6d, 0x2e, 0x58, 0xcc, 0x32,
	0x9d, 0xbe, 0x0a, 0x75, 0xab, 0xd0, 0xad, 0x19, 0xa5, 0x63, 0x5b, 0x2f, 0x00, 0xaa, 0xf9, 0x41,
	0xbb, 0x50, 0x4b, 0x59, 0x56, 0x29, 0xbb, 0x81, 0x5a, 0x6b, 0x18, 0x5f, 0x57, 0x2a, 0x0a, 0xc6,
	0xd7, 0xad, 0x63, 0x80, 0x6a, 0x1e, 0x50, 0x13, 0x6a, 0xef, 0xe9, 0xd4, 0xfc, 0xd1, 0x03, 0xf5,
	0x89, 0x1e, 0x43, 0xfd, 0x0a, 0x27, 0xa5, 0xb5, 0x0f, 0xcc, 0xe2, 0x85, 0x7b, 0xec, 0xec, 0xff,
	0x0c, 0x50, 0x35, 0x0d, 0x3d, 0x85, 0x27, 0xa3, 0x8b, 0xf3, 0xfe, 0xeb, 0x70, 0x74, 0xd1, 0xbd,
	0xf0, 0xc3, 0xcb, 0x37, 0xa3, 0xa1, 0xdf, 0x1f, 0x9c, 0x0c, 0xfc, 0x97, 0xcd, 0x15, 0xb4, 0x09,
	0xeb, 0x83, 0x37, 0xa1, 0xe6, 0x9b, 0x2b, 0xa8, 0x09, 0x9b, 0xe7, 0x97, 0x17, 0xe1, 0xf9, 0x89,
	0x45, 0x1c, 0xc5, 0x0f, 0x03, 0xff, 0x3c, 0x78, 0xe9, 0x07, 0x4d, 0x17, 0x6d, 0xc1, 0x46, 0xaf,
	0xdb, 0x7f, 0x6d, 0x96, 0xb5, 0x96, 0xdb, 0x74, 0x7a, 0x0f, 0xa0, 0xae, 0x0f, 0x61, 0xff, 0x2d,
	0xd4, 0xf5, 0x21, 0xab, 0x72, 0x54, 0x37, 0x17, 0xde, 0x1e, 0xb5, 0x46, 0x4f, 0x61, 0x6d, 0x42,
	0x59, 0x3c, 0x91, 0x3a, 0xdf, 0xba, 0xed, 0x85, 0x81, 0xd4, 0xb3, 0xf0, 0x2b, 0x8b, 0xe4, 0x44,
	0x3f, 0x0b, 0x96, 0x33, 0x48, 0xef, 0x2f, 0x07, 0xbe, 0x24, 0x3c, 0xbd, 0x7f, 0xaf, 0x87, 0xce,
	0x8f, 0x3f, 0xd9, 0xe0, 0x98, 0xab, 0x7f, 0x75, 0x9b, 0x8b, 0xb8, 0x13, 0xd3, 0x4c, 0xbf, 0x94,
	0x1d, 0x43, 0xe1, 0x9c, 0x15, 0xf7, 0x78, 0xa0, 0xbf, 0xbb, 0x8b, 0xfc, 0xc3, 0xad, 0x07, 0x7e,
	0xbf, 0x3b, 0xf8, 0xd3, 0x7d, 0xf6, 0xca, 0xf8, 0xf4, 0x75, 0x52, 0xc1, 0x07, 0xb1, 0xbe, 0x49,
	0xea, 0xed, 0x61, 0x4f, 0x09, 0x8d, 0xd7, 0xb4, 0xfb, 0xd7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x21, 0xe0, 0xf5, 0x06, 0xa2, 0x08, 0x00, 0x00,
}
