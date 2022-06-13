package goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

/* Proses transfer data pada channel secara default dilakukan dengan cara un-buffered
Buffered channel sedikit berbeda. Pada channel jenis ini, ditentukan angka jumlah buffer-nya
Selama jumlah data yang dikirim tidak melebihi jumlah buffer,
maka pengiriman akan berjalan asynchronous (tidak blocking)

Ketika jumlah data yang dikirim sudah melewati batas buffer
maka pengiriman data hanya bisa dilakukan ketika salah satu data yang sudah terkirim
sehingga ada slot channel yang kosong

Pengiriman data indeks ke 0, 1, 2 dan 3 akan berjalan secara asynchronous,
hal ini karena channel ditentukan nilai buffer-nya sebanyak 3 (ingat, jika nilai buffer adalah 3, maka 4 data yang akan di-buffer).
*/

func TestBufferedChannel(t *testing.T) {

	runtime.GOMAXPROCS(2)

	messages := make(chan int, 3)

	//anonymous
	go func() {
		for {
			i := <-messages
			fmt.Println("menerima data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("mengirim data", i)
		messages <- i
	}

}
