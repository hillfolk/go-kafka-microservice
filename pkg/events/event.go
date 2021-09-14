package events

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type EVENT_TYPE string

var (
	INVEST_REQUEST_EVENT   EVENT_TYPE = "INVEST_REQUEST_EVENT"
	INVEST_SUCCESS_EVENT   EVENT_TYPE = "INVEST_SUCCESS_EVENT"
	INVEST_FAIL_EVENT      EVENT_TYPE = "INVEST_FAIL_EVENT"
	WITHDRAW_REQUEST_EVENT EVENT_TYPE = "WITHDRAW_REQUEST_EVENT"
	WITHDRAW_SUCCESS_EVENT EVENT_TYPE = "WITHDRAW_SUCCESS_EVENT"
	WITHDRAW_FAIL_EVENT    EVENT_TYPE = "WITHDRAW_FAIL_EVENT"
)

/* each event expressed as its own struct */

type Event struct {
	EvtId     string
	Type      EVENT_TYPE
	CreatedAt time.Time
}

type InvestRequestEvent struct {
	Event
	Project int64
	Amount  int64
	Qty     int64
}

type InvestSuccessEvent struct {
	Event
	Project int64
	Amount  int64
	Qty     int64
}
type InvestFailEvent struct {
	Event
	FailReason string
}

type WithdrawRequestEvent struct {
	Event
	AccountNumber string
	ExpireTime    time.Time
	Amount        int
}

type WithdrawSuccessEvent struct {
	Event
	AccountNumber    string
	WithdrawDatetime time.Time
	Amount           int
}
type WithdrawFailEvent struct {
	Event
	FailReason    string
	AccountNumber string
	Amount        int
}

/* helper to create events */

func getUUID() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}

func NewInvestRequestEvent() InvestRequestEvent {
	event := new(InvestRequestEvent)
	event.Type = INVEST_REQUEST_EVENT
	event.EvtId = getUUID()
	return *event
}

func NewWithdrawRequestEvent(accNumber string, amt int) WithdrawRequestEvent {
	event := new(WithdrawRequestEvent)
	event.Type = "DepositEvent"
	event.EvtId = getUUID()
	event.Amount = amt
	event.AccountNumber = accNumber
	return *event
}

func NewWithdrawSuccessEvent(accNumber string, amt int, withdrawDateTime time.Time) WithdrawSuccessEvent {
	event := new(WithdrawSuccessEvent)
	event.Type = WITHDRAW_SUCCESS_EVENT
	event.EvtId = getUUID()
	event.Amount = amt
	event.AccountNumber = accNumber
	event.WithdrawDatetime = withdrawDateTime
	return *event
}

func NewWithdrawFailEvent(accNumber string, amt int, reason string, withdrawDateTime time.Time) WithdrawFailEvent {
	event := new(WithdrawFailEvent)
	event.Type = WITHDRAW_SUCCESS_EVENT
	event.EvtId = getUUID()
	event.Amount = amt
	event.AccountNumber = accNumber
	event.FailReason = reason
	return *event
}
