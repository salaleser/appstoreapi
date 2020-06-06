package appstoreapi

// PageData2 это костыль для конвертации appIds []string в []int.
// !!! в Room adamIds может быть []string или []int
type PageData2 struct {
	AdamID                    int                       `json:"adamId"`
	AdamIDs                   []string                  `json:"adamIds"`
	AllCategoriesLink         AllCategoriesLink         `json:"allCategoriesLink"`
	AppRatingsLearnMoreURL    string                    `json:"appRatingsLearnMoreUrl"`
	Bubbles                   Bubbles                   `json:"bubbles"`
	CmcSwooshData             CmcSwooshData             `json:"cmcSwooshData"`
	ComponentName             string                    `json:"componentName"`
	ContentID                 string                    `json:"contentId"`
	CustomerReviewsURL        string                    `json:"customerReviewsUrl"`
	CustomersAlsoBoughtApps   []string                  `json:"customersAlsoBoughtApps"`
	DoNotFilter               bool                      `json:"doNotFilter"`
	FcKind                    string                    `json:"fcKind"`
	FcStructure               FcStructure               `json:"fcStructure"`
	GenreID                   int                       `json:"genreId"`
	KindExtID                 string                    `json:"kindExtId"`
	KindID                    int                       `json:"kindId"`
	KindName                  string                    `json:"kindName"`
	MetricsBase               MetricsBase               `json:"metricsBase"`
	Metrics                   Metrics                   `json:"metrics"`
	MoreByThisDeveloper       []string                  `json:"moreByThisDeveloper"`
	Mt                        int                       `json:"mt"`
	NfaSwooshData             NfaSwooshData             `json:"nfaSwooshData"`
	Today                     Today                     `json:"today"`
	TotalNumberOfReviews      int                       `json:"totalNumberOfReviews"`
	ID                        string                    `json:"id"`
	IsFatBinary               int                       `json:"isFatBinary"`
	IsNewsstand               bool                      `json:"isNewsstand"`
	PageTitle                 string                    `json:"pageTitle"`
	PageType                  string                    `json:"pageType"`
	PurchaseURL               string                    `json:"purchaseUrl"`
	RadioURL                  string                    `json:"radioUrl"`
	RatingAndAdvisories       RatingAndAdvisories       `json:"rating-and-advisories"`
	RecommendationsSwooshData RecommendationsSwooshData `json:"recommendationsSwooshData"`
	SellerLabel               string                    `json:"sellerLabel"`
	Sf6ResourceImagePath      string                    `json:"sf6ResourceImagePath"`
	Term                      string                    `json:"term"`
	TopApps                   TopApps                   `json:"topApps"`
	VersionHistory            VersionHistory            `json:"versionHistory"`
	UnAvailableContentIds     map[string]int            `json:"unAvailableContentIds"`
	UserReviewsSortOptions    UserReviewsSortOptions    `json:"userReviewsSortOptions"`
	UserReviewList            UserReviewList            `json:"userReviewList"`
	URL                       string                    `json:"url"`
}
