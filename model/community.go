package model

// Community is the community datasource skelton
type Community struct {
	tableName struct{} `sql:"communities"`
	PrimaryID
	OwnerID   uint   `json:"owner_id"`
	Owner     User   `json:"owner"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private"`
	CommonDates
}
