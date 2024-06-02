package domain

type Pagination struct {
	RecordPerPage int // 20
	CurrentPage   int // 1
	TotalPage     int // total count
	NextPage      bool
	TotalCount    int
}

type EmpolyeeResponseDto struct {
	Empolyee []*Empolyee `json:"empolyee"`
	MetaData Pagination  `json:"meta_data"`
}
