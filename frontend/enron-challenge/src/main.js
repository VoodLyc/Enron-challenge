import { createApp } from 'vue'
import App from './App.vue'
import './style.css'
import PrimeVue from 'primevue/config';
import Wind from './presets/wind';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';

const app = createApp(App);
app.use(PrimeVue, {
    unstyled: true,
    pt: Wind
});
app.component('DataTable', DataTable);
app.component('Column', Column);
app.component('InputGroup', InputGroup);
app.component('InputGroupAddon', InputGroupAddon);
app.component('Button', Button);
app.component('InputText', InputText);
app.mount('#app');