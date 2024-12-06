// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/Luxc0/onvif"
	"github.com/Luxc0/onvif/sdk"
	"github.com/Luxc0/onvif/media"
)

// Call_GetVideoEncoderConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoEncoderConfigurationsResponse.
func Call_GetVideoEncoderConfigurations(ctx context.Context, dev *onvif.Device, request media.GetVideoEncoderConfigurations) (media.GetVideoEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoEncoderConfigurationsResponse media.GetVideoEncoderConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoEncoderConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetVideoEncoderConfigurations")
		return reply.Body.GetVideoEncoderConfigurationsResponse, errors.Annotate(err, "reply")
	}
}