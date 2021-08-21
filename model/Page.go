package model

type Receipt struct {
	TotalPrice              string
	TotalVatCost            string
	TotalTaxCost            string
	TotalShippingCost       string
	CurrencyCode            string
	ReceiptId               int
	OrderId                 int
	CreationTsz             int
	Name                    string
	FirsLine                string
	SecondLine              string
	City                    string
	State                   string
	Zip                     string
	FormattedAddress        string
	BuyerEmail              string
	SellerEmail             string
	DiscountAmount          string
	CouponDiscountAmt       string
	SubTotal                string
	GrandTotal              string
	AdjustedGrandtotal      string
	BuyerAdjustedGrandTotal string
}

type User struct {
	Name string `json:"name"`
}
