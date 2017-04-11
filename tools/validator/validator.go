package validator

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/krostar/nebulo-golib/router/httperror"

	validator "gopkg.in/validator.v2"
)

type checkFunction func(toCheck string, options string) (err error)

type checkDefinition struct {
	checkFct  checkFunction
	omitError bool
}
type checkMap map[string]checkDefinition

func init() {
	var err error
	// tell to the validator lib that we have some function to use for our custom validators
	if err = validator.SetValidationFunc("file", File); err != nil {
		panic(fmt.Errorf("unable to set validation function %q: %v", "file", err))
	}
	if err = validator.SetValidationFunc("string", String); err != nil {
		panic(fmt.Errorf("unable to set validation function %q: %v", "string", err))
	}
}

func validate(v interface{}, checksToCall string, checksMapping checkMap) (err error) {
	// get the value to test
	st := reflect.ValueOf(v)
	if st.Kind() != reflect.String {
		return validator.ErrUnsupported
	}
	toCheck := st.String()

	// checks we want to make are separated with a +
	checks := strings.Split(checksToCall, "+")
	for _, check := range checks {
		optionsIndex := strings.Index(check, ":")
		options := ""
		if optionsIndex > -1 {
			options = check[optionsIndex+1:]
			check = check[:optionsIndex]
		}
		checkDef, ok := checksMapping[check]
		if ok {
			if err = checkDef.checkFct(toCheck, options); err != nil {
				err = fmt.Errorf("validation check failed: %v", err)
			}
		} else {
			err = fmt.Errorf("unable to find check named :'%s'", check)
		}

		if ok && err != nil && checkDef.omitError {
			return nil
		} else if err != nil {
			return err
		}

	}

	return nil
}

// HTTPErrors transform lib-errors into send-able HTTP error
func HTTPErrors(err error) *httperror.HTTPErrors {
	// get the lib-error-map
	vErrs, ok := err.(validator.ErrorMap)
	if !ok {
		return nil
	}

	// transform the map into a list of httperror
	var httpErrors *httperror.HTTPErrors
	for key, errs := range vErrs {
		var erra []string
		for _, err := range errs {
			erra = append(erra, err.Error())
		}
		key = strings.ToLower(key)
		if httpErrors == nil {
			httpErrors = httperror.New(http.StatusBadRequest, key, httperror.BadParam(erra))
		} else {
			httpErrors = httpErrors.Add(key, httperror.BadParam(erra))
		}
	}

	return httpErrors
}
