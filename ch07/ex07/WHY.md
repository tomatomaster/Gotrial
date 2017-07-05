func CelsiusFlagのデフォルト値の型はCelsius型だから。

```
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```
