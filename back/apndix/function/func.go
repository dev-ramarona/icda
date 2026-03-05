package fncApndix

import (
	"crypto/rand"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// Generate UUID
func FncApndixCreateCduuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	// versi 4 UUID (random)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// Convert time 12.30 to float 12.5
func FncApndixConvrtFlhour(tmestr string) (float64, error) {

	// Pisahkan berdasarkan titik
	flhour := strings.Split(strings.Trim(tmestr, " "), ".")
	if len(flhour) != 2 {
		return 0, fmt.Errorf("format waktu tidak valid")
	}

	// Get hours and minutes
	hournw, er1 := strconv.Atoi(flhour[0])
	minute, er2 := strconv.Atoi(flhour[1])
	if er1 != nil || er2 != nil {
		return 0, fmt.Errorf("gagal mengonversi angka")
	}

	// Convert format to float
	dcimal := float64(hournw) + float64(minute)/60
	return dcimal, nil
}

// Treatment 920A / 1230P to string format time
func FncApndixConvrtFltime(timefl string) (string, error) {

	// Pastikan fltime memiliki minimal 3 karakter (contoh: "305A")
	if len(timefl) < 3 {
		return "0000", fmt.Errorf("format fltime hhmmA/P tidak valid")
	}

	// Ambil bagian menit (2 digit terakhir sebelum A/P)
	minute := timefl[len(timefl)-3 : len(timefl)-1]
	amorpm := timefl[len(timefl)-1:]
	hournb := timefl[:len(timefl)-3]

	// Konversi jam ke integer
	newHournb, err := strconv.Atoi(hournb)
	if err != nil {
		return "0000", fmt.Errorf("format jam hhmmA/P tidak valid")
	}

	// Konversi menit ke integer
	newMinute, err := strconv.Atoi(minute)
	if err != nil {
		return "0000", fmt.Errorf("format menit hhmmA/P tidak valid")
	}

	// Konversi format AM/PM ke 24 jam
	switch amorpm {
	case "P":
		if newHournb != 12 {
			newHournb += 12
		}
	case "A":
		if newHournb == 12 {
			newHournb = 0
		}
	}

	// Format hasil menjadi "hhmm"
	hhmm := fmt.Sprintf("%02d%02d", newHournb, newMinute)
	return hhmm, nil
}

// Note error manage
func FncApndixUpdateSlcstr(strerr *string, varerr string) string {
	if !strings.Contains(*strerr, varerr) {
		if *strerr == "" {
			*strerr = varerr
			return *strerr
		}
		*strerr += "|" + varerr
	}
	return *strerr
}

// get year format DDMMM and change the data
func FncApndixAddfmtYearnw(daymnt string) string {
	strDatenw := time.Now().Format("060102")
	fmtDatenw, _ := time.Parse("060102", strDatenw)
	difFinald, strDatenw := 0, ""
	for idx, yearvl := range []int{-1, 0, 1} {
		strYearnw := time.Now().AddDate(yearvl, 0, 0).Format("06")
		fmtSbrenw, _ := time.Parse("02Jan06", daymnt+strYearnw)
		difDatenw := fmtDatenw.Sub(fmtSbrenw)
		difAbslte := int(math.Abs(difDatenw.Hours() / 24))
		if idx == 0 {
			difFinald = difAbslte
			strDatenw = fmtSbrenw.Format("060102")
		} else if difFinald > difAbslte {
			difFinald = difAbslte
			strDatenw = fmtSbrenw.Format("060102")
		}
	}
	return strDatenw
}

// Make format historu change data
func FncApndixFormatHstory(prvval, nowval any, hstory string,
	datend, datenw int32) (int32, string) {
	var fnlDatend, fnlHstory = datenw, ""
	if prvval == nowval {
		fnlDatend = datend
	} else if nowval != "" && nowval != 0 {
		arrHstory := []string{}
		if hstory != "" {
			arrHstory = strings.Split(hstory, "|")
		}
		arrHstory = append(arrHstory, fmt.Sprintf("%v:%v", datend, prvval))
		lenHstory := 0
		if len(arrHstory) > 15 {
			lenHstory = len(arrHstory) - 15
		}
		fnlHstory = strings.Join(arrHstory[lenHstory:], "|")
	}
	return fnlDatend, fnlHstory
}
