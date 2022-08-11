package entity

import (
	"fmt"
	"regexp"
	"strings"
)

type PlaceHolderInvalidArgsError struct {
	placeHolder *PlaceHolder
	args        []string
}

var _ = error(new(PlaceHolderInvalidArgsError))

func (p *PlaceHolderInvalidArgsError) Error() string {
	return fmt.Sprintf("%v is invalid args for \"%s\"", p.args, p.placeHolder.format)
}

func (p *PlaceHolder) invalidArgsError(args []string) error {
	return &PlaceHolderInvalidArgsError{
		placeHolder: p,
		args:        args,
	}
}

type PlaceHolder struct {
	format    string
	numOfArgs int
}

var placeHolderPattern = regexp.MustCompile(`%[0-9]*\$s`)

func NewPlaceHolder(src string) *PlaceHolder {
	format := placeHolderPattern.ReplaceAllString(src, "%s")
	numOfArgs := strings.Count(format, "%s")
	return &PlaceHolder{
		format:    format,
		numOfArgs: numOfArgs,
	}
}

func (p *PlaceHolder) Apply(args []string) (string, error) {
	if len(args) != p.numOfArgs {
		return "", p.invalidArgsError(args)
	}
	params := make([]any, 0, len(args))
	for _, arg := range args {
		params = append(params, arg)
	}
	return fmt.Sprintf(p.format, params...), nil
}

func (p *PlaceHolder) MustApply(args []string) string {
	applied, err := p.Apply(args)
	if err != nil {
		panic(err)
	}
	return applied
}

func (p *PlaceHolder) GetArgs(src string) ([]string, error) {
	ae := newArgExtractor(p, src)
	return ae.ExtractArgs()
}

type ArgExtractNotMatchError struct {
	argExtractor *argExtractor
}

var _ = error(new(ArgExtractNotMatchError))

func (a *ArgExtractNotMatchError) Error() string {
	format := strings.Join(a.argExtractor.formatWords, " ")
	formed := strings.Join(a.argExtractor.formedWords, " ")
	return fmt.Sprintf("\"%s\" is not match \"%s\"", formed, format)
}

func (a *argExtractor) notMatchError() error {
	return &ArgExtractNotMatchError{
		argExtractor: a,
	}
}

type argExtractor struct {
	formatWords      []string
	formedWords      []string
	formatWordsIndex int
	formedWordsIndex int
	result           []string
}

func newArgExtractor(p *PlaceHolder, src string) *argExtractor {
	return &argExtractor{
		formatWords:      strings.Split(p.format, " "),
		formedWords:      strings.Split(src, " "),
		formatWordsIndex: 0,
		formedWordsIndex: 0,
		result:           make([]string, 0, p.numOfArgs),
	}
}

func (a *argExtractor) currentFormedWord() string {
	if len(a.formedWords) <= a.formedWordsIndex {
		return ""
	}
	return a.formedWords[a.formedWordsIndex]
}

func (a *argExtractor) currentFormatWord() string {
	if len(a.formatWords) <= a.formatWordsIndex {
		return ""
	}
	return a.formatWords[a.formatWordsIndex]
}

func (a *argExtractor) getArg(arg string) (string, error) {
	a.formedWordsIndex++
	if a.currentFormedWord() == a.currentFormatWord() {
		return arg, nil
	}
	if len(a.formedWords) <= a.formedWordsIndex {
		return "", a.notMatchError()
	}
	next, err := a.getArg(a.currentFormedWord())
	if err != nil {
		return "", err
	}
	return arg + " " + next, nil
}

func (a *argExtractor) ExtractArgs() ([]string, error) {
	if len(a.formedWords) <= a.formedWordsIndex && len(a.formatWords) <= a.formatWordsIndex {
		return a.result, nil
	}
	if len(a.formedWords) <= a.formedWordsIndex || len(a.formatWords) <= a.formatWordsIndex {
		return nil, a.notMatchError()
	}

	if a.currentFormatWord() == "%s" {
		a.formatWordsIndex++
		arg, err := a.getArg(a.currentFormedWord())
		if err != nil {
			return nil, err
		}
		a.result = append(a.result, arg)
	}

	if a.currentFormedWord() == a.currentFormatWord() {
		a.formedWordsIndex++
		a.formatWordsIndex++
		return a.ExtractArgs()
	}
	return nil, a.notMatchError()
}
