# Maps In Golang

Go provides another important data type named map which maps unique keys to values.

* A `key` is an object that you use to retrieve a `value` at a later date.
* Given a key and a value, you can store the value in a Map object.
* After the value is stored, you can retrieve it by using its key.
* Maps are also called Go's built-in associative data type (sometimes called hashes or dicts in other languages).


## How To Define A Map ?

* Like slices, `make()` function is used to create a map as well.

    **Syntax:**
    ```
    // create map
    make(map[key-type]val-type)

    // set key/value pairs
    name[key] = val
    ```

    **example:**
    ```go
    studentsAndRollNum := make(map[int]string)

    studentsAndRollNum[23] = "Akshay"

    // print whole map
    fmt.Println("\nStudents and roll num:", studentsAndRollNum)

    // print value for a specific key
    fmt.Println("Student for roll num 23:", studentsAndRollNum[23])
    ```

* We can also declare and initialize a new map in the same line.
    
    **Syntax:**
    ```
    // create map
    map[key-type]val-type{key1: val1, key2: val2}
    ```

    **example:**
    ```go
    countryCapital := map[string]string{"India": "New_Delhi", "Japan": "Tokyo"}

    fmt.Println("Country and capital:", countryCapital)
    ```
* If the key doesn’t exist, the zero value of the value type is returned.
    - zero value for `int` is `0`.
    - zero value for `bool` is `false`.
    - zero value for `string` is `""`.


## `len()`
The builtin `len()` function returns the number of **key/value pairs** when called on a map.

**example:**
```go
lengthVar := len(studentsAndRollNum)
fmt.Println("studentsAndRollNum length:", lengthVar)
```


## Deleting A Value From A Map 

* `delete()` function is used to delete an entry from a map.
* It requires the map and the corresponding key which is to be deleted.

**Syntax**
```
delete(map, "key-name")
```

**example:**
```go
// this will delete an entry from a map where key is 23
delete(studentsAndRollNum, 23)
```


## Delete A Whole Map
To remove all *key/value* pairs from a map, use the `clear()` builtin function.

**example:**

```go
// this will clear the whole map
clear(studentsAndRollNum)
```


## Loop Through A Map
* We can use loop to print a whole map

    **example:**
    ```go
    for key, val := range studentsAndRollNum {
        fmt.Printf("key: %v, value: %v\n", key, val)
    }
    ```
* We can also ignore the *key* while looping through a map by using **comma ok syntax**.

    **example:**
    ```go
    for _, val := range studentsAndRollNum {
        fmt.Printf("value: %v\n", key, val)
    }
    ```


## Checking If A Value Is Present In A Map

* For this can use **comma ok syntax**.
* The optional second return value when getting a value from a map indicates if the key was present in the map or not.
* This can be used to disambiguate between missing keys and keys with zero values like `0` or `""`. 

    **example:**
    ```go
    _, present := studentsAndRollNum[67]
    fmt.Println("Is Present", present) // false
    ```
