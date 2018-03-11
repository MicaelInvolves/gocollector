import Vue from 'vue'
import App from './App.vue'

import axios from 'axios'

// resources
import 'bulma/css/bulma.css'

Vue.prototype.$http = axios;

new Vue({
  el: '#app',
  components: { App },
  template: '<App/>'
})
