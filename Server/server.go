package main

import (
	"fmt"
	"github.com/james226/smscemu/server/pdu"
	"github.com/james226/smscemu/server/logger"
	smpp "github.com/destel/smpp34"
)

func main() {
	log := logger.Init()

    log.Info().Print("Starting up...")

	listener, err := smpp.Listen("localhost", 2775)

	if err != nil {
		fmt.Println("Connection Err:", err)
		return
	}

	go AcceptConnection(log, listener)

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

func AcceptConnection(log logger.Log, listener smpp.Listener) {
	log.Info().Print("Awaiting connection...")
	session, err := listener.Accept()
	log.Info().Print("Client connected")

    go AcceptConnection(log, listener)

	if err != nil {
		log.Error().Printf("Session Err:", err)
		return
	}

	for {
		pdu, err := session.Read()
		if err != nil {
			break
		}

        ProcessPdu(log, session, pdu);
	}

	log.Info().Print("Client disconnected")
}

func ProcessPdu(log logger.Log, session *smpp.Smpp, receivedPdu smpp.Pdu) {
    switch receivedPdu.GetHeader().Id {
        case smpp.BIND_TRANSCEIVER:
            pdu.ProcessTranceiverRequest(log, session, receivedPdu)
        break

        case smpp.SUBMIT_SM:
            pdu.ProcessSubmitSm(log, session, receivedPdu)
        break

        case smpp.ENQUIRE_LINK:
            pdu.ProcessEnquireLink(log, session, receivedPdu)
        break

        default:
            log.Warn().Printf("Unknown PDU received: %v", receivedPdu.GetHeader().Id)
        break
    }
}
