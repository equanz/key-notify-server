<template>
  <section>
    <h2>{{room_name}}</h2>
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
    height: 15rem;
    width: @content-small-size;
    text-align: center;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;

    .screen-sm({
      width: 50vw;
    });

    .screen-md({
      width: 35vw;
    });

    .screen-lg({
      width: 35vw;
    });
  }

  h2{
    margin: 0 auto;
    font-size: 1.5rem;
    font-weight: normal;
    font-style: normal;
  }

  p{
    margin: 0 auto;
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

