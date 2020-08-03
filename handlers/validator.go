package handlers

import (
	"log"
	"mintscan/schema"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"mintscan/client"
	"mintscan/db"
	"mintscan/errors"
	"mintscan/utils"

	cmtypes "github.com/binance-chain/go-sdk/common/types"
)

// Validator is a validator handler
type Validator struct {
	l      *log.Logger
	client *client.Client
	db     *db.Database
	nt     cmtypes.ChainNetwork
}

// NewValidator creates a new validator handler with the given params
func NewValidator(l *log.Logger, client *client.Client, db *db.Database, network cmtypes.ChainNetwork) *Validator {
	return &Validator{l, client, db, network}
}

// GetValidators returns validators on the active chain
func (v *Validator) GetValidators(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	//vals, err := v.db.QueryValidators()
	//if err != nil {
	//	v.l.Printf("failed to query validators: %s", err)
	//	return
	//}
	tmpVals, err := v.client.Validators()
	if err != nil {
		v.l.Printf("failed to query validators: %s", err)
		return
	}

	var vals []*schema.Validator
	for i, val := range tmpVals {
		val := &schema.Validator{
			ID:                      int32(i),
			Moniker:                 val.Description.Moniker,
			AccountAddress:          val.AccountAddress,
			OperatorAddress:         val.OperatorAddress,
			ConsensusAddress:        val.ConsensusAddress,
			Jailed:                  val.Jailed,
			Status:                  val.Status,
			Tokens:                  val.Tokens,
			VotingPower:             val.Power,
			DelegatorShares:         val.DelegatorShares,
			BondHeight:              val.BondHeight,
			BondIntraTxCounter:      val.BondIntraTxCounter,
			UnbondingHeight:         val.UnbondingHeight,
			UnbondingTime:           val.UnbondingTime.String(),
			CommissionRate:          val.Commission.Rate,
			CommissionMaxRate:       val.Commission.MaxRate,
			CommissionMaxChangeRate: val.Commission.MaxChangeRate,
			CommissionUpdateTime:    val.Commission.UpdateTime.String(),
			Timestamp:               time.Now(),
		}
		vals = append(vals, val)
	}

	utils.Respond(rw, vals)
	return
}

// GetValidator returns validator information on the active chain
func (v *Validator) GetValidator(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	address := vars["address"]

	if address == "" {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "address is required")
		return
	}

	val, err := v.client.Validator(address)
	if err != nil {
		v.l.Printf("failed to query validator by address: %s", err)
		return
	}

	validator := &schema.Validator{
		ID:                      int32(0),
		Moniker:                 val.Description.Moniker,
		AccountAddress:          val.AccountAddress,
		OperatorAddress:         val.OperatorAddress,
		ConsensusAddress:        val.ConsensusAddress,
		Jailed:                  val.Jailed,
		Status:                  val.Status,
		Tokens:                  val.Tokens,
		VotingPower:             val.Power,
		DelegatorShares:         val.DelegatorShares,
		BondHeight:              val.BondHeight,
		BondIntraTxCounter:      val.BondIntraTxCounter,
		UnbondingHeight:         val.UnbondingHeight,
		UnbondingTime:           val.UnbondingTime.String(),
		CommissionRate:          val.Commission.Rate,
		CommissionMaxRate:       val.Commission.MaxRate,
		CommissionMaxChangeRate: val.Commission.MaxChangeRate,
		CommissionUpdateTime:    val.Commission.UpdateTime.String(),
		Timestamp:               time.Now(),
	}

	utils.Respond(rw, validator)
	//switch {
	//case strings.HasPrefix(address, v.nt.Bech32ValidatorAddrPrefix()):
	//	result, err := v.db.QueryValidatorByOperAddr(address)
	//	if err != nil {
	//		v.l.Printf("failed to query validator by operator address: %s", err)
	//		return
	//	}
	//	utils.Respond(rw, result)
	//	return
	//case strings.HasPrefix(address, v.nt.Bech32Prefixes()):
	//	result, err := v.db.QueryValidatorByAccountAddr(address)
	//	if err != nil {
	//		v.l.Printf("failed to query validator by account address: %s", err)
	//		return
	//	}
	//	utils.Respond(rw, result)
	//	return
	//case len(address) == 40:
	//	result, err := v.db.QueryValidatorByConsAddr(address)
	//	if err != nil {
	//		v.l.Printf("failed to query validator by consensus address: %s", err)
	//		return
	//	}
	//	utils.Respond(rw, result)
	//	return
	//default:
	//	result, err := v.db.QueryValidatorByMoniker(address)
	//	if err != nil {
	//		v.l.Printf("failed to query validator by moniker: %s", err)
	//		return
	//	}
	//	utils.Respond(rw, result)
	//	return
	//}
}
