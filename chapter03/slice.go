package main

import "fmt"

func main() {
	var slicea = []int{10,20,30}
	fmt.Println(slicea)
	fmt.Println(len(slicea))

	var sliceb = []int{1, 5:4, 6, 10:100, 15}
	fmt.Println(sliceb)
	fmt.Println(len(sliceb))

	sliceb[0] = 10
	fmt.Println(sliceb)
	fmt.Println(len(sliceb))

	var slicec []int
	fmt.Println(slicec)
	fmt.Println(slicec == nil)

	// append and cao
	var sliced = []int{1, 2, 3}
	fmt.Println(sliced)
	fmt.Println(len(sliced))
	fmt.Println(cap(sliced))

	sliced = append(sliced, 4)
	fmt.Println(sliced)
	fmt.Println(len(sliced))

	sliced = append(sliced, 5, 6, 7)
	fmt.Println(sliced)
	fmt.Println(len(sliced))
	fmt.Println(cap(sliced))

	slicee := []int{20, 30, 40}
	fmt.Println(slicee)
	fmt.Println(len(slicee))
	sliced = append(sliced, slicee...)
	fmt.Println(sliced)
	fmt.Println(len(sliced))
	fmt.Println(cap(sliced))

	// make
	fmt.Println("\tThese is \"slicef\"")
	slicef := make([]int, 5)
	fmt.Println(slicef)
	fmt.Println(len(slicef))
	fmt.Println(cap(slicef))

	slicef = append(slicef, 10)
	fmt.Println(slicef)
	fmt.Println(len(slicef))
	fmt.Println(cap(slicef))

	sliceg := make([]int, 5, 10)
	fmt.Println("\tThese is \"sliceg\"")
	fmt.Println(sliceg)
	fmt.Println(len(sliceg))
	fmt.Println(cap(sliceg))

	sliceg = append(sliceg, 10)
	fmt.Println(sliceg)
	fmt.Println(len(sliceg))
	fmt.Println(cap(sliceg))

	// slice of slice
	sliceh := []int{1, 2, 3, 4}
	slicei := sliceh[:2]
	slicej := sliceh[1:]
	slicek := sliceh[1:3]
	slicel := sliceh[:]
	fmt.Println("sliceh:", sliceh)
	fmt.Println("slicei:", slicei)
	fmt.Println("slicej:", slicej)
	fmt.Println("slicek:", slicek)
	fmt.Println("slicel:", slicel)

	fmt.Println("Change sliceh")
	sliceh[1] = 20
	fmt.Println("sliceh:", sliceh)
	fmt.Println("slicei:", slicei)
	fmt.Println("slicej:", slicej)
	fmt.Println("slicek:", slicek)
	fmt.Println("slicel:", slicel)

	fmt.Println("Change slicei")
	slicei[0] = 10
	fmt.Println("sliceh:", sliceh)
	fmt.Println("slicei:", slicei)
	fmt.Println("slicej:", slicej)
	fmt.Println("slicek:", slicek)
	fmt.Println("slicel:", slicel)

	fmt.Println("Change slicej")
	slicej[1] = 30
	fmt.Println("sliceh:", sliceh)
	fmt.Println("slicei:", slicei)
	fmt.Println("slicej:", slicej)
	fmt.Println("slicek:", slicek)
	fmt.Println("slicel:", slicel)

	// array to slice
	arrayx := [...]int{5, 6, 7, 8}
	slicem := arrayx[:2]
	slicen := arrayx[2:]
	sliceo := arrayx[:]
	arrayx[0] = 10

	fmt.Println("arrayx:", arrayx)
	fmt.Println("slicem:", slicem)
	fmt.Println("slicen:", slicen)
	fmt.Println("sliceo:", sliceo)

    // copy slice without memory share
    slicep := []int{1, 2, 3, 4}
    sliceq := make([]int, 2)
    fmt.Println(slicep, sliceq)

    copy(sliceq, slicep)
    fmt.Println(slicep, sliceq)

    slicep[0] = 10
    fmt.Println(slicep, sliceq)

    sliceq[1] = 20
    fmt.Println(slicep, sliceq)
}
