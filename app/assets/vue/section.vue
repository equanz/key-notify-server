<template>
  <section>
    <h1>
      <p>{{section_name}}</p>
    </h1>
    <slot></slot>
  </section>
</template>

<script>
  export default {
    props: ['section_name']
  }
</script>

<style lang="less" scoped>
  @import "../less/main.less";

  section{
    background-color: @article-back-color;
    width: 45vw;
    height: auto;
  }

  h1{
    margin: 0;
    margin-bottom: 10px;
    background-color: @title-back-color;
  }

  p{
    margin: 0;
    margin-left: 1vw;
    font-weight: normal;
    font-style: normal;
    font-size: 1rem;
  }
</style>

