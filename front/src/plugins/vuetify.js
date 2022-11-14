import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';
import colors from 'vuetify/lib/util/colors'

Vue.use(Vuetify);
let options = {
    theme: {
        dark: true,
        themes: {
            dark: {
                primary: '#1976D2',
                secondary: '#424242',
                accent: '#82B1FF',
                error: '#FF5252',
                info: '#2196F3',
                success: '#4CAF50',
                warning: '#FB8C00',
                navigationDrawer: "#263238",
                topNav: "#263238"
                

            },
            light: {
                navigationDrawer: "#CFD8DC",
                navigationDrawerSheet: "#90A4AE",
                topNav: "#CFD8DC",
                mainSheet: "#ECEFF1"

            }
        }
    }
}

export default new Vuetify({
    ...options
});
