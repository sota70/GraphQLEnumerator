// Main Package
package main

import (
	"flag"
	"fmt"

	"example.com/graphqlenumerator/commandargs"
	"example.com/graphqlenumerator/enumeration"
	"example.com/graphqlenumerator/query"
)

/*
Handle function

# Overview

handles subcommands depending on command arguments (which are Enumeration Mode, Query Mode, etc...)

if no valid mode is found in command arguments, print the usage

# Parameters

args commandargs.CommandArgs:

	Command arguments (like -v, -q, etc...)

# Return

returns nothing
*/
func handle(args commandargs.CommandArgs) {
	switch {
	case *args.E:
		fmt.Println(enumeration.Enumerate(args))
	case *args.Q:
		fmt.Println(query.Query(args))
	case *args.V:
		fmt.Printf("GraphQL Visualizer: https://graphql-kit.com/graphql-voyager\n1: Go to the website\n2: Click CHANGE SCHEMA\n3: Go to INTROSPECTION\n4: Copy the introspection query provided in this section\n5: Run the query and retrieve the result\n6: Paste the result in INTROSPECTION section and click DISPLAY\n")
	default:
		fmt.Println("Usage: ./graphqlenumerator [-e | -q] [-u url] [-query query]")
	}
}

/*
 * Main function
 * Retrieves command arguments from user input
 * and then executes it
 *
 * return: Nothing
 */
func main() {
	var (
		e     *bool   = flag.Bool("e", false, "Enumeration Mode")
		q     *bool   = flag.Bool("q", false, "Query Mode")
		v     *bool   = flag.Bool("v", false, "Shows graphql visualizer")
		c     *bool   = flag.Bool("c", false, "Copy the result to clipboard")
		query *string = flag.String("query", "{}", "GraphQL Query")
		u     *string = flag.String("u", "", "GraphQL endpoint URL")
	)
	flag.Parse()
	handle(commandargs.CommandArgs{e, q, v, c, u, query})
}
