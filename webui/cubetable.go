package webui

import (
	"github.com/kubex/proto-go/application"
	"strconv"
	"encoding/base64"
	"strings"
)

type CubeTable struct {
	PerPage    int64
	PageNumber int64
	Sort       map[int]CubeTableSortItem
}

type CubeTableSortDirection string

const (
	// CubeTableSortDirectionAsc Ascending
	CubeTableSortDirectionAsc CubeTableSortDirection = "ASC"
	// CubeTableSortDirectionDesc Descending
	CubeTableSortDirectionDesc CubeTableSortDirection = "DESC"
)

type CubeTableSortItem struct {
	Column    string
	Direction CubeTableSortDirection
}

func (tbl *CubeTable) Offset() int64 {
	return (tbl.PageNumber - 1) * tbl.PerPage
}

func NewCubeTable(request *application.HTTPRequest) *CubeTable {
	tbl := &CubeTable{}
	tbl.Sort = make(map[int]CubeTableSortItem, 0)
	url := GetUrl(request)
	query := url.Query()

	tbl.PerPage, _ = strconv.ParseInt(query.Get("count"), 10, 64)
	tbl.PageNumber, _ = strconv.ParseInt(query.Get("page"), 10, 64)

	sortQ, sortErr := base64.StdEncoding.DecodeString(query.Get("sort"))
	if sortErr == nil {
		for i, srt := range strings.Split(string(sortQ), ",") {
			if len(srt) > 1 {
				sort := CubeTableSortItem{}
				if srt[:1] == "-" {
					sort.Direction = CubeTableSortDirectionDesc
					sort.Column = srt[1:]
				} else {
					sort.Direction = CubeTableSortDirectionAsc
					sort.Column = srt
				}
				tbl.Sort[i] = sort
			}
		}
	}

	return tbl
}
