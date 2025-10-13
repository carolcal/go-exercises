package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, c:= range log {
        switch c {
            case 10071:		return "recommendation"
            case 128269:	return "search"
            case 9728:		return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log to the caller.
func Replace(log string, oldRune, newRune rune) string {
    newLog := ""
	for _, c:= range log {
        if c == oldRune{
            newLog += string(newRune)
        } else {
            newLog += string(c)
        }
    }
    return newLog
}

// WithinLimit determines whether or not the number of characters in log is within the limit.
func WithinLimit(log string, limit int) bool {
	count := 0
    for range log {
        count ++
    }
    if count <= limit {
        return true
    }
    return false
}
