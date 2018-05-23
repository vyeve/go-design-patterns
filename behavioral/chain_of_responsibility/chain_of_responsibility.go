/*
The single responsibility principle implies that a type, function, method, or any
similar abstraction must have one single responsibility only and it must do it quite well.
This way, we can apply many functions that achieve one specific thing each to some struct,
slice, map, and so on.
When we apply many of these abstractions in a logical way very often, we can chain them to
execute in order such as, for example, a logging chain.
*/

package chor

type ChainLogger interface {
	Next(string)
}
