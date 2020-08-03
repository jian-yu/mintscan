package handlers

import (
	"log"
	"net/http"

	"mintscan/client"
	"mintscan/db"
	"mintscan/utils"
)

// Fee is a fee handler
type Fee struct {
	l      *log.Logger
	client *client.Client
	db     *db.Database
}

// NewFee creates a new fee handler with the given params
func NewFee(l *log.Logger, client *client.Client, db *db.Database) *Fee {
	return &Fee{l, client, db}
}

// GetFees returns current fee on the active chain
func (f *Fee) GetFees(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	fees, err := f.client.TxMsgFees()
	if err != nil {
		f.l.Printf("failed to fetch tx msg fees: %s", err)
		return
	}

	utils.Respond(rw, fees)
	return
}
