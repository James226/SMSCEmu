package pdu

import (
    "github.com/james226/smscemu/server/logger"
    smpp "github.com/destel/smpp34"
    )

func ProcessSubmitSm(log logger.Log, session *smpp.Smpp, submitSm smpp.Pdu) {
    log.Info().Printf("Submit SM received.");

    response, _ := smpp.NewSubmitSmResp(
        &smpp.Header{
            Id:       smpp.SUBMIT_SM_RESP,
            Sequence: submitSm.GetHeader().Sequence,
            },
            []byte{},
    )

    SendPdu(log, session, response)
}
