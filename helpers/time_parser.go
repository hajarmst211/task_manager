package helpers

import(
	"time"
)

const TimeLayout = "2006-01-02"

func DateParser(inputDate string) (time.Time, error) {
   
    date, err := time.Parse(TimeLayout, inputDate)
    if err != nil {
        return time.Time{}, err
    }
    return date, nil
}

