// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"
	"log"

	"bufio"
	"os"

	"strings"

	"strconv"

	"./eval"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		formula := scanner.Text()
		expr, err := eval.Parse(formula)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Enter the variation {var}={val}")
		fmt.Println("Example: x=3")
		envSeed := make(map[eval.Var]float64)
		scanner.Scan()
		text := scanner.Text()
		for _, v := range strings.Split(text, ",") {
			varAndval := strings.Split(v, "=")
			if len(varAndval) != 2 {
				log.Fatal("Please input folloing order {var}={val}")
			}
			key := varAndval[0]
			val := varAndval[1]
			intVal, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			envSeed[eval.Var(key)] = float64(intVal)
		}
		env := eval.Env(envSeed)
		answer := expr.Eval(env)
		fmt.Printf("%s = %f\n", formula, answer)
	}
}
