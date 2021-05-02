package main

import "fmt"

func printFactorsWorker(workerId int, jobs <-chan int64, results chan<- int64) {
	for nr := range jobs {
		fmt.Printf("Worker ID : %d; Processing : %d\n", workerId, nr)
		if nr < 1 {
			fmt.Println("\nFactors of", nr, "not computed")
			return
		}

		fmt.Printf("\nFactors of %d: ", nr)
		fs := make([]int64, 1)
		fs[0] = 1
		apf := func(p int64, e int) {
			n := len(fs)
			for i, pp := 0, p; i < e; i, pp = i+1, pp*p {
				for j := 0; j < n; j++ {
					fs = append(fs, fs[j]*pp)
				}
			}
		}
		e := 0
		for ; nr&1 == 0; e++ {
			nr >>= 1
		}
		apf(2, e)
		for d := int64(3); nr > 1; d += 2 {
			if d*d > nr {
				d = nr
			}
			for e = 0; nr%d == 0; e++ {
				nr /= d
			}
			if e > 0 {
				apf(d, e)
			}
		}
		fmt.Println(fs)
		fmt.Println("Number of factors =", len(fs))
		fmt.Println()
		results <- int64(len(fs))
	}
}

func main() {
	// var input int64 = 134217728
	// var input int64 = 262144
	var input int64 = 128
	var x int64
	const WORKER_TOTAL = 100 // - total of worker that will execute the job

	jobs := make(chan int64, input)
	results := make(chan int64, input)

	// - worker waiting for assignment
	for w := 1; w <= WORKER_TOTAL; w++ {
		go printFactorsWorker(w, jobs, results) // - assign the job to the worker
	}

	for x = 1; x <= input; x++ {
		jobs <- x
	}
	close(jobs)

	counter := 0
	for x := 1; x <= int(input); x++ {
		res := <-results
		if res == 6 {
			counter++
		}
	}

	fmt.Println("Total : ", counter)
}
