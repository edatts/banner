package banner

import (
	"errors"
	"fmt"
)

var (
	ErrNoText          = errors.New("no text provided")
	ErrUnsupportedRune = errors.New("unsupported rune")
)

func unsupportedRuneError(char rune) error {
	return fmt.Errorf("%w: %v", ErrUnsupportedRune, string(char))
}
