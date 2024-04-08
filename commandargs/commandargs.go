/*
Command arguments package

defines command arguments struct
*/
package commandargs

/*
CommandArgs struct

# Overview

flag E indicates that it is on enumeration mode

flag Q indicates that it is on query mode

flag V shows GraphQL mindmap URL

flag C allows it to copy the result to clipboard

flag U specifies GraphQL endpoint URL

flag Query specifies GraphQL query
*/
type CommandArgs struct {
	E     *bool
	Q     *bool
	V     *bool
	C     *bool
	U     *string
	Query *string
}
