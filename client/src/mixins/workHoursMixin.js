import { parseISO } from 'date-fns'

export const workHoursMixin = {

  methods: {
    // requires on parent component to define:
    // this.holidays
    getHolidays (from, to) {
      return this.holidays.filter(d => parseISO(d.date) > from && parseISO(d.date) < to).length
      // switch (period) {
      //   case 'week':

      //   case 'month':
      //     if (yesterday === null) {
      //       return this.holidays.filter(d => parseISO(d.date).getMonth() === m).length * this.dailyWorkingHours
      //     } else {
      //       const yesterday = addDays(new Date(), -1)
      //       const beginning = startOfMonth(yesterday)
      //       return this.holidays.filter(d => parseISO(d.date) >= beginning && parseISO(d.date) <= yesterday).length * this.dailyWorkingHours
      //     }
      //   default:
      //     console.log('workingTimeOverview unknown period:', period) /* eslint-disable-line no-console */
      // }
    }
  }
}
