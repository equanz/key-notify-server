<script>
  import {Bar} from 'vue-chartjs'

  export default {
    extends: Bar,
    props: ['raw_data', 'fiscal_year'],
    mounted() {
      this.renderChart({
        labels: createYearLabel(this.fiscal_year),
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
   * compile to plot data from raw data
   * @param raw_data {Array} - raw data about status
   * @param fiscal_year {Number} - fiscal year by raw_data
   * @return {Array} return plot data
   */
  let compilePlotData = (raw_data, fiscal_year) => {
    // property
    let plot_data = Array(getWeekVal(fiscal_year)).fill(0) //initialize
    let raw_data_index = 0 // initialize
    let index_date = new Date(fiscal_year, 3, 1) // initialize to 4/1

    // const
    const BACK_COLOR =  '#f87979'

    // copy array
    raw_data = raw_data.concat()

    // set index_date to first sunday
    index_date.setDate(index_date.getDate() - index_date.getDay())


    for(let i = 0; i < plot_data.length; i++){
      let plot_value = 0

      while(raw_data_index < raw_data.length){
        let time_date = new Date(raw_data[raw_data_index].Time)
        // state is ON
        if(raw_data[raw_data_index].State == "ON"){
          if(isSameWeek(time_date, index_date) == true){
            let off_time_date = {} // initialize
            if(raw_data_index < raw_data.length - 1){
              off_time_date = new Date(raw_data[raw_data_index + 1].Time)
            } else{
              let now_time_date = new Date()
              let now_fiscal_year = 0

              // calc fiscal year
              if(now_time_date.getMonth() < 3){
                now_fiscal_year = now_time_date.getFullYear() + 1
              } else{
                now_fiscal_year = now_time_date.getFullYear()
              }
              // now year
              if(time_date.getFullYear() == now_fiscal_year){
                off_time_date = new Date()
              } else{ // not now year
                off_time_date = new Date(time_date.getFullYear(), 2, 31, 23, 59, 59, 999) // 3/31 23:59:59;999
                // not saturday
                if(off_time_date.getDay() != 6){
                  off_time_date.setDate(off_time_date.getDate() - off_time_date.getDay() - 1) // last day in fiscal year
                }
                raw_data.push(off_time_date)
              }
            }

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
    }

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

  /**
   * if last day by fiscal year is sunday, return 53, otherwise return 52
   * @param year {Number} - fiscal year
   * @return {Number} return week value
   */
  let getWeekVal = (year) => {
    // const
    const LEAP_WEEK = 53
    const NORMAL_WEEK = 53
    const LAST_DAY_FISCAL_YEAR = new Date(year + 1, 2, 31)
    if(isLeapYear(year) == true){
      // sunday
      if(LAST_DAY_FISCAL_YEAR.getDay() == 0){
        return LEAP_WEEK
      } else{ // not sunday
        return NORMAL_WEEK
      }

    } else{ // normal year
      return NORMAL_WEEK
    }
  }

  let createYearLabel = (year) => {
    // property
    let index_date = new Date(year, 3, 1) // initialize to 4/1
    let year_label = Array(getWeekVal(year)).fill("")
    let before_month = -1

    // set index_date to first sunday
    index_date.setDate(index_date.getDate() - index_date.getDay())

    for(let i = 0; i < year_label.length; i++){
      if(index_date.getMonth() != before_month){
        year_label[i] = `${index_date.getMonth() + 1}`
        before_month = index_date.getMonth()
      }

      // increment
      index_date.setDate(index_date.getDate() +7)
    }

    return year_label
  }
</script>
