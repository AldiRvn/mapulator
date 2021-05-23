# mapulator / Map Manipulator
Used for set value of a field in nested key map

Example:
```golang
splitBySpace := strings.Split("Robert Juan May Thornton", " ")
newMap := map[string]interface{}{}

path := "user"
Set(newMap, path+".age", 48)
Set(newMap, path+".active", true)
Set(newMap, path+".hobby", []string{"Fishing", "Drawing"})

path += ".name"
Set(newMap, path, map[string]interface{}{
  "first": splitBySpace[0],
  "last":  splitBySpace[len(splitBySpace)-1],
})
Set(newMap, path+".full", map[string]interface{}{
  "length": len(value),
  "value":  value,
})
```
