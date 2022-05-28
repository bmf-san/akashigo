package types

// TypeStamp is a type for stamp.
type TypeStamp int

const (
	// Attendance is a type of Attendance.
	Attendance TypeStamp = 11
	// LeaveWork is a type of leaving work.
	LeaveWork TypeStamp = 12
	// GoStraight is a type of going straight.
	GoStraight TypeStamp = 21
	// Break is a type of break.
	Break TypeStamp = 31
	// BreakBack is a type of break back.
	BreakBack TypeStamp = 32
)

var stampText = map[TypeStamp]string{
	Attendance: "出勤",
	LeaveWork:  "退勤",
	GoStraight: "直行",
	Break:      "休憩入",
	BreakBack:  "休憩戻",
}

// StampText returns a text of stamp.
func StampText(t TypeStamp) string {
	return stampText[t]
}
