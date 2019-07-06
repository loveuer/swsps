<template>
    <div id="container">
        <sps-menu :path="'/'"></sps-menu>
        <div style="height:50px;width:100%;zoom:100%;"></div>
        <div class="spname-space">
            <div>{{ sp.name }}</div>
        </div>
        <div>
            <component :is="onespComponents" :onesp="sp" v-on:chgmod="chgComponent"></component>
        </div>
    </div>
</template>

<script>
import SpsMenu from "../Menu.vue";
import SPImf from "./OneSP-Imf.vue";
import SPEdit from './OneSP-Edit.vue';
import SPHis from './OneSP-His.vue';

export default {
    data() {
        return {
            sp: {name: 'Name ...', img1: "None", img2: "None"},
            onespComponents: 'sp-imf',
        };
    },
    methods: {
        chgComponent: function(which) {
            this.onespComponents = which;
        },
        dealRespData: function(data) {
            let sp = data;
            if (sp.img1 !== 'None') {
                sp.img1 = `http://192.168.0.19:8820/static/sp_img/${sp.img1}`;
                // sp.img1 = `http://${window.location.host}/static/sp_img/${sp.img1}`;
            };
            if (sp.img2 !== 'None') {
                sp.img2 = `http://192.168.0.19:8820/static/sp_img/${sp.img2}`;
                // sp.img2 = `http://${window.location.host}/static/sp_img/${sp.img2}`;
            };
            this.sp = sp;
        },
    },
    computed: {
        
    },
    mounted() {
        this.$http.get(`/api/sps/one/${this.$route.params.spid}`)
            .then(resp => {
                this.dealRespData(resp.data);
            })
            .catch(err => {
                switch (err.response.status) {
                    case 401:
                        this.$router.push("/login");
                        break;
                };
            });
    },
    components: {
        "sps-menu": SpsMenu,
        'sp-imf': SPImf,
        'sp-edit': SPEdit,
        'sp-his': SPHis,
    },
};
</script>

<style scoped>
.spname-space {
    width: 100%;
    height: 30px;
    display: flex;
    justify-content: center;
}
.spname-space > div {
    width: 80%;
    height: 30px;
    font-size: 2rem;
    font-family: hack;
    border-bottom: 5px solid #409EFF;
    line-height: 30px;
}
@font-face {
    font-family: hack;
    src: url("~@/assets/fonts/Hack-BoldItalic.ttf");
}
</style>
