package utils

import (
	"time"
	// "github.com/aws/aws-sdk-go/service/codebuild"
)

type PaymentStruct struct {
	PaymentId        string `json:"paymentId,omitempty"`
	UserId           string `json:"userId"`
	PaymentAmount    string `json:"paymentAmount"`
	TransactionType  string `json:"transactionType"`
	Paymentfrequency string `json:"paymentFrequency"`
}

type UserStruct struct {
	Username      string `json:"username"`
	EmailAddress  string `json:"emailAddress"`
	Password      string `json:"password"`
	Guid          string `json:"guid,omitempty"`
	AccountStatus string `json:"accountStatus"`
	MobileNumber  string `json:"mobileNumber"`
	// UserDetails    UserDetailsStruct       `json:"userDetails"`
	UserDetails     map[string]interface{} `json:"userDetails"`
	SportPreference []string               `json:"sportPreference"`
	MyGoats         []string               `json:"myGoats"`
	MyFans          []string               `json:"myFans"`
	UserStatus      string                 `json:"userStatus"`
	CreatedDate     time.Time              `json:"createdDate,omitempty"`
}

// type UserDetailsStruct struct {
//     AccountBalance    string                 `json:"accountBalance"`
//     IsProUser     	  string                 `json:"isProUser"`
//     PointsEarned      string                 `json:"pointsEarned"`
//     ProfileDP         string                 `json:"profileDP"`

// }

type ChatStruct struct {
	ChatId      string                 `json:"chatId"`
	RoomName    string                 `json:"roomName"`
	ChatDetails map[string]interface{} `json:"chatDetails"`
	CreatedDate time.Time              `json:"createdDate,omitempty"`
}
type PvtChatStruct struct {
	ChatId      string                 `json:"chatId"`
	PersonsGUID string                 `json:"personsGUID"`
	ChatDetails map[string]interface{} `json:"chatDetails"`
	CreatedDate time.Time              `json:"createdDate,omitempty"`
}

type PickStruct struct {
	// PickId      string            `json:"pickId,omitempty"`
	UserId      string        `json:"userId,omitempty"`
	CreatedDate time.Time     `json:"createdDate,omitempty"`
	UpdateDate  time.Time     `json:"updateDate"`
	PickDetails []PickDetails `json:"pickDetails"`
}
type PickDetails struct {
	Picks           []Picks   `json:"picks"`
	OddDataUnits    int       `json:"oddDataUnits"`
	OddDataWins     int       `json:"oddDataWins"`
	PickId          string    `json:"pickId"`
	Gamestatus      []string  `json:"gamestatus"`
	PickStatus      string    `json:"pickStatus"`
	Picktype        string    `json:"picktype"`
	PickCreatedDate time.Time `json:"pickCreatedDate,omitempty"`
}
type Picks struct {
	EventId          string `json:"eventId"`
	TeamName         string `json:"teamName"`
	OppTeamName      string `json:"oppTeamName"`
	Sport            string `json:"sport"`
	OddDataMoneyLine int    `json:"oddDataMoneyLine"`
}
type AwsPick struct {
	GameStatus      bool          `json:"gamestatus.NULL"`
	OddDataUnits    int           `json:"oddDataUnits.N"`
	OddDataWins     int           `json:"oddDataWins.N"`
	PickCreatedDate time.Time     `json:"pickCreatedDate.S"`
	PickId          string        `json:"pickId.S"`
	PickStatus      string        `json:"pickStatus.S"`
	Picks           []AwsPickItem `json:"picks.L"`
	PickType        string        `json:"picktype.S"`
}

type AwsPickItem struct {
	EventId          string `json:"eventId.S"`
	OddDataMoneyLine int    `json:"oddDataMoneyLine.N"`
	Sport            string `json:"sport.S"`
	TeamName         string `json:"teamname.S"`
}

type SportStruct struct {
	SportsId        string            `json:"sportsId,omitempty"`
	SportsName      string            `json:"sportsName"`
	UpcomingGames   map[string]string `json:"upcomingGames"`
	ChatroomHistory map[string]string `json:"chatroomHistory"`
}
type CheckUsernameResp struct {
	IsUserNameAlreadyExist bool `json:"isUserNameAlreadyExist"`
}
type CheckUserEmailResp struct {
	IsUserEmailAlreadyExist bool `json:"isUserEmailAlreadyExist"`
}
type CheckPickResp struct {
	IsPickExist bool `json:"isPickExist"`
}
type OddsBySportParam struct {
	Sport string `json:"sport"`
}
type Score struct {
	Name  string `json:"name"`
	Score string `json:"score"`
}

type Game struct {
	ID           string    `json:"id"`
	SportKey     string    `json:"sport_key"`
	SportTitle   string    `json:"sport_title"`
	CommenceTime time.Time `json:"commence_time"`
	Completed    bool      `json:"completed"`
	HomeTeam     string    `json:"home_team"`
	AwayTeam     string    `json:"away_team"`
	Scores       []Score   `json:"scores"`
	LastUpdate   time.Time `json:"last_update"`
}

type Pricing struct {
	Code   string `json:"code"`
	Points string `json:"points"`
	Price  string `json:"price"`
}

type Plans struct {
	Code           string `json:"code"`
	TotalPrice     string `json:"totalPrice"`
	PerDayPrice    string `json:"perDayPrice"`
	BonusPoints    string `json:"bonusPoints"`
	SavePercentage string `json:"savePercentage"`
}

type Avatar struct {
	Base64string string `json:"base64"`
}
type Statistics struct {
	PickerId   string         `json:"pickerId,omitempty"`
	Summary    PickStatistics `json:"summary"`
	Yesterday  PickStatistics `json:"yesterday"`
	Last7Days  PickStatistics `json:"last7Days"`
	Last30Days PickStatistics `json:"last30Days"`
	Last60Days PickStatistics `json:"last60Days"`
	UpdateDate time.Time      `json:"updateDate"`
}

type PickStatistics struct {
	Followers         int    `json:"followers"`
	PicksCount        int    `json:"picksCount"`
	WinningUnits      int    `json:"winningUnits"`
	WinningPercentage int    `json:"winningPercentage"`
	Last30DaysWU      int    `json:"last30DaysWU"`
	Win               int    `json:"win"`
	Loss              int    `json:"loss"`
	WinLoss           string `json:"winLoss"`
}
type OddsBySport struct {
	DataId string `json:"dataId"`
	Data   string `json:"data"`
}
