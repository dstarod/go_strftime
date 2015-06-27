package go_strftime

import(
	"fmt"
	"time"
	"regexp"
)

func Strftime(layout string, t time.Time) string {
	// Implemented most useful of http://www.cplusplus.com/reference/ctime/strftime/
	// Skipped: %U, %W (not ISO week number) and modifiers E and O
	// How to use : fmt.Println( Strftime("%Y-%m-%d", time.Now()) )
	
	// Prepared data: time zone, utc offset, week based year, week number
	time_zone, utc_offset := t.Zone()
	minutes_offset := utc_offset/60
	h := minutes_offset/60
	m := minutes_offset%60
	sign := "+"
	if utc_offset<0{
		sign = "-"
	}
	iso_offset := fmt.Sprintf("%v%v%02v", sign, h, m)
	
	week_based_year, week_num := t.ISOWeek()

	// Group patterns	
	patterns_regexp := regexp.MustCompile("%[DxcFrRTX]{1}")
	patterns_positions := patterns_regexp.FindAllStringIndex(layout, 1000)
	
	var pattern string;
	var newstr string;
	pos := 0
	
	for _, p := range patterns_positions{
		newstr += layout[pos:p[0]]
		pos = p[1]
		
		pattern = layout[p[0]:p[1]]
		switch pattern{
			case "%D", "%x":
				newstr += "%m/%d/%y"
			case "%c":
				newstr += "%a %b %e %H:%M:%S %Y"
			case "%F":
				newstr += "%Y-%m-%d"
			case "%r":
				newstr += "%I:%M:%S %p"
			case "%R":
				newstr += "%H:%M"
			case "%T", "%X":
				newstr += "%H:%M"
			default:
				newstr += ""
		}
	}
	
	if len(newstr) > 0{
		newstr += layout[pos:]
		layout = newstr
	}
	
	// Single patterns
	patterns_regexp = regexp.MustCompile("%[ZzCjGgyYmBbhdeVAauwHIMSpnt%]{1}")
	patterns_positions = patterns_regexp.FindAllStringIndex(layout, 1000)
	
	newstr = ""
	pos = 0
	
	for _, p := range patterns_positions{
		newstr += layout[pos:p[0]]
		pos = p[1]
		
		pattern = layout[p[0]:p[1]]
		switch pattern{
			case "%Z":
				// Timezone
				newstr += time_zone
			case "%z":
				// UTC offset
				newstr += iso_offset
			case "%C":
				// Century
				newstr += fmt.Sprintf("%v", int(t.Year()/100))
			case "%j":
				// Day of year
				newstr += fmt.Sprintf("%03v", t.YearDay())
			case "%G":
				// Week based year (4 digits)
				newstr += fmt.Sprintf("%04v", week_based_year)
			case "%g":
				// Week based year (last 2 digits)
				newstr += fmt.Sprintf("%04v", week_based_year)[2:4]
			case "%y":
				// Year (last 2 digits)
				newstr += fmt.Sprintf("%04v", t.Year())[2:4]
			case "%Y":
				// Year
				newstr += fmt.Sprintf("%04v", t.Year())
			case "%m":
				// Month (2 digit)
				newstr += fmt.Sprintf("%02v", int(t.Month()))
			case "%B":
				// Month full name
				newstr += t.Month().String()
			case "%b", "%h":
				// Month short name
				switch t.Month(){
					case time.June, time.July, time.September:
						newstr += t.Month().String()[:4]
					default:
						newstr += t.Month().String()[:3]
				}
			case "%d":
				// Day of month (2 digit)
				newstr += fmt.Sprintf("%02v", t.Day())
			case "%e":
				// Day of month
				newstr += fmt.Sprintf("%v", t.Day())
			case "%V":
				// ISO 8601 week number
				newstr += fmt.Sprintf("%v", week_num)
			case "%A":
				// Weekday
				newstr += t.Weekday().String()
			case "%a":
				// Weekday short name
				newstr += t.Weekday().String()[:3]
			case "%u":
				// Weekday number (1-7)
				newstr += fmt.Sprintf("%v", int(t.Weekday()))
			case "%w":
				// Weekday numder (0-6)
				newstr += fmt.Sprintf("%v", int(t.Weekday())-1)
			case "%H":
				// Hour (00-24)
				newstr += fmt.Sprintf("%02v", t.Hour())
			case "%I":
				// Hour (00-12)
				newstr += fmt.Sprintf("%02v", t.Hour() - (t.Hour()/12)*12)
			case "%M":
				// Minute (00-59)
				newstr += fmt.Sprintf("%02v", t.Minute())
			case "%S":
				// Second (00-59)
				newstr += fmt.Sprintf("%02v", t.Second())
			case "%p":
				// AM/PM
				if t.Hour()/12 > 0{
					newstr += "PM"
				} else {
					newstr += "AM"
				}
			case "%n":
				newstr += "\n"
			case "%t":
				newstr += "\t"
			case "%%":
				newstr += "%"
			default:
				newstr += ""
		}
	}
	
	if len(newstr) > 0{
		newstr += layout[pos:]
		layout = newstr
	}

	return layout
}