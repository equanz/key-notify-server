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
  function GetGtatistics(first_date){
    return new Promise(function(resolve,reject){
      var request = new XMLHttpRequest()
      // format Date
      if(first_date.getDay() != 0){
        first_date.setDate(first_date.getDate() - first_date.getDay())
      }
      var month = first_date.getMonth() + 1
      if(month <= 9){
        month = "0" + month
      }
      var url = "http://localhost:8080/api/statistics?first_date=" + first_date.getFullYear() + "-" + month + "-" + first_date.getDate() + " 00:00:00"
      // send request
      request.responseType = 'json'
      request.open("GET", url)
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
