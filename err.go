package asfgo

import "errors"

var (
	ErrNoAppAndSub = errors.New("no appIDs and subIDs provided")
	ErrInvalidType = errors.New("invalid type")
)
