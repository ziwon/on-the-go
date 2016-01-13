// All the validators to validate the profile page
package validator

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ProfilePage struct {
	Form *url.Values
}

type Errors struct {
	Errors []error
}

func (p *ProfilePage) GetErrors() Errors {
	errs := make([]error, 0, 10)
	if *p.Form == nil || len(*p.Form) < 1 {
		errs = append(errs, errors.New("No data was received. Please submit from the profile page."))
	}

	for name, val := range *p.Form {
		if fn, ok := stringValidator[name]; ok {
			if err := fn(strings.Join(val, "")); err != nil {
				errs = append(errs, err)
			}
		} else {
			if fn, ok := stringsValidator[name]; ok {
				if err := fn(val); err != nil {
					errs = append(errs, err)
				}
			}
		}
	}

	return Errors{errs}
}

const (
	mmddyyyyForm = "01/02/2006"
	yyyymmddForm = "2006-01-02"
)

var stringValidator map[string]func(string) error = map[string]func(string) error{
	"age":         checkAge,
	"birthday":    checkDate,
	"chineseName": checkChineseName,
	"email":       checkEmail,
	"gender":      checkGender,
	"shirtsize":   checkShirtSize,
	"username":    checkUsername,
}

var stringsValidator map[string]func([]string) error = map[string]func([]string) error{
	"sibling": checkSibling,
}

func doSlicesIntersect(s1, s2 []string) bool {
	if s1 == nil || s2 == nil {
		return false
	}

	for _, str := range s1 {
		if isElementInSlice(str, s2) {
			return true
		}
	}
	return false
}

func isElementInSlice(str string, sl []string) bool {
	if sl == nil || str == "" {
		return false
	}

	for _, v := range sl {
		if v == str {
			return true
		}
	}
	return false
}

func checkChineseName(str string) error {
	if str != "" {
		if m, _ := regexp.MatchString("[\\x{4e00}-\\x{9fa5}]+$", strings.Trim(str, " ")); !m {
			return errors.New("Please make sure that the chinese name only contains chinese characters.")
		}

	}
	return nil
}

func checkUsername(str string) error {
	if strings.Trim(str, " ") == "" {
		return errors.New("Please enter a username.")
	}
	return nil
}

func checkAge(str string) error {
	age, err := strconv.Atoi(str)
	if str == "" || err != nil {
		return errors.New("Please enter a valid age.")
	}

	if age < 13 {
		return errors.New("You must be at least 13 years of age to submit.")
	}

	if age > 130 {
		return errors.New("You're too old to register, granpa")
	}
	return nil
}

func checkEmail(str string) error {
	if m, err := regexp.MatchString(`^[^@]+@[^@]+$`, str); !m {
		fmt.Println("err = ", err)
		return errors.New("Please enter a valid email address.")
	}
	return nil
}

func checkDate(str string) error {
	_, err := time.Parse(mmddyyyyForm, str)
	if err != nil {
		_, err = time.Parse(yyyymmddForm, str)
	}
	if str == "" || err != nil {
		return errors.New("Please enter a valid Date.")
	}
	return nil
}

func checkGender(str string) error {
	if str == "" {
		return nil
	}
	siblings := []string{"m", "f", "na"}
	if !isElementInSlice(str, siblings) {
		return errors.New("Please select a valid gender.")
	}
	return nil
}

func checkSibling(strs []string) error {
	if strs == nil || len(strs) < 1 {
		return nil
	}
	siblings := []string{"m", "f"}
	if siblings != nil && !doSlicesIntersect(siblings, strs) {
		return errors.New("Please select a valid sibling")
	}
	return nil
}

func checkShirtSize(str string) error {
	if str == "" {
		return nil
	}
	shirts := []string{"s", "m", "l", "xl", "xxl"}
	if !isElementInSlice(str, shirts) {
		return errors.New("Please select a valid shirt size")
	}
	return nil
}
