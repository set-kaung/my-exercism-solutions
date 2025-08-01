package wordy

import (
	"strconv"
	"strings"
)

type TOKEN_TYPE string

const (
	NUMBER   TOKEN_TYPE = "number"
	OPERATOR TOKEN_TYPE = "op"
)

const (
	ADD      = "plus"
	SUB      = "minus"
	MULTIPLY = "multiplied"
	DIVIDE   = "divided"
)

type token struct {
	tok_type TOKEN_TYPE
	value    string
}

func Answer(question string) (int, bool) {
	tokens, ok := lex(question)
	if !ok {
		return -1, false
	}
	return parseAndEval(tokens)
}

func parseAndEval(tokens []token) (int, bool) {
	if len(tokens) == 0 {
		return -1, false
	}
	total, err := strconv.Atoi(tokens[0].value)
	if err != nil {
		return -1, false
	}
	prev := NUMBER
	prevOps := ""
	for i := 1; i < len(tokens); i += 1 {
		currentTok := tokens[i]
		if currentTok.tok_type == prev {
			return -1, false
		}
		if currentTok.tok_type == NUMBER {
			val, err := strconv.Atoi(currentTok.value)
			if err != nil {
				return -1, false
			}
			switch prevOps {
			case "+":
				total += val
			case "-":
				total -= val
			case "*":
				total *= val
			case "/":
				total /= val
			default:
				return -1, false
			}

			prev = NUMBER
			prevOps = ""
		} else {
			prev = OPERATOR
			prevOps = currentTok.value
		}
	}
	if prevOps != "" {
		return -1, false
	}

	return total, true
}

func lex(question string) ([]token, bool) {

	question = strings.ToLower(question)
	question = strings.TrimRight(question, "?")
	parts := strings.Split(question, " ")
	if len(parts) < 3 {
		return nil, false
	}
	if parts[0] != "what" && parts[1] != "is" {
		return nil, false
	}
	tokens := make([]token, 0, len(parts)-2)

	i := 2
	for i < len(parts) {
		current := parts[i]
		switch current {
		case ADD:
			tokens = append(tokens, token{tok_type: OPERATOR, value: "+"})
			i += 1
		case SUB:
			tokens = append(tokens, token{tok_type: OPERATOR, value: "-"})
			i += 1
		case MULTIPLY:
			tokens = append(tokens, token{tok_type: OPERATOR, value: "*"})
			i += 2
		case DIVIDE:
			tokens = append(tokens, token{tok_type: OPERATOR, value: "/"})
			i += 2
		default:
			tokens = append(tokens, token{tok_type: NUMBER, value: current})
			i += 1
		}
	}
	return tokens, true
}
