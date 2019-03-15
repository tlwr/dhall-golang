
package parser

import (
"bytes"
"errors"
"fmt"
"io"
"io/ioutil"
"math"
"os"
"strconv"
"strings"
"unicode"
"unicode/utf8"
)
import "github.com/philandstuff/dhall-golang/ast"


var g = &grammar {
	rules: []*rule{
{
	name: "DhallFile",
	pos: position{line: 21, col: 1, offset: 180},
	expr: &actionExpr{
	pos: position{line: 21, col: 13, offset: 194},
	run: (*parser).callonDhallFile1,
	expr: &seqExpr{
	pos: position{line: 21, col: 13, offset: 194},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 21, col: 13, offset: 194},
	label: "e",
	expr: &ruleRefExpr{
	pos: position{line: 21, col: 15, offset: 196},
	name: "CompleteExpression",
},
},
&ruleRefExpr{
	pos: position{line: 21, col: 34, offset: 215},
	name: "EOF",
},
	},
},
},
},
{
	name: "CompleteExpression",
	pos: position{line: 23, col: 1, offset: 238},
	expr: &actionExpr{
	pos: position{line: 23, col: 22, offset: 261},
	run: (*parser).callonCompleteExpression1,
	expr: &seqExpr{
	pos: position{line: 23, col: 22, offset: 261},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 23, col: 22, offset: 261},
	label: "e",
	expr: &ruleRefExpr{
	pos: position{line: 23, col: 24, offset: 263},
	name: "Expression",
},
},
&ruleRefExpr{
	pos: position{line: 23, col: 35, offset: 274},
	name: "_",
},
	},
},
},
},
{
	name: "EOL",
	pos: position{line: 25, col: 1, offset: 295},
	expr: &choiceExpr{
	pos: position{line: 25, col: 7, offset: 303},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 25, col: 7, offset: 303},
	val: "\n",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 25, col: 14, offset: 310},
	val: "\r\n",
	ignoreCase: false,
},
	},
},
},
{
	name: "BlockComment",
	pos: position{line: 27, col: 1, offset: 318},
	expr: &seqExpr{
	pos: position{line: 27, col: 16, offset: 335},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 27, col: 16, offset: 335},
	val: "{-",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 27, col: 21, offset: 340},
	name: "BlockCommentContinue",
},
	},
},
},
{
	name: "BlockCommentChunk",
	pos: position{line: 29, col: 1, offset: 362},
	expr: &choiceExpr{
	pos: position{line: 30, col: 5, offset: 388},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 30, col: 5, offset: 388},
	name: "BlockComment",
},
&charClassMatcher{
	pos: position{line: 31, col: 5, offset: 405},
	val: "[\\t\\u0020-\\U0010ffff]",
	chars: []rune{'\t',},
	ranges: []rune{' ','\U0010ffff',},
	ignoreCase: false,
	inverted: false,
},
&ruleRefExpr{
	pos: position{line: 32, col: 5, offset: 431},
	name: "EOL",
},
	},
},
},
{
	name: "BlockCommentContinue",
	pos: position{line: 34, col: 1, offset: 436},
	expr: &choiceExpr{
	pos: position{line: 34, col: 24, offset: 461},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 34, col: 24, offset: 461},
	val: "-}",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 34, col: 31, offset: 468},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 34, col: 31, offset: 468},
	name: "BlockCommentChunk",
},
&ruleRefExpr{
	pos: position{line: 34, col: 49, offset: 486},
	name: "BlockCommentContinue",
},
	},
},
	},
},
},
{
	name: "NotEOL",
	pos: position{line: 36, col: 1, offset: 508},
	expr: &charClassMatcher{
	pos: position{line: 36, col: 10, offset: 519},
	val: "[\\t\\u0020-\\U0010ffff]",
	chars: []rune{'\t',},
	ranges: []rune{' ','\U0010ffff',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "LineComment",
	pos: position{line: 38, col: 1, offset: 542},
	expr: &actionExpr{
	pos: position{line: 38, col: 15, offset: 558},
	run: (*parser).callonLineComment1,
	expr: &seqExpr{
	pos: position{line: 38, col: 15, offset: 558},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 38, col: 15, offset: 558},
	val: "--",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 38, col: 20, offset: 563},
	label: "content",
	expr: &actionExpr{
	pos: position{line: 38, col: 29, offset: 572},
	run: (*parser).callonLineComment5,
	expr: &zeroOrMoreExpr{
	pos: position{line: 38, col: 29, offset: 572},
	expr: &ruleRefExpr{
	pos: position{line: 38, col: 29, offset: 572},
	name: "NotEOL",
},
},
},
},
&ruleRefExpr{
	pos: position{line: 38, col: 68, offset: 611},
	name: "EOL",
},
	},
},
},
},
{
	name: "WhitespaceChunk",
	pos: position{line: 41, col: 1, offset: 640},
	expr: &choiceExpr{
	pos: position{line: 41, col: 19, offset: 660},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 41, col: 19, offset: 660},
	val: " ",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 41, col: 25, offset: 666},
	val: "\t",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 41, col: 32, offset: 673},
	name: "EOL",
},
&ruleRefExpr{
	pos: position{line: 41, col: 38, offset: 679},
	name: "LineComment",
},
&ruleRefExpr{
	pos: position{line: 41, col: 52, offset: 693},
	name: "BlockComment",
},
	},
},
},
{
	name: "_",
	pos: position{line: 43, col: 1, offset: 707},
	expr: &zeroOrMoreExpr{
	pos: position{line: 43, col: 5, offset: 713},
	expr: &ruleRefExpr{
	pos: position{line: 43, col: 5, offset: 713},
	name: "WhitespaceChunk",
},
},
},
{
	name: "HexDig",
	pos: position{line: 45, col: 1, offset: 731},
	expr: &charClassMatcher{
	pos: position{line: 45, col: 10, offset: 742},
	val: "[0-9a-f]i",
	ranges: []rune{'0','9','a','f',},
	ignoreCase: true,
	inverted: false,
},
},
{
	name: "SimpleLabel",
	pos: position{line: 47, col: 1, offset: 753},
	expr: &actionExpr{
	pos: position{line: 47, col: 15, offset: 769},
	run: (*parser).callonSimpleLabel1,
	expr: &seqExpr{
	pos: position{line: 47, col: 15, offset: 769},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 47, col: 15, offset: 769},
	expr: &ruleRefExpr{
	pos: position{line: 47, col: 16, offset: 770},
	name: "KeywordRaw",
},
},
&charClassMatcher{
	pos: position{line: 48, col: 13, offset: 793},
	val: "[A-Za-z_]",
	chars: []rune{'_',},
	ranges: []rune{'A','Z','a','z',},
	ignoreCase: false,
	inverted: false,
},
&zeroOrMoreExpr{
	pos: position{line: 48, col: 23, offset: 803},
	expr: &charClassMatcher{
	pos: position{line: 48, col: 23, offset: 803},
	val: "[A-Za-z0-9_/-]",
	chars: []rune{'_','/','-',},
	ranges: []rune{'A','Z','a','z','0','9',},
	ignoreCase: false,
	inverted: false,
},
},
	},
},
},
},
{
	name: "Label",
	pos: position{line: 52, col: 1, offset: 867},
	expr: &actionExpr{
	pos: position{line: 52, col: 9, offset: 877},
	run: (*parser).callonLabel1,
	expr: &seqExpr{
	pos: position{line: 52, col: 9, offset: 877},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 52, col: 9, offset: 877},
	name: "_",
},
&labeledExpr{
	pos: position{line: 52, col: 11, offset: 879},
	label: "label",
	expr: &ruleRefExpr{
	pos: position{line: 52, col: 17, offset: 885},
	name: "SimpleLabel",
},
},
	},
},
},
},
{
	name: "EscapedChar",
	pos: position{line: 56, col: 1, offset: 953},
	expr: &actionExpr{
	pos: position{line: 57, col: 3, offset: 971},
	run: (*parser).callonEscapedChar1,
	expr: &seqExpr{
	pos: position{line: 57, col: 3, offset: 971},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 57, col: 3, offset: 971},
	val: "\\",
	ignoreCase: false,
},
&choiceExpr{
	pos: position{line: 58, col: 5, offset: 980},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 58, col: 5, offset: 980},
	val: "\"",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 59, col: 10, offset: 993},
	val: "$",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 60, col: 10, offset: 1006},
	val: "\\",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 61, col: 10, offset: 1020},
	val: "/",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 62, col: 10, offset: 1033},
	val: "b",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 63, col: 10, offset: 1046},
	val: "f",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 64, col: 10, offset: 1059},
	val: "n",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 65, col: 10, offset: 1072},
	val: "r",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 66, col: 10, offset: 1085},
	val: "t",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 67, col: 10, offset: 1098},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 67, col: 10, offset: 1098},
	val: "u",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 67, col: 14, offset: 1102},
	name: "HexDig",
},
&ruleRefExpr{
	pos: position{line: 67, col: 21, offset: 1109},
	name: "HexDig",
},
&ruleRefExpr{
	pos: position{line: 67, col: 28, offset: 1116},
	name: "HexDig",
},
&ruleRefExpr{
	pos: position{line: 67, col: 35, offset: 1123},
	name: "HexDig",
},
	},
},
	},
},
	},
},
},
},
{
	name: "DoubleQuoteChunk",
	pos: position{line: 88, col: 1, offset: 1566},
	expr: &choiceExpr{
	pos: position{line: 89, col: 6, offset: 1592},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 89, col: 6, offset: 1592},
	run: (*parser).callonDoubleQuoteChunk2,
	expr: &seqExpr{
	pos: position{line: 89, col: 6, offset: 1592},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 89, col: 6, offset: 1592},
	val: "${",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 89, col: 11, offset: 1597},
	label: "e",
	expr: &ruleRefExpr{
	pos: position{line: 89, col: 13, offset: 1599},
	name: "CompleteExpression",
},
},
&litMatcher{
	pos: position{line: 89, col: 32, offset: 1618},
	val: "}",
	ignoreCase: false,
},
	},
},
},
&ruleRefExpr{
	pos: position{line: 90, col: 6, offset: 1645},
	name: "EscapedChar",
},
&charClassMatcher{
	pos: position{line: 91, col: 6, offset: 1662},
	val: "[\\x20-\\x21]",
	ranges: []rune{' ','!',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 92, col: 6, offset: 1679},
	val: "[\\x23-\\x5b]",
	ranges: []rune{'#','[',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 93, col: 6, offset: 1696},
	val: "[\\x5d-\\U0010ffff]",
	ranges: []rune{']','\U0010ffff',},
	ignoreCase: false,
	inverted: false,
},
	},
},
},
{
	name: "DoubleQuoteLiteral",
	pos: position{line: 95, col: 1, offset: 1715},
	expr: &actionExpr{
	pos: position{line: 95, col: 22, offset: 1738},
	run: (*parser).callonDoubleQuoteLiteral1,
	expr: &seqExpr{
	pos: position{line: 95, col: 22, offset: 1738},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 95, col: 22, offset: 1738},
	val: "\"",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 95, col: 26, offset: 1742},
	label: "chunks",
	expr: &zeroOrMoreExpr{
	pos: position{line: 95, col: 33, offset: 1749},
	expr: &ruleRefExpr{
	pos: position{line: 95, col: 33, offset: 1749},
	name: "DoubleQuoteChunk",
},
},
},
&litMatcher{
	pos: position{line: 95, col: 51, offset: 1767},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "TextLiteral",
	pos: position{line: 112, col: 1, offset: 2251},
	expr: &actionExpr{
	pos: position{line: 112, col: 15, offset: 2267},
	run: (*parser).callonTextLiteral1,
	expr: &seqExpr{
	pos: position{line: 112, col: 15, offset: 2267},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 112, col: 15, offset: 2267},
	name: "_",
},
&labeledExpr{
	pos: position{line: 112, col: 17, offset: 2269},
	label: "t",
	expr: &ruleRefExpr{
	pos: position{line: 112, col: 19, offset: 2271},
	name: "DoubleQuoteLiteral",
},
},
	},
},
},
},
{
	name: "ReservedRaw",
	pos: position{line: 114, col: 1, offset: 2309},
	expr: &choiceExpr{
	pos: position{line: 114, col: 15, offset: 2325},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 114, col: 15, offset: 2325},
	run: (*parser).callonReservedRaw2,
	expr: &litMatcher{
	pos: position{line: 114, col: 15, offset: 2325},
	val: "Bool",
	ignoreCase: false,
},
},
&litMatcher{
	pos: position{line: 115, col: 5, offset: 2361},
	val: "Optional",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 116, col: 5, offset: 2376},
	val: "None",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 117, col: 5, offset: 2387},
	run: (*parser).callonReservedRaw6,
	expr: &litMatcher{
	pos: position{line: 117, col: 5, offset: 2387},
	val: "Natural",
	ignoreCase: false,
},
},
&actionExpr{
	pos: position{line: 118, col: 5, offset: 2429},
	run: (*parser).callonReservedRaw8,
	expr: &litMatcher{
	pos: position{line: 118, col: 5, offset: 2429},
	val: "Integer",
	ignoreCase: false,
},
},
&actionExpr{
	pos: position{line: 119, col: 5, offset: 2471},
	run: (*parser).callonReservedRaw10,
	expr: &litMatcher{
	pos: position{line: 119, col: 5, offset: 2471},
	val: "Double",
	ignoreCase: false,
},
},
&actionExpr{
	pos: position{line: 120, col: 5, offset: 2511},
	run: (*parser).callonReservedRaw12,
	expr: &litMatcher{
	pos: position{line: 120, col: 5, offset: 2511},
	val: "Text",
	ignoreCase: false,
},
},
&actionExpr{
	pos: position{line: 121, col: 5, offset: 2547},
	run: (*parser).callonReservedRaw14,
	expr: &litMatcher{
	pos: position{line: 121, col: 5, offset: 2547},
	val: "List",
	ignoreCase: false,
},
},
&actionExpr{
	pos: position{line: 122, col: 5, offset: 2583},
	run: (*parser).callonReservedRaw16,
	expr: &litMatcher{
	pos: position{line: 122, col: 5, offset: 2583},
	val: "True",
	ignoreCase: false,
},
},
&actionExpr{
	pos: position{line: 123, col: 5, offset: 2619},
	run: (*parser).callonReservedRaw18,
	expr: &litMatcher{
	pos: position{line: 123, col: 5, offset: 2619},
	val: "False",
	ignoreCase: false,
},
},
&actionExpr{
	pos: position{line: 124, col: 5, offset: 2657},
	run: (*parser).callonReservedRaw20,
	expr: &litMatcher{
	pos: position{line: 124, col: 5, offset: 2657},
	val: "NaN",
	ignoreCase: false,
},
},
&actionExpr{
	pos: position{line: 125, col: 5, offset: 2709},
	run: (*parser).callonReservedRaw22,
	expr: &litMatcher{
	pos: position{line: 125, col: 5, offset: 2709},
	val: "Infinity",
	ignoreCase: false,
},
},
&actionExpr{
	pos: position{line: 126, col: 5, offset: 2767},
	run: (*parser).callonReservedRaw24,
	expr: &litMatcher{
	pos: position{line: 126, col: 5, offset: 2767},
	val: "Type",
	ignoreCase: false,
},
},
&actionExpr{
	pos: position{line: 127, col: 5, offset: 2803},
	run: (*parser).callonReservedRaw26,
	expr: &litMatcher{
	pos: position{line: 127, col: 5, offset: 2803},
	val: "Kind",
	ignoreCase: false,
},
},
&actionExpr{
	pos: position{line: 128, col: 5, offset: 2839},
	run: (*parser).callonReservedRaw28,
	expr: &litMatcher{
	pos: position{line: 128, col: 5, offset: 2839},
	val: "Sort",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "Reserved",
	pos: position{line: 130, col: 1, offset: 2872},
	expr: &actionExpr{
	pos: position{line: 130, col: 12, offset: 2885},
	run: (*parser).callonReserved1,
	expr: &seqExpr{
	pos: position{line: 130, col: 12, offset: 2885},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 130, col: 12, offset: 2885},
	name: "_",
},
&labeledExpr{
	pos: position{line: 130, col: 14, offset: 2887},
	label: "r",
	expr: &ruleRefExpr{
	pos: position{line: 130, col: 16, offset: 2889},
	name: "ReservedRaw",
},
},
	},
},
},
},
{
	name: "KeywordRaw",
	pos: position{line: 132, col: 1, offset: 2920},
	expr: &choiceExpr{
	pos: position{line: 132, col: 14, offset: 2935},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 132, col: 14, offset: 2935},
	val: "if",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 133, col: 5, offset: 2944},
	val: "then",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 134, col: 5, offset: 2955},
	val: "else",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 135, col: 5, offset: 2966},
	val: "let",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 136, col: 5, offset: 2976},
	val: "in",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 137, col: 5, offset: 2985},
	val: "as",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 138, col: 5, offset: 2994},
	val: "using",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 139, col: 5, offset: 3006},
	val: "merge",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 140, col: 5, offset: 3018},
	val: "constructors",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 141, col: 5, offset: 3037},
	val: "Some",
	ignoreCase: false,
},
	},
},
},
{
	name: "If",
	pos: position{line: 143, col: 1, offset: 3045},
	expr: &seqExpr{
	pos: position{line: 143, col: 6, offset: 3052},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 143, col: 6, offset: 3052},
	name: "_",
},
&litMatcher{
	pos: position{line: 143, col: 8, offset: 3054},
	val: "if",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 143, col: 13, offset: 3059},
	name: "WhitespaceChunk",
},
	},
},
},
{
	name: "Then",
	pos: position{line: 144, col: 1, offset: 3075},
	expr: &seqExpr{
	pos: position{line: 144, col: 8, offset: 3084},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 144, col: 8, offset: 3084},
	name: "_",
},
&litMatcher{
	pos: position{line: 144, col: 10, offset: 3086},
	val: "then",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 144, col: 17, offset: 3093},
	name: "WhitespaceChunk",
},
	},
},
},
{
	name: "Else",
	pos: position{line: 145, col: 1, offset: 3109},
	expr: &seqExpr{
	pos: position{line: 145, col: 8, offset: 3118},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 145, col: 8, offset: 3118},
	name: "_",
},
&litMatcher{
	pos: position{line: 145, col: 10, offset: 3120},
	val: "else",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 145, col: 17, offset: 3127},
	name: "WhitespaceChunk",
},
	},
},
},
{
	name: "Let",
	pos: position{line: 146, col: 1, offset: 3143},
	expr: &seqExpr{
	pos: position{line: 146, col: 7, offset: 3151},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 146, col: 7, offset: 3151},
	name: "_",
},
&litMatcher{
	pos: position{line: 146, col: 9, offset: 3153},
	val: "let",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 146, col: 15, offset: 3159},
	name: "WhitespaceChunk",
},
	},
},
},
{
	name: "In",
	pos: position{line: 147, col: 1, offset: 3175},
	expr: &seqExpr{
	pos: position{line: 147, col: 6, offset: 3182},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 147, col: 6, offset: 3182},
	name: "_",
},
&litMatcher{
	pos: position{line: 147, col: 8, offset: 3184},
	val: "in",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 147, col: 13, offset: 3189},
	name: "WhitespaceChunk",
},
	},
},
},
{
	name: "As",
	pos: position{line: 148, col: 1, offset: 3205},
	expr: &seqExpr{
	pos: position{line: 148, col: 6, offset: 3212},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 148, col: 6, offset: 3212},
	name: "_",
},
&litMatcher{
	pos: position{line: 148, col: 8, offset: 3214},
	val: "as",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 148, col: 13, offset: 3219},
	name: "WhitespaceChunk",
},
	},
},
},
{
	name: "Using",
	pos: position{line: 149, col: 1, offset: 3235},
	expr: &seqExpr{
	pos: position{line: 149, col: 9, offset: 3245},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 149, col: 9, offset: 3245},
	name: "_",
},
&litMatcher{
	pos: position{line: 149, col: 11, offset: 3247},
	val: "using",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 149, col: 19, offset: 3255},
	name: "WhitespaceChunk",
},
	},
},
},
{
	name: "Merge",
	pos: position{line: 150, col: 1, offset: 3271},
	expr: &seqExpr{
	pos: position{line: 150, col: 9, offset: 3281},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 150, col: 9, offset: 3281},
	name: "_",
},
&litMatcher{
	pos: position{line: 150, col: 11, offset: 3283},
	val: "merge",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 150, col: 19, offset: 3291},
	name: "WhitespaceChunk",
},
	},
},
},
{
	name: "Some",
	pos: position{line: 151, col: 1, offset: 3307},
	expr: &seqExpr{
	pos: position{line: 151, col: 8, offset: 3316},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 151, col: 8, offset: 3316},
	name: "_",
},
&litMatcher{
	pos: position{line: 151, col: 10, offset: 3318},
	val: "Some",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 151, col: 17, offset: 3325},
	name: "WhitespaceChunk",
},
	},
},
},
{
	name: "Optional",
	pos: position{line: 152, col: 1, offset: 3341},
	expr: &seqExpr{
	pos: position{line: 152, col: 12, offset: 3354},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 152, col: 12, offset: 3354},
	name: "_",
},
&litMatcher{
	pos: position{line: 152, col: 14, offset: 3356},
	val: "Optional",
	ignoreCase: false,
},
	},
},
},
{
	name: "Text",
	pos: position{line: 153, col: 1, offset: 3367},
	expr: &seqExpr{
	pos: position{line: 153, col: 8, offset: 3376},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 153, col: 8, offset: 3376},
	name: "_",
},
&litMatcher{
	pos: position{line: 153, col: 10, offset: 3378},
	val: "Text",
	ignoreCase: false,
},
	},
},
},
{
	name: "List",
	pos: position{line: 154, col: 1, offset: 3385},
	expr: &seqExpr{
	pos: position{line: 154, col: 8, offset: 3394},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 154, col: 8, offset: 3394},
	name: "_",
},
&litMatcher{
	pos: position{line: 154, col: 10, offset: 3396},
	val: "List",
	ignoreCase: false,
},
	},
},
},
{
	name: "Equal",
	pos: position{line: 156, col: 1, offset: 3404},
	expr: &seqExpr{
	pos: position{line: 156, col: 9, offset: 3414},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 156, col: 9, offset: 3414},
	name: "_",
},
&litMatcher{
	pos: position{line: 156, col: 11, offset: 3416},
	val: "=",
	ignoreCase: false,
},
	},
},
},
{
	name: "Plus",
	pos: position{line: 157, col: 1, offset: 3420},
	expr: &seqExpr{
	pos: position{line: 157, col: 8, offset: 3429},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 157, col: 8, offset: 3429},
	name: "_",
},
&litMatcher{
	pos: position{line: 157, col: 10, offset: 3431},
	val: "+",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 157, col: 14, offset: 3435},
	name: "WhitespaceChunk",
},
	},
},
},
{
	name: "Times",
	pos: position{line: 158, col: 1, offset: 3451},
	expr: &seqExpr{
	pos: position{line: 158, col: 9, offset: 3461},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 158, col: 9, offset: 3461},
	name: "_",
},
&litMatcher{
	pos: position{line: 158, col: 11, offset: 3463},
	val: "*",
	ignoreCase: false,
},
	},
},
},
{
	name: "Dot",
	pos: position{line: 159, col: 1, offset: 3467},
	expr: &seqExpr{
	pos: position{line: 159, col: 7, offset: 3475},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 159, col: 7, offset: 3475},
	name: "_",
},
&litMatcher{
	pos: position{line: 159, col: 9, offset: 3477},
	val: ".",
	ignoreCase: false,
},
	},
},
},
{
	name: "OpenBrace",
	pos: position{line: 160, col: 1, offset: 3481},
	expr: &seqExpr{
	pos: position{line: 160, col: 13, offset: 3495},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 160, col: 13, offset: 3495},
	name: "_",
},
&litMatcher{
	pos: position{line: 160, col: 15, offset: 3497},
	val: "{",
	ignoreCase: false,
},
	},
},
},
{
	name: "CloseBrace",
	pos: position{line: 161, col: 1, offset: 3501},
	expr: &seqExpr{
	pos: position{line: 161, col: 14, offset: 3516},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 161, col: 14, offset: 3516},
	name: "_",
},
&litMatcher{
	pos: position{line: 161, col: 16, offset: 3518},
	val: "}",
	ignoreCase: false,
},
	},
},
},
{
	name: "OpenBracket",
	pos: position{line: 162, col: 1, offset: 3522},
	expr: &seqExpr{
	pos: position{line: 162, col: 15, offset: 3538},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 162, col: 15, offset: 3538},
	name: "_",
},
&litMatcher{
	pos: position{line: 162, col: 17, offset: 3540},
	val: "[",
	ignoreCase: false,
},
	},
},
},
{
	name: "CloseBracket",
	pos: position{line: 163, col: 1, offset: 3544},
	expr: &seqExpr{
	pos: position{line: 163, col: 16, offset: 3561},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 163, col: 16, offset: 3561},
	name: "_",
},
&litMatcher{
	pos: position{line: 163, col: 18, offset: 3563},
	val: "]",
	ignoreCase: false,
},
	},
},
},
{
	name: "Comma",
	pos: position{line: 164, col: 1, offset: 3567},
	expr: &seqExpr{
	pos: position{line: 164, col: 9, offset: 3577},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 164, col: 9, offset: 3577},
	name: "_",
},
&litMatcher{
	pos: position{line: 164, col: 11, offset: 3579},
	val: ",",
	ignoreCase: false,
},
	},
},
},
{
	name: "OpenParens",
	pos: position{line: 165, col: 1, offset: 3583},
	expr: &seqExpr{
	pos: position{line: 165, col: 14, offset: 3598},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 165, col: 14, offset: 3598},
	name: "_",
},
&litMatcher{
	pos: position{line: 165, col: 16, offset: 3600},
	val: "(",
	ignoreCase: false,
},
	},
},
},
{
	name: "CloseParens",
	pos: position{line: 166, col: 1, offset: 3604},
	expr: &seqExpr{
	pos: position{line: 166, col: 15, offset: 3620},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 166, col: 15, offset: 3620},
	name: "_",
},
&litMatcher{
	pos: position{line: 166, col: 17, offset: 3622},
	val: ")",
	ignoreCase: false,
},
	},
},
},
{
	name: "At",
	pos: position{line: 167, col: 1, offset: 3626},
	expr: &seqExpr{
	pos: position{line: 167, col: 6, offset: 3633},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 167, col: 6, offset: 3633},
	name: "_",
},
&litMatcher{
	pos: position{line: 167, col: 8, offset: 3635},
	val: "@",
	ignoreCase: false,
},
	},
},
},
{
	name: "Colon",
	pos: position{line: 168, col: 1, offset: 3639},
	expr: &seqExpr{
	pos: position{line: 168, col: 9, offset: 3649},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 168, col: 9, offset: 3649},
	name: "_",
},
&litMatcher{
	pos: position{line: 168, col: 11, offset: 3651},
	val: ":",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 168, col: 15, offset: 3655},
	name: "WhitespaceChunk",
},
	},
},
},
{
	name: "Lambda",
	pos: position{line: 170, col: 1, offset: 3672},
	expr: &seqExpr{
	pos: position{line: 170, col: 10, offset: 3683},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 170, col: 10, offset: 3683},
	name: "_",
},
&choiceExpr{
	pos: position{line: 170, col: 13, offset: 3686},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 170, col: 13, offset: 3686},
	val: "\\",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 170, col: 20, offset: 3693},
	val: "λ",
	ignoreCase: false,
},
	},
},
	},
},
},
{
	name: "Forall",
	pos: position{line: 171, col: 1, offset: 3699},
	expr: &seqExpr{
	pos: position{line: 171, col: 10, offset: 3710},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 171, col: 10, offset: 3710},
	name: "_",
},
&choiceExpr{
	pos: position{line: 171, col: 13, offset: 3713},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 171, col: 13, offset: 3713},
	val: "forall",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 171, col: 24, offset: 3724},
	val: "∀",
	ignoreCase: false,
},
	},
},
	},
},
},
{
	name: "Arrow",
	pos: position{line: 172, col: 1, offset: 3731},
	expr: &seqExpr{
	pos: position{line: 172, col: 9, offset: 3741},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 172, col: 9, offset: 3741},
	name: "_",
},
&choiceExpr{
	pos: position{line: 172, col: 12, offset: 3744},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 172, col: 12, offset: 3744},
	val: "->",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 172, col: 19, offset: 3751},
	val: "→",
	ignoreCase: false,
},
	},
},
	},
},
},
{
	name: "Exponent",
	pos: position{line: 174, col: 1, offset: 3759},
	expr: &seqExpr{
	pos: position{line: 174, col: 12, offset: 3772},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 174, col: 12, offset: 3772},
	val: "e",
	ignoreCase: true,
},
&zeroOrOneExpr{
	pos: position{line: 174, col: 17, offset: 3777},
	expr: &charClassMatcher{
	pos: position{line: 174, col: 17, offset: 3777},
	val: "[+-]",
	chars: []rune{'+','-',},
	ignoreCase: false,
	inverted: false,
},
},
&oneOrMoreExpr{
	pos: position{line: 174, col: 23, offset: 3783},
	expr: &charClassMatcher{
	pos: position{line: 174, col: 23, offset: 3783},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
	},
},
},
{
	name: "DoubleLiteralRaw",
	pos: position{line: 176, col: 1, offset: 3791},
	expr: &actionExpr{
	pos: position{line: 176, col: 20, offset: 3812},
	run: (*parser).callonDoubleLiteralRaw1,
	expr: &seqExpr{
	pos: position{line: 176, col: 20, offset: 3812},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 176, col: 20, offset: 3812},
	expr: &charClassMatcher{
	pos: position{line: 176, col: 20, offset: 3812},
	val: "[+-]",
	chars: []rune{'+','-',},
	ignoreCase: false,
	inverted: false,
},
},
&oneOrMoreExpr{
	pos: position{line: 176, col: 26, offset: 3818},
	expr: &charClassMatcher{
	pos: position{line: 176, col: 26, offset: 3818},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
&choiceExpr{
	pos: position{line: 176, col: 35, offset: 3827},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 176, col: 35, offset: 3827},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 176, col: 35, offset: 3827},
	val: ".",
	ignoreCase: false,
},
&oneOrMoreExpr{
	pos: position{line: 176, col: 39, offset: 3831},
	expr: &charClassMatcher{
	pos: position{line: 176, col: 39, offset: 3831},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 176, col: 46, offset: 3838},
	expr: &ruleRefExpr{
	pos: position{line: 176, col: 46, offset: 3838},
	name: "Exponent",
},
},
	},
},
&ruleRefExpr{
	pos: position{line: 176, col: 58, offset: 3850},
	name: "Exponent",
},
	},
},
	},
},
},
},
{
	name: "DoubleLiteral",
	pos: position{line: 184, col: 1, offset: 4010},
	expr: &actionExpr{
	pos: position{line: 184, col: 17, offset: 4028},
	run: (*parser).callonDoubleLiteral1,
	expr: &seqExpr{
	pos: position{line: 184, col: 17, offset: 4028},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 184, col: 17, offset: 4028},
	name: "_",
},
&labeledExpr{
	pos: position{line: 184, col: 19, offset: 4030},
	label: "d",
	expr: &ruleRefExpr{
	pos: position{line: 184, col: 21, offset: 4032},
	name: "DoubleLiteralRaw",
},
},
	},
},
},
},
{
	name: "NaturalLiteralRaw",
	pos: position{line: 186, col: 1, offset: 4068},
	expr: &actionExpr{
	pos: position{line: 186, col: 21, offset: 4090},
	run: (*parser).callonNaturalLiteralRaw1,
	expr: &oneOrMoreExpr{
	pos: position{line: 186, col: 21, offset: 4090},
	expr: &charClassMatcher{
	pos: position{line: 186, col: 21, offset: 4090},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
},
},
{
	name: "NaturalLiteral",
	pos: position{line: 191, col: 1, offset: 4183},
	expr: &actionExpr{
	pos: position{line: 191, col: 18, offset: 4202},
	run: (*parser).callonNaturalLiteral1,
	expr: &seqExpr{
	pos: position{line: 191, col: 18, offset: 4202},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 191, col: 18, offset: 4202},
	name: "_",
},
&labeledExpr{
	pos: position{line: 191, col: 20, offset: 4204},
	label: "n",
	expr: &ruleRefExpr{
	pos: position{line: 191, col: 22, offset: 4206},
	name: "NaturalLiteralRaw",
},
},
	},
},
},
},
{
	name: "IntegerLiteralRaw",
	pos: position{line: 193, col: 1, offset: 4243},
	expr: &actionExpr{
	pos: position{line: 193, col: 21, offset: 4265},
	run: (*parser).callonIntegerLiteralRaw1,
	expr: &seqExpr{
	pos: position{line: 193, col: 21, offset: 4265},
	exprs: []interface{}{
&charClassMatcher{
	pos: position{line: 193, col: 21, offset: 4265},
	val: "[+-]",
	chars: []rune{'+','-',},
	ignoreCase: false,
	inverted: false,
},
&oneOrMoreExpr{
	pos: position{line: 193, col: 25, offset: 4269},
	expr: &charClassMatcher{
	pos: position{line: 193, col: 25, offset: 4269},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
	},
},
},
},
{
	name: "IntegerLiteral",
	pos: position{line: 201, col: 1, offset: 4417},
	expr: &actionExpr{
	pos: position{line: 201, col: 18, offset: 4436},
	run: (*parser).callonIntegerLiteral1,
	expr: &seqExpr{
	pos: position{line: 201, col: 18, offset: 4436},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 201, col: 18, offset: 4436},
	name: "_",
},
&labeledExpr{
	pos: position{line: 201, col: 20, offset: 4438},
	label: "i",
	expr: &ruleRefExpr{
	pos: position{line: 201, col: 22, offset: 4440},
	name: "IntegerLiteralRaw",
},
},
	},
},
},
},
{
	name: "DeBruijn",
	pos: position{line: 203, col: 1, offset: 4477},
	expr: &actionExpr{
	pos: position{line: 203, col: 12, offset: 4490},
	run: (*parser).callonDeBruijn1,
	expr: &seqExpr{
	pos: position{line: 203, col: 12, offset: 4490},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 203, col: 12, offset: 4490},
	name: "At",
},
&labeledExpr{
	pos: position{line: 203, col: 15, offset: 4493},
	label: "index",
	expr: &ruleRefExpr{
	pos: position{line: 203, col: 21, offset: 4499},
	name: "NaturalLiteral",
},
},
	},
},
},
},
{
	name: "Identifier",
	pos: position{line: 205, col: 1, offset: 4559},
	expr: &actionExpr{
	pos: position{line: 205, col: 14, offset: 4574},
	run: (*parser).callonIdentifier1,
	expr: &seqExpr{
	pos: position{line: 205, col: 14, offset: 4574},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 205, col: 14, offset: 4574},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 205, col: 19, offset: 4579},
	name: "Label",
},
},
&labeledExpr{
	pos: position{line: 205, col: 25, offset: 4585},
	label: "index",
	expr: &zeroOrOneExpr{
	pos: position{line: 205, col: 31, offset: 4591},
	expr: &ruleRefExpr{
	pos: position{line: 205, col: 31, offset: 4591},
	name: "DeBruijn",
},
},
},
	},
},
},
},
{
	name: "IdentifierReservedPrefix",
	pos: position{line: 213, col: 1, offset: 4762},
	expr: &actionExpr{
	pos: position{line: 214, col: 10, offset: 4800},
	run: (*parser).callonIdentifierReservedPrefix1,
	expr: &seqExpr{
	pos: position{line: 214, col: 10, offset: 4800},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 214, col: 10, offset: 4800},
	name: "_",
},
&labeledExpr{
	pos: position{line: 214, col: 12, offset: 4802},
	label: "name",
	expr: &actionExpr{
	pos: position{line: 214, col: 18, offset: 4808},
	run: (*parser).callonIdentifierReservedPrefix5,
	expr: &seqExpr{
	pos: position{line: 214, col: 18, offset: 4808},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 214, col: 18, offset: 4808},
	name: "ReservedRaw",
},
&oneOrMoreExpr{
	pos: position{line: 214, col: 30, offset: 4820},
	expr: &charClassMatcher{
	pos: position{line: 214, col: 30, offset: 4820},
	val: "[A-Za-z0-9/_-]",
	chars: []rune{'/','_','-',},
	ranges: []rune{'A','Z','a','z','0','9',},
	ignoreCase: false,
	inverted: false,
},
},
	},
},
},
},
&labeledExpr{
	pos: position{line: 215, col: 10, offset: 4876},
	label: "index",
	expr: &zeroOrOneExpr{
	pos: position{line: 215, col: 16, offset: 4882},
	expr: &ruleRefExpr{
	pos: position{line: 215, col: 16, offset: 4882},
	name: "DeBruijn",
},
},
},
	},
},
},
},
{
	name: "LetBinding",
	pos: position{line: 227, col: 1, offset: 5239},
	expr: &actionExpr{
	pos: position{line: 227, col: 14, offset: 5254},
	run: (*parser).callonLetBinding1,
	expr: &seqExpr{
	pos: position{line: 227, col: 14, offset: 5254},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 227, col: 14, offset: 5254},
	name: "Let",
},
&labeledExpr{
	pos: position{line: 227, col: 18, offset: 5258},
	label: "label",
	expr: &ruleRefExpr{
	pos: position{line: 227, col: 24, offset: 5264},
	name: "Label",
},
},
&labeledExpr{
	pos: position{line: 227, col: 30, offset: 5270},
	label: "a",
	expr: &zeroOrOneExpr{
	pos: position{line: 227, col: 32, offset: 5272},
	expr: &ruleRefExpr{
	pos: position{line: 227, col: 32, offset: 5272},
	name: "Annotation",
},
},
},
&ruleRefExpr{
	pos: position{line: 227, col: 44, offset: 5284},
	name: "Equal",
},
&labeledExpr{
	pos: position{line: 227, col: 50, offset: 5290},
	label: "v",
	expr: &ruleRefExpr{
	pos: position{line: 227, col: 52, offset: 5292},
	name: "Expression",
},
},
	},
},
},
},
{
	name: "Expression",
	pos: position{line: 242, col: 1, offset: 5611},
	expr: &choiceExpr{
	pos: position{line: 243, col: 7, offset: 5632},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 243, col: 7, offset: 5632},
	run: (*parser).callonExpression2,
	expr: &seqExpr{
	pos: position{line: 243, col: 7, offset: 5632},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 243, col: 7, offset: 5632},
	name: "Lambda",
},
&ruleRefExpr{
	pos: position{line: 243, col: 14, offset: 5639},
	name: "OpenParens",
},
&labeledExpr{
	pos: position{line: 243, col: 25, offset: 5650},
	label: "label",
	expr: &ruleRefExpr{
	pos: position{line: 243, col: 31, offset: 5656},
	name: "Label",
},
},
&ruleRefExpr{
	pos: position{line: 243, col: 37, offset: 5662},
	name: "Colon",
},
&labeledExpr{
	pos: position{line: 243, col: 43, offset: 5668},
	label: "t",
	expr: &ruleRefExpr{
	pos: position{line: 243, col: 45, offset: 5670},
	name: "Expression",
},
},
&ruleRefExpr{
	pos: position{line: 243, col: 56, offset: 5681},
	name: "CloseParens",
},
&ruleRefExpr{
	pos: position{line: 243, col: 68, offset: 5693},
	name: "Arrow",
},
&labeledExpr{
	pos: position{line: 243, col: 74, offset: 5699},
	label: "body",
	expr: &ruleRefExpr{
	pos: position{line: 243, col: 79, offset: 5704},
	name: "Expression",
},
},
	},
},
},
&actionExpr{
	pos: position{line: 246, col: 7, offset: 5833},
	run: (*parser).callonExpression15,
	expr: &seqExpr{
	pos: position{line: 246, col: 7, offset: 5833},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 246, col: 7, offset: 5833},
	name: "If",
},
&labeledExpr{
	pos: position{line: 246, col: 10, offset: 5836},
	label: "cond",
	expr: &ruleRefExpr{
	pos: position{line: 246, col: 15, offset: 5841},
	name: "Expression",
},
},
&ruleRefExpr{
	pos: position{line: 246, col: 26, offset: 5852},
	name: "Then",
},
&labeledExpr{
	pos: position{line: 246, col: 31, offset: 5857},
	label: "t",
	expr: &ruleRefExpr{
	pos: position{line: 246, col: 33, offset: 5859},
	name: "Expression",
},
},
&ruleRefExpr{
	pos: position{line: 246, col: 44, offset: 5870},
	name: "Else",
},
&labeledExpr{
	pos: position{line: 246, col: 49, offset: 5875},
	label: "f",
	expr: &ruleRefExpr{
	pos: position{line: 246, col: 51, offset: 5877},
	name: "Expression",
},
},
	},
},
},
&actionExpr{
	pos: position{line: 249, col: 7, offset: 5979},
	run: (*parser).callonExpression26,
	expr: &seqExpr{
	pos: position{line: 249, col: 7, offset: 5979},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 249, col: 7, offset: 5979},
	label: "bindings",
	expr: &zeroOrMoreExpr{
	pos: position{line: 249, col: 16, offset: 5988},
	expr: &ruleRefExpr{
	pos: position{line: 249, col: 16, offset: 5988},
	name: "LetBinding",
},
},
},
&ruleRefExpr{
	pos: position{line: 249, col: 28, offset: 6000},
	name: "In",
},
&labeledExpr{
	pos: position{line: 249, col: 31, offset: 6003},
	label: "b",
	expr: &ruleRefExpr{
	pos: position{line: 249, col: 33, offset: 6005},
	name: "Expression",
},
},
	},
},
},
&actionExpr{
	pos: position{line: 256, col: 7, offset: 6261},
	run: (*parser).callonExpression34,
	expr: &seqExpr{
	pos: position{line: 256, col: 7, offset: 6261},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 256, col: 7, offset: 6261},
	name: "Forall",
},
&ruleRefExpr{
	pos: position{line: 256, col: 14, offset: 6268},
	name: "OpenParens",
},
&labeledExpr{
	pos: position{line: 256, col: 25, offset: 6279},
	label: "label",
	expr: &ruleRefExpr{
	pos: position{line: 256, col: 31, offset: 6285},
	name: "Label",
},
},
&ruleRefExpr{
	pos: position{line: 256, col: 37, offset: 6291},
	name: "Colon",
},
&labeledExpr{
	pos: position{line: 256, col: 43, offset: 6297},
	label: "t",
	expr: &ruleRefExpr{
	pos: position{line: 256, col: 45, offset: 6299},
	name: "Expression",
},
},
&ruleRefExpr{
	pos: position{line: 256, col: 56, offset: 6310},
	name: "CloseParens",
},
&ruleRefExpr{
	pos: position{line: 256, col: 68, offset: 6322},
	name: "Arrow",
},
&labeledExpr{
	pos: position{line: 256, col: 74, offset: 6328},
	label: "body",
	expr: &ruleRefExpr{
	pos: position{line: 256, col: 79, offset: 6333},
	name: "Expression",
},
},
	},
},
},
&actionExpr{
	pos: position{line: 259, col: 7, offset: 6454},
	run: (*parser).callonExpression47,
	expr: &seqExpr{
	pos: position{line: 259, col: 7, offset: 6454},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 259, col: 7, offset: 6454},
	label: "o",
	expr: &ruleRefExpr{
	pos: position{line: 259, col: 9, offset: 6456},
	name: "OperatorExpression",
},
},
&ruleRefExpr{
	pos: position{line: 259, col: 28, offset: 6475},
	name: "Arrow",
},
&labeledExpr{
	pos: position{line: 259, col: 34, offset: 6481},
	label: "e",
	expr: &ruleRefExpr{
	pos: position{line: 259, col: 36, offset: 6483},
	name: "Expression",
},
},
	},
},
},
&ruleRefExpr{
	pos: position{line: 260, col: 7, offset: 6555},
	name: "AnnotatedExpression",
},
	},
},
},
{
	name: "Annotation",
	pos: position{line: 262, col: 1, offset: 6576},
	expr: &actionExpr{
	pos: position{line: 262, col: 14, offset: 6591},
	run: (*parser).callonAnnotation1,
	expr: &seqExpr{
	pos: position{line: 262, col: 14, offset: 6591},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 262, col: 14, offset: 6591},
	name: "Colon",
},
&labeledExpr{
	pos: position{line: 262, col: 20, offset: 6597},
	label: "a",
	expr: &ruleRefExpr{
	pos: position{line: 262, col: 22, offset: 6599},
	name: "Expression",
},
},
	},
},
},
},
{
	name: "AnnotatedExpression",
	pos: position{line: 264, col: 1, offset: 6629},
	expr: &choiceExpr{
	pos: position{line: 265, col: 5, offset: 6657},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 265, col: 5, offset: 6657},
	name: "EmptyList",
},
&actionExpr{
	pos: position{line: 266, col: 5, offset: 6671},
	run: (*parser).callonAnnotatedExpression3,
	expr: &seqExpr{
	pos: position{line: 266, col: 5, offset: 6671},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 266, col: 5, offset: 6671},
	label: "e",
	expr: &ruleRefExpr{
	pos: position{line: 266, col: 7, offset: 6673},
	name: "OperatorExpression",
},
},
&labeledExpr{
	pos: position{line: 266, col: 26, offset: 6692},
	label: "a",
	expr: &zeroOrOneExpr{
	pos: position{line: 266, col: 28, offset: 6694},
	expr: &ruleRefExpr{
	pos: position{line: 266, col: 28, offset: 6694},
	name: "Annotation",
},
},
},
	},
},
},
	},
},
},
{
	name: "EmptyList",
	pos: position{line: 271, col: 1, offset: 6811},
	expr: &actionExpr{
	pos: position{line: 271, col: 13, offset: 6825},
	run: (*parser).callonEmptyList1,
	expr: &seqExpr{
	pos: position{line: 271, col: 13, offset: 6825},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 271, col: 13, offset: 6825},
	name: "OpenBracket",
},
&ruleRefExpr{
	pos: position{line: 271, col: 25, offset: 6837},
	name: "CloseBracket",
},
&ruleRefExpr{
	pos: position{line: 271, col: 38, offset: 6850},
	name: "Colon",
},
&ruleRefExpr{
	pos: position{line: 271, col: 44, offset: 6856},
	name: "List",
},
&labeledExpr{
	pos: position{line: 271, col: 49, offset: 6861},
	label: "a",
	expr: &ruleRefExpr{
	pos: position{line: 271, col: 51, offset: 6863},
	name: "ImportExpression",
},
},
	},
},
},
},
{
	name: "OperatorExpression",
	pos: position{line: 275, col: 1, offset: 6934},
	expr: &ruleRefExpr{
	pos: position{line: 275, col: 22, offset: 6957},
	name: "ImportAltExpression",
},
},
{
	name: "ImportAltExpression",
	pos: position{line: 277, col: 1, offset: 6978},
	expr: &ruleRefExpr{
	pos: position{line: 277, col: 23, offset: 7002},
	name: "PlusExpression",
},
},
{
	name: "MorePlus",
	pos: position{line: 279, col: 1, offset: 7018},
	expr: &actionExpr{
	pos: position{line: 279, col: 12, offset: 7031},
	run: (*parser).callonMorePlus1,
	expr: &seqExpr{
	pos: position{line: 279, col: 12, offset: 7031},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 279, col: 12, offset: 7031},
	name: "Plus",
},
&labeledExpr{
	pos: position{line: 279, col: 17, offset: 7036},
	label: "e",
	expr: &ruleRefExpr{
	pos: position{line: 279, col: 19, offset: 7038},
	name: "TimesExpression",
},
},
	},
},
},
},
{
	name: "PlusExpression",
	pos: position{line: 280, col: 1, offset: 7072},
	expr: &actionExpr{
	pos: position{line: 281, col: 7, offset: 7097},
	run: (*parser).callonPlusExpression1,
	expr: &seqExpr{
	pos: position{line: 281, col: 7, offset: 7097},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 281, col: 7, offset: 7097},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 281, col: 13, offset: 7103},
	name: "TimesExpression",
},
},
&labeledExpr{
	pos: position{line: 281, col: 29, offset: 7119},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 281, col: 34, offset: 7124},
	expr: &ruleRefExpr{
	pos: position{line: 281, col: 34, offset: 7124},
	name: "MorePlus",
},
},
},
	},
},
},
},
{
	name: "MoreTimes",
	pos: position{line: 290, col: 1, offset: 7364},
	expr: &actionExpr{
	pos: position{line: 290, col: 13, offset: 7378},
	run: (*parser).callonMoreTimes1,
	expr: &seqExpr{
	pos: position{line: 290, col: 13, offset: 7378},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 290, col: 13, offset: 7378},
	name: "Times",
},
&labeledExpr{
	pos: position{line: 290, col: 19, offset: 7384},
	label: "e",
	expr: &ruleRefExpr{
	pos: position{line: 290, col: 21, offset: 7386},
	name: "ApplicationExpression",
},
},
	},
},
},
},
{
	name: "TimesExpression",
	pos: position{line: 291, col: 1, offset: 7426},
	expr: &actionExpr{
	pos: position{line: 292, col: 7, offset: 7452},
	run: (*parser).callonTimesExpression1,
	expr: &seqExpr{
	pos: position{line: 292, col: 7, offset: 7452},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 292, col: 7, offset: 7452},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 292, col: 13, offset: 7458},
	name: "ApplicationExpression",
},
},
&labeledExpr{
	pos: position{line: 292, col: 35, offset: 7480},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 292, col: 40, offset: 7485},
	expr: &ruleRefExpr{
	pos: position{line: 292, col: 40, offset: 7485},
	name: "MoreTimes",
},
},
},
	},
},
},
},
{
	name: "MoreApp",
	pos: position{line: 301, col: 1, offset: 7727},
	expr: &actionExpr{
	pos: position{line: 301, col: 11, offset: 7739},
	run: (*parser).callonMoreApp1,
	expr: &seqExpr{
	pos: position{line: 301, col: 11, offset: 7739},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 301, col: 11, offset: 7739},
	name: "WhitespaceChunk",
},
&labeledExpr{
	pos: position{line: 301, col: 27, offset: 7755},
	label: "e",
	expr: &ruleRefExpr{
	pos: position{line: 301, col: 29, offset: 7757},
	name: "ImportExpression",
},
},
	},
},
},
},
{
	name: "ApplicationExpression",
	pos: position{line: 302, col: 1, offset: 7792},
	expr: &actionExpr{
	pos: position{line: 302, col: 25, offset: 7818},
	run: (*parser).callonApplicationExpression1,
	expr: &seqExpr{
	pos: position{line: 302, col: 25, offset: 7818},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 302, col: 25, offset: 7818},
	label: "f",
	expr: &ruleRefExpr{
	pos: position{line: 302, col: 27, offset: 7820},
	name: "ImportExpression",
},
},
&labeledExpr{
	pos: position{line: 302, col: 44, offset: 7837},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 302, col: 49, offset: 7842},
	expr: &ruleRefExpr{
	pos: position{line: 302, col: 49, offset: 7842},
	name: "MoreApp",
},
},
},
	},
},
},
},
{
	name: "ImportExpression",
	pos: position{line: 311, col: 1, offset: 8075},
	expr: &ruleRefExpr{
	pos: position{line: 311, col: 20, offset: 8096},
	name: "SelectorExpression",
},
},
{
	name: "SelectorExpression",
	pos: position{line: 313, col: 1, offset: 8116},
	expr: &actionExpr{
	pos: position{line: 313, col: 22, offset: 8139},
	run: (*parser).callonSelectorExpression1,
	expr: &seqExpr{
	pos: position{line: 313, col: 22, offset: 8139},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 313, col: 22, offset: 8139},
	label: "e",
	expr: &ruleRefExpr{
	pos: position{line: 313, col: 24, offset: 8141},
	name: "PrimitiveExpression",
},
},
&labeledExpr{
	pos: position{line: 313, col: 44, offset: 8161},
	label: "ls",
	expr: &zeroOrMoreExpr{
	pos: position{line: 313, col: 47, offset: 8164},
	expr: &seqExpr{
	pos: position{line: 313, col: 48, offset: 8165},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 313, col: 48, offset: 8165},
	name: "Dot",
},
&ruleRefExpr{
	pos: position{line: 313, col: 52, offset: 8169},
	name: "Label",
},
	},
},
},
},
	},
},
},
},
{
	name: "PrimitiveExpression",
	pos: position{line: 323, col: 1, offset: 8407},
	expr: &choiceExpr{
	pos: position{line: 324, col: 7, offset: 8437},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 324, col: 7, offset: 8437},
	name: "DoubleLiteral",
},
&ruleRefExpr{
	pos: position{line: 325, col: 7, offset: 8457},
	name: "NaturalLiteral",
},
&ruleRefExpr{
	pos: position{line: 326, col: 7, offset: 8478},
	name: "IntegerLiteral",
},
&actionExpr{
	pos: position{line: 327, col: 7, offset: 8499},
	run: (*parser).callonPrimitiveExpression5,
	expr: &litMatcher{
	pos: position{line: 327, col: 7, offset: 8499},
	val: "-Infinity",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 328, col: 7, offset: 8561},
	name: "TextLiteral",
},
&actionExpr{
	pos: position{line: 329, col: 7, offset: 8579},
	run: (*parser).callonPrimitiveExpression8,
	expr: &seqExpr{
	pos: position{line: 329, col: 7, offset: 8579},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 329, col: 7, offset: 8579},
	name: "OpenBrace",
},
&labeledExpr{
	pos: position{line: 329, col: 17, offset: 8589},
	label: "r",
	expr: &ruleRefExpr{
	pos: position{line: 329, col: 19, offset: 8591},
	name: "RecordTypeOrLiteral",
},
},
&ruleRefExpr{
	pos: position{line: 329, col: 39, offset: 8611},
	name: "CloseBrace",
},
	},
},
},
&ruleRefExpr{
	pos: position{line: 330, col: 7, offset: 8646},
	name: "NonEmptyListLiteral",
},
&ruleRefExpr{
	pos: position{line: 331, col: 7, offset: 8672},
	name: "IdentifierReservedPrefix",
},
&ruleRefExpr{
	pos: position{line: 332, col: 7, offset: 8703},
	name: "Reserved",
},
&ruleRefExpr{
	pos: position{line: 333, col: 7, offset: 8718},
	name: "Identifier",
},
&actionExpr{
	pos: position{line: 334, col: 7, offset: 8735},
	run: (*parser).callonPrimitiveExpression18,
	expr: &seqExpr{
	pos: position{line: 334, col: 7, offset: 8735},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 334, col: 7, offset: 8735},
	name: "OpenParens",
},
&labeledExpr{
	pos: position{line: 334, col: 18, offset: 8746},
	label: "e",
	expr: &ruleRefExpr{
	pos: position{line: 334, col: 20, offset: 8748},
	name: "Expression",
},
},
&ruleRefExpr{
	pos: position{line: 334, col: 31, offset: 8759},
	name: "CloseParens",
},
	},
},
},
	},
},
},
{
	name: "RecordTypeOrLiteral",
	pos: position{line: 336, col: 1, offset: 8790},
	expr: &choiceExpr{
	pos: position{line: 337, col: 7, offset: 8820},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 337, col: 7, offset: 8820},
	run: (*parser).callonRecordTypeOrLiteral2,
	expr: &ruleRefExpr{
	pos: position{line: 337, col: 7, offset: 8820},
	name: "Equal",
},
},
&ruleRefExpr{
	pos: position{line: 338, col: 7, offset: 8885},
	name: "NonEmptyRecordType",
},
&ruleRefExpr{
	pos: position{line: 339, col: 7, offset: 8910},
	name: "NonEmptyRecordLiteral",
},
&actionExpr{
	pos: position{line: 340, col: 7, offset: 8938},
	run: (*parser).callonRecordTypeOrLiteral6,
	expr: &litMatcher{
	pos: position{line: 340, col: 7, offset: 8938},
	val: "",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "RecordTypeField",
	pos: position{line: 342, col: 1, offset: 8992},
	expr: &actionExpr{
	pos: position{line: 342, col: 19, offset: 9012},
	run: (*parser).callonRecordTypeField1,
	expr: &seqExpr{
	pos: position{line: 342, col: 19, offset: 9012},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 342, col: 19, offset: 9012},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 342, col: 24, offset: 9017},
	name: "Label",
},
},
&ruleRefExpr{
	pos: position{line: 342, col: 30, offset: 9023},
	name: "Colon",
},
&labeledExpr{
	pos: position{line: 342, col: 36, offset: 9029},
	label: "expr",
	expr: &ruleRefExpr{
	pos: position{line: 342, col: 41, offset: 9034},
	name: "Expression",
},
},
	},
},
},
},
{
	name: "MoreRecordType",
	pos: position{line: 345, col: 1, offset: 9091},
	expr: &actionExpr{
	pos: position{line: 345, col: 18, offset: 9110},
	run: (*parser).callonMoreRecordType1,
	expr: &seqExpr{
	pos: position{line: 345, col: 18, offset: 9110},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 345, col: 18, offset: 9110},
	name: "Comma",
},
&labeledExpr{
	pos: position{line: 345, col: 24, offset: 9116},
	label: "f",
	expr: &ruleRefExpr{
	pos: position{line: 345, col: 26, offset: 9118},
	name: "RecordTypeField",
},
},
	},
},
},
},
{
	name: "NonEmptyRecordType",
	pos: position{line: 346, col: 1, offset: 9150},
	expr: &actionExpr{
	pos: position{line: 347, col: 7, offset: 9179},
	run: (*parser).callonNonEmptyRecordType1,
	expr: &seqExpr{
	pos: position{line: 347, col: 7, offset: 9179},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 347, col: 7, offset: 9179},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 347, col: 13, offset: 9185},
	name: "RecordTypeField",
},
},
&labeledExpr{
	pos: position{line: 347, col: 29, offset: 9201},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 347, col: 34, offset: 9206},
	expr: &ruleRefExpr{
	pos: position{line: 347, col: 34, offset: 9206},
	name: "MoreRecordType",
},
},
},
	},
},
},
},
{
	name: "RecordLiteralField",
	pos: position{line: 357, col: 1, offset: 9618},
	expr: &actionExpr{
	pos: position{line: 357, col: 22, offset: 9641},
	run: (*parser).callonRecordLiteralField1,
	expr: &seqExpr{
	pos: position{line: 357, col: 22, offset: 9641},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 357, col: 22, offset: 9641},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 357, col: 27, offset: 9646},
	name: "Label",
},
},
&ruleRefExpr{
	pos: position{line: 357, col: 33, offset: 9652},
	name: "Equal",
},
&labeledExpr{
	pos: position{line: 357, col: 39, offset: 9658},
	label: "expr",
	expr: &ruleRefExpr{
	pos: position{line: 357, col: 44, offset: 9663},
	name: "Expression",
},
},
	},
},
},
},
{
	name: "MoreRecordLiteral",
	pos: position{line: 360, col: 1, offset: 9720},
	expr: &actionExpr{
	pos: position{line: 360, col: 21, offset: 9742},
	run: (*parser).callonMoreRecordLiteral1,
	expr: &seqExpr{
	pos: position{line: 360, col: 21, offset: 9742},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 360, col: 21, offset: 9742},
	name: "Comma",
},
&labeledExpr{
	pos: position{line: 360, col: 27, offset: 9748},
	label: "f",
	expr: &ruleRefExpr{
	pos: position{line: 360, col: 29, offset: 9750},
	name: "RecordLiteralField",
},
},
	},
},
},
},
{
	name: "NonEmptyRecordLiteral",
	pos: position{line: 361, col: 1, offset: 9785},
	expr: &actionExpr{
	pos: position{line: 362, col: 7, offset: 9817},
	run: (*parser).callonNonEmptyRecordLiteral1,
	expr: &seqExpr{
	pos: position{line: 362, col: 7, offset: 9817},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 362, col: 7, offset: 9817},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 362, col: 13, offset: 9823},
	name: "RecordLiteralField",
},
},
&labeledExpr{
	pos: position{line: 362, col: 32, offset: 9842},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 362, col: 37, offset: 9847},
	expr: &ruleRefExpr{
	pos: position{line: 362, col: 37, offset: 9847},
	name: "MoreRecordLiteral",
},
},
},
	},
},
},
},
{
	name: "MoreList",
	pos: position{line: 372, col: 1, offset: 10265},
	expr: &actionExpr{
	pos: position{line: 372, col: 12, offset: 10278},
	run: (*parser).callonMoreList1,
	expr: &seqExpr{
	pos: position{line: 372, col: 12, offset: 10278},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 372, col: 12, offset: 10278},
	name: "Comma",
},
&labeledExpr{
	pos: position{line: 372, col: 18, offset: 10284},
	label: "e",
	expr: &ruleRefExpr{
	pos: position{line: 372, col: 20, offset: 10286},
	name: "Expression",
},
},
	},
},
},
},
{
	name: "NonEmptyListLiteral",
	pos: position{line: 374, col: 1, offset: 10314},
	expr: &actionExpr{
	pos: position{line: 375, col: 7, offset: 10344},
	run: (*parser).callonNonEmptyListLiteral1,
	expr: &seqExpr{
	pos: position{line: 375, col: 7, offset: 10344},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 375, col: 7, offset: 10344},
	name: "OpenBracket",
},
&labeledExpr{
	pos: position{line: 375, col: 19, offset: 10356},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 375, col: 25, offset: 10362},
	name: "Expression",
},
},
&labeledExpr{
	pos: position{line: 375, col: 36, offset: 10373},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 375, col: 41, offset: 10378},
	expr: &ruleRefExpr{
	pos: position{line: 375, col: 41, offset: 10378},
	name: "MoreList",
},
},
},
&ruleRefExpr{
	pos: position{line: 375, col: 51, offset: 10388},
	name: "CloseBracket",
},
	},
},
},
},
{
	name: "EOF",
	pos: position{line: 385, col: 1, offset: 10689},
	expr: &notExpr{
	pos: position{line: 385, col: 7, offset: 10697},
	expr: &anyMatcher{
	line: 385, col: 8, offset: 10698,
},
},
},
	},
}
func (c *current) onDhallFile1(e interface{}) (interface{}, error) {
 return e, nil 
}

