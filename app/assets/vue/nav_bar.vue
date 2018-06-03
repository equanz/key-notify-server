<template>
  <nav>
    <p>Key Status Dashboard</p>
  </nav>
</template>

<script>
  export default {
  }
</script>

<style lang="less" scoped>
  @import "../less/main.less";

  nav{
    display: flex;
    width: 100vw;
    height: 7vh;
    background-color: @nav-color;
    align-items: center;
  }

  p{
    margin: 0;
    margin-left: 1vw;
    color: @text-white-color;
    font-size: 1.5rem;
  }
</style>

