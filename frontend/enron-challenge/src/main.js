import { createApp } from 'vue'
import App from './App.vue'
import './style.css'
import PrimeVue from 'primevue/config';
import Wind from './presets/wind';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';


const app = createApp(App);
app.use(PrimeVue, {
    unstyled: true,
    pt: Wind
});
app.component('DataTable', DataTable);
app.component('Column', Column);
app.mount('#app');