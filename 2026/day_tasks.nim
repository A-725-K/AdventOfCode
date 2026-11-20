import std/tables
export tables

type DayTaskProc* = proc(lines: seq[string], day: int)

var tasks* = initTable[string, DayTaskProc]()
