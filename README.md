# mapulator / Map Manipulator
Used for set value of a field in nested key map, this repo was inspired by [sjson](github.com/tidwall/sjson) and [jsonparser](github.com/buger/jsonparser)

## Installation
`go get github.com/AldiRvn/mapulator`

## Usage:
```golang
value := "Chris Casey"
splitBySpace := strings.Split(value, " ")
newMap := map[string]interface{}{}

path := "user"
mapulator.Set(newMap, path+".age", 48)
fmt.Println(newMap) // map[user:map[age:48]]
mapulator.Set(newMap, path+".active", true)
fmt.Println(newMap) // map[user:map[active:true age:48]]
mapulator.Set(newMap, path+".hobby", []string{"Fishing", "Drawing"})
fmt.Println(newMap) // map[user:map[active:true age:48 hobby:[Fishing Drawing]]]

path += ".name"
mapulator.Set(newMap, path, map[string]interface{}{
  "first": splitBySpace[0],
  "last":  splitBySpace[len(splitBySpace)-1],
})
fmt.Println(newMap) // map[user:map[active:true age:48 hobby:[Fishing Drawing] name:map[first:Chris last:Casey]]]
mapulator.Set(newMap, path+".full", map[string]interface{}{
  "length": len(value),
  "value":  value,
})
fmt.Println(newMap) // map[user: ... full:map[length:11 value:Chris Casey] last:Casey]]]
```
