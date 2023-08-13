package validator

import "strings"

const (
	ErrFieldIsEmpty     = "نام فیلد خالی است"
	ErrAlreadyHasID     = "فیلد آیدی موجود است . از آپدیت استفاده کنید"
	ErrProductExists    = "این محصول قبلا ثبت شده است"
	ErrProductNotFound  = "محصول مورد نظر یافت نشد"
	ErrCategoryNotFound = "دسته بندی مورد نظر یافت نشد"
	ErrCategoryExists   = "این دسته بندی موجود است"
)

type (
	Rule func() error

	Validator interface {
		AddRule(Rule) Validator
		Validate() error
	}

	ValidatorError struct {
		errors  []error
		strings []string
	}

	validator struct {
		rules []Rule
	}
)

func New() Validator {
	return &validator{
		rules: make([]Rule, 0),
	}
}

func (v *validator) AddRule(rule Rule) Validator {
	v.rules = append(v.rules, rule)
	return v
}

func (v *validator) Validate() error {
	ret := new(ValidatorError)
	for _, rule := range v.rules {
		if err := rule(); err != nil {
			ret.errors = append(ret.errors, err)
			ret.strings = append(ret.strings, err.Error())
		}
	}

	if len(ret.errors) == 0 {
		return nil
	}

	return ret
}

func (v *ValidatorError) Error() string {
	return strings.Join(v.strings, ", ")
}
