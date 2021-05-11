package model

import (
	"fmt"
	"time"
)

type OptionInfo struct {
	Deleted     bool   `xml:"Delete" json:"Delete"`
	Description string `xml:"Description" json:"Description"`
	Group       string `xml:"Group" json:"Group"`
	Name        string `xml:"Name" json:"Name"`
	Platform    string `xml:"Platform" json:"Platform"`
	SupplierID  string `xml:"SupplierID" json:"SupplierID"`
	Type        string `xml:"Type" json:"Type"`
	Value       string `xml:"Value" json:"Value"`
	Version     int    `xml:"Version" json:"Version"`
}

//OrderItem
type OrderItem struct {
	ItemID          int     `xml:"ItemID" json:"ItemID"`
	ProductID       string  `xml:"ProductID" json:"ProductID"`
	Description     string  `xml:"Description" json:"Description"`
	Quantity        float32 `xml:"Quantity" json:"Quantity"`
	QtyRequired     float32 `xml:"" json:",omitempty"`
	QtyCredited     float32 `xml:"QtyCredited" json:"QtyCredited,omitempty"`
	QtyDelivered    float32 `xml:"" json:",omitempty"`
	QtyInvoiced     float32 `xml:"QtyInvoiced" json:"QtyInvoiced,omitempty"`
	Gross           float32 `xml:"Gross" json:"Gross"`
	Discount        float32 `xml:"Discount" json:"Discount"` //the ,string allows incoming string to be mapped to float32
	Nett            float32 `xml:"Nett" json:"Nett"`
	Value           float32 `xml:"Value" json:"Value"`
	RepNett         float32 `xml:"RepNett" json:"RepNett,omitempty"`
	RepDiscount     float32 `xml:"RepDiscount" json:"RepDiscount,omitempty"`
	RepChangedPrice bool    `xml:"RepChangedPrice" json:"RepChangedPrice,omitempty"`
	ValueUnit       float32 `xml:"ValueUnit" json:"ValueUnit,omitempty"`
	VAT             float32 `xml:"VAT" json:"VAT,omitempty"`
	Unit            string  `xml:"Unit" json:"Unit,omitempty"` //unit
	Warehouse       string  `xml:"Warehouse" json:"Warehouse,omitempty"`
	UserField01     string  `xml:"UserField01" json:"UserField01,omitempty"`   //Actual nett weight
	UserField02     string  `xml:"UserField02" json:"UserField02,omitempty"`   //Weight Unit
	UserField03     string  `xml:"UserField03" json:"UserField03,omitempty"`   //Reversal reason
	UserField04     string  `xml:"UserField04" json:"UserField04,omitempty"`   //Check in quantity
	UserField05     string  `xml:"UserField05" json:"UserField05,omitempty"`   //Check in pallets
	UserAmount01    float32 `xml:"UserAmount01" json:"UserAmount01"`           //Actual nett weight
	UserAmount02    float32 `xml:"UserAmount02" json:"UserAmount02,omitempty"` //Weight Unit
	UserAmount03    float32 `xml:"UserAmount03" json:"UserAmount03,omitempty"` //Reversal reason
	UserAmount04    float32 `xml:"UserAmount04" json:"UserAmount04,omitempty"` //Check in quantity
	UserAmount05    float32 `xml:"UserAmount05" json:"UserAmount05,omitempty"` //Check in pallets
	Reason          string  `xml:"Reason" json:"Reason,omitempty"`
	Reasons         []struct {
		Description string `xml:"Description" json:"description"`
		Quantity    int    `xml:"Quantity" json:"quantity"`
	} `xml:"Reasons" json:"Reasons,omitempty"`
	DiscountID     string    `xml:"DiscountID" json:"DiscountID,omitempty"`
	PromoID        string    `xml:"PromoID" json:"PromoID,omitempty"`
	Deal           string    `xml:"Deal" json:"Deal,omitempty"`
	UserDate01     time.Time `xml:"UserDate01" json:"UserDate01,omitempty"`
	CrossReference []struct {
		AccountID   string
		AccountName string
		BranchID    string
		UserID      string
		OrderID     string
		ItemID      int
		Quantity    int
	} `xml:"CrossReference" json:"CrossReference,omitempty"`
}
type CrossReference struct {
	OrderID     string
	Type        string
	AccountID   string
	AccountName string
	BranchID    string
	UserID      string
}

