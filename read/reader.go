package read

import "regexp"

const MALEXP = "[\\s,]*(~@|[\\[\\]{}()'`~^@]|\"(?:\\.|[^\\\"])*\"?|;.*|[^\\s\\[\\]{}('\"`,;)]*)"

func tokenize(in string) []string {
	re := regexp.MustCompile(MALEXP)
	submatches := re.FindAllStringSubmatch(in, -1)

	tokens := make([]string, len(submatches))
	for i, submatch := range submatches {
		tokens[i] = submatch[1]
	}
	return tokens
}

type reader interface {
	next() string // Returns next available token.
	peek() string // Returns next available token without consuming it.
}

type readerImpl struct {
	position int
	tokens   []string
}

func (r readerImpl) next() string {
	currToken := r.peek()
	r.position += 1
	return currToken
}

func (r readerImpl) peek() string {
	return r.tokens[r.position]
}

func newReader(in string) reader {
	return readerImpl{
		tokens: tokenize(in),
	}
}
