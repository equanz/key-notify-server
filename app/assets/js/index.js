import Vue from 'vue'
import App from '../vue/index.vue'
import NavBar from '../vue/nav_bar.vue'
import Footer from '../vue/footer.vue'

Vue.config.productionTip = false

// nav bar
new Vue({
  el: '#nav-bar',
  template: '<NavBar/>',
  components:
  {
    NavBar: NavBar
  }
})

// main app
new Vue({
  el: '#app',
  template: '<App/>',
  components:
  {
    App: App
  }
})

// footer
new Vue({
  el: '#footer',
  template: '<Footer/>',
  components:
  {
    Footer: Footer
  }
})


