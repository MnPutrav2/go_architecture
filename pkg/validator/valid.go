package validator

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/MnPutrav2/go_architecture/pkg/decoder"
)

func Validate(validator []string, payload any) error {
	payloadType := reflect.TypeOf(payload)
	payloadValue := reflect.ValueOf(payload)
	tipeRange := len(validator)
	var errs []string

	if payloadType.Kind() == reflect.Pointer {
		payloadType = payloadType.Elem()
	}

	if tipeRange != payloadValue.NumField() {
		return fmt.Errorf("invalid validator format")
	}

	for i := range tipeRange {
		validRange := strings.Split(validator[i], "|")
		var er []string
		for x := range len(validRange) {

			// Nullabel
			if strings.Contains(validRange[x], "null") {
				return nil
			}

			// Required
			if strings.Contains(validRange[x], "required") {
				c := payloadValue.Field(i).String()
				if c == "" {
					er = append(er, "required value")
				}
			}

			// Max value
			if strings.Contains(validRange[x], "max") {
				h := strings.Split(validRange[x], ":")
				c := payloadValue.Field(i).String()
				x := len(c)

				m, _ := strconv.Atoi(h[1])
				if x > m {
					er = append(er, fmt.Sprintf("maximux %s chacter", h[1]))
				}
			}

			// Min value
			if strings.Contains(validRange[x], "min") {
				h := strings.Split(validRange[x], ":")
				c := payloadValue.Field(i).String()
				x := len(c)

				m, _ := strconv.Atoi(h[1])
				if x < m {
					er = append(er, fmt.Sprintf("minimal %s chacter", h[1]))
				}
			}

		}

		if len(er) == 0 {
			return nil
		}

		errs = append(errs, fmt.Sprintf("Field %s (%s)", payloadType.Field(i).Tag.Get("json"), strings.Join(er, ",")))
	}

	if len(errs) == 0 {
		return nil
	}

	return fmt.Errorf("%s", strings.Join(errs, ", "))
}

func ValidatePayload[T any](r *http.Request) (T, error) {
	var e T
	payload, err := decoder.BodyDecoder[T](r)
	if err != nil {
		return e, fmt.Errorf("failed decode body")
	}

	payloadType := reflect.TypeOf(payload)
	payloadValue := reflect.ValueOf(payload)
	var errs []string

	if payloadType.Kind() == reflect.Pointer {
		payloadType = payloadType.Elem()
	}

	for i := range payloadType.NumField() {
		validRange := strings.Split(payloadType.Field(i).Tag.Get("validate"), ";")

		fmt.Println(validRange)
		var er []string
		for x := range len(validRange) {

			// Nullabel
			if strings.Contains(validRange[x], "null") {
				return e, nil
			}

			// Required
			if strings.Contains(validRange[x], "required") {
				c := payloadValue.Field(i).String()
				if c == "" {
					er = append(er, "required value")
				}
			}

			// Max value
			if strings.Contains(validRange[x], "max:") {
				h := strings.Split(validRange[x], ":")
				c := payloadValue.Field(i).String()
				x := len(c)

				m, _ := strconv.Atoi(h[1])
				if x > m {
					er = append(er, fmt.Sprintf("maximux %s character", h[1]))
				}
			}

			// Min value
			if strings.Contains(validRange[x], "min:") {
				h := strings.Split(validRange[x], ":")
				c := payloadValue.Field(i).String()
				x := len(c)

				m, _ := strconv.Atoi(h[1])
				if x < m {
					er = append(er, fmt.Sprintf("minimal %s character", h[1]))
				}
			}

		}

		if len(er) == 0 {
			return e, nil
		}

		errs = append(errs, fmt.Sprintf("Field %s (%s)", payloadType.Field(i).Tag.Get("json"), strings.Join(er, ",")))
	}

	if len(errs) == 0 {
		return e, nil
	}

	return e, fmt.Errorf("%s", strings.Join(errs, ", "))
}
