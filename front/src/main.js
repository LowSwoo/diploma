import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import Vuetify from 'vuetify/lib'
import VueResource from "vue-resource"
import marked from 'marked'
import lodash from 'lodash'

Vue.config.productionTip = false
Vue.use(VueResource)


new Vue({
  router,
  vuetify,
  VueResource,
  render: h => h(App)
}).$mount('#app')
