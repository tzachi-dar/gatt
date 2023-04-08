package service

import (
	"fmt"
	"log"
	"time"
	"bufio"
	"os"
        b64 "encoding/base64"
        "encoding/hex"

	"github.com/paypal/gatt"
)


func fromInts(data  []int) string {
  var s  string
  for _, d := range data {
	    if(d < 0) {
		    d+= 256;
		}
	    hh := fmt.Sprintf("%02x", d )
		s+=hh
	    //fmt.Println(d, hh)
	}
  
  
  return s
}




func readData() []string {
    file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    ret:= []string{};

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
//        fmt.Println(scanner.Text())
	ret = append(ret, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
//    fmt.Println(ret)
    return ret
    
}


func NewCountService() *gatt.Service {
	//n := 0
	s := gatt.NewService(gatt.MustParseUUID("0000fde3-0000-1000-8000-00805f9b34fb"))
        log.Println("aaaa", gatt.MustParseUUID("0000fde3-0000-1000-8000-00805f9b34fb"));
        
	//s.AddCharacteristic(gatt.MustParseUUID("0000f001-0000-1000-8000-00805f9b34fb")).HandleReadFunc(
	//	func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
	//		fmt.Fprintf(rsp, "count: %d", n)
	//		n++
	//	})

	s.AddCharacteristic(gatt.MustParseUUID("0000f001-0000-1000-8000-00805f9b34fb")).HandleWriteFunc(
		func(r gatt.Request, data []byte) (status byte) {
                        uEnc := b64.URLEncoding.EncodeToString([]byte(data))
                        fmt.Println(uEnc)
			//log.Println("Wrote:", string(data))
			return gatt.StatusSuccess
		})

	s.AddCharacteristic(gatt.MustParseUUID("0000f002-0000-1000-8000-00805f9b34fb")).HandleNotifyFunc(
		func(r gatt.Request, n gatt.Notifier) {
//                               	data := []int {-51, -52, 119, -121, 17, -78, -86, 77, -102, 108, -73, -80, -20, -44, -69, 27, -118, 59, 91, 16, 104, 80, 46, 3, -107, -39, 22, 16, -36, -74, 2, -50, -6, 123, -117, -112, 78, 12, 89, 65, -90, 26, -7, -68, -118, 97}

                        	board := []string{
//"d780b09d3bbf6f9c3eae44f7cfed78d6d68f0b2b9fab08eaf5b5404e49c60b2b51dfa73e80724c19b4a95255d4fe",
					"00002299d1093d164bb59cee6d7ce0e29b7f2b1eebdef75099a1ee60271ff40ea2f799bacfac460ff87501939e4c",
					"0000929e750822164bb5eeee6f7c1eed9b7f6f5ef4def710963116c25daef40ea2f799bacfac460ff8750693c7cb",
					"000092de72088fd1e9b4e2ee687c356d9a7f9151f4def750d83deb5c08bdf40ea2f799bacfac460ff87507931078",
					"000092de72083fd64db5fdee687c476d987fbad1f5def750d83dee607c4ff40ea2f799bacfac460ff8750493b87e",
					"0000925e70083f964ab55029ca7d4b6d9f7fc8d1f7de1e98733cee2073dff40ea2f799bacfac460ff8750593fcf3",
					"0000955e73083f964ab5e02e6e7c546d9f7fc4d1f0deaad8693cee603dd3f40ea2f799bacfac460ff8751a93a130",
					"0000985e71083f1648b5e06e697cf9aa3d7edbd1f0deee98763cee603dd3f806d9df9403103cf439a1441b93b91f",
					"0000959e720838164bb5e06e697c49ad997f761652df1097763c07a896d2f806d9df9403103cf439a1441893f850",
					"00009e5e7008351649b5e0ee6b7c49ed9e7fc611f6de3b17773cb3e88cd2f806d9df9403103cf439a1441993fcdd",
					"00009e5e710838d64ab5e7ee687c49ed9e7fc651f1de4917753cf7a893d2f806d9df9403103cf439a1441e93a9e6",		
//                                    fromInts(data),
	                        }
				board = readData()
			cnt64 := time.Now().UnixNano() / int64(time.Second)/ 60
                        cnt := int(cnt64)
                        log.Println("cnt = ", cnt)
			for !n.Done() {
				//fmt.Fprintf(n, "Count: %d", cnt)
                                //arr :=[46]byte{'a', 'b'}
                                //arr[0] = 1
                                //arr[1] = 200
                                //arr[2] = 100
                                //ss := string(arr[:])
	                        //s := "a2f2d2bad43d7e0804c3e7a13ff4a4966e9058e3ab1fe7214d2a2311596e789a0ba14b7da5382f8eb252e0f96138"
                                s := board[cnt % len(board)]

                               	data, err := hex.DecodeString(s)
	                        if err != nil {
	                            panic(err)
	                        }
                                ss := string(data[:20])
                                fmt.Fprintf(n, ss)
                                fmt.Fprintf(n, string(data[20:38]))
                                fmt.Fprintf(n, string(data[38:46]))
				log.Println("Writing data")
 
				cnt++
				time.Sleep(time.Second * 60)
			}
	})

	return s
}
