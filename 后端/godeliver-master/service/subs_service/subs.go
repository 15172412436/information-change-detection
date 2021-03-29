package subs_service

import (
	"fmt"
	"godeliver/models/results"
	"godeliver/models/subs"
	"godeliver/pkg/logging"
	"strings"
)

type Subs struct {
	SubType        string `json:"sub_type"`
	SubscriptionId string `json:"subscription_id"`
	UID            uint   `json:"uid"`
	Content        string `json:"content"`
	SetTime        int    `json:"set_time"`
	Email          string `json:"email"`
	Frequency      string `json:"frequency"`
	Keyword        string `json:"keyword"`
	ReturnType     string `json:"return_type"`
	Subject        string `json:"subject"`
	Fanwei         string `json:"fanwei"`
}

// 增加
func (s *Subs) Add() (string, error) {
	sub := map[string]interface{}{
		"sub_type":        s.SubType,
		"subscription_id": s.SubscriptionId,
		"uid":             s.UID,
		"content":         s.Content,
		"frequency":       s.Frequency,
		"set_time":        s.SetTime,
		"email":           s.Email,
		"fanwei":          s.Fanwei,
	}
	var subscriptionId string
	if s.SubType == "fcdy" {
		sub["subject"] = s.Subject
		sid, err := subs.AddFcdy(sub)
		if err != nil {
			logging.Error(err)
			return "0", err
		}
		subscriptionId = sid
	}
	if s.SubType == "dtdy" {
		sub["subject"] = s.Subject
		sid, err := subs.AddDtdy(sub)
		if err != nil {
			logging.Error(err)
			return "0", err
		}
		subscriptionId = sid
	}
	if s.SubType == "ztdy" {
		sub["keyword"] = s.Keyword
		sub["return_type"] = s.ReturnType
		fmt.Println(sub["keyword"], "------")
		sid, err := subs.AddZtdy(sub)
		if err != nil {
			logging.Error(err)
			return "0", err
		}
		subscriptionId = sid
	}

	return subscriptionId, nil
}

// 获取 返回所有订阅fcdy 和 dtdy
func (s *Subs) GetAll() (map[string]interface{}, error) {

	// fcdy
	subsFcdy, err := subs.GetFcdy(s.UID)
	if err != nil {
		logging.Error(err)
	}
	ret := make(map[string]interface{})

	ret["fcdy"] = subsFcdy

	// dtdy
	subsDtdy, err := subs.GetDtdy(s.UID)
	if err != nil {
		logging.Error(err)
	}
	ret["dtdy"] = subsDtdy

	// ztdy
	subsZtdy, err := subs.GetZtdy(s.UID)
	if err != nil {
		logging.Error(err)
	}
	ret["ztdy"] = subsZtdy
	return ret, nil
}

// 修改
func (s *Subs) Edit() error {
	if strings.Split(s.SubscriptionId, "_")[1] == "fcdy" {
		return subs.EditFcdy(s.SubscriptionId, map[string]interface{}{
			"url":       s.Content,
			"frequency": s.Frequency,
			"set_time":  s.SetTime,
			"email":     s.Email,
			"subject":   s.Subject,
		})
	}
	//修改 dtdy
	if strings.Split(s.SubscriptionId, "_")[1] == "dtdy" {
		return subs.EditDtdy(s.SubscriptionId, map[string]interface{}{
			"url":       s.Content,
			"frequency": s.Frequency,
			"set_time":  s.SetTime,
			"email":     s.Email,
			"subject":   s.Subject,
			"fanwei":    s.Fanwei,
		})
	}

	//修改 ztdy
	if strings.Split(s.SubscriptionId, "_")[1] == "ztdy" {

		return subs.EditZtdy(s.SubscriptionId, map[string]interface{}{
			"subject":     s.Content,
			"frequency":   s.Frequency,
			"set_time":    s.SetTime,
			"email":       s.Email,
			"keyword":     s.Keyword,
			"return_type": s.ReturnType,
		})
	}

	return nil
}

// 删除
func (s *Subs) Delete() error {
	// fcdy
	if strings.Split(s.SubscriptionId, "_")[1] == "fcdy" {
		//删除对应的订阅结果
		err := results.DeleteFcdyResult(s.SubscriptionId)
		if err != nil {
			logging.Error("[删除订阅->删除分词订阅结果]", err)
		}
		return subs.DeleteFcdy(s.SubscriptionId)
	}
	// dtdy
	if strings.Split(s.SubscriptionId, "_")[1] == "dtdy" {
		err := results.DeleteDtdyResult(s.SubscriptionId)
		if err != nil {
			logging.Error("[删除订阅->删除动态订阅结果]", err)
		}
		return subs.DeleteDtdy(s.SubscriptionId)
	}
	// ztdy
	if strings.Split(s.SubscriptionId, "_")[1] == "ztdy" {
		err := results.DeleteZtdyResult(s.SubscriptionId)
		if err != nil {
			logging.Error("[删除订阅->删除主题订阅结果]", err)
		}
		return subs.DeleteZtdy(s.SubscriptionId)
	}
	return nil
}
