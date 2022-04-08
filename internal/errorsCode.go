package internal

import "errors"

var ID = errors.New("bad ID")
var NotFound = errors.New("not found")
var Decode = errors.New("decode error")
var BadRequest = errors.New("bad request")
var Conflict = errors.New("conflict request")
