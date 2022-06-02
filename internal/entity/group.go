package entity

type Group struct {
	Id      int32
	OwnerId int32
	Name    string
	Members []int32
}
