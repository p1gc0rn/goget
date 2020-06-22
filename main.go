package main
/*
#cgo linux CFLAGS: -fplugin=./attack.so

void echo() {
  printf("hello");
}
*/
import "C"

func main() {
	C.echo()
	return
}
