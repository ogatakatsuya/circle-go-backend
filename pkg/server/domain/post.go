package domain

import "time"

type TagEnum string

const (
	BeginnerFriendly         TagEnum = "初心者多め"
	AllowMultipleMemberships TagEnum = "兼部・兼サー可能"
	OfficialCircle           TagEnum = "公認サークル"
	AnnualFeeUnder10K        TagEnum = "年会費1万円以下"
	Intercollegiate          TagEnum = "インカレ"
	ManagerWelcome           TagEnum = "マネージャー歓迎"
)

type CampusEnum string

const (
	Suita    CampusEnum = "吹田キャンパス"
	Minoh    CampusEnum = "箕面キャンパス"
	Toyonaka CampusEnum = "豊中キャンパス"
)

type ActivityFrequencyEnum string

const (
	Irregular      ActivityFrequencyEnum = "不規則"
	Online         ActivityFrequencyEnum = "オンライン"
	UpToWeekly     ActivityFrequencyEnum = "〜週1"
	WeeklyTo3Times ActivityFrequencyEnum = "週1〜週3"
	ThreeTo5Times  ActivityFrequencyEnum = "週3〜週5"
	Over5Times     ActivityFrequencyEnum = "週5〜"
)

type MainCategoryEnum string

const (
	Activity MainCategoryEnum = "活動・学生団体"
	Music    MainCategoryEnum = "音楽系"
	Culture  MainCategoryEnum = "文化系"
	Sports   MainCategoryEnum = "体育会系"
)

type SubCategoryEnum string

const (
	// 活動・学生団体
	EventManagement SubCategoryEnum = "イベント運営"
	International   SubCategoryEnum = "国際系"
	Volunteer       SubCategoryEnum = "ボランティア"
	OtherActivity   SubCategoryEnum = "その他"

	// 音楽系
	Band       SubCategoryEnum = "バンド"
	Dance      SubCategoryEnum = "ダンス"
	Orchestra  SubCategoryEnum = "楽団・演奏"
	Choir      SubCategoryEnum = "合唱・アカペラ"
	OtherMusic SubCategoryEnum = "その他音楽系"

	// 文化系
	Club         SubCategoryEnum = "同好会・研究会"
	CultureMedia SubCategoryEnum = "文化・メディア"
	Study        SubCategoryEnum = "勉強・学術"
	Computer     SubCategoryEnum = "パソコン"
	Hobby        SubCategoryEnum = "趣味・特技・エンタメ"
	Creation     SubCategoryEnum = "制作"
	Theater      SubCategoryEnum = "演劇"
	Outdoor      SubCategoryEnum = "アウトドア"
	OtherCulture SubCategoryEnum = "その他文化系"

	// 体育会系
	MartialArts   SubCategoryEnum = "武道・格闘"
	WinterSports  SubCategoryEnum = "ウィンタースポーツ"
	TrackAndField SubCategoryEnum = "陸上・体操"
	IndoorBall    SubCategoryEnum = "球技(屋内)"
	OutdoorBall   SubCategoryEnum = "球技(屋外)"
	WaterSports   SubCategoryEnum = "水上競技"
	RacketSports  SubCategoryEnum = "ラケットスポーツ"
	OtherSports   SubCategoryEnum = "その他体育会系"
)

type Media struct {
	Twitter   *string `json:"twitter,omitempty" dynamodbav:"twitter"`
	Instagram *string `json:"instagram,omitempty" dynamodbav:"instagram"`
	HomePage  *string `json:"home_page,omitempty" dynamodbav:"home_page"`
	Line      *string `json:"line,omitempty" dynamodbav:"line"`
}

type Atmosphere struct {
	Participant int `json:"participant" dynamodbav:"participant"`
	Engagement  int `json:"engagement" dynamodbav:"engagement"`
}

type Schedule struct {
	ScheduleItems []*string `json:"schedule_items" dynamodbav:"schedule_items"`
}

// PostRequest
type Post struct {
	Media           Media                 `json:"media" dynamodbav:"media"`
	ActivityContent string                `json:"activity_content" dynamodbav:"activity_content"`
	ActivityPlace   string                `json:"activity_place" dynamodbav:"activity_place"`
	ActivityDate    string                `json:"activity_date" dynamodbav:"activity_date"`
	Atmpsphere      Atmosphere            `json:"atmpsphere" dynamodbav:"atmpsphere"`
	MemberNum       *int                  `json:"member_num" dynamodbav:"member_num"`
	ParticipantFee  *string               `json:"participant_fee" dynamodbav:"participant_fee"`
	MaleRate        *int                  `json:"male_rate" dynamodbav:"male_rate"`
	ScienceRate     *int                  `json:"science_rate" dynamodbav:"science_rate"`
	AppealPoint     *string               `json:"appeal_point" dynamodbav:"appeal_point"`
	Schedule        Schedule              `json:"schedule" dynamodbav:"schedule"`
	Tags            []TagEnum             `json:"tags" dynamodbav:"tags"`
	Campus          CampusEnum            `json:"campus" dynamodbav:"campus"`
	ActivityFreq    ActivityFrequencyEnum `json:"activity_frequency" dynamodbav:"activity_frequency"`
	MainCategory    MainCategoryEnum      `json:"main_category" dynamodbav:"main_category"`
	SubCategories   []SubCategoryEnum     `json:"sub_categories" dynamodbav:"sub_categories"`
}

// GetResponse
type GetResponse struct {
	Post       *Post     `json:"post"`
	UserID     string    `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	IconURL    *string   `json:"icon_url"`
	HeaderURLs []*string `json:"header_urls"`
	Name       string    `json:"name"`
}
