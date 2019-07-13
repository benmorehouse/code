package main

import(
	"io/ioutil"
	"os"
	"bufio"
)

type cron struct{ // these conditions will be written to a file and held
	filename string
	condition int // 0 if initial state, 1 if it is initialized to something and non-binary if it is broken
	condition_name string
}


//This will reset cron to the initial stance that it was supposed to be before... call this many a times!

func reset_cron()error{
	var reset_val string
	reset_val = "cron.txt\n0\nINITIAL_CONDITION"
	writer := ioutil.WriteFile("cron_conditions.txt",[]byte(reset_val),0566)
	return writer
}

func update(reset_val string)error{
	writer := ioutil.WriteFile("cron_conditions.txt",[]byte(reset_val),0566)
	return writer
}

func (scheduler cron) parse_cron()(error){
	// this will go in and make the neccessary edits to cron like we wish 
	// pass in the file flag which we will use as t
	file ,err = os.Open("cron.txt")
	var cron_reader_writers = [3] bufio.NewReadWriter(file,file) // an array of scanners for each of the goroutines
	newscraper := scheduler.condition_name //updating tool
	var cron_channel =  [3]chan string {
		for i:= range cron_channel{
			cron_channel[i] = make(chan string)
		}
	}

	go func(){
		for i:=0;i<13;i++{
			cron_reader_writers[0].Reader.Scan()
		} // will scan to line 14
		cron_reader_writers[0].Writer.Write([]byte("/*")
		for temp != ""{


		}()

	select{
		case: <-

}



