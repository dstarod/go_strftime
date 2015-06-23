package go_strftime

import(
	"fmt"
	"time"
	"strings"
)

func Strftime(layout string, t time.Time) string {
	// Implemented most useful of http://www.cplusplus.com/reference/ctime/strftime/
	// How to use : fmt.Println( Strftime("%Y-%m-%d", time.Now()) )

	// mm/dd/yy
	layout = strings.Replace(layout, "%D", "%m/%d/%y" , 100)	
	layout = strings.Replace(layout, "%x", "%m/%d/%y" , 100)
	// YYYY-mm-dd
	layout = strings.Replace(layout, "%F", "%Y-%m-%d" , 100)
	// HH:MM:SS AM/PM
	layout = strings.Replace(layout, "%r", "%I:%M:%S %p" , 100)
	// HH:MM
	layout = strings.Replace(layout, "%R", "%H:%M" , 100)
	// HH:MM:SS
	layout = strings.Replace(layout, "%T", "%H:%M:%S" , 100)
	layout = strings.Replace(layout, "%X", "%H:%M:%S" , 100)
	// Century
	layout = strings.Replace(layout, "%C", fmt.Sprintf("%v", int(t.Year()/100)), 100)
	// Day of year
	layout = strings.Replace(layout, "%j", fmt.Sprintf("%03v", t.YearDay()), 100)
	// Year
	layout = strings.Replace(layout, "%Y", fmt.Sprintf("%04v", t.Year()), 100)
	layout = strings.Replace(layout, "%y", fmt.Sprintf("%04v", t.Year())[2:4], 100)
	// Month
	layout = strings.Replace(layout, "%m", fmt.Sprintf("%02v", int(t.Month())), 100)
	layout = strings.Replace(layout, "%B", t.Month().String(), 100)
	// Day	
	layout = strings.Replace(layout, "%d", fmt.Sprintf("%02v", t.Day()), 100)
	layout = strings.Replace(layout, "%e", fmt.Sprintf("%v", t.Day()), 100)
	// Weekday
	layout = strings.Replace(layout, "%A", t.Weekday().String(), 100)
	layout = strings.Replace(layout, "%u", fmt.Sprintf("%v", int(t.Weekday())), 100)
	layout = strings.Replace(layout, "%w", fmt.Sprintf("%v", int(t.Weekday())-1 ), 100)
	// Hour
	layout = strings.Replace(layout, "%H", fmt.Sprintf("%02v", t.Hour()), 100)
	layout = strings.Replace(layout, "%I", fmt.Sprintf("%02v", t.Hour() - (t.Hour()/12)*12 ), 100)
	// Minute
	layout = strings.Replace(layout, "%M", fmt.Sprintf("%02v", t.Minute()), 100)
	// Second
	layout = strings.Replace(layout, "%S", fmt.Sprintf("%02v", t.Second()), 100)
	// AM/PM
	if t.Hour()/12 > 0{
		layout = strings.Replace(layout, "%p", "PM", 100)
	} else {
		layout = strings.Replace(layout, "%p", "AM", 100)
	}
	// Symbols
	layout = strings.Replace(layout, "%n", "\n", 100)
	layout = strings.Replace(layout, "%t", "\t", 100)
	layout = strings.Replace(layout, "%%", "%", 100)
		
	return layout
}