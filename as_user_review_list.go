package appstoreapi

import "time"

type UserReviewList []struct {
	UserReviewID             string               `json:"userReviewId"`
	Body                     string               `json:"body"`
	Date                     time.Time            `json:"date"`
	Name                     string               `json:"name"`
	Rating                   int                  `json:"rating"`
	Title                    string               `json:"title"`
	VoteCount                int                  `json:"voteCount"`
	VoteSum                  int                  `json:"voteSum"`
	IsEdited                 bool                 `json:"isEdited"`
	ViewUsersUserReviewsURL  string               `json:"viewUsersUserReviewsUrl"`
	VoteURL                  string               `json:"voteUrl"`
	ReportConcernURL         string               `json:"reportConcernUrl"`
	ReportConcernExplanation string               `json:"reportConcernExplanation"`
	CustomerType             string               `json:"customerType"`
	DeveloperResponse        DeveloperResponse    `json:"developerResponse"`
	ReportConcernReasons     ReportConcernReasons `json:"reportConcernReasons"`
}
