<template>
  <section>
    <p>{{room_name}}</p>
    <figure v-if="status=='ON'" id="status-on"></figure>
    <figure v-else id="status-off"></figure>
    <p>from {{from_time}}</p>
  </section>
</template>

<script>
  export default {
    props: ['room_name', 'from_time', 'status']
  }
</script>

<style lang="less" scoped>
  @import "../less/main.less";

  section{
    display: flex;
    background-color: @article-back-color;
    width: 35vw;
    height: 15rem;
    text-align: center;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
  }

  p{
    margin: 0 auto;
    color: @text-black-color;
    font-size: 1.5rem;
  }

  // mixin
  .figure{
    margin: 0 auto;
    border-radius: 50%;
    height: 100px;
    width: 100px;
  }

  figure#status-on{
    .figure;
    background-color: @state-on-color;
  }
  figure#status-off{
    .figure;
    background-color: @state-off-color;
  }
</style>

