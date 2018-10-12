package models

import (
	// "gopkg.in/mgo.v2"
)

type BankLog struct {
    MemberId   int  `bson: "member_id"`
    SerialNumber    string     `bson: "serial_number"`
    TransCode    string  `bson: "trans_code"`
    RequestData    string     `bson: "request_data"`
    BankCode    string     `bson: "bank_code"`
    CreatedAt    string     `bson: "created_at"`
    UpdatedAt    string     `bson: "updated_at"`
}

type Per struct {
    Per    []BankLog
}