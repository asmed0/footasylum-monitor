package main

import (
	"time"

	http "github.com/useflyent/fhttp"
)

type MonitorSession struct {
	oauthData   *TokenResponse
	newResponse *ProductsResponse
	oldResponse *ProductsResponse
	Client      *http.Client
}

type TokenResponse struct {
	AccessToken                string `json:"access_token"`
	TokenType                  string `json:"token_type"`
	ExpiresIn                  int    `json:"expires_in"`
	UserName                   string `json:"userName"`
	Issued                     string `json:".issued"`
	Expires                    string `json:".expires"`
	OauthVerificationToken     string `json:"oauth-verification-token"`
	OauthVerificationStatus    string `json:"oauth-verification-status"`
	OauthVerificationChallenge string `json:"oauth-verification-challenge"` //related
}

type ProductsResponse struct {
	Tracking struct {
		Userid      string    `json:"userid"`
		Session     string    `json:"session"`
		Sessiondate time.Time `json:"sessiondate"`
		Campaign    int       `json:"campaign"`
		Variation   int       `json:"variation"`
		Debug       bool      `json:"debug"`
	} `json:"tracking"`
	Data struct {
		Type     string `json:"type"`
		Key      string `json:"key"`
		Products struct {
			Products []struct {
				Type    string `json:"type"`
				Product struct {
					Display        int     `json:"display"`
					Productid      string  `json:"productid"`
					Name           string  `json:"name"`
					Image          string  `json:"image"`
					URL            string  `json:"url"`
					Dateintroduced string  `json:"dateintroduced"`
					Group          string  `json:"group"`
					Score          float64 `json:"score"`
					Hero           int     `json:"hero"`
					Scores         struct {
						Newness       int `json:"newness"`
						Stock         int `json:"stock"`
						CategoryBoost int `json:"category_boost"`
						Sales         int `json:"sales"`
						Ctr           int `json:"ctr"`
						Conversion    int `json:"conversion"`
						SalesAdjusted int `json:"sales_adjusted"`
					} `json:"scores"`
					Attributes struct {
						Alternative                 []string `json:"alternative"`
						Alternatives                []string `json:"alternatives"`
						SwatchImage                 []string `json:"swatch_image"`
						Brand                       []string `json:"brand"`
						BrandName                   []string `json:"brand_name"`
						ShowSaleInFullpriceCategory []string `json:"show_sale_in_fullprice_category"`
						Defaultcategory             []string `json:"defaultcategory"`
						Producttype                 []string `json:"producttype"`
						Gender                      []string `json:"gender"`
						Colour                      []string `json:"colour"`
						Style                       []string `json:"style"`
						ShortDescription            []string `json:"short_description"`
						Description                 []string `json:"description"`
						ProductCategory             []string `json:"product_category"`
						CategoryListAll             []string `json:"category_list_all"`
						HasCloudVideo               []string `json:"has_cloud_video"`
						HasCloud360                 []string `json:"has_cloud_360"`
						Outlet                      []string `json:"outlet"`
						SeasonCode                  []string `json:"season_code"`
						VendorID                    []string `json:"vendor_id"`
						CategoryDefaultCodes        []string `json:"category_default_codes"`
					} `json:"attributes"`
					Variantattributes struct {
						Size []string `json:"size"`
					} `json:"variantattributes"`
					Associations struct {
						Upsells []interface{} `json:"upsells"`
					} `json:"associations"`
					Categorylist     []string `json:"categorylist"`
					Categorykeywords []string `json:"categorykeywords"`
					Options          []struct {
						Sku                    string  `json:"sku"`
						Onsale                 bool    `json:"onsale"`
						Price                  string  `json:"price"`
						Priceasnumber          float64 `json:"priceasnumber"`
						Promotionprice         string  `json:"promotionprice"`
						Promotionpriceasnumber float64 `json:"promotionpriceasnumber"`
						Default                bool    `json:"default"`
						Startdate              string  `json:"startdate"`
						Enddate                string  `json:"enddate"`
						Attributes             struct {
							Size []string `json:"size"`
						} `json:"attributes"`
						Stock        float64 `json:"stock"`
						Sequence     float64 `json:"sequence"`
						Savingamount string  `json:"savingamount"`
					} `json:"options"`
					Option struct {
						Sku                    string  `json:"sku"`
						Onsale                 bool    `json:"onsale"`
						Price                  string  `json:"price"`
						Priceasnumber          float64 `json:"priceasnumber"`
						Promotionprice         string  `json:"promotionprice"`
						Promotionpriceasnumber float64 `json:"promotionpriceasnumber"`
						Default                bool    `json:"default"`
						Startdate              string  `json:"startdate"`
						Enddate                string  `json:"enddate"`
						Attributes             struct {
							Size []string `json:"size"`
						} `json:"attributes"`
						Stock        float64 `json:"stock"`
						Sequence     float64 `json:"sequence"`
						Savingamount string  `json:"savingamount"`
					} `json:"option"`
					Productgroups []struct {
						Productid  string `json:"productid"`
						Image      string `json:"image"`
						URL        string `json:"url"`
						Name       string `json:"name"`
						Onsale     bool   `json:"onsale"`
						Default    bool   `json:"default"`
						Attributes struct {
							Colour []string `json:"colour"`
						} `json:"attributes"`
					} `json:"productgroups"`
					Totalstock      int  `json:"totalstock"`
					Sales           int  `json:"sales"`
					Instock         bool `json:"instock"`
					Stockpercentage int  `json:"stockpercentage"`
					Stockband       int  `json:"stockband"`
				} `json:"product"`
				Badges []struct {
					Badgeid string `json:"badgeid"`
					Name    string `json:"name"`
					Meta    struct {
						Class      string `json:"class"`
						Childclass string `json:"childclass"`
						Position   string `json:"position"`
					} `json:"meta"`
				} `json:"badges"`
				Position int `json:"position"`
			} `json:"products"`
			Aggregations []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				Data []struct {
					Key         string `json:"key"`
					Displayname string `json:"displayname"`
					Count       int    `json:"count"`
					Sequence    int    `json:"sequence"`
					Selected    bool   `json:"selected"`
					Include     bool   `json:"include"`
				} `json:"data,omitempty"`
				Rule   string `json:"rule,omitempty"`
				Prefix string `json:"prefix,omitempty"`
				Range  struct {
					Min   int `json:"min"`
					Max   int `json:"max"`
					Count int `json:"count"`
				} `json:"range,omitempty"`
			} `json:"aggregations"`
			Sorting []struct {
				Value    string `json:"value"`
				Text     string `json:"text"`
				Selected bool   `json:"selected"`
			} `json:"sorting"`
			Template     int    `json:"template"`
			Sort         int    `json:"sort"`
			Paging       string `json:"paging"`
			Page         int    `json:"page"`
			Total        int    `json:"total"`
			Columns      int    `json:"columns"`
			Itemsperpage int    `json:"itemsperpage"`
			Itemsfrom    int    `json:"itemsfrom"`
			Itemssize    int    `json:"itemssize"`
			Properties   struct {
				Total        int `json:"total"`
				Maxpages     int `json:"maxpages"`
				Page         int `json:"page"`
				Itemsperpage int `json:"itemsperpage"`
				Prev         int `json:"prev"`
				Next         int `json:"next"`
				Viewed       int `json:"viewed"`
			} `json:"properties"`
			Didyoumean bool `json:"didyoumean"`
		} `json:"products"`
		Content struct {
			Content    []interface{} `json:"content"`
			Template   int           `json:"template"`
			Sort       int           `json:"sort"`
			Page       int           `json:"page"`
			Total      int           `json:"total"`
			Columns    int           `json:"columns"`
			Itemsfrom  int           `json:"itemsfrom"`
			Itemssize  int           `json:"itemssize"`
			Didyoumean bool          `json:"didyoumean"`
		} `json:"content"`
		Category struct {
			Categoryid string `json:"categoryid"`
			Display    int    `json:"display"`
			Parentid   string `json:"parentid"`
			Name       string `json:"name"`
			Image      string `json:"image"`
			URL        string `json:"url"`
			Keywords   string `json:"keywords"`
			Attributes struct {
				Children    []string `json:"children"`
				Descendants []string `json:"descendants"`
				Description []string `json:"description"`
			} `json:"attributes"`
			Breadcrumb []struct {
				Categoryid string `json:"categoryid"`
				Image      string `json:"image"`
				Name       string `json:"name"`
				URL        string `json:"url"`
				Display    int    `json:"display"`
			} `json:"breadcrumb"`
		} `json:"category"`
		Contentpersonalisation bool `json:"contentpersonalisation"`
		Isvisualsearch         bool `json:"isvisualsearch"`
		Image                  int  `json:"image"`
		Pricelist              struct {
			Pricelist string `json:"pricelist"`
			Currency  string `json:"currency"`
			Prefix    string `json:"prefix"`
		} `json:"pricelist"`
	} `json:"data"`
}
