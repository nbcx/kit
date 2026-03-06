package nums

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Decimal struct {
	dec decimal.Decimal
}

// NewDecimal 精度计算
func NewDecimal(d1 interface{}, d2 ...interface{}) *Decimal {
	return &Decimal{decimalAdd(d1, d2...)}
}

// Mul 乘
func (d *Decimal) Mul(d1 interface{}, d2 ...interface{}) *Decimal {
	d.dec = d.dec.Mul(decimalAdd(d1, d2...))
	return d
}

// Add 加
func (d *Decimal) Add(d1 interface{}, d2 ...interface{}) *Decimal {
	d.dec = d.dec.Add(decimalAdd(d1, d2...))
	return d
}

// Div 除
func (d *Decimal) Div(d1 interface{}, d2 ...interface{}) *Decimal {
	d.dec = d.dec.Div(decimalAdd(d1, d2...))
	return d
}

// Sub 减
func (d *Decimal) Sub(d1 interface{}, d2 ...interface{}) *Decimal {
	d.dec = d.dec.Sub(decimalAdd(d1, d2...))
	return d
}

// String 得到运算结果，字符串类型
func (d *Decimal) String() string {
	return d.dec.String()

}

// Int64 得到Int64类型的运算结果
func (d *Decimal) Int64() int64 {
	return d.dec.IntPart()
}

// Int32 得到Int32类型的运算结果
func (d *Decimal) Int32() int32 {
	return int32(d.dec.IntPart()) // nolint
}

// Int 得到Int类型的运算结果
func (d *Decimal) Int() int {
	return int(d.dec.IntPart()) // nolint
}

// Float 得到Float64类型的运算结果
func (d *Decimal) Float() float64 {
	return d.dec.InexactFloat64()
}

// Float32 得到Float32类型的运算结果
func (d *Decimal) Float32() float32 {
	return float32(d.Float())
}

func decimalAdd(d1 interface{}, d2 ...interface{}) decimal.Decimal {
	add, err := decimal.NewFromString(fmt.Sprint(d1))
	if err != nil {
		panic(fmt.Errorf("decimal add err %v case %v", err, d1))
	}
	for _, v := range d2 {
		vv, err := decimal.NewFromString(fmt.Sprint(v))
		if err != nil {
			panic(fmt.Errorf("decimal for add err %v case %v", err, v))
		}
		add = add.Add(vv)
	}
	return add
}
