package pdu

import (
    "github.com/james226/smscemu/server/logger"
    smpp "github.com/destel/smpp34"
    )

func ProcessTranceiverRequest(log logger.Log, session *smpp.Smpp, bindRequest smpp.Pdu) {
    log.Info().Printf("Client '%v' established connection.\n", bindRequest.Fields()["system_id"])

    response, _ := session.BindResp(smpp.BIND_TRANSCEIVER_RESP, bindRequest.GetHeader().Sequence, smpp.ESME_ROK, bindRequest.Fields()["system_id"].String())

    SendPdu(log, session, response)
}
