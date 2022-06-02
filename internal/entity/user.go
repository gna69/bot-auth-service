package entity

type User struct {
	Id           int32
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	IsBot        bool
	ChatId       uint64
}