func (p *parser) callonDhallFile1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDhallFile1(stack["e"])
}

func (c *current) onCompleteExpression1(e interface{}) (interface{}, error) {
 return e, nil 
}

func (p *parser) callonCompleteExpression1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCompleteExpression1(stack["e"])
}

func (c *current) onLineComment5() (interface{}, error) {
 return string(c.text), nil
}

func (p *parser) callonLineComment5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLineComment5()
}

func (c *current) onLineComment1(content interface{}) (interface{}, error) {
 return content, nil 
}

func (p *parser) callonLineComment1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLineComment1(stack["content"])
}

func (c *current) onSimpleLabel1() (interface{}, error) {
 return string(c.text), nil 
}

func (p *parser) callonSimpleLabel1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSimpleLabel1()
}

func (c *current) onLabel1(label interface{}) (interface{}, error) {
 return label, nil 
}

func (p *parser) callonLabel1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabel1(stack["label"])
}

func (c *current) onEscapedChar1() (interface{}, error) {
    switch c.text[1] {
    case 'b':
        return []byte("\b"), nil
    case 'f':
        return []byte("\f"), nil
    case 'n':
        return []byte("\n"), nil
    case 'r':
        return []byte("\r"), nil
    case 't':
        return []byte("\t"), nil
    case 'u':
        i, err := strconv.ParseInt(string(c.text[2:]), 16, 32)
        return []byte(string([]rune{rune(i)})), err
    }
    return c.text[1:2], nil
}

