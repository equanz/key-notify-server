<template>
  <div id="app">
    <h1>My Key-Notify App!</h1>
    <CommitChart></CommitChart>
  </div>
</template>

<script>
  import CommitChart from './chart.vue'
  export default {
    components: {
      CommitChart
    }
  }
</script>

<style>
  h1 {
    text-align: center;
  }
</style>
