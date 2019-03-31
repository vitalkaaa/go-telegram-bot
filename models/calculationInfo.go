package models

import (
	"math"
)

var userCalculationInfo = map[int64]*CalculationInfo{}


type CalculationInfo struct {
	Money uint32
	Days  uint
	FZ string
	State string
}

type CalculationInfoRepository map[int64]*CalculationInfo

func (c *CalculationInfo) CheckMoney(money int) bool {
	if money > 1000000 {
		return false
	}
	return true
}

func (c *CalculationInfo) Calculate() uint64 {
	return uint64(float64(c.Money) + math.Round(0.07*float64(c.Money*uint32(c.Days)/365)))
}

func GetCalculationInfo(chatID int64) *CalculationInfo {
	if !IsExists(chatID) {
		userCalculationInfo[chatID] = &CalculationInfo{}
	}

	return userCalculationInfo[chatID]
}

func IsExists(chatID int64) bool {
	return userCalculationInfo[chatID] != nil
}

