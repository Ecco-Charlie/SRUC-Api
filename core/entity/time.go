package entity

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const CTimeFormat = "15:04:05"

type CTime time.Time

func NewCTime(hour, min, sec int) CTime {
	t := time.Date(0, time.January, 1, hour, min, sec, 0, time.UTC)
	return CTime(t)
}

func (t *CTime) Scan(value any) error {
	switch v := value.(type) {
	case []byte:
		return t.UnmarshalText(string(v))
	case string:
		return t.UnmarshalText(v)
	case time.Time:
		*t = CTime(v)
	case nil:
		*t = CTime{}
	default:
		return fmt.Errorf("cannot sql.Scan() MyTime from: %#v", v)
	}
	return nil
}

func (t CTime) Value() (driver.Value, error) {
	return driver.Value(time.Time(t).Format(CTimeFormat)), nil
}

func (t *CTime) UnmarshalText(value string) error {
	dd, err := time.Parse(CTimeFormat, value)
	if err != nil {
		return err
	}
	*t = CTime(dd)
	return nil
}

func (CTime) GormDataType() string {
	return "TIME"
}
