import { createApp } from 'vue'
import './style.css'
import PrimeVue from 'primevue/config';
import Aura from '@primevue/themes/aura';
import App from './App.vue'
import { createMemoryHistory, createRouter } from 'vue-router'
import ListAulas from './components/view/tab/aulas/ListAulas.vue'
import ListProfessores from './components/view/tab/professores/ListProfessores.vue'
import ProfessoresView from './components/view/ProfessoresView.vue'

const routes = [
  { path: '/', component: ProfessoresView },
  { path: '/aulas', component: ListAulas },
  { path: '/professores', component: ListProfessores },
]

const router = createRouter({
  history: createMemoryHistory(),
  routes,
})


createApp(App).use(router).mount('#app')
