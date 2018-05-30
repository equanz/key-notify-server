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
</script>

<style>
  h1 {
    text-align: center;
  }
</style>
