// Package schema provides the Chrome Debugging Protocol
// commands, types, and events for the Schema domain.
//
// Provides information about the protocol schema.
//
// Generated by the chromedp-gen command.
package schema

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
)

// GetDomainsParams returns supported domains.
type GetDomainsParams struct{}

// GetDomains returns supported domains.
func GetDomains() *GetDomainsParams {
	return &GetDomainsParams{}
}

// GetDomainsReturns return values.
type GetDomainsReturns struct {
	Domains []*Domain `json:"domains,omitempty"` // List of supported domains.
}

// Do executes Schema.getDomains against the provided context and
// target handler.
//
// returns:
//   domains - List of supported domains.
func (p *GetDomainsParams) Do(ctxt context.Context, h cdp.Handler) (domains []*Domain, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandSchemaGetDomains, cdp.Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return nil, cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r GetDomainsReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return nil, cdp.ErrInvalidResult
			}

			return r.Domains, nil

		case error:
			return nil, v
		}

	case <-ctxt.Done():
		return nil, ctxt.Err()
	}

	return nil, cdp.ErrUnknownResult
}
