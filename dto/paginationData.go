package dto

import "test-pr/anywr-test-studentProject/payload"

type PaginationData struct {
	PageNb       string `json:"page_nb"`
	RecordByPage string `json:"record_by_page"`
	DataRecordNb int64  `json:"data_record_nb"`
	FilterBy     string `json:"filter_by"`
}

// -----------------------------------------------------------------
// re-construct with dto

func ReconstructPaginationData(filterData payload.FilterByTeacherAndClass, nb int64, text string) PaginationData {

	return PaginationData{
		PageNb:       filterData.Page,
		RecordByPage: filterData.PageSize,
		DataRecordNb: nb,
		FilterBy:     text,
	}
}

// -----------------------------------------------------------------
