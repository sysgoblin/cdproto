// Package io provides the Chrome Debugging Protocol
// commands, types, and events for the IO domain.
//
// Input/Output operations for streams produced by DevTools.
//
// Generated by the chromedp-gen command.
package io

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
)

// ReadParams read a chunk of the stream.
type ReadParams struct {
	Handle StreamHandle `json:"handle"`           // Handle of the stream to read.
	Offset int64        `json:"offset,omitempty"` // Seek to the specified offset before reading (if not specificed, proceed with offset following the last read).
	Size   int64        `json:"size,omitempty"`   // Maximum number of bytes to read (left upon the agent discretion if not specified).
}

// Read read a chunk of the stream.
//
// parameters:
//   handle - Handle of the stream to read.
func Read(handle StreamHandle) *ReadParams {
	return &ReadParams{
		Handle: handle,
	}
}

// WithOffset seek to the specified offset before reading (if not specificed,
// proceed with offset following the last read).
func (p ReadParams) WithOffset(offset int64) *ReadParams {
	p.Offset = offset
	return &p
}

// WithSize maximum number of bytes to read (left upon the agent discretion
// if not specified).
func (p ReadParams) WithSize(size int64) *ReadParams {
	p.Size = size
	return &p
}

// ReadReturns return values.
type ReadReturns struct {
	Data string `json:"data,omitempty"` // Data that were read.
	EOF  bool   `json:"eof,omitempty"`  // Set if the end-of-file condition occured while reading.
}

// Do executes IO.read against the provided context and
// target handler.
//
// returns:
//   data - Data that were read.
//   eof - Set if the end-of-file condition occured while reading.
func (p *ReadParams) Do(ctxt context.Context, h cdp.Handler) (data string, eof bool, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return "", false, err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandIORead, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return "", false, cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r ReadReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return "", false, cdp.ErrInvalidResult
			}

			return r.Data, r.EOF, nil

		case error:
			return "", false, v
		}

	case <-ctxt.Done():
		return "", false, ctxt.Err()
	}

	return "", false, cdp.ErrUnknownResult
}

// CloseParams close the stream, discard any temporary backing storage.
type CloseParams struct {
	Handle StreamHandle `json:"handle"` // Handle of the stream to close.
}

// Close close the stream, discard any temporary backing storage.
//
// parameters:
//   handle - Handle of the stream to close.
func Close(handle StreamHandle) *CloseParams {
	return &CloseParams{
		Handle: handle,
	}
}

// Do executes IO.close against the provided context and
// target handler.
func (p *CloseParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandIOClose, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}
