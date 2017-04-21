
```go
for py := 0; py < height; py++ {
	y := float64(py)/height*(ymax-ymin) + ymin
	fmt.Printf("py=%v y=%v \n", py, y)
}
```

```
py=0 y=-2 
py=1 y=-1.99609375 
py=2 y=-1.9921875 
py=3 y=-1.98828125 
py=4 y=-1.984375 
py=5 y=-1.98046875 
```