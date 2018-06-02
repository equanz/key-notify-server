<template>
  <div id="app">
    <h1>WeeklyChart</h1>
    <WeeklyChart></WeeklyChart>

    <h1>YearlyChart</h1>
    <YearlyChart></YearlyChart>
  </div>
</template>

<script>
  import WeeklyChart from './weekly_chart.vue'
  import YearlyChart from './yearly_chart.vue'
  export default {
    components: { // add to component
      WeeklyChart,
      YearlyChart
    }
  }
  /**
   * APIを用い指定した日付からの鍵の情報を取得
   * @param first_date{Date} -  get from this param
   * @return {Promise} - Promise object
   */
  function GetStatistics(first_date){
    return new Promise(function(resolve,reject){
      const url = "https://key-notify-server.herokuapp.com/api/statistics"
      let from_date = new Date(first_date.getTime())
      let request = new XMLHttpRequest()
      // format Date
      if(from_date.getDay() != 0){
        from_date.setDate(from_date.getDate() - from_date.getDay())
      }
      let month = from_date.getMonth() + 1
      if(month <= 9){
        month = "0" + month
      }
      let url_param = url + "?first_date=" + from_date.getFullYear() + "-" + month + "-" + from_date.getDate() + " 00:00:00"
      console.log(url_param)
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
