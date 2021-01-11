/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func main() {

	var nilai = 70
	var absensi = 80

	var lulusNilai = nilai >= 70
	var lulusAbsensi = absensi >70

	var lulus = lulusNilai && lulusAbsensi
	fmt.Println(lulus)

}
