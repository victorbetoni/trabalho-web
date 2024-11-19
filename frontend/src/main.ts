import { createApp } from 'vue'
import './style.css'
import PrimeVue from 'primevue/config';
import Aura from '@primevue/themes/aura';
import App from './App.vue'
import { createMemoryHistory, createRouter } from 'vue-router'
import ListAulas from './components/view/tab/aulas/ListAulas.vue'
import ListProfessores from './components/view/tab/professores/ListProfessores.vue'
import ProfessoresView from './components/view/ProfessoresView.vue'
import Login from './components/view/Login.vue';
import Home from './components/view/Home.vue';

const routes = [
  { path: '/', component: Login },
  { path: '/login', component: Login },
  { path: '/aulas', component: ListAulas },
  { path: '/professores', component: ProfessoresView },
]

const router = createRouter({
  history: createMemoryHistory(),
  routes,
})


createApp(App).use(router).mount('#app')
