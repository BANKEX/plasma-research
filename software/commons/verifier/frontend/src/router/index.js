import Vue from 'vue'
import Router from 'vue-router'
import Verifier from '@/components/Verifier'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Verifier',
      component: Verifier
    }
  ]
})
