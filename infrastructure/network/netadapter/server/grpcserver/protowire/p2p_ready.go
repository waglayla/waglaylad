package protowire

import (
	"github.com/waglayla/waglaylad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *WaglayladMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "WaglayladMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *WaglayladMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
