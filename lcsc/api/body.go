package api

type ProductSearchBody struct {
	CurrentPage       int   `json:"currentPage"`
	PageSize          int   `json:"pageSize"`
	CatalogID         []int `json:"catalogIdList,omitempty"`
	ParamNameValueMap struct {
	} `json:"paramNameValueMap"`
	BrandID       []any    `json:"brandIdList,omitempty"`
	IsStock       bool     `json:"isStock"`
	IsEnvironment bool     `json:"isEnvironment"`
	IsHot         bool     `json:"isHot"`
	IsDiscount    bool     `json:"isDiscount"`
	EncapValue    []string `json:"encapValueList,omitempty"`
}

type AutoGenerated1 struct {
	CurrentPage       int   `json:"currentPage"`
	PageSize          int   `json:"pageSize"`
	CatalogIDList     []int `json:"catalogIdList"`
	ParamNameValueMap struct {
		Tolerance  []string `json:"Tolerance"`
		Resistance []string `json:"Resistance"`
		PowerWatts []string `json:"Power(Watts)"`
	} `json:"paramNameValueMap"`
	BrandIDList    []any    `json:"brandIdList"`
	IsStock        bool     `json:"isStock"`
	IsEnvironment  bool     `json:"isEnvironment"`
	IsHot          bool     `json:"isHot"`
	IsDiscount     bool     `json:"isDiscount"`
	EncapValueList []string `json:"encapValueList"`
}