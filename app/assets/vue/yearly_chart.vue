<script>
  import {Bar, mixins} from 'vue-chartjs'

  export default {
    extends: Bar,
    mixins: [mixins.reactiveData],
    props: ['raw_data', 'fiscal_year', 'select_date'],
    data() {
      return {
        options: { // options
          scales: {
            xAxes: [{
              scaleLabel: {
                display: true,
                labelString: "date[months]"
              }
            }],
            yAxes: [{
              scaleLabel: {
                display: true,
                labelString: "active time[hours]"
              }
            }]
          },
          tooltips: {
            enabled: false
          },
                   legend: {
                     display: false
                   }
        }
      }
    },
    mounted() {
      this.chartData = {
        labels: this.createYearLabel(this.fiscal_year),
        datasets: this.datasets
      }
    },
    watch: {
      raw_data: function() {
        // run update
        this.updatePlotData()
      }
    },
    methods: {
      updatePlotData: function() {
        let newChartData = Object.assign({}, this.chartData)
        newChartData.datasets = this.compilePlotData(this.raw_data, this.fiscal_year)
        this.chartData = newChartData
      },
      /**
       * compile to plot data from raw data
       * @param raw_data {Array} - raw data about status
       * @param fiscal_year {Number} - fiscal year by raw_data
       * @return {Array} return plot data
       */
      compilePlotData: function(raw_data, fiscal_year) {
        // property
        let plot_data = Array(this.getWeekVal(fiscal_year)).fill(0) //initialize
        let raw_data_index = 0 // initialize
        let index_date = new Date(fiscal_year, 3, 1) // initialize to 4/1

        // const
        const BACK_COLOR =  '#27AE60'

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
              if(this.isSameWeek(time_date, index_date) == true){
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

                if(this.isSameWeek(off_time_date, index_date) == true){
                  // add value
                  plot_value += this.getDiffHour(time_date, off_time_date)
                  raw_data_index += 2 // increment
                } else{ // different week
                  let last_week_date = new Date(index_date.getTime())
                  last_week_date.setDate(last_week_date.getDate() + (6 - last_week_date.getDay())) // calc last week
                  last_week_date.setHours(23, 59, 59, 999) // calc last time

                  // add value
                  plot_value += this.getDiffHour(time_date, last_week_date)
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
              if(this.isSameWeek(time_date, index_date) == true){
                // add value
                plot_value += this.getDiffHour(index_date, time_date)
                raw_data_index += 1 // increment
              } else{
                let last_week_date = new Date(index_date.getTime())
                last_week_date.setDate(last_week_date.getDate() + (6 - last_week_date.getDay())) // calc last week
                last_week_date.setHours(23, 59, 59, 999) // calc last time

                // add value
                plot_value += this.getDiffHour(index_date, last_week_date)

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
      },

      /**
       * check leap year
       * @param year {Number} - fiscal year
       * @return {Bool} when input fiscal year is leap year, return true
       */
      isLeapYear: function(year) {
        if(((year % 4 == 0) && (year % 100 != 0)) || (year % 400 == 0)){
          return true
        } else{
          return false
        }
      },

      /**
       * check same week
       * @param date_one {Date} - one of check date
       * @param date_two {Date} - one of check date
       * @return {Bool} when input date is same week, return true
       */
      isSameWeek: function(date_one, date_two) {
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
      },

      /**
       * calc diff hour
       * @param date_one {Date} - one of calc date
       * @param date_two {Date} - one of calc date
       * @return {Number} return diff hour
       */
      getDiffHour: function(date_one, date_two) {
        // property
        let diff_millsec = 0
        let diff_hour = 0

        // calc diff
        diff_millsec = Math.abs(date_one.getTime() - date_two.getTime())
        diff_hour = Math.floor(diff_millsec / (1000 * 60 * 60))

        return diff_hour
      },

      /**
       * if last day by fiscal year is sunday, return 53, otherwise return 52
       * @param year {Number} - fiscal year
       * @return {Number} return week value
       */
      getWeekVal: function(year) {
        // const
        const LEAP_WEEK = 53
        const NORMAL_WEEK = 52
        const LAST_DAY_FISCAL_YEAR = new Date(year + 1, 2, 31)
        if(this.isLeapYear(year) == true){
          // sunday
          if(LAST_DAY_FISCAL_YEAR.getDay() == 0){
            return LEAP_WEEK
          } else{ // not sunday
            return NORMAL_WEEK
          }

        } else{ // normal year
          return NORMAL_WEEK
        }
      },

      createYearLabel: function(year) {
        // property
        let index_date = new Date(year, 3, 1) // initialize to 4/1
        let year_label = Array(this.getWeekVal(year)).fill("")
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
    }
  }

</script>
