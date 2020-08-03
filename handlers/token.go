package handlers

import (
	"log"
	"net/http"
	"strconv"

	"mintscan/client"
	"mintscan/db"
	"mintscan/errors"
	"mintscan/utils"
)

// Token is a token handler
type Token struct {
	l      *log.Logger
	client *client.Client
	db     *db.Database
}

// NewToken creates a new token handler with the given params
func NewToken(l *log.Logger, client *client.Client, db *db.Database) *Token {
	return &Token{l, client, db}
}

// GetTokens returns assets based upon the request params
func (t *Token) GetTokens(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	limit := 100
	offset := 0

	if len(r.URL.Query()["limit"]) > 0 {
		limit, _ = strconv.Atoi(r.URL.Query()["limit"][0])
	}

	if len(r.URL.Query()["offset"]) > 0 {
		offset, _ = strconv.Atoi(r.URL.Query()["offset"][0])
	}

	if limit > 1000 {
		errors.ErrOverMaxLimit(rw, http.StatusUnauthorized)
		return
	}

	tks, _ := t.client.Tokens(limit, offset)

	utils.Respond(rw, tks)
	return
}
