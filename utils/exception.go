/*
Package Utils implements exception functions to make
- Block struct
- Exception interface
- Throw
- Do
*/
package utils

type Exception interface{}

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

// Throw is an exported method and is accessible in other packages to get if return nil
func Throw(up Exception) {
	panic(up)
}

// Do is an exported method and is accessible in other packages to handle process exception at every package
func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}

	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}

	tcf.Try()
}
