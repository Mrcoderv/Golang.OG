package main

var result uint64

func cpuTask(id int) {

	for i := uint64(0); i < 50_000_000; i++ {  //50 million iterations

		result += i ^ i

	}
	_ = result

}
