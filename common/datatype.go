package common

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type DateOnly time.Time

func (d DateOnly) Value() (driver.Value, error) {
	t := time.Time(d)
	return t.Format("2006-01-02"), nil
}

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = DateOnly(t)
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(d).Format("2006-01-02") + `"`), nil
}

func (d *DateOnly) Scan(value any) error {
	s, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("cannot scan type %T into DateOnly", value)
	}
	*d = DateOnly(s)
	return nil
}

func StructToMap(u interface{}) map[string]any {
	result := make(map[string]any)
	val := reflect.ValueOf(u)
	typ := reflect.TypeOf(u)

	for i := 0; i < val.NumField(); i++ { // loop semua field
		field := val.Field(i)
		fieldType := typ.Field(i)
		gormTag := fieldType.Tag.Get("gorm")
		columnName := parseGormColumn(gormTag)
		if columnName == "" {
			continue
		}
		result[columnName] = field.Interface()
	}
	return result
}

func parseGormColumn(tag string) string {
	// tag bisa berupa "primaryKey;column:cust_id"
	parts := strings.Split(tag, ";")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if strings.HasPrefix(p, "column:") {
			return strings.TrimPrefix(p, "column:")
		}
	}
	return ""
}
