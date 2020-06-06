package appstoreapi

type UserRating struct {
	Value               float32 `json:"value"`
	RatingCount         int     `json:"ratingCount"`
	RatingCountList     []int   `json:"ratingCountList"`
	AriaLabelForRatings string  `json:"ariaLabelForRatings"`
}
