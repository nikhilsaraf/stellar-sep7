package sep7

import (
	"log"

	"github.com/stellar/go/build"
	"github.com/stellar/go/xdr"
)

const maxUint256Bytes = 32

// SentinelSourceAccount represents a source account that is not necessarily valid but can still mutate a Transaction
// hide inside struct so it can only be used to create the empty byte array
type SentinelSourceAccount struct {
	sentinel [maxUint256Bytes]byte
}

// accountID is a helper method to get the accountID as PublicKeyTypePublicKeyTypeEd25519
func (s SentinelSourceAccount) accountID() (xdr.AccountId, error) {
	return xdr.NewAccountId(xdr.PublicKeyTypePublicKeyTypeEd25519, xdr.Uint256(s.sentinel))
}

// MutateTransaction for SentinelSourceAccount sets the SourceAccount on the TransactionBuilder's xdr.Transaction to the bytes
// stored in it without checking validity of the bytes. This is needed for SEP-7 compatibility
func (s SentinelSourceAccount) MutateTransaction(o *build.TransactionBuilder) error {
	var e error
	o.TX.SourceAccount, e = s.accountID()
	return e
}

// Address returns the address corresponding to the sentinel value along with a possible error
func (s SentinelSourceAccount) Address() (string, error) {
	aid, e := s.accountID()
	if e != nil {
		return "", e
	}
	return aid.Address(), nil
}

// MustAddress returns the address corresponding to the sentinel value, panicing if there is an error
func (s SentinelSourceAccount) MustAddress() string {
	address, e := s.Address()
	if e != nil {
		log.Panic(e)
	}
	return address
}

// Blank is a factory method that creates a SentinelSourceAccount with an empty byte array
func Blank() SentinelSourceAccount {
	return SentinelSourceAccount{}
}
