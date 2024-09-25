package protowire

import (
	"github.com/waglayla/waglaylad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *WaglayladMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "WaglayladMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *WaglayladMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
