<template>
  <div id="app">
    <h1>WeeklyChart</h1>
    <WeeklyChart></WeeklyChart>
  </div>
</template>

<script>
  import WeeklyChart from './weekly_chart.vue'
  export default {
    components: {
      WeeklyChart
    }
  }
</script>

<style>
  h1 {
    text-align: center;
  }
</style>
