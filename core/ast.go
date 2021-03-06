package core

import (
	"fmt"
	"math"

	"github.com/philandstuff/dhall-golang/term"
)

// A Value is a Dhall value in beta-normal form.  You can think of
// Values as the subset of Dhall which cannot be beta-reduced any
// further.  Valid Values include 3, "foo" and 5 + x.
//
// You create a Value by calling Eval() on a Term.  You can convert a
// Value back to a Term with Quote().
type Value interface {
	isValue()
}

// A Universe is a type of types.
type Universe int

// These are the valid Universes.
const (
	Type Universe = iota
	Kind
	Sort
)

// Builtin is the type of Dhall's builtin constants.
type Builtin string

// These are the Builtins.
const (
	Double  Builtin = "Double"
	Text    Builtin = "Text"
	Bool    Builtin = "Bool"
	Natural Builtin = "Natural"
	Integer Builtin = "Integer"
)

// A BoolLit is a Dhall boolean literal.
type BoolLit bool

func (BoolLit) isValue() {}

// Naturally, it is True or False.
const (
	True  BoolLit = true
	False BoolLit = false
)

type (
	naturalBuild struct{}
	naturalEven  struct{}
	naturalFold  struct {
		n    Value
		typ  Value
		succ Value
		// zero Value
	}
	naturalIsZero   struct{}
	naturalOdd      struct{}
	naturalShow     struct{}
	naturalSubtract struct {
		a Value
		// b Value
	}
	naturalToInteger struct{}

	integerClamp    struct{}
	integerNegate   struct{}
	integerShow     struct{}
	integerToDouble struct{}

	doubleShow struct{}

	optional      struct{}
	optionalBuild struct{ typ Value }
	optionalFold  struct {
		typ1 Value
		opt  Value
		typ2 Value
		some Value
		// none Value
	}
	none struct{}

	textShow struct{}

	list      struct{}
	listBuild struct {
		typ Value
		// fn  Value
	}
	listFold struct {
		typ1 Value
		list Value
		typ2 Value
		cons Value
		// empty Value
	}
	listLength  struct{ typ Value }
	listHead    struct{ typ Value }
	listLast    struct{ typ Value }
	listIndexed struct{ typ Value }
	listReverse struct{ typ Value }
)

func (naturalBuild) isValue()     {}
func (naturalEven) isValue()      {}
func (naturalFold) isValue()      {}
func (naturalIsZero) isValue()    {}
func (naturalOdd) isValue()       {}
func (naturalShow) isValue()      {}
func (naturalSubtract) isValue()  {}
func (naturalToInteger) isValue() {}

func (integerClamp) isValue()    {}
func (integerNegate) isValue()   {}
func (integerShow) isValue()     {}
func (integerToDouble) isValue() {}

func (doubleShow) isValue() {}

func (optional) isValue()      {}
func (optionalBuild) isValue() {}
func (optionalFold) isValue()  {}
func (none) isValue()          {}

func (textShow) isValue() {}

func (list) isValue()        {}
func (listBuild) isValue()   {}
func (listFold) isValue()    {}
func (listLength) isValue()  {}
func (listHead) isValue()    {}
func (listLast) isValue()    {}
func (listIndexed) isValue() {}
func (listReverse) isValue() {}

type (
	// OptionalOf is the Value version of `Optional a`
	OptionalOf struct{ Type Value }

	// ListOf is the Value version of `List a`
	ListOf struct{ Type Value }

	// NoneOf is the Value version of `None a`
	NoneOf struct{ Type Value }
)

func (OptionalOf) isValue() {}
func (ListOf) isValue()     {}
func (NoneOf) isValue()     {}

// A freeVar is a free variable.  It can appear in a Value where we
// Eval() a sub-Term within a whole, larger Term.
type freeVar struct {
	Name  string
	Index int
}

type (
	// A localVar is an internal sentinel value used by TypeOf() in
	// the process of typechecking the body of lambdas and pis.
	// Essentially it lets us convert de Bruijn indices (which count
	// how many binders we need to pass from the variable to the
	// correct binder) to de Bruijn levels (which count how many
	// binders we passed before binding this variable)
	localVar struct {
		Name  string
		Index int
	}

	// A quoteVar is an internal sentinel value used by Quote() in the
	// process of converting Values back to Terms.
	quoteVar struct {
		Name  string
		Index int
	}
)

