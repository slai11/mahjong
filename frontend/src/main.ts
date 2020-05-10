import Vue from 'vue'
import App from './App.vue'
import Home from './Home.vue'
import store from './store'
import vuetify from './plugins/vuetify';
import VueRouter from 'vue-router'

Vue.config.productionTip = false
Vue.use(VueRouter)

const router = new VueRouter({
  routes: [
    { path: '/', component: Home },
    { path: '/:gameid', component: Home, props: true }
  ],
  mode: 'history'
});

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
