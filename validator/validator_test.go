package validator

import (
	"log"
	"testing"
)

func TestPhoneNumber(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		for i, phone := range []string{
			"09121234567",
			"09129876543",
			"09365432109",
			"09217998990",
		} {
			validator := New()
			err := validator.AddRule(Phone(phone)).Validate()
			if err != nil {
				t.Fatalf("[test case %d] expected nil got %v", i, err)
			}
		}
	})

	t.Run("invalid", func(t *testing.T) {
		for i, phone := range []string{
			"",
			"0912",
			"0936",
			"1234567",
			"0214567895",
			"sdf32412sdf",
		} {
			validator := New()
			err := validator.AddRule(Phone(phone)).Validate()
			if err == nil {
				t.Fatalf("[test case %d] expected nil got %v", i, err)
			}
		}
	})

}

func TestValidEmail(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		for i, email := range []string{
			"mostafa.solati@gmail.com",
			"mostafasolati@gmail.com",
			"a@b.c",
		} {
			validator := New()
			err := validator.AddRule(Email(email)).Validate()
			if err != nil {
				log.Fatalf("[test case %d] expected to be valid email but got: %v", i, err)
			}

		}
	})

	t.Run("invalid", func(t *testing.T) {
		for i, email := range []string{
			"mostafa.solati.com",
			"@gmail.com",
			"mostafa@com",
			"mostafa@.com",
		} {
			validator := New()
			err := validator.AddRule(Email(email)).Validate()
			if err == nil || err.Error() != "ایمیل معتبر نیست" {
				log.Fatalf("[test case %d] expected to be invalid email but got nil", i)
			}

		}
	})
}

func TestValidatorEmptyString(t *testing.T) {

	type testcase struct {
		field string
		value string
		err   string
	}

	for i, test := range []testcase{
		{field: "", value: "", err: ErrFieldIsEmpty},
		{field: " ", value: "", err: ErrFieldIsEmpty},
		{field: "", value: "bar", err: ErrFieldIsEmpty},
		{field: " ", value: "bar", err: ErrFieldIsEmpty},
		{field: "نام", value: "", err: "نام خالی است"},
		{field: "نام خانوادگی", value: " ", err: "نام خانوادگی خالی است"},
	} {
		v := New()
		err := v.
			AddRule(String(test.field, test.value)).
			Validate()
		if err.Error() != test.err {
			t.Fatalf("Expected %s got %s in test case %d", test.err, err.Error(), i)
		}
	}
}

func TestValidatorEmptyNumber(t *testing.T) {

	type testcase struct {
		field string
		value int
		err   string
	}

	for i, test := range []testcase{
		{field: "", value: 0, err: ErrFieldIsEmpty},
		{field: " ", value: 0, err: ErrFieldIsEmpty},
		{field: "", value: 5, err: ErrFieldIsEmpty},
		{field: " ", value: 3, err: ErrFieldIsEmpty},
		{field: "وزن", value: 0, err: "وزن خالی است"},
		{field: "قیمت", value: 0, err: "قیمت خالی است"},
	} {
		v := New()
		err := v.
			AddRule(Number(test.field, test.value)).
			Validate()
		if err.Error() != test.err {
			t.Fatalf("Expected %s got %s in test case %d", test.err, err.Error(), i)
		}
	}
}

func TestComplex(t *testing.T) {

	t.Run("valid", func(t *testing.T) {
		suites := [][]Rule{
			{
				Number("قیمت", 15000),
			},
			{
				Number("قیمت", 3700),
				Number("وزن", 0.1),
			},
			{
				String("آدرس", "ایران زیبا کوی بهشت"),
				String("توضیحات", "لطفا روی حلیم شکر بریزید"),
			},
			{
				Number("قیمت", 15000),
				String("نام", "مریم ببریان"),
				Number("وزن", 0.1),
			},
		}

		for i, suit := range suites {
			validator := New()
			for _, rule := range suit {
				validator.AddRule(rule)
			}
			err := validator.Validate()
			if err != nil {
				log.Fatalf("[test case %d] Expected nil got %v", i, err)
			}
		}

	})

	t.Run("not valid", func(t *testing.T) {
		suites := [][]Rule{
			{
				Number("قیمت", 0),
			},
			{
				Number("قیمت", 0),
				Number("وزن", 0.1),
			},
			{
				String("آدرس", " "),
				String("توضیحات", "لطفا روی حلیم شکر بریزید"),
			},
			{
				Number("قیمت", 15000),
				String("نام", "مریم خوانچه‌رو"),
				Number("وزن", 0),
			},
		}

		for i, suit := range suites {
			validator := New()
			for _, rule := range suit {
				validator.AddRule(rule)
			}
			err := validator.Validate()
			if err == nil {
				log.Fatalf("[test case %d] Expected %s got nil", i, err)
			}
		}

	})

}
