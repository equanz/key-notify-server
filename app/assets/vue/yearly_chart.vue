<script>
  import {Bar} from 'vue-chartjs'

  export default {
    extends: Bar,
    props: ['raw_data', 'fiscal_year'],
    mounted() {
      this.renderChart({
        labels: Array(52).fill(0),
        datasets: compilePlotData(this.raw_data, this.fiscal_year),
      }, { // options
        tooltips: {
          enabled: false
        },
           legend: {
             display: false
           }
      })
    }
  }

  /**
   * compile to plot data from raw data   (caution: not implemented)
   * @param raw_data {Array} - raw data about status
   * @param fiscal_year {Number} - fiscal year by raw_data
   * @return {Array} return plot data
   */
  let compilePlotData = (raw_data, fiscal_year) => {
    // property
    let plot_data = [] //initialize
    let raw_data_index = 0 // initialize
    let index_date = new Date(fiscal_year, 3, 1) // initialize to 4/1

    // const
    const BACK_COLOR =  '#f87979'

    // allocate plot_data
    // leap year
    if(isLeapYear(fiscal_year) == true){
      // if last day by fiscal year is sunday, to allocate 53
      const LAST_DAY_FISCAL_YEAR = new Date(fiscal_year + 1, 2, 31)

      // sunday
      if(LAST_DAY_FISCAL_YEAR.getDay() == 0){
        plot_data = Array(53).fill(0)
      } else{ // not sunday
        plot_data = Array(52).fill(0)
      }

    } else{ // normal year
      plot_data = Array(52).fill(0)
    }

    // set index_date to first sunday
    index_date.setDate(index_date.getDate() - index_date.getDay())

    for(let i = 0; i < plot_data.length; i++){
      let plot_value = 0

      while(raw_data_index < raw_data.length - 1){
        let time_date = new Date(raw_data[raw_data_index].Time)
        // state is ON
        if(raw_data[raw_data_index].State == "ON"){
          console.log('hoge')
          console.log(time_date)
          console.log(index_date)
          if(isSameWeek(time_date, index_date) == true){
            let off_time_date = new Date(raw_data[raw_data_index + 1].Time)
            if(isSameWeek(off_time_date, index_date) == true){
              // add value
              plot_value += getDiffHour(time_date, off_time_date)
              raw_data_index += 2 // increment
            } else{ // different week
              let last_week_date = new Date(index_date.getTime())
              last_week_date.setDate(last_week_date.getDate() + (6 - last_week_date.getDay())) // calc last week
              last_week_date.setHours(23, 59, 59, 999) // calc last time

              // add value
              plot_value += getDiffHour(time_date, last_week_date)
              raw_data_index += 1 // increment

              // increment
              index_date.setDate(index_date.getDate() + 7)
              break
            }
          } else{
            // increment
            index_date.setDate(index_date.getDate() + 7)
            break
          }
        } else{ // state is OFF
          if(isSameWeek(time_date, index_date) == true){
            // add value
            plot_value += getDiffHour(index_date, time_date)
            raw_data_index += 1 // increment
          } else{
            let last_week_date = new Date(index_date.getTime())
            last_week_date.setDate(last_week_date.getDate() + (6 - last_week_date.getDay())) // calc last week
            last_week_date.setHours(23, 59, 59, 999) // calc last time

            // add value
            plot_value += getDiffHour(index_date, last_week_date)

            // increment
            index_date.setDate(index_date.getDate() + 7)
            break
          }
        }
      }

      plot_data[i] = plot_value
      // insert to plot_data
      /* plot_data[i] = {
       *   backgroundColor: BACK_COLOR,
       *   data: [plot_value]
       * } */
    }

    console.log([{
      backgroundColor: BACK_COLOR,
      data: plot_data
    }])

    return [{
      backgroundColor: BACK_COLOR,
      data: plot_data
    }]
  }

  /**
   * check leap year
   * @param year {Number} - fiscal year
   * @return {Bool} when input fiscal year is leap year, return true
   */
  let isLeapYear = (year) => {
    if(((year % 4 == 0) && (year % 100 != 0)) || (year % 400 == 0)){
      return true
    } else{
      return false
    }
  }

  /**
   * check same week
   * @param date_one {Date} - one of check date
   * @param date_two {Date} - one of check date
   * @return {Bool} when input date is same week, return true
   */
  let isSameWeek = (date_one, date_two) => {
    // property
    let date_one_sunday = new Date(date_one.getTime())
    let date_two_sunday = new Date(date_two.getTime())

    // calc sunday
    date_one_sunday.setDate(date_one_sunday.getDate() - date_one_sunday.getDay())
    date_two_sunday.setDate(date_two_sunday.getDate() - date_two_sunday.getDay())

    // same week
    if(date_one_sunday.getFullYear() == date_two_sunday.getFullYear() &&
       date_one_sunday.getMonth() == date_two_sunday.getMonth() &&
       date_one_sunday.getDate() == date_two_sunday.getDate()){
      return true
    } else{ // different week
      return false
    }
  }

  /**
   * calc diff hour
   * @param date_one {Date} - one of calc date
   * @param date_two {Date} - one of calc date
   * @return {Number} return diff hour
   */
  let getDiffHour = (date_one, date_two) => {
    // property
    let diff_millsec = 0
    let diff_hour = 0

    // calc diff
    diff_millsec = Math.abs(date_one.getTime() - date_two.getTime())
    diff_hour = Math.floor(diff_millsec / (1000 * 60 * 60))

    return diff_hour
  }
</script>
