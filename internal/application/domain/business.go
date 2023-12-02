package domain

type Business struct {
	Title    string
	Id       string
	Type     string
	CreateAt string
	UpdateAt string

	Config Config
}

type Config struct {
	Id         string
	BusinessId string

	DataSourceConfig []string
}