func (p *parser) callonEscapedChar1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedChar1()
}

func (c *current) onDoubleQuoteChunk2(e interface{}) (interface{}, error) {
 return e, nil 
}

func (p *parser) callonDoubleQuoteChunk2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDoubleQuoteChunk2(stack["e"])
}

func (c *current) onDoubleQuoteLiteral1(chunks interface{}) (interface{}, error) {
    var str strings.Builder
    var outChunks ast.Chunks
    for _, chunk := range chunks.([]interface{}) {
        switch e := chunk.(type) {
        case []byte:
                str.Write(e)
        case ast.Expr:
                outChunks = append(outChunks, ast.Chunk{str.String(), e})
                str.Reset()
        default:
                return nil, errors.New("can't happen")
        }
    }
    return ast.TextLit{Chunks: outChunks, Suffix: str.String()}, nil
}

func (p *parser) callonDoubleQuoteLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDoubleQuoteLiteral1(stack["chunks"])
}

func (c *current) onTextLiteral1(t interface{}) (interface{}, error) {
 return t, nil 
}

func (p *parser) callonTextLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTextLiteral1(stack["t"])
}

func (c *current) onReservedRaw2() (interface{}, error) {
 return ast.Bool, nil 
}

func (p *parser) callonReservedRaw2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReservedRaw2()
}

