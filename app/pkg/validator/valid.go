package validator

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/MnPutrav2/go_architecture/app/pkg/decoder"
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
	var zero T

	payload, err := decoder.BodyDecoder[T](r)
	if err != nil {
		return zero, fmt.Errorf("failed decode body")
	}

	payloadType := reflect.TypeOf(payload)
	payloadValue := reflect.ValueOf(payload)

	// Jika suatu saat decoder mengembalikan pointer
	if payloadType.Kind() == reflect.Pointer {
		payloadType = payloadType.Elem()
		payloadValue = payloadValue.Elem()
	}

	if payloadType.Kind() != reflect.Struct {
		return zero, fmt.Errorf("payload must be struct")
	}

	var errs []string

	for i := 0; i < payloadType.NumField(); i++ {

		field := payloadType.Field(i)
		value := payloadValue.Field(i)

		rules := strings.Split(field.Tag.Get("validate"), ";")

		var fieldErr []string

		for _, rule := range rules {

			switch {

			// nullable
			case rule == "null":
				continue

			// required
			case rule == "required":
				switch value.Kind() {
				case reflect.String:
					if strings.TrimSpace(value.String()) == "" {
						fieldErr = append(fieldErr, "required")
					}
				}

			// min:x
			case strings.HasPrefix(rule, "min:"):
				if value.Kind() != reflect.String {
					continue
				}

				n, _ := strconv.Atoi(strings.TrimPrefix(rule, "min:"))
				if len(value.String()) < n {
					fieldErr = append(fieldErr, fmt.Sprintf("minimum %d characters", n))
				}

			// max:x
			case strings.HasPrefix(rule, "max:"):
				if value.Kind() != reflect.String {
					continue
				}

				n, _ := strconv.Atoi(strings.TrimPrefix(rule, "max:"))
				if len(value.String()) > n {
					fieldErr = append(fieldErr, fmt.Sprintf("maximum %d characters", n))
				}
			}
		}

		if len(fieldErr) > 0 {

			name := field.Tag.Get("json")
			if name == "" {
				name = field.Name
			}

			errs = append(errs,
				fmt.Sprintf("%s (%s)",
					name,
					strings.Join(fieldErr, ", "),
				),
			)
		}
	}

	if len(errs) > 0 {
		return zero, fmt.Errorf("%s", strings.Join(errs, "; "))
	}

	return payload, nil
}
