// 1154. Day of the Year
// Given a string date representing a Gregorian calendar date formatted as YYYY-MM-DD, return the day number of the year.

 

// Example 1:

// Input: date = "2019-01-09"
// Output: 9
// Explanation: Given date is the 9th day of the year in 2019.
// Example 2:

// Input: date = "2019-02-10"
// Output: 41
 

// Constraints:

// date.length == 10
// date[4] == date[7] == '-', and all other date[i]'s are digits
// date represents a calendar date between Jan 1st, 1900 and Dec 31th, 2019.

func dayOfYear(date string) int {
    dateArr := strings.Split(date, "-")
    year, _ := strconv.Atoi(dateArr[0])
    monthNum, _ := strconv.Atoi(dateArr[1])
    month := time.Month(monthNum)
    day, _ := strconv.Atoi(dateArr[2])
    t := time.Date(year, month, day, 12, 00, 00, 0, time.UTC)
    yearDay := t.YearDay()
    return yearDay 
}