func (c *current) onReservedRaw6() (interface{}, error) {
 return ast.Natural, nil 
}

func (p *parser) callonReservedRaw6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReservedRaw6()
}

func (c *current) onReservedRaw8() (interface{}, error) {
 return ast.Integer, nil 
}

func (p *parser) callonReservedRaw8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReservedRaw8()
}

func (c *current) onReservedRaw10() (interface{}, error) {
 return ast.Double, nil 
}

func (p *parser) callonReservedRaw10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReservedRaw10()
}

func (c *current) onReservedRaw12() (interface{}, error) {
 return ast.Text, nil 
}

func (p *parser) callonReservedRaw12() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReservedRaw12()
}

func (c *current) onReservedRaw14() (interface{}, error) {
 return ast.List, nil 
}

func (p *parser) callonReservedRaw14() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReservedRaw14()
}

func (c *current) onReservedRaw16() (interface{}, error) {
 return ast.True, nil 
}

func (p *parser) callonReservedRaw16() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReservedRaw16()
}

func (c *current) onReservedRaw18() (interface{}, error) {
 return ast.False, nil 
}

func (p *parser) callonReservedRaw18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReservedRaw18()
}

func (c *current) onReservedRaw20() (interface{}, error) {
 return ast.DoubleLit(math.NaN()), nil 
}

