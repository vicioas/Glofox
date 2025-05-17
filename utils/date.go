package utils

import (
 "fmt"
 "time"
)

func GenerateDateRange(startStr, endStr string) ([]string, error) {
 const layout = "2006-01-02"

 start, err1 := time.Parse(layout, startStr)
 if err1 != nil {
 return nil, fmt.Errorf("invalid start_date: %v", err1)
 }

 end, err2 := time.Parse(layout, endStr)
 if err2 != nil {
 return nil, fmt.Errorf("invalid end_date: %v", err2)
 }

 if end.Before(start) {
 return nil, fmt.Errorf("end_date must not be before start_date")
 }
  var dates []string
 for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
 dates = append(dates, d.Format(layout))
 }
 return dates, nil
}
