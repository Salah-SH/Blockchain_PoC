package types

import "strings"

const QueryListCagnotte = "list-cagnotte"
const QueryResolveName = "resolve-name"
const QueryGetCagnotte = "get-cagnotte"
const QueryAdmin = "get-admin"

// QueryResResolve Queries Result Payload for a resolve query
type QueryResResolve struct {
	Value string `json:"value"`
}

// implement fmt.Stringer
func (r QueryResResolve) String() string {
	return r.Value
}

// QueryResNames Queries Result Payload for a names query
type QueryResNames []string

// implement fmt.Stringer
func (n QueryResNames) String() string {
	return strings.Join(n[:], "\n")
}
