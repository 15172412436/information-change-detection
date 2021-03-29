package results_service

import (
	"godeliver/models/results"
	"godeliver/pkg/logging"
	"time"
)

// 分词订阅
type FcdyResult struct {
	SubscriptionId string
	WordFrequency  string
	ReadMark       string
}

// 动态订阅
type DtdyResult struct {
	SubscriptionId string
	Uid            uint
	ExecTime       time.Time
	Title          string
	Url            string
	Content        string
	ReadMark       string
}

// 主题订阅
type ZtdyResult struct {
	SubscriptionId string
	url            string
	ExecTime       time.Time
	ReadMark       string
}

// todo
// 获取分词订阅
func (f *FcdyResult) Get() (map[string]interface{}, error) {
	//unRead, err := results.GetUnreadFcdyResult(f.SubscriptionId)
	//if err != nil {
	//	logging.Error(err)
	//	unRead = []*results.FcdyResult{}
	//}
	//read, err := results.GetReadFcdyResult(f.SubscriptionId)
	//if err != nil {
	//	logging.Error(err)
	//	read = []*results.FcdyResult{}
	//}
	all, err := results.GetAllFcdyResult(f.SubscriptionId)
	if err != nil {
		logging.Error(err)
	}
	ret := make(map[string]interface{})
	ret["result"] = all

	//ret["unRead"] = unRead
	//ret["Read"] = read
	return ret, nil
}

// 获取动态订阅
func (d *DtdyResult) Get() (map[string]interface{}, error) {
	//unRead, err := results.GetUnreadDtdyResult(d.SubscriptionId)
	//if err != nil {
	//	logging.Info(err)
	//	unRead = []*results.DtdyResult{}
	//}
	//read, err := results.GetReadDtdyResult(d.SubscriptionId)
	//if err != nil {
	//	logging.Info(err)
	//	read = []*results.DtdyResult{}
	//}
	all, err := results.GetAllDtdyResult(d.SubscriptionId)
	if err != nil {
		logging.Error(err)
	}
	ret := make(map[string]interface{})
	ret["result"] = all
	//ret["unRead"] = unRead
	//ret["Read"] = read
	return ret, nil
}

// todo
// 获取主题订阅
func (z *ZtdyResult) Get() (map[string]interface{}, error) {
	//unRead, err := results.GetUnreadZtdyResult(z.SubscriptionId)
	//if err != nil {
	//	logging.Error(err)
	//	unRead = []*results.ZtdyResult{}
	//}
	//read, err := results.GetReadZtdyResult(z.SubscriptionId)
	//if err != nil {
	//	logging.Error(err)
	//	read = []*results.ZtdyResult{}
	//}
	all, err := results.GetAllZtdyResult(z.SubscriptionId)
	if err != nil {
		logging.Error(err)
	}
	ret := make(map[string]interface{})
	ret["result"] = all
	//ret["unRead"] = unRead
	//ret["Read"] = read

	return ret, nil
}
