package types

/*
类型: 启用状态
*/
const (
	STATUS_DISABLE, DISABLE uint8 = iota, iota
	STATUS_ENABLE, ENABLE
)

/*
类型: 状态开关
*/
const (
	SWITCH_OFF, OFF uint8 = iota, iota
	SWITCH_ON, ON
)

/*
类型: 星期
*/
const (
	SUNDAY, SUM uint8 = iota, iota
	MONDAY, MON
	TUESDAY, TUE
	SEDNESDAY, SED
	THURSDAY, THU
	FRIDAY, FRI
	SATUREDAY, SAT
)

const (
	JANUARY uint8 = iota + 1
	FEBRUARY
	MARCH
	APRIL
	MAY
	JUNE
	JULY
	AUGUST
	SEPTEMBER
	OCTOBER
	NOVEMBER
	DECEMBER
)

/*
类型: 性别
*/
type Gender uint8

const (
	GENDER_FEMALE  Gender = iota + 1 // 男
	GENDER_MALE                      // 女
	GENDER_UNKNOWN                   // 未知
)

/*
类型: 内存容量
*/
type MemSize uint64

const (
	MEM_SIZE_BYTE, BYTE MemSize = 1 << (iota * 10), 1 << (iota * 10)
	MEM_SIZE_KB, KB
	MEM_SIZE_MB, MB
	MEM_SIZE_GB, GB
	MEM_SIZE_TB, TB
	MEM_SIZE_PB, PB
)
