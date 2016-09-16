// Package vbs prints text to stdout if Verbose is set to true. This package
// isn't threadsafe. Use a mutex if you're going to be changing Verbose in
// goroutines.
package vbs

import "fmt"

// Verbose indicates whether or not something should be printed.
var Verbose bool

// Println prints to stdout if Verbose is true.
func Println(a ...interface{}) {
	if Verbose {
		fmt.Println(a...)
	}
}

// Printf prints to stdout if Verbose is true.
func Printf(format string, a ...interface{}) {
	if Verbose {
		fmt.Printf(format, a...)
	}
}