func (p *parser) callonReservedRaw20() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReservedRaw20()
}

func (c *current) onReservedRaw22() (interface{}, error) {
 return ast.DoubleLit(math.Inf(1)), nil 
}

func (p *parser) callonReservedRaw22() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReservedRaw22()
}

func (c *current) onReservedRaw24() (interface{}, error) {
 return ast.Type, nil 
}

func (p *parser) callonReservedRaw24() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReservedRaw24()
}

func (c *current) onReservedRaw26() (interface{}, error) {
 return ast.Kind, nil 
}

func (p *parser) callonReservedRaw26() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReservedRaw26()
}

func (c *current) onReservedRaw28() (interface{}, error) {
 return ast.Sort, nil 
}

func (p *parser) callonReservedRaw28() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReservedRaw28()
}

func (c *current) onReserved1(r interface{}) (interface{}, error) {
 return r, nil 
}

func (p *parser) callonReserved1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReserved1(stack["r"])
}

func (c *current) onDoubleLiteralRaw1() (interface{}, error) {
      d, err := strconv.ParseFloat(string(c.text), 64)
      if err != nil {
         return nil, err
      }
      return ast.DoubleLit(d), nil
}

func (p *parser) callonDoubleLiteralRaw1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDoubleLiteralRaw1()
}

