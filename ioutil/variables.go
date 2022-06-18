package ioutil

import "io"

// Discard is an io.Writer on which all Write calls succeed
// without doing anything.
var Discard io.Writer = io.Discard
