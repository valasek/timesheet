import { parseISO, isWeekend } from 'date-fns'

export const workHoursMixin = {

  methods: {
    // requires on parent component to define:
    // this.holidays
    // return number of holidays out of working days
    getHolidays (from, to) {
      const h = this.holidays.filter(d => parseISO(d.date) > from && parseISO(d.date) < to)
      let w = 0
      h.forEach(function (e) {
        if (isWeekend(parseISO(e.date))) {
          w = w + 1
        }
      })
      return h.length - w
    }
  }
}
