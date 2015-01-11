package pdu

import (
    "testing"
    "github.com/james226/smscemu/server/logger"
    smpp "github.com/destel/smpp34"
    )

type StubSession struct {
    writeCalls uint
}

func (s *StubSession) Write(pdu smpp.Pdu) error {
    s.writeCalls++;
    return nil
}


func TestSendPduShouldWriteToSession(t *testing.T) {
    session := &StubSession{}
    log := logger.Init()

    pdu, _ := smpp.NewSubmitSmResp(
        &smpp.Header{
            Id:       smpp.SUBMIT_SM_RESP,
            Sequence: 1,
            },
            []byte{},
                )

    SendPdu(log, session, pdu)

    if session.writeCalls != 1 {
        t.Errorf("Expected '1' call to Write, received '%d' calls.", session.writeCalls);
    }
}
