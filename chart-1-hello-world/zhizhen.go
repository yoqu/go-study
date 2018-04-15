package main

import "fmt"

func main()  {
  x :=1
  p :=&x
  fmt.Println(*p)
  *p=2
  fmt.Println(x)
  i :=1
  fmt.Println(incr(&i))
  fmt.Println(incr(&i))
}

func incr(p *int) int{
  *p++
  return *p
}