func (c *current) onDoubleLiteral1(d interface{}) (interface{}, error) {
 return d, nil 
}

func (p *parser) callonDoubleLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDoubleLiteral1(stack["d"])
}

func (c *current) onNaturalLiteralRaw1() (interface{}, error) {
      i, err := strconv.Atoi(string(c.text))
      return ast.NaturalLit(i), err
}

func (p *parser) callonNaturalLiteralRaw1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNaturalLiteralRaw1()
}

func (c *current) onNaturalLiteral1(n interface{}) (interface{}, error) {
 return n, nil 
}

func (p *parser) callonNaturalLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNaturalLiteral1(stack["n"])
}

func (c *current) onIntegerLiteralRaw1() (interface{}, error) {
      i, err := strconv.Atoi(string(c.text))
      if err != nil {
         return nil, err
      }
      return ast.IntegerLit(i), nil
}

func (p *parser) callonIntegerLiteralRaw1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIntegerLiteralRaw1()
}

func (c *current) onIntegerLiteral1(i interface{}) (interface{}, error) {
 return i, nil 
}

func (p *parser) callonIntegerLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIntegerLiteral1(stack["i"])
}

func (c *current) onDeBruijn1(index interface{}) (interface{}, error) {
 return int(index.(ast.NaturalLit)), nil 
}

