package pdu

import (
    "github.com/james226/smscemu/server/logger"
    smpp "github.com/destel/smpp34"
    )

func ProcessEnquireLink(log logger.Log, session *smpp.Smpp, submitSm smpp.Pdu) {
    log.Info().Printf("Enquire Link received.");

    response, _ := session.EnquireLinkResp(submitSm.GetHeader().Sequence)

    SendPdu(log, session, response)
}
