<template>
  <footer id="footer">
    <p>
      Copyright (c) 2018 <a href="https://github.com/equanz">Yuri Mizushima</a>, <a href="https://github.com/amataku">Amano Takumi</a>
    </p>
  </footer>
</template>

<style lang="less" scoped>
  @import "../less/main.less";

  footer{
    margin-top: 20px;
    display: flex;
    width: 100vw;
    height: 50px;
    background-color: @title-back-color;
    align-items: center;
    justify-content: center;
  }

  a, a:link, a:visited, a:hover, a:active{
    color: @text-black-color;
  }
</style>