// OrderInfo contains the order
type OrderInfo struct {
	SupplierID           string           `xml:"SupplierID" json:"SupplierID,omitempty"`
	SortKey              string           `xml:"SortKey" json:"SortKey,omitempty"`
	OrderID              string           `xml:"OrderID" json:"OrderID"`
	BranchID             string           `xml:"BranchID" json:"BranchID"`
	AccountID            string           `xml:"AccountID" json:"AccountID"`
	AccountName          string           `xml:"AccountName" json:"AccountName,omitempty"`
	UserID               string           `xml:"UserID" json:"UserID"`
	RepID                string           `xml:"RepID" json:"RepID,omitempty"`
	Type                 string           `xml:"Type" json:"Type"`
	CreateDate           time.Time        `xml:"CreateDate" json:"CreateDate"`
	RequiredByDate       time.Time        `xml:"RequiredByDate" json:"RequiredByDate"`
	Reference            string           `xml:"Reference" json:"Reference"`
	Comments             string           `xml:"Comments" json:"Comments"`
	Route                string           `xml:"Route" json:"Route,omitempty"`
	Status               string           `xml:"Status" json:"Status,omitempty"`
	Longitude            string           `xml:"Longitude" json:"Longitude,omitempty"`
	Latitude             string           `xml:"Latitude" json:"Latitude,omitempty"`
	TotalExcl            float32          `xml:"TotalExcl" json:"TotalExcl,omitempty"`
	RepChangedPrice      bool             `xml:"RepChangedPrice" json:"RepChangedPrice,omitempty"`
	ClientOrderID        string           `xml:"ClientOrderID" json:"ClientOrderID,omitempty"`
	DeliveryName         string           `xml:"DeliveryName" json:"DeliveryName"`
	DeliveryAddress1     string           `xml:"DeliveryAddress1" json:"DeliveryAddress1,omitempty"`
	DeliveryAddress2     string           `xml:"DeliveryAddress2" json:"DeliveryAddress2,omitempty"`
	DeliveryAddress3     string           `xml:"DeliveryAddress3" json:"DeliveryAddress3,omitempty"`
	DeliveryPostCode     string           `xml:"DeliveryPostCode" json:"DeliveryPostCode,omitempty"`
	DeliveryMethod       string           `xml:"DeliveryMethod" json:"DeliveryMethod,omitempty"`
	RouteID              string           `xml:"RouteID" json:"RouteID,omitempty"`
	ShipmentID           string           `xml:"ShipmentID" json:"ShipmentID,omitempty"`
	PostedToERP          bool             `xml:"PostedToERP" json:"PostedToERP"`
	ERPOrderNumber       string           `xml:"ERPOrderNumber" json:"ERPOrderNumber"`
	ERPStatus            string           `xml:"ERPStatus" json:"ERPStatus,omitempty"`
	Email                string           `xml:"Email" json:"Email,omitempty"`
	Value                float32          `xml:"Value" json:"Value,omitempty"`
	Locked               bool             `xml:"Locked" json:"Locked,omitempty"`
	LockedBy             string           `xml:"LockedBy" json:"LockedBy,omitempty"`
	LockedDate           time.Time        `xml:"LockedDate" json:"LockedDate,omitempty"`
	PaymentDate          time.Time        `xml:"PaymentDate" json:"PaymentDate,omitempty"`
	HaveNotifiedCreator  time.Time        `xml:"HaveNotifiedCreator" json:"HaveNotifiedCreator,omitempty"`
	HaveNotifiedCustomer time.Time        `xml:"HaveNotifiedCustomer" json:"HaveNotifiedCustomer,omitempty"`
	WorkflowAllowed      bool             `xml:"WorkflowAllowed" json:"WorkflowAllowed,omitempty"`
	UpdateStockAllowed   bool             `xml:"UpdateStockAllowed" json:"UpdateStockAllowed,omitempty"`
	UserField01          string           `xml:"UserField01" json:"UserField01,omitempty"`
	UserField02          string           `xml:"UserField02" json:"UserField02,omitempty"`
	UserField03          string           `xml:"UserField03" json:"UserField03,omitempty"`
	UserField04          string           `xml:"UserField04" json:"UserField04,omitempty"`
	UserField05          string           `xml:"UserField05" json:"UserField05,omitempty"`
	UserField06          string           `xml:"UserField06" json:"UserField06,omitempty"`
	UserField07          string           `xml:"UserField07" json:"UserField07,omitempty"`
	UserField08          string           `xml:"UserField08" json:"UserField08,omitempty"`
	UserField09          string           `xml:"UserField09" json:"UserField09,omitempty"`
	UserField10          string           `xml:"UserField10" json:"UserField10,omitempty"`
	UserAmount01         float32          `xml:"UserAmount01" json:"UserAmount01,omitempty"`
	UserAmount02         float32          `xml:"UserAmount02" json:"UserAmount02,omitempty"`
	UserAmount03         float32          `xml:"UserAmount03" json:"UserAmount03,omitempty"`
	UserAmount04         float32          `xml:"UserAmount04" json:"UserAmount04,omitempty"`
	UserAmount05         float32          `xml:"UserAmount05" json:"UserAmount05,omitempty"`
	UserAmount06         float32          `xml:"UserAmount06" json:"UserAmount06,omitempty"`
	UserAmount07         float32          `xml:"UserAmount07" json:"UserAmount07,omitempty"`
	UserAmount08         float32          `xml:"UserAmount08" json:"UserAmount08,omitempty"`
	UserAmount09         float32          `xml:"UserAmount09" json:"UserAmount09,omitempty"`
	UserAmount10         float32          `xml:"UserAmount10" json:"UserAmount10,omitempty"`
	UserAmount11         float32          `xml:"UserAmount11" json:"UserAmount11,omitempty"`
	UserAmount12         float32          `xml:"UserAmount12" json:"UserAmount12,omitempty"`
	UserAmount13         float32          `xml:"UserAmount13" json:"UserAmount13,omitempty"`
	UserAmount14         float32          `xml:"UserAmount14" json:"UserAmount14,omitempty"`
	UserAmount15         float32          `xml:"UserAmount15" json:"UserAmount15,omitempty"`
	PaymentTerms         string           `xml:"PaymentTerms" json:"PaymentTerms,omitempty"`
	ShipTo               string           `xml:"ShipTo" json:"ShipTo,omitempty"`
	Payor                string           `xml:"Payor" json:"Payor,omitempty"`
	PayorOrderNumb       string           `xml:"PayorOrderNumb" json:"PayorOrderNumb,omitempty"`
	BillTo               string           `xml:"BillTo" json:"BillTo,omitempty"`
	Comments2            string           `xml:"Comments2" json:"Comments2,omitempty"`
	SMSText              string           `xml:"SMSText" json:"SMSText,omitempty"`
	SMSAuth              string           `xml:"SMSAuth" json:"SMSAuth,omitempty"`
	SMSNum               string           `xml:"SMSNum" json:"SMSNum,omitempty"`
	InvoiceID            string           `xml:"InvoiceID" json:"InvoiceID,omitempty"`
	OrderItems           []OrderItem      `xml:"OrderItems>OrderItemInfo" json:"OrderItems"`
	CrossReference       []CrossReference `xml:"" json:",omitempty"`
	// Approvals            []ApprovalInfo
	Deleted bool `xml:"Deleted" json:"Deleted"`
}

