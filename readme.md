# Stockpile

This is a example library (with tests) showing how to modify an existing slice type when a pointer of the type is the receiver type of a function.

I often find myself wanting to iterate a slice and change the values of the type, and forget how to do that in a range loop.

`hint:`
```
for k, v := range *s {
    if v.Name == stock.Name {
        (*s)[k].Quantity += stock.Quantity
        added = true
    }
}
```