package gomatchers

import (
	"fmt"
	"reflect"
	"strings"
)

type structEqualWithExceptionMatcher struct {
	sampleValue      reflect.Value
	exceptFieldNames []string
}

func (m structEqualWithExceptionMatcher) Matches(x interface{}) bool {
	preappValue := reflect.ValueOf(x)
	if preappValue.Kind() == reflect.Ptr {
		preappValue = preappValue.Elem()
	}

fields:
	for i, n := 0, preappValue.NumField(); i < n; i++ {
		fieldName := preappValue.Type().Field(i).Name
		for i := range m.exceptFieldNames {
			if fieldName == m.exceptFieldNames[i] {
				continue fields
			}
		}
		if !reflect.DeepEqual(preappValue.Field(i).Interface(), m.sampleValue.Field(i).Interface()) {
			return false
		}
	}

	return true
}

func (m structEqualWithExceptionMatcher) String() string {
	return fmt.Sprintf("matches struct %+v except fields: %s", m.sampleValue.Interface(), strings.Join(m.exceptFieldNames, ", "))
}

func StrutEqualExceptFields(sample interface{}, exceptFieldNames ...string) structEqualWithExceptionMatcher {
	sampleValue := reflect.ValueOf(sample)
	if sampleValue.Kind() == reflect.Ptr {
		sampleValue = sampleValue.Elem()
	}
	return structEqualWithExceptionMatcher{
		sampleValue:      sampleValue,
		exceptFieldNames: exceptFieldNames,
	}
}
