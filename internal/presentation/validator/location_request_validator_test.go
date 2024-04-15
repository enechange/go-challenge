package validators

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fieldLevelMock struct {
	field reflect.Value
	tag   string
	param string
}

func (fl fieldLevelMock) Top() reflect.Value {
	return reflect.Value{}
}

func (fl fieldLevelMock) Parent() reflect.Value {
	return reflect.Value{}
}

func (fl fieldLevelMock) Field() reflect.Value {
	return fl.field
}

func (fl fieldLevelMock) FieldName() string {
	return ""
}

func (fl fieldLevelMock) StructFieldName() string {
	return ""
}

func (fl fieldLevelMock) Param() string {
	return fl.param
}

func (fl fieldLevelMock) GetTag() string {
	return fl.tag
}

func (fl fieldLevelMock) ExtractType(field reflect.Value) (reflect.Value, reflect.Kind, bool) {
	if !field.IsValid() {
		return reflect.Value{}, reflect.Invalid, false
	}
	kind := field.Kind()
	nullable := kind == reflect.Ptr || kind == reflect.Interface || kind == reflect.Slice || kind == reflect.Map || kind == reflect.Chan
	return field, kind, nullable
}

func (fl fieldLevelMock) GetStructFieldOK() (reflect.Value, reflect.Kind, bool) {
	return reflect.Value{}, reflect.Invalid, false
}

func (fl fieldLevelMock) GetStructFieldOKAdvanced(val reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool) {
	return reflect.Value{}, reflect.Invalid, false
}

func (fl fieldLevelMock) GetStructFieldOK2() (reflect.Value, reflect.Kind, bool, bool) {
	return reflect.Value{}, reflect.Invalid, false, false
}

func (fl fieldLevelMock) GetStructFieldOKAdvanced2(val reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool, bool) {
	return reflect.Value{}, reflect.Invalid, false, false
}

func TestLatitudeValidation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid Latitude", "34.1234567", true},
		{"Invalid Latitude Too Low", "-91.1234567", false},
		{"Invalid Latitude Too High", "91.1234567", false},
		{"Valid Negative Latitude", "-34.1234567", true},
		{"Invalid Format", "34,123456", false},
		{"Excessive Length", "34.12345678", false},
		{"Boundary Valid Lower", "-90.0000000", true},
		{"Boundary Valid Upper", "90.0000000", true},
		{"Boundary Invalid Lower", "-90.00000001", false},
		{"Boundary Invalid Upper", "90.00000001", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := fieldLevelMock{field: reflect.ValueOf(tt.input)}
			result := LatitudeValidation(fl)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLongitudeValidation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid Longitude", "150.1234567", true},
		{"Invalid Longitude Too Low", "-181.1234567", false},
		{"Invalid Longitude Too High", "181.1234567", false},
		{"Valid Negative Longitude", "-150.1234567", true},
		{"Invalid Format", "150.123", true},
		{"Excessive Length", "-150.12345678", false},
		{"Edge Case Valid Lower", "-180.1234567", false},
		{"Edge Case Valid Upper", "180.1234567", false},
		{"Edge Case Invalid Lower", "-180.12345671", false},
		{"Edge Case Invalid Upper", "180.12345671", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := fieldLevelMock{field: reflect.ValueOf(tt.input)}
			result := LongitudeValidation(fl)
			assert.Equal(t, tt.expected, result)
		})
	}
}
