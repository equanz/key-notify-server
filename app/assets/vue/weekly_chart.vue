<script>
  import {HorizontalBar, mixins} from 'vue-chartjs'

  export default {
    extends: HorizontalBar,
    mixins: [mixins.reactiveData],
    props: ['raw_data', 'last_week_state', 'select_date'],
    data() {
      return { // options
        options: {
          scales: {
            xAxes: [{
              stacked: true,
              ticks: {
                beginAtZero: true,
                max: 24 // max value in x-axes
              }
            }],
            yAxes: [{
              stacked: true
            }]
          },
          tooltips: {
            enabled: false // hide tooltips
          },
          legend: {
            display: false // hide legend
          }
        }
      }
    },
    mounted() {
      this.chartData = {
        labels: ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'], // day label
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
      updatePlotData: function() { // update datasets
        let newChartData = Object.assign({}, this.chartData)
        newChartData.datasets = this.compilePlotData(this.raw_data, this.last_week_state)
        this.chartData = newChartData
      },
      /**
       * compile to plot data from raw data
       * @param raw_data {Array} - raw data about status
       * @param last_week_state {String} - last week state
       * @return {Array} return plot data
       */
      compilePlotData: function(raw_data, last_week_state){
        // property
        let plot_data = [] //initialize

        // const
        const BACK_COLOR = {
          ON: '#27AE60',
          OFF: '#BDC3C7'
        }


        for(let i = 0; i < raw_data.length; i++){
          let background_color = "" // initialize
          let plot_value = 0 // time
          let time_date = new Date(raw_data[i].Time) // convert to type Date from String
          let plot_array = Array(7).fill(0) // initialize(quota(7) is amount of day)

          // not first data
          if(i > 0){
            let before_time_date = new Date(raw_data[i - 1].Time)
            // same day
            if(before_time_date.getFullYear() == time_date.getFullYear() &&
               before_time_date.getMonth() == time_date.getMonth() &&
               before_time_date.getDate() == time_date.getDate()){
              plot_value = time_date.getHours() - before_time_date.getHours()
            } else{ // different day
              plot_value = time_date.getHours()

              // fill by last status to before day
              let plot_array_before = Array(7).fill(0) // initialize(quota(7) is amount of day)
              plot_array_before[before_time_date.getDay()] = 24 - before_time_date.getHours()
              for(let i = before_time_date.getDay() + 1; i < time_date.getDay(); i++){
                plot_array_before[i] = 24 // fill about all day
              }
              plot_data.push({
                backgroundColor: BACK_COLOR[raw_data[i - 1].State],
                data: plot_array_before
              })
            }
          } else{ // first data
            // not sunday
            if(time_date.getDay() > 0){
              // fill before day
              let background_color_before = ""
              if(raw_data[i].State == "ON"){
                background_color_before = BACK_COLOR.OFF
              } else{
                background_color_before = BACK_COLOR.ON
              }
              for(let j = 0; j < time_date.getDay(); j++){
                let plot_array_before = Array(7).fill(0) // initialize(quota(7) is amount of day)
                plot_array_before[j] = 24

                plot_data.push({
                  backgroundColor: background_color_before,
                  data: plot_array_before
                })
              }

              plot_value = time_date.getHours()
            } else{ //sunday
              plot_value = time_date.getHours()
            }
          }

          plot_array[time_date.getDay()] = plot_value // insert plot_value into selected index

          if(raw_data[i].State == "ON"){
            background_color = BACK_COLOR.OFF
          } else{
            background_color = BACK_COLOR.ON
          }

          plot_data.push({
            backgroundColor: background_color,
            data: plot_array
          })
        }

        // have action in week
        if(raw_data.length > 0){
          // when input week and now week is not same, fill by last status
          let plot_array = Array(7).fill(0) // initialize(quota(7) is amount of day)
          let plot_array_now = Array(7).fill(0) // initialize(quota(7) is amount of day)
          let now_time_date = new Date()
          let now_time_date_sunday = new Date()
          let last_row = raw_data[raw_data.length - 1]
          let last_row_time_date = new Date(raw_data[raw_data.length - 1].Time)
          let last_row_time_date_sunday = new Date(raw_data[raw_data.length - 1].Time)

          // get sunday and set
          now_time_date_sunday.setDate(now_time_date_sunday.getDate() - now_time_date_sunday.getDay())
          last_row_time_date_sunday.setDate(last_row_time_date_sunday.getDate() - last_row_time_date_sunday.getDay())

          // same week
          if(last_row_time_date_sunday.getFullYear() == now_time_date_sunday.getFullYear() &&
             last_row_time_date_sunday.getMonth() == now_time_date_sunday.getMonth() &&
             last_row_time_date_sunday.getDate() == now_time_date_sunday.getDate()){
            // same day
            if(now_time_date.getDay() - last_row_time_date.getDay() == 0){
              plot_array[now_time_date.getDay()] = now_time_date.getHours() - last_row_time_date.getHours()

              plot_data.push({
                backgroundColor: BACK_COLOR[last_row.State],
                data: plot_array
              })
            } else{ // different day
              plot_array[last_row_time_date.getDay()] = 24 - last_row_time_date.getHours()

              plot_data.push({
                backgroundColor: BACK_COLOR[last_row.State],
                data: plot_array
              })

              for(let i = 0; i < now_time_date.getDay() - last_row_time_date.getDay() - 1; i++){
                let plot_array_after = Array(7).fill(0) // initialize(quota(7) is amount of day)
                plot_array_after[last_row_time_date.getDay() + i + 1] = 24

                plot_data.push({
                  backgroundColor: BACK_COLOR[last_row.State],
                  data: plot_array_after
                })
              }
              plot_array_now[now_time_date.getDay()] = now_time_date.getHours()

              plot_data.push({
                backgroundColor: BACK_COLOR[last_row.State],
                data: plot_array_now
              })
            }
          } else{ //different week
            plot_array[last_row_time_date.getDay()] = 24 - last_row_time_date.getHours()

            plot_data.push({
              backgroundColor: BACK_COLOR[last_row.State],
              data: plot_array
            })

            for(let i = 0; i < 6 - last_row_time_date.getDay(); i++){
              let plot_array_after = Array(7).fill(0) // initialize(quota(7) is amount of day)
              plot_array_after[i + last_row_time_date.getDay() + 1] = 24

              plot_data.push({
                backgroundColor: BACK_COLOR[last_row.State],
                data: plot_array_after
              })
            }
          }
        } else{ // no action in week
          let now_time_date = new Date()
          let now_time_date_sunday = new Date()
          let select_date_sunday = new Date(this.select_date.getTime())

          // get sunday and set
          now_time_date_sunday.setDate(now_time_date_sunday.getDate() - now_time_date_sunday.getDay())
          select_date_sunday.setDate(select_date_sunday.getDate() - select_date_sunday.getDay())

          // same week
          if(select_date_sunday.getFullYear() == now_time_date_sunday.getFullYear() &&
             select_date_sunday.getMonth() == now_time_date_sunday.getMonth() &&
             select_date_sunday.getDate() == now_time_date_sunday.getDate()){
            for(let i = 0; i < 7; i++){ // 7 is amount of day
              let plot_array_after = Array(7).fill(0) // initialize(quota(7) is amount of day)
              // different day
              if(i < now_time_date.getDay()){
                plot_array_after[i] = 24
              } else if(i == now_time_date.getDay()){ // same day
                plot_array_after[i] = now_time_date.getHours()
              } else{ // future day
                plot_array_after[i] = 0
              }

              plot_data.push({
                backgroundColor: BACK_COLOR[last_week_state],
                data: plot_array_after
              })
            }
          } else{ // different week
            for(let i = 0; i < 7; i++){ // 7 is amount of day
              let plot_array_after = Array(7).fill(0) // initialize(quota(7) is amount of day)
              plot_array_after[i] = 24

              plot_data.push({
                backgroundColor: BACK_COLOR[last_week_state],
                data: plot_array_after
              })
            }
          }

        }
        return plot_data
      }
    }
  }
</script>
