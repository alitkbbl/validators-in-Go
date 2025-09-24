package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func registerValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("nationalid", validNationalID)
		v.RegisterValidation("birthdate", validBirthDate)
		v.RegisterValidation("phonestart", validPhoneStart)
		v.RegisterStructValidation(validUserStruct, User{})
	}
}

func validUserStruct(sl validator.StructLevel) {
	user := sl.Current().Interface().(User)

	birthDate, err := time.Parse("2006/01/02", user.BirthDate)
	if err == nil {
		age := time.Now().Year() - birthDate.Year()
		if time.Now().YearDay() < birthDate.YearDay() {
			age--
		}
		if age < 7 {
			sl.ReportError(user.BirthDate, "birth_date", "BirthDate", "age", "User must be at least 7 years old")
		}
		if age > 100 {
			sl.ReportError(user.BirthDate, "birth_date", "BirthDate", "age", "User must be less than 100 years old")
		}
	}

	if user.Username == user.Email {
		sl.ReportError(user.Username, "username", "Username", "unique", "Username and email must be different")
	}
}

var validNationalID validator.Func = func(fl validator.FieldLevel) bool {
	nationalID := fl.Field().String()
	if len(nationalID) != 10 {
		return false
	}

	matched, _ := regexp.MatchString(`^\d{10}$`, nationalID)
	if !matched {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(nationalID[i]))
		sum += digit * (10 - i)
	}

	remainder := sum % 11
	controlDigit, _ := strconv.Atoi(string(nationalID[9]))

	if remainder < 2 {
		return controlDigit == remainder
	} else {
		return controlDigit == (11 - remainder)
	}
}

var validBirthDate validator.Func = func(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	birthDate, err := time.Parse("2006/01/02", dateStr)
	if err != nil {
		return false
	}

	now := time.Now()
	sevenYearsAgo := now.AddDate(-7, 0, 0)
	hundredYearsAgo := now.AddDate(-100, 0, 0)

	return birthDate.Before(sevenYearsAgo) && birthDate.After(hundredYearsAgo)
}

var validPhoneStart validator.Func = func(fl validator.FieldLevel) bool {
	num := fl.Field().String()
	return strings.HasPrefix(num, "09")
}
