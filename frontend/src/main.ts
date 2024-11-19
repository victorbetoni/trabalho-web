import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { createMemoryHistory, createRouter, createWebHashHistory } from 'vue-router'
import ListAulas from './components/view/tab/aulas/ListAulas.vue'
import ProfessoresView from './components/view/ProfessoresView.vue'
import Login from './components/view/Login.vue';
import Home from './components/view/Home.vue';

const routes = [
  { path: '/', component: Home },
  { path: '/login', component: Login },
  { path: '/aulas', component: ListAulas },
  { path: '/professores', component: ProfessoresView },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})


createApp(App).use(router).mount('#app')
