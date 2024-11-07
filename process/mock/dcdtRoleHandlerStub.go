package mock

import "github.com/kalyan3104/k-chain-go/state"

// DCDTRoleHandlerStub -
type DCDTRoleHandlerStub struct {
	CheckAllowedToExecuteCalled func(account state.UserAccountHandler, tokenID []byte, action []byte) error
}

// CheckAllowedToExecute -
func (e *DCDTRoleHandlerStub) CheckAllowedToExecute(account state.UserAccountHandler, tokenID []byte, action []byte) error {
	if e.CheckAllowedToExecuteCalled != nil {
		return e.CheckAllowedToExecuteCalled(account, tokenID, action)
	}

	return nil
}

// IsInterfaceNil -
func (e *DCDTRoleHandlerStub) IsInterfaceNil() bool {
	return e == nil
}
