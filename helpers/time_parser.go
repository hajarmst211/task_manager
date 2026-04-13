package helpers

import(
	"time"
)

func DateParser(inputDate string) (time.Time, error) {
    const layout = "2006-01-02"
    date, err := time.Parse(layout, inputDate)
    if err != nil {
        return time.Time{}, err
    }
    return date, nil
}

