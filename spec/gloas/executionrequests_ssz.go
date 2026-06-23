// Code generated manually for Gloas (EIP-8282). DO NOT EDIT via sszgen.
// Gloas ParentExecutionRequests has 5 variable-length fields; first SSZ offset = 20.
package gloas

import (
	"github.com/attestantio/go-eth2-client/spec/electra"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the ExecutionRequests object.
func (e *ExecutionRequests) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ExecutionRequests object to a target array.
func (e *ExecutionRequests) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(20) // 5 variable fields × 4 bytes each

	// Offset (0) 'Deposits'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.Deposits) * 192

	// Offset (1) 'Withdrawals'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.Withdrawals) * 76

	// Offset (2) 'Consolidations'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.Consolidations) * 116

	// Offset (3) 'BuilderDeposits'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.BuilderDeposits) * 184

	// Offset (4) 'BuilderExits'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.BuilderExits) * 68

	// Field (0) 'Deposits'
	if size := len(e.Deposits); size > 8192 {
		err = ssz.ErrListTooBigFn("ExecutionRequests.Deposits", size, 8192)
		return
	}
	for ii := 0; ii < len(e.Deposits); ii++ {
		if dst, err = e.Deposits[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (1) 'Withdrawals'
	if size := len(e.Withdrawals); size > 16 {
		err = ssz.ErrListTooBigFn("ExecutionRequests.Withdrawals", size, 16)
		return
	}
	for ii := 0; ii < len(e.Withdrawals); ii++ {
		if dst, err = e.Withdrawals[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (2) 'Consolidations'
	if size := len(e.Consolidations); size > 2 {
		err = ssz.ErrListTooBigFn("ExecutionRequests.Consolidations", size, 2)
		return
	}
	for ii := 0; ii < len(e.Consolidations); ii++ {
		if dst, err = e.Consolidations[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (3) 'BuilderDeposits'
	if size := len(e.BuilderDeposits); size > 256 {
		err = ssz.ErrListTooBigFn("ExecutionRequests.BuilderDeposits", size, 256)
		return
	}
	for ii := 0; ii < len(e.BuilderDeposits); ii++ {
		if dst, err = e.BuilderDeposits[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (4) 'BuilderExits'
	if size := len(e.BuilderExits); size > 16 {
		err = ssz.ErrListTooBigFn("ExecutionRequests.BuilderExits", size, 16)
		return
	}
	for ii := 0; ii < len(e.BuilderExits); ii++ {
		if dst, err = e.BuilderExits[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the ExecutionRequests object.
func (e *ExecutionRequests) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 20 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o2, o3, o4 uint64

	// Offset (0) 'Deposits'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}
	if o0 < 20 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'Withdrawals'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Offset (2) 'Consolidations'
	if o2 = ssz.ReadOffset(buf[8:12]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Offset (3) 'BuilderDeposits'
	if o3 = ssz.ReadOffset(buf[12:16]); o3 > size || o2 > o3 {
		return ssz.ErrOffset
	}

	// Offset (4) 'BuilderExits'
	if o4 = ssz.ReadOffset(buf[16:20]); o4 > size || o3 > o4 {
		return ssz.ErrOffset
	}

	// Field (0) 'Deposits'
	{
		buf = tail[o0:o1]
		num, err := ssz.DivideInt2(len(buf), 192, 8192)
		if err != nil {
			return err
		}
		e.Deposits = make([]*electra.DepositRequest, num)
		for ii := 0; ii < num; ii++ {
			if e.Deposits[ii] == nil {
				e.Deposits[ii] = new(electra.DepositRequest)
			}
			if err = e.Deposits[ii].UnmarshalSSZ(buf[ii*192 : (ii+1)*192]); err != nil {
				return err
			}
		}
	}

	// Field (1) 'Withdrawals'
	{
		buf = tail[o1:o2]
		num, err := ssz.DivideInt2(len(buf), 76, 16)
		if err != nil {
			return err
		}
		e.Withdrawals = make([]*electra.WithdrawalRequest, num)
		for ii := 0; ii < num; ii++ {
			if e.Withdrawals[ii] == nil {
				e.Withdrawals[ii] = new(electra.WithdrawalRequest)
			}
			if err = e.Withdrawals[ii].UnmarshalSSZ(buf[ii*76 : (ii+1)*76]); err != nil {
				return err
			}
		}
	}

	// Field (2) 'Consolidations'
	{
		buf = tail[o2:o3]
		num, err := ssz.DivideInt2(len(buf), 116, 2)
		if err != nil {
			return err
		}
		e.Consolidations = make([]*electra.ConsolidationRequest, num)
		for ii := 0; ii < num; ii++ {
			if e.Consolidations[ii] == nil {
				e.Consolidations[ii] = new(electra.ConsolidationRequest)
			}
			if err = e.Consolidations[ii].UnmarshalSSZ(buf[ii*116 : (ii+1)*116]); err != nil {
				return err
			}
		}
	}

	// Field (3) 'BuilderDeposits'
	{
		buf = tail[o3:o4]
		num, err := ssz.DivideInt2(len(buf), 184, 256)
		if err != nil {
			return err
		}
		e.BuilderDeposits = make([]*BuilderDeposit, num)
		for ii := 0; ii < num; ii++ {
			if e.BuilderDeposits[ii] == nil {
				e.BuilderDeposits[ii] = new(BuilderDeposit)
			}
			if err = e.BuilderDeposits[ii].UnmarshalSSZ(buf[ii*184 : (ii+1)*184]); err != nil {
				return err
			}
		}
	}

	// Field (4) 'BuilderExits'
	{
		buf = tail[o4:]
		num, err := ssz.DivideInt2(len(buf), 68, 16)
		if err != nil {
			return err
		}
		e.BuilderExits = make([]*BuilderExit, num)
		for ii := 0; ii < num; ii++ {
			if e.BuilderExits[ii] == nil {
				e.BuilderExits[ii] = new(BuilderExit)
			}
			if err = e.BuilderExits[ii].UnmarshalSSZ(buf[ii*68 : (ii+1)*68]); err != nil {
				return err
			}
		}
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ExecutionRequests object.
func (e *ExecutionRequests) SizeSSZ() (size int) {
	size = 20 // 5 offsets × 4 bytes

	// Field (0) 'Deposits'
	size += len(e.Deposits) * 192

	// Field (1) 'Withdrawals'
	size += len(e.Withdrawals) * 76

	// Field (2) 'Consolidations'
	size += len(e.Consolidations) * 116

	// Field (3) 'BuilderDeposits'
	size += len(e.BuilderDeposits) * 184

	// Field (4) 'BuilderExits'
	size += len(e.BuilderExits) * 68

	return
}

// HashTreeRoot ssz hashes the ExecutionRequests object.
func (e *ExecutionRequests) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ExecutionRequests object with a hasher.
func (e *ExecutionRequests) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Deposits'
	{
		subIndx := hh.Index()
		num := uint64(len(e.Deposits))
		if num > 8192 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range e.Deposits {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 8192)
	}

	// Field (1) 'Withdrawals'
	{
		subIndx := hh.Index()
		num := uint64(len(e.Withdrawals))
		if num > 16 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range e.Withdrawals {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 16)
	}

	// Field (2) 'Consolidations'
	{
		subIndx := hh.Index()
		num := uint64(len(e.Consolidations))
		if num > 2 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range e.Consolidations {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 2)
	}

	// Field (3) 'BuilderDeposits'
	{
		subIndx := hh.Index()
		num := uint64(len(e.BuilderDeposits))
		if num > 256 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range e.BuilderDeposits {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 256)
	}

	// Field (4) 'BuilderExits'
	{
		subIndx := hh.Index()
		num := uint64(len(e.BuilderExits))
		if num > 16 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range e.BuilderExits {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 16)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ExecutionRequests object.
func (e *ExecutionRequests) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(e)
}
