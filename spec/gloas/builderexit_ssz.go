// Code generated manually for Gloas (EIP-8282). DO NOT EDIT via sszgen.
package gloas

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BuilderExit object.
func (b *BuilderExit) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BuilderExit object to a target array.
func (b *BuilderExit) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'SourceAddress'
	dst = append(dst, b.SourceAddress[:]...)

	// Field (1) 'Pubkey'
	dst = append(dst, b.Pubkey[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the BuilderExit object.
func (b *BuilderExit) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 68 {
		return ssz.ErrSize
	}

	// Field (0) 'SourceAddress'
	copy(b.SourceAddress[:], buf[0:20])

	// Field (1) 'Pubkey'
	copy(b.Pubkey[:], buf[20:68])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BuilderExit object.
func (b *BuilderExit) SizeSSZ() (size int) {
	return 68
}

// HashTreeRoot ssz hashes the BuilderExit object.
func (b *BuilderExit) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BuilderExit object with a hasher.
func (b *BuilderExit) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'SourceAddress'
	hh.PutBytes(b.SourceAddress[:])

	// Field (1) 'Pubkey'
	hh.PutBytes(b.Pubkey[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BuilderExit object.
func (b *BuilderExit) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
