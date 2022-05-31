package domain

type UrlAlias struct {
	Id    string
	Alias string
	Url   string
}

func (UrlAlias) TableName() string {
	return "public.url_alias"
}
