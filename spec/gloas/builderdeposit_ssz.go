// Code generated manually for Gloas (EIP-8282). DO NOT EDIT via sszgen.
package gloas

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BuilderDeposit object.
func (b *BuilderDeposit) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BuilderDeposit object to a target array.
func (b *BuilderDeposit) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Pubkey'
	dst = append(dst, b.Pubkey[:]...)

	// Field (1) 'WithdrawalCredentials'
	if size := len(b.WithdrawalCredentials); size != 32 {
		err = ssz.ErrBytesLengthFn("BuilderDeposit.WithdrawalCredentials", size, 32)
		return
	}
	dst = append(dst, b.WithdrawalCredentials...)

	// Field (2) 'Amount'
	dst = ssz.MarshalUint64(dst, uint64(b.Amount))

	// Field (3) 'Signature'
	dst = append(dst, b.Signature[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the BuilderDeposit object.
func (b *BuilderDeposit) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 184 {
		return ssz.ErrSize
	}

	// Field (0) 'Pubkey'
	copy(b.Pubkey[:], buf[0:48])

	// Field (1) 'WithdrawalCredentials'
	if cap(b.WithdrawalCredentials) == 0 {
		b.WithdrawalCredentials = make([]byte, 0, 32)
	}
	b.WithdrawalCredentials = append(b.WithdrawalCredentials[:0], buf[48:80]...)

	// Field (2) 'Amount'
	b.Amount = phase0.Gwei(ssz.UnmarshallUint64(buf[80:88]))

	// Field (3) 'Signature'
	copy(b.Signature[:], buf[88:184])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BuilderDeposit object.
func (b *BuilderDeposit) SizeSSZ() (size int) {
	return 184
}

// HashTreeRoot ssz hashes the BuilderDeposit object.
func (b *BuilderDeposit) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BuilderDeposit object with a hasher.
func (b *BuilderDeposit) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Pubkey'
	hh.PutBytes(b.Pubkey[:])

	// Field (1) 'WithdrawalCredentials'
	if size := len(b.WithdrawalCredentials); size != 32 {
		err = ssz.ErrBytesLengthFn("BuilderDeposit.WithdrawalCredentials", size, 32)
		return
	}
	hh.PutBytes(b.WithdrawalCredentials)

	// Field (2) 'Amount'
	hh.PutUint64(uint64(b.Amount))

	// Field (3) 'Signature'
	hh.PutBytes(b.Signature[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BuilderDeposit object.
func (b *BuilderDeposit) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
