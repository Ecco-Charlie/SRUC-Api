package entity

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const CTimeFormat = "2006-01-02 15:04:05"

type CTime time.Time

func NewCTime(year, month, day, hour, min, sec int) CTime {
	t := time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC)
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
		return fmt.Errorf("cannot sql.Scan() CTime from: %#v", v)
	}
	return nil
}

func (t *CTime) Value() (driver.Value, error) {
	return driver.Value(time.Time(*t).Format(CTimeFormat)), nil
}

func (t *CTime) UnmarshalText(value string) error {
	parsed, err := time.Parse(CTimeFormat, value)
	if err != nil {
		return err
	}
	*t = CTime(parsed)
	return nil
}

func (t *CTime) String() string {
	return time.Time(*t).Format(CTimeFormat)
}

func (CTime) GormDataType() string {
	return "DATETIME"
}