func (p *parser) callonDeBruijn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDeBruijn1(stack["index"])
}

func (c *current) onIdentifier1(name, index interface{}) (interface{}, error) {
    if index != nil {
        return ast.Var{Name:name.(string), Index:index.(int)}, nil
    } else {
        return ast.Var{Name:name.(string)}, nil
    }
}

func (p *parser) callonIdentifier1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIdentifier1(stack["name"], stack["index"])
}

func (c *current) onIdentifierReservedPrefix5() (interface{}, error) {
 return string(c.text),nil 
}

func (p *parser) callonIdentifierReservedPrefix5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIdentifierReservedPrefix5()
}

func (c *current) onIdentifierReservedPrefix1(name, index interface{}) (interface{}, error) {
    if index != nil {
        return ast.Var{Name:name.(string), Index:index.(int)}, nil
    } else {
        return ast.Var{Name:name.(string)}, nil
    }
}

func (p *parser) callonIdentifierReservedPrefix1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIdentifierReservedPrefix1(stack["name"], stack["index"])
}

func (c *current) onLetBinding1(label, a, v interface{}) (interface{}, error) {
    if a != nil {
        return ast.Binding{
            Variable: label.(string),
            Annotation: a.(ast.Expr),
            Value: v.(ast.Expr),
        }, nil
    } else {
        return ast.Binding{
            Variable: label.(string),
            Value: v.(ast.Expr),
        }, nil
    }
}

