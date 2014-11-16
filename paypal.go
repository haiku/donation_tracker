package main

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

const PayPalDateFormat = "2006-01-02T15:04:05Z"

type PayPalTxn struct {
	Timestamp time.Time
	Type string
	Email string
	Name string
	TransactionID string
	Status string
	Amt float32
	FeeAmt float32
	NetAmt float32
	CurrencyCode string
}

func NewPayPalTxn(tran map[string]string) *PayPalTxn {
	result := new(PayPalTxn)

	elem := reflect.ValueOf(result).Elem()

	for name, value := range(tran) {
		field := elem.FieldByNameFunc(func(n string) bool {
			return name == strings.ToUpper(n)
		})
		if field.IsValid() && field.CanSet() {
			switch field.Kind() {
			case reflect.String:
				field.SetString(value)
			case reflect.Float32:
				f, _ := strconv.ParseFloat(value, 32)
				field.SetFloat(f)
			default:
				// its the Timestamp field
				timestamp, _ := time.Parse(PayPalDateFormat, value)
				field.Set(reflect.ValueOf(timestamp))
			}
		}
	}

	return result
}

type PayPalTxns []*PayPalTxn

func PayPalTxnsFromNvp(nvp *NvpResult) PayPalTxns {
	result := make(PayPalTxns, len(nvp.List))

	for i, item := range(nvp.List) {
		result[i] = NewPayPalTxn(item)
		// fmt.Println(result[i])
	}

	return result
}

// Sorting

func (p PayPalTxns) Len() int      { return len(p) }
func (p PayPalTxns) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type ByDate struct { PayPalTxns }

func (s ByDate) Less(i, j int) bool {
	return s.PayPalTxns[i].Timestamp.Before(s.PayPalTxns[j].Timestamp)
}

func (p PayPalTxns) Sort() {
	sort.Sort(ByDate{p})
}
