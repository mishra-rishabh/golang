# Time Package

The time package in Go provides functions for measuring and displaying time. It is a standard library package that is included in all Go installations.


## Functions For Measuring Time

* **`Now()`:** Returns the current time as a `time.Time` value.
* **`Since()`:** Returns the elapsed time between two `time.Time` values.
* **`Until()`:** Returns the time remaining until a specified `time.Time` value.
* **`Add()`** and **`Sub()`:** Add or subtract a duration from a `time.Time` value.


## Functions For Displaying Time

* **`Format()`** and **`Parse()`**: Format and parse a `time.Time` value according to a given layout.
* **`Weekday()`** and **`Month()`**: Return the weekday and month of a `time.Time` value.
* **`Hour()`**, **`Minute()`**, and **`Second()`**: Return the hour, minute, and second of a `time.Time` value.

**NOTE:** When using `Format()` function we have to give layout in which we want the time. <br/>
The string `2006-01-02 15:04:05`, is a layout string that can be used with the `Format()` function in the `time` package in Go to format a `time.Time` value as follows:
```
YYYY-MM-DD HH:MM:SS
```
where: <br/>

* **YYYY** is the four-digit year, represented by `2006`.
* **MM** is the two-digit month, represented by `01`.
* **DD** is the two-digit day of the month, represented by `02`.
* **HH** is the two-digit hour of the day (in 24-hour format), represented by `15`.
* **MM** is the two-digit minute of the hour, represented by `04`.
* **SS** is the two-digit second of the minute, represented by `05`.
* Layout for day name is `Monday`

We can try multiple layout with reference to above given layout like `02-01-2006 15:04:05 Monday`, `2006-01-02 15:04:05 Monday`


## Package Link

[Reference Link for time Package](https://pkg.go.dev/time)