func (p *parser) callonLetBinding1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLetBinding1(stack["label"], stack["a"], stack["v"])
}

func (c *current) onExpression2(label, t, body interface{}) (interface{}, error) {
          return &ast.LambdaExpr{Label:label.(string), Type:t.(ast.Expr), Body: body.(ast.Expr)}, nil
      
}

func (p *parser) callonExpression2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExpression2(stack["label"], stack["t"], stack["body"])
}

func (c *current) onExpression15(cond, t, f interface{}) (interface{}, error) {
          return ast.BoolIf{cond.(ast.Expr),t.(ast.Expr),f.(ast.Expr)},nil
      
}

func (p *parser) callonExpression15() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExpression15(stack["cond"], stack["t"], stack["f"])
}

func (c *current) onExpression26(bindings, b interface{}) (interface{}, error) {
        bs := make([]ast.Binding, len(bindings.([]interface{})))
        for i, binding := range bindings.([]interface{}) {
            bs[i] = binding.(ast.Binding)
        }
        return ast.MakeLet(b.(ast.Expr), bs...), nil
      
}

func (p *parser) callonExpression26() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExpression26(stack["bindings"], stack["b"])
}

func (c *current) onExpression34(label, t, body interface{}) (interface{}, error) {
          return &ast.Pi{Label:label.(string), Type:t.(ast.Expr), Body: body.(ast.Expr)}, nil
      
}

func (p *parser) callonExpression34() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExpression34(stack["label"], stack["t"], stack["body"])
}

func (c *current) onExpression47(o, e interface{}) (interface{}, error) {
 return &ast.Pi{"_",o.(ast.Expr),e.(ast.Expr)}, nil 
}

func (p *parser) callonExpression47() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExpression47(stack["o"], stack["e"])
}

func (c *current) onAnnotation1(a interface{}) (interface{}, error) {
 return a, nil 
}

func (p *parser) callonAnnotation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAnnotation1(stack["a"])
}

func (c *current) onAnnotatedExpression3(e, a interface{}) (interface{}, error) {
        if a == nil { return e, nil }
        return ast.Annot{e.(ast.Expr), a.(ast.Expr)}, nil
    
}

func (p *parser) callonAnnotatedExpression3() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAnnotatedExpression3(stack["e"], stack["a"])
}

func (c *current) onEmptyList1(a interface{}) (interface{}, error) {
          return ast.EmptyList{a.(ast.Expr)},nil
}

func (p *parser) callonEmptyList1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEmptyList1(stack["a"])
}

func (c *current) onMorePlus1(e interface{}) (interface{}, error) {
 return e, nil 
}

func (p *parser) callonMorePlus1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMorePlus1(stack["e"])
}

func (c *current) onPlusExpression1(first, rest interface{}) (interface{}, error) {
          a := first.(ast.Expr)
          if rest == nil { return a, nil }
          for _, b := range rest.([]interface{}) {
              a = ast.NaturalPlus{L: a, R: b.(ast.Expr)}
          }
          return a, nil
      
}

func (p *parser) callonPlusExpression1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPlusExpression1(stack["first"], stack["rest"])
}

func (c *current) onMoreTimes1(e interface{}) (interface{}, error) {
 return e, nil 
}

func (p *parser) callonMoreTimes1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMoreTimes1(stack["e"])
}

func (c *current) onTimesExpression1(first, rest interface{}) (interface{}, error) {
          a := first.(ast.Expr)
          if rest == nil { return a, nil }
          for _, b := range rest.([]interface{}) {
              a = ast.NaturalTimes{L: a, R: b.(ast.Expr)}
          }
          return a, nil
      
}

func (p *parser) callonTimesExpression1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTimesExpression1(stack["first"], stack["rest"])
}

func (c *current) onMoreApp1(e interface{}) (interface{}, error) {
 return e, nil 
}

func (p *parser) callonMoreApp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMoreApp1(stack["e"])
}

func (c *current) onApplicationExpression1(f, rest interface{}) (interface{}, error) {
          e := f.(ast.Expr)
          if rest == nil { return e, nil }
          for _, arg := range rest.([]interface{}) {
              e = &ast.App{Fn:e, Arg: arg.(ast.Expr)}
          }
          return e,nil
      
}

func (p *parser) callonApplicationExpression1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onApplicationExpression1(stack["f"], stack["rest"])
}

func (c *current) onSelectorExpression1(e, ls interface{}) (interface{}, error) {
    expr := e.(ast.Expr)
    labels := ls.([]interface{})
    for _, labelSelector := range labels {
        label := labelSelector.([]interface{})[1]
        expr = ast.Field{expr, label.(string)}
    }
    return expr, nil
}

func (p *parser) callonSelectorExpression1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSelectorExpression1(stack["e"], stack["ls"])
}

func (c *current) onPrimitiveExpression5() (interface{}, error) {
 return ast.DoubleLit(math.Inf(-1)), nil 
}

func (p *parser) callonPrimitiveExpression5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPrimitiveExpression5()
}

func (c *current) onPrimitiveExpression8(r interface{}) (interface{}, error) {
 return r, nil 
}

func (p *parser) callonPrimitiveExpression8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPrimitiveExpression8(stack["r"])
}

func (c *current) onPrimitiveExpression18(e interface{}) (interface{}, error) {
 return e, nil 
}

func (p *parser) callonPrimitiveExpression18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPrimitiveExpression18(stack["e"])
}

func (c *current) onRecordTypeOrLiteral2() (interface{}, error) {
 return ast.RecordLit(map[string]ast.Expr{}), nil 
}

func (p *parser) callonRecordTypeOrLiteral2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRecordTypeOrLiteral2()
}

func (c *current) onRecordTypeOrLiteral6() (interface{}, error) {
 return ast.Record(map[string]ast.Expr{}), nil 
}

func (p *parser) callonRecordTypeOrLiteral6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRecordTypeOrLiteral6()
}

func (c *current) onRecordTypeField1(name, expr interface{}) (interface{}, error) {
    return []interface{}{name, expr}, nil
}

func (p *parser) callonRecordTypeField1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRecordTypeField1(stack["name"], stack["expr"])
}

func (c *current) onMoreRecordType1(f interface{}) (interface{}, error) {
return f, nil
}

func (p *parser) callonMoreRecordType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMoreRecordType1(stack["f"])
}

func (c *current) onNonEmptyRecordType1(first, rest interface{}) (interface{}, error) {
          fields := rest.([]interface{})
          content := make(map[string]ast.Expr, len(fields)+1)
          content[first.([]interface{})[0].(string)] = first.([]interface{})[1].(ast.Expr)
          for _, field := range(fields) {
              content[field.([]interface{})[0].(string)] = field.([]interface{})[1].(ast.Expr)
          }
          return ast.Record(content), nil
      
}

func (p *parser) callonNonEmptyRecordType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNonEmptyRecordType1(stack["first"], stack["rest"])
}

func (c *current) onRecordLiteralField1(name, expr interface{}) (interface{}, error) {
    return []interface{}{name, expr}, nil
}

func (p *parser) callonRecordLiteralField1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRecordLiteralField1(stack["name"], stack["expr"])
}

func (c *current) onMoreRecordLiteral1(f interface{}) (interface{}, error) {
return f, nil
}

func (p *parser) callonMoreRecordLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMoreRecordLiteral1(stack["f"])
}

func (c *current) onNonEmptyRecordLiteral1(first, rest interface{}) (interface{}, error) {
          fields := rest.([]interface{})
          content := make(map[string]ast.Expr, len(fields)+1)
          content[first.([]interface{})[0].(string)] = first.([]interface{})[1].(ast.Expr)
          for _, field := range(fields) {
              content[field.([]interface{})[0].(string)] = field.([]interface{})[1].(ast.Expr)
          }
          return ast.RecordLit(content), nil
      
}

func (p *parser) callonNonEmptyRecordLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNonEmptyRecordLiteral1(stack["first"], stack["rest"])
}

func (c *current) onMoreList1(e interface{}) (interface{}, error) {
return e, nil
}

func (p *parser) callonMoreList1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMoreList1(stack["e"])
}

func (c *current) onNonEmptyListLiteral1(first, rest interface{}) (interface{}, error) {
          exprs := rest.([]interface{})
          content := make([]ast.Expr, len(exprs)+1)
          content[0] = first.(ast.Expr)
          for i, expr := range(exprs) {
              content[i+1] = expr.(ast.Expr)
          }
          return ast.NonEmptyList(content), nil
      
}

func (p *parser) callonNonEmptyListLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNonEmptyListLiteral1(stack["first"], stack["rest"])
}


var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule          = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch         = errors.New("no match found")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos    position
	expr   interface{}
	run    func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs: new(errList),
		data: b,
		pt: savepoint{position: position{line: 1}},
		recover: true,
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v interface{}
	b bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug bool
	depth  int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules  map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth) + ">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth) + "<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n == 1 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth) + "MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth) + "MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	// can't match EOF
	if cur == utf8.RuneError {
		return nil, false
	}
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}