// OrderEnvelope holds the SOAP envelope
type OrderEnvelope struct {
	X           string    `xml:"x,attr"`
	Header      string    `xml:"Header"`
	PackageID   string    `xml:"Body>ModifyAll>packageID"`      // NumRows holds the number of rows per file that will be sent. Set to 0 if only one file
	BpackageID  string    `xml:"Body>ModifyAllBegin>packageID"` // NumRows holds the number of rows per file that will be sent. Set to 0 if only one file
	CallbackURL string    `xml:"Body>Reload>callbackURL"`       // This indicates the file number, ie. if we send 500 rows at a time, then this counter starts at 1
	Order       OrderInfo `xml:"Body>Modify>OrderInfo"`
	BOrder      OrderInfo `xml:"Body>ModifyBegin>OrderInfo"`
}

type StockInfo struct {
	Deleted            bool    `xml:"Deleted" json:"Deleted"`
	SupplierID         string  `xml:"SupplierID" json:"SupplierID"`
	ProductID          string  `xml:"ProductID" json:"ProductID"`
	Warehouse          string  `xml:"Warehouse" json:"Warehouse"`
	Stock              float32 `xml:"Stock" json:"Stock, string"`
	WarehouseProductID string
	Bin                string    `xml:"Bin" json:"Bin,omitempty"`
	Lot                string    `xml:"Lot" json:"Lot,omitempty"`
	ExpiryDate         time.Time `xml:"ExpiryDate" json:"ExpiryDate,omitempty"`
}

type OrderStateInput struct {
	OrderID    string
	SupplierID string
	Status     string
	S3File     string
	Payload    string
}

// OrderTaskInput gets passed from statemachine into lambda
type OrderTaskInput struct {
	// the input order structure
	Input OrderStateInput
	// the execution context
	ExecutionContext struct {
		Execution struct {
			Id        string
			Input     OrderStateInput
			Name      string
			RoleArn   string
			StartTime time.Time
		}
		State struct {
			EnteredTime time.Time
			Name        string
			RetryCount  int
		}
		StateMachine struct {
			Id  string
			Nme string
		}
		Task struct {
			Token string
		}
	}
}

func main() {

}

func Helloworld() {

	fmt.Println("hello world")

}
