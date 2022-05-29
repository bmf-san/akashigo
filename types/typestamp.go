package types

// TypeStamp is a type for stamp.
type TypeStamp int

const (
	// Attendance is a type of attendance.
	Attendance TypeStamp = 11
	// LeaveWork is a type of leaving work.
	LeaveWork TypeStamp = 12
	// GoStraight is a type of going straight.
	GoStraight TypeStamp = 21
	// Break is a type of break.
	Break TypeStamp = 31
	// BreakBack is a type of break back.
	BreakBack TypeStamp = 32

	// TextAttendance is a text of attendance.
	TextAttendance = "出勤"
	// TextLeaveWork is a text of leaving work.
	TextLeaveWork = "退勤"
	// TextGoStraight is a text of going straight.
	TextGoStraight = "直行"
	// TextBreak is a text of break.
	TextBreak = "休憩入"
	// TextBreakBack is a text of break back.
	TextBreakBack = "休憩戻"
)

var stampText = map[TypeStamp]string{
	Attendance: TextAttendance,
	LeaveWork:  TextLeaveWork,
	GoStraight: TextGoStraight,
	Break:      TextBreak,
	BreakBack:  TextBreakBack,
}

var stampNumber = map[string]TypeStamp{
	TextAttendance: Attendance,
	TextLeaveWork:  LeaveWork,
	TextGoStraight: GoStraight,
	TextBreak:      Break,
	TextBreakBack:  BreakBack,
}

// StampText returns a text of stamp.
func StampText(t TypeStamp) string {
	return stampText[t]
}

func StampNumber(s string) TypeStamp {
	return stampNumber[s]
}
