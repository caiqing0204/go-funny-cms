package base

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `json:"id"`

	CreatedAt TimeAt `json:"created_at,omitempty" gorm:"-"`
	UpdatedAt TimeAt `json:"updated_at,omitempty" gorm:"-"`
}

type Result struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
}

type TimeAt time.Time

func (t *TimeAt) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = TimeAt(t1)
	return err
}

func (t TimeAt) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t TimeAt) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *TimeAt) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = TimeAt(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *TimeAt) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}

type IntJson []int

func (v IntJson) Value() (driver.Value, error) {
	return json.Marshal(v)
}

func (v *IntJson) Scan(data interface{}) error {
	bytes, _ := data.([]byte)
	r := IntJson{}
	err := json.Unmarshal(bytes, &r)
	*v = IntJson(r)

	return err
}