func (Universe) isValue() {}

func (Builtin) isValue() {}

func (freeVar) isValue() {}

func (localVar) isValue() {}

func (quoteVar) isValue() {}

// Callable is a function Value that can be called with one Value
// argument.  Call() may return nil if normalization isn't possible
// (for example, `Natural/even x` does not normalize).  ArgType()
// returns the declared type of Call()'s parameter.
type Callable interface {
	Value
	Call(Value) Value
	ArgType() Value
}

func (l lambda) Call(a Value) Value {
	return l.Fn(a)
}

func (l lambda) ArgType() Value {
	return l.Domain
}

var (
	_ Callable = lambda{}
)

type (
	// A lambda is a go function representing a Dhall function
	// which has not yet been applied to its argument
	lambda struct {
		Label  string
		Domain Value
		Fn     func(Value) Value
	}

	// A Pi is the value of a Dhall Pi type.  Domain is the type
	// of the domain, and Range is a go function which returns the
	// type of the range, given the type of the domain.
	Pi struct {
		Label  string
		Domain Value
		Range  func(Value) Value
	}

	// An app is the Value of a Fn applied to an Arg.  Note that
	// this is only a valid Value if Fn is a free variable (such as f
	// 3, with f free), or if Fn is a builtin function which hasn't
	// been applied to enough arguments (such as Natural/subtract 3).
	app struct {
		Fn  Value
		Arg Value
	}

	oper struct {
		OpCode term.OpCode
		L      Value
		R      Value
	}
)

func (lambda) isValue() {}

func (Pi) isValue() {}

func (app) isValue() {}

func (oper) isValue() {}

// NewPi returns a new pi Value.
func NewPi(label string, d Value, r func(Value) Value) Pi {
	return Pi{
		Label:  label,
		Domain: d,
		Range:  r,
	}
}

// NewFnType returns a non-dependent function type Value.
func NewFnType(l string, d Value, r Value) Pi {
	return NewPi(l, d, func(Value) Value { return r })
}

type (
	// A NaturalLit is a literal of type Natural.
	NaturalLit uint

	// An EmptyList is an empty list literal Value of the given type.
	EmptyList struct{ Type Value }

	// A NonEmptyList is a non-empty list literal Value with the given contents.
	NonEmptyList []Value

	Chunk struct {
		Prefix string
		Expr   Value
	}
	Chunks  []Chunk
	TextLit struct {
		Chunks Chunks
		Suffix string
	}

	ifVal struct {
		Cond Value
		T    Value
		F    Value
	}

	// A DoubleLit is a literal of type Double.
	DoubleLit float64

	// A IntegerLit is a literal of type Integer.
	IntegerLit int

	// Some represents a Value which is present in an Optional type.
	Some struct{ Val Value }

	RecordType map[string]Value

	RecordLit map[string]Value

	toMap struct {
		Record Value
		Type   Value // optional
	}

	field struct {
		Record    Value
		FieldName string
	}

	project struct {
		Record     Value
		FieldNames []string
	}

	// no projectType because it cannot be in a normal form so cannot
	// be a Value

	UnionType map[string]Value

	merge struct {
		Handler    Value
		Union      Value
		Annotation Value // optional
	}

	assert struct{ Annotation Value }
)

func (NaturalLit) isValue() {}

func (EmptyList) isValue()    {}
func (NonEmptyList) isValue() {}

func (TextLit) isValue() {}

func (ifVal) isValue() {}

func (DoubleLit) isValue()  {}
func (IntegerLit) isValue() {}

func (d DoubleLit) String() string {
	f := float64(d)
	if math.IsInf(f, 1) {
		return "Infinity"
	}
	if math.IsInf(f, -1) {
		return "-Infinity"
	}
	// if we have a whole number, we need to append .0 to it so we get a valid
	// Double literal
	if f == float64(int64(f)) {
		return fmt.Sprintf("%#v.0", float64(d))
	}
	return fmt.Sprintf("%#v", float64(d))
}

func (Some) isValue() {}

func (RecordType) isValue() {}
func (RecordLit) isValue()  {}
func (toMap) isValue()      {}
func (field) isValue()      {}
func (project) isValue()    {}
func (UnionType) isValue()  {}
func (merge) isValue()      {}
func (assert) isValue()     {}
