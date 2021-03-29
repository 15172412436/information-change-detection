package results_service

import (
	"godeliver/models/results"
)

type Result struct {
	Id      uint
	SubType string
	Mark    string
}

// 修改
func (r *Result) Edit() error {
	switch r.SubType {
	case "fcdy":
		return results.EditFcdyReadMark(r.Id, r.Mark)
	case "dtdy":
		return results.EditDtdyReadMark(r.Id, r.Mark)
	case "ztdy":
		return results.EditZtdyReadMark(r.Id, r.Mark)
	default:
		return nil
	}

}
