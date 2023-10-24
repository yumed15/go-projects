package parser

type expressionRange int

const (
	minuteRange     expressionRange = iota
	hourRange       expressionRange = iota
	dayOfMonthRange expressionRange = iota
	monthRange      expressionRange = iota
	dayOfWeekRange  expressionRange = iota
)

func (r expressionRange) getMin() int {
	switch r {
	case minuteRange, hourRange:
		return 0
	case dayOfWeekRange, dayOfMonthRange, monthRange:
		return 1
	default:
		panic("invalid expression")
	}
}

func (r expressionRange) getMax() int {
	switch r {
	case minuteRange:
		return 59
	case hourRange:
		return 23
	case dayOfWeekRange:
		return 7
	case dayOfMonthRange:
		return 31
	case monthRange:
		return 12
	default:
		panic("invalid expression")
	}
}

func (r expressionRange) formatStepValue(value int) int {
	switch r {
	case minuteRange:
		if value > 60 {
			return r.getMin()
		}
	case hourRange:
		if value > 24 {
			return r.getMin()
		}
	case dayOfWeekRange:
		if value > 7 {
			return r.getMin()
		}
	case dayOfMonthRange:
		if value > 12 {
			return r.getMin()
		}
	case monthRange:
		if value > 31 {
			return r.getMin()
		}
	default:
		panic("invalid expression")
	}
	return value
}
