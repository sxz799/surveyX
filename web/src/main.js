import { createApp } from 'vue'

import App from './App.vue'
import router from './utils/router'
import 'element-plus/es/components/message/style/css'

const app = createApp(App)

app.use(router)
app.mount('#app')
