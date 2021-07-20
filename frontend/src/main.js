import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';
import vuetify from './plugins/vuetify'
import router from './router'

Wails.Init(() => {
	new Vue({
        vuetify,
        router,
        render: h => h(App)
    }).$mount('#app');
});
