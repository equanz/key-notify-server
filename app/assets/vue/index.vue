<template>
  <div id="app">
    <h1>WeeklyChart</h1>
    <WeeklyChart :raw_data=weekly_raw_data :last_week_state=last_week_state></WeeklyChart>

    <h1>YearlyChart</h1>
    <YearlyChart :raw_data=yearly_raw_data :fiscal_year=fiscal_year></YearlyChart>
  </div>
</template>

<script>
  import WeeklyChart from './weekly_chart.vue'
  import YearlyChart from './yearly_chart.vue'
  export default {
    components: { // add to component
      WeeklyChart,
      YearlyChart
    },
    data() {
      return {
        // test data
        weekly_raw_data: [{"Time":"2018-05-28 15:52:56","State":"ON","Key_info_id":1},{"Time":"2018-05-28 17:43:55","State":"OFF","Key_info_id":2},{"Time":"2018-05-29 15:52:56","State":"ON","Key_info_id":3},{"Time":"2018-05-29 17:43:55","State":"OFF","Key_info_id":4},{"Time":"2018-05-30 15:52:56","State":"ON","Key_info_id":5},{"Time":"2018-05-30 17:43:55","State":"OFF","Key_info_id":6},{"Time":"2018-05-31 15:52:56","State":"ON","Key_info_id":7},{"Time":"2018-05-31 17:43:55","State":"OFF","Key_info_id":8},{"Time":"2018-06-01 16:50:06","State":"ON","Key_info_id":9}],
        last_week_state: "ON",
        yearly_raw_data: [{"Time":"2018-04-25 15:52:56","State":"ON","Key_info_id":1},{"Time":"2018-04-25 17:43:55","State":"OFF","Key_info_id":2},{"Time":"2018-05-29 15:52:56","State":"ON","Key_info_id":3},{"Time":"2018-05-29 17:43:55","State":"OFF","Key_info_id":4},{"Time":"2018-05-30 15:52:56","State":"ON","Key_info_id":5},{"Time":"2018-05-30 17:43:55","State":"OFF","Key_info_id":6},{"Time":"2018-05-31 15:52:56","State":"ON","Key_info_id":7},{"Time":"2018-05-31 17:43:55","State":"OFF","Key_info_id":8},{"Time":"2018-06-01 16:50:06","State":"ON","Key_info_id":9}],
        fiscal_year: 2018
      }
    }
  }

  /**
   * APIを用い指定した日付の週の鍵の情報を取得
   * @param date{Date} -  get information this param week
   * @return {Promise} - Promise object
   */
  function GetStatistics_Week(date){
    return new Promise(function(resolve, reject){
      const url = "https://key-notify-server.herokuapp.com/api/statistics"
      let from_date = new Date(date.getTime())
      let by_date = new Date(date.getTime())
      let request = new XMLHttpRequest()
      // format Date
      from_date.setDate(from_date.getDate() - from_date.getDay())
      if(by_date.getDay() != 6){
        by_date.setDate(by_date.getDate() + (6 - from_date.getDay()))
      }
      let from_month = from_date.getMonth() + 1
      if(from_month <= 9){
        from_month = "0" + from_month
      }
      let by_month = by_date.getMonth() + 1
      if(by_month <= 9){
        by_month = "0" + by_month
      }
      let from_day = from_date.getDate()
      if(from_day <= 9){
        from_day = "0" + from_day
      }
      let by_day = by_date.getDate()
      if(by_day <= 9){
        by_day = "0" + by_day
      }
      let url_param = url + "?first_date=" + from_date.getFullYear() + "-" + from_month + "-" + from_day + " 00:00:00&end_date=" + by_date.getFullYear() + "-" + by_month + "-" + by_day + " 23:59:59"
      // send request
      request.responseType = 'json'
      request.open("GET", url_param)
      request.addEventListener("load", (event) => {
        // get server error
        if(event.target.status != 200){
          reject()
        } else {
          resolve(event.target.response)
        }
      })
      // get connect error
      request.addEventListener("error", () => {
        reject()
      })
      request.send()
    })
  }

  /**
   * APIを用い指定した日付の年度の鍵の情報を取得
   * @param date{Date} -  get information this param fiscal year
   * @return {Promise} - Promise object
   */
  function GetStatistics_Year(date){
    return new Promise(function(resolve, reject){
      const url = "https://key-notify-server.herokuapp.com/api/statistics"
      let from_date = new Date(date.getTime())
      let by_date = new Date(date.getTime())
      let request = new XMLHttpRequest()
      // form from month
      if(from_date.getMonth() > 3){
        from_date.setMonth(from_date.getMonth() - (from_date.getMonth() - 3))
      } else if (from_date.getMonth() < 3) {
        from_date.setMonth(from_date.getMonth() + (3 - from_date.getMonth()))
        from_date.setYear(from_date.getFullYear() - 1)
      }
      // form by month
      if(by_date.getMonth() > 2){
        by_date.setMonth(by_date.getMonth() - (by_date.getMonth() - 2))
        by_date.setYear(by_date.getFullYear() + 1)
      } else if (by_date.getMonth() < 2) {
        by_date.setMonth(by_date.getMonth() + (2 - by_date.getMonth()))
      }
      // set from day
      from_date.setDate(1)
      from_date.setDate(from_date.getDate() - from_date.getDay())
      // set by day
      by_date.setDate(31)
      if(by_date.getDay() != 6){
        by_date.setDate(by_date.getDate() - (by_date.getDay() + 1))
      }
      // form url
      let from_month = from_date.getMonth() + 1
      let by_month = by_date.getMonth() + 1
      let from_day = from_date.getDate()
      if(from_day <= 9){
        from_day = "0" + from_day
      }
      let by_day = by_date.getDate()
      if(by_day <= 9){
        by_day = "0" + by_day
      }
      let url_param = url + "?first_date=" + from_date.getFullYear() + "-0" + from_month + "-" + from_day + " 00:00:00&end_date=" + by_date.getFullYear() + "-0" + by_month + "-" + by_day + " 23:59:59"
      // send request
      request.responseType = 'json'
      request.open("GET", url_param)
      request.addEventListener("load", (event) => {
        // get server error
        if(event.target.status != 200){
          reject()
        } else {
          resolve(event.target.response)
        }
      })
      // get connect error
      request.addEventListener("error", () => {
        reject()
      })
      request.send()
    })
  }
</script>

<style>
  h1 {
    text-align: center;
  }
</style>
