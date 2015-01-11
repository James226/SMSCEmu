package pdu

import (
    "github.com/james226/smscemu/server/logger"
    smpp "github.com/destel/smpp34"
    )

type Session interface {
    Write(pdu smpp.Pdu) error
}

func SendPdu(log logger.Log, session Session, pdu smpp.Pdu) {
    err := session.Write(pdu)
    if err != nil {
        log.Warn().Printf("An error occured", err)
    }
}
