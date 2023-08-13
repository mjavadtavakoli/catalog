package validator

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/exp/constraints"
)

func Phone(value string) Rule {
	return func() error {
		if err := String("تلفن", value)(); err != nil {
			return err
		}

		value = strings.Map(func(r rune) rune {
			if unicode.IsDigit(r) {
				return r
			}
			return -1
		}, value)

		if strings.HasPrefix(value, "09") && len(value) == 11 {
			return nil
		}

		return errors.New("شماره تلفن معتبر نیست")
	}
}

func String(field, value string) Rule {
	return func() error {
		switch {
		case strings.TrimSpace(field) == "":
			return errors.New(ErrFieldIsEmpty)
		case strings.TrimSpace(value) == "":
			return fmt.Errorf("%s خالی است", field)
		default:
			return nil
		}
	}
}

func Number[T constraints.Integer | constraints.Float](field string, value T) Rule {
	var zero T
	return func() error {
		switch {
		case strings.TrimSpace(field) == "":
			return errors.New(ErrFieldIsEmpty)
		case value == zero:
			return fmt.Errorf("%s خالی است", field)
		default:
			return nil
		}
	}
}

func Email(value string) Rule {
	return func() error {
		if err := String("ایمیل", value)(); err != nil {
			return err
		}
		parts := strings.Split(value, "@")
		if len(parts) > 1 && parts[0] != "" && parts[1] != "" {
			domainParts := strings.Split(parts[1], ".")
			if len(domainParts) > 1 && domainParts[0] != "" && domainParts[1] != "" {
				return nil
			}
		}

		return errors.New("ایمیل معتبر نیست")
	}
}
