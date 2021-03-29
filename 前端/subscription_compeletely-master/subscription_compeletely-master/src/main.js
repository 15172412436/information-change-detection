// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './app.vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import echarts from 'echarts'

// 网络的路由库
import router from './router'

// 网络请求get,post所需要的库
import axios from 'axios'
import VueAxios from 'vue-axios'
import qs from 'qs'
import {Addemail, Deleteemail, AddSub, Whenload, DeleteRow, EditRowOpen, PushEdit} from './js/func'

Vue.prototype.$Addemail = Addemail
Vue.prototype.$Deleteemail = Deleteemail
Vue.prototype.$AddSub = AddSub
Vue.prototype.$Whenload = Whenload
Vue.prototype.$DeleteRow = DeleteRow
Vue.prototype.$EditRowOpen = EditRowOpen
Vue.prototype.$PushEdit = PushEdit
Vue.use(VueAxios, axios)
Vue.use(qs)
Vue.use(ElementUI)

Vue.config.productionTip = false
Vue.prototype.$echarts = echarts;

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
