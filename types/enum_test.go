package types

import (
	"fmt"
	"testing"
)

func Test_Status(t *testing.T) {
	fmt.Printf("STATUS_DISABLE: %T, %[1]d; DISABLE: %[2]T, %[2]d\n", STATUS_DISABLE, DISABLE)
	fmt.Printf("STATUS_ENABLE: %T, %[1]d; ENABLE: %[2]T, %[2]d\n", STATUS_ENABLE, ENABLE)
}

func Test_Switch(t *testing.T) {
	fmt.Printf("SWITCH_OFF: %T, %[1]d; OFF: %[2]T, %[2]d\n", SWITCH_OFF, OFF)
	fmt.Printf("SWITCH_ON: %T, %[1]d; ON: %[2]T, %[2]d\n", SWITCH_ON, ON)
}

func Test_Week(t *testing.T) {
	fmt.Printf("SUNDAY: %T, %[1]d; SUM: %[2]T, %[2]d;\n", SUNDAY, SUM)
	fmt.Printf("MONDAY: %T, %[1]d; MON: %[2]T, %[2]d;\n", MONDAY, MON)
	fmt.Printf("TUESDAY: %T, %[1]d; TUE: %[2]T, %[2]d;\n", TUESDAY, TUE)
	fmt.Printf("SEDNESDAY: %T, %[1]d; SED: %[2]T, %[2]d;\n", SEDNESDAY, SED)
	fmt.Printf("THURSDAY: %T, %[1]d; THU: %[2]T, %[2]d;\n", THURSDAY, THU)
	fmt.Printf("FRIDAY: %T, %[1]d; FRI: %[2]T, %[2]d;\n", FRIDAY, FRI)
	fmt.Printf("SATUREDAY: %T, %[1]d; SAT: %[2]T, %[2]d;\n", SATUREDAY, SAT)
}

func Test_Month(t *testing.T) {
	fmt.Printf("JANUARY: %T, %[1]d;\n", JANUARY)
	fmt.Printf("FEBRUARY: %T, %[1]d;\n", FEBRUARY)
	fmt.Printf("MARCH: %T, %[1]d;\n", MARCH)
	fmt.Printf("APRIL: %T, %[1]d;\n", APRIL)
	fmt.Printf("MAY: %T, %[1]d;\n", MAY)
	fmt.Printf("JUNE: %T, %[1]d;\n", JUNE)
	fmt.Printf("JULY: %T, %[1]d;\n", JULY)
	fmt.Printf("AUGUST: %T, %[1]d;\n", AUGUST)
	fmt.Printf("SEPTEMBER: %T, %[1]d;\n", SEPTEMBER)
	fmt.Printf("OCTOBER: %T, %[1]d;\n", OCTOBER)
	fmt.Printf("NOVEMBER: %T, %[1]d;\n", NOVEMBER)
	fmt.Printf("DECEMBER: %T, %[1]d;\n", DECEMBER)
}

func Test_GENDER(t *testing.T) {
	fmt.Printf("GENDER_FEMALE: %T, %[1]d\n", GENDER_FEMALE)
	fmt.Printf("GENDER_MALE: %T, %[1]d\n", GENDER_MALE)
	fmt.Printf("GENDER_UNKNOWN: %T, %[1]d\n", GENDER_UNKNOWN)
}

func Test_MEMSIZE(t *testing.T) {
	fmt.Printf("MEM_SIZE_BYTE: %T, %[1]d BYTE: %[2]T, %[2]d\n", MEM_SIZE_BYTE, BYTE)
	fmt.Printf("MEM_SIZE_KB: %T, %[1]d; KB: %[2]T, %[2]d\n", MEM_SIZE_KB, KB)
	fmt.Printf("MEM_SIZE_MB: %T, %[1]d; MB: %[2]T, %[2]d\n", MEM_SIZE_MB, MB)
	fmt.Printf("MEM_SIZE_GB: %T, %[1]d; GB: %[2]T, %[2]d\n", MEM_SIZE_GB, GB)
	fmt.Printf("MEM_SIZE_TB: %T, %[1]d; TB: %[2]T, %[2]d\n", MEM_SIZE_TB, TB)
	fmt.Printf("MEM_SIZE_PB: %T, %[1]d; PB: %[2]T, %[2]d\n", MEM_SIZE_PB, PB)
}
