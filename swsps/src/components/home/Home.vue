<template>
    <div id="container">
        <sps-menu :path="'/'"></sps-menu>
        <div style="width:100%;height:30px;zoom:100%;"></div>
        <div class="home-container">
            <div>
                <div>
                    <div>
                        <home-search></home-search>
                    </div>
                </div>
                <div>
                    <div>
                        <home-history></home-history>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import SpsMenu from "../Menu.vue";
import HomeSearch from "./Home-Search.vue";
import HomeHistory from "./Home-History.vue";

export default {
    data() {
        return {};
    },
    methods: {
       
    },
    components: {
        "sps-menu": SpsMenu,
        "home-search": HomeSearch,
        "home-history": HomeHistory,
    },
    mounted() {
        
    },
};
</script>

<style scoped>
.home-container {
    width: 100%;
    display: flex;
    justify-content: center;
}
.home-container > div {
    width: 95%;
    display: flex;
}
.home-container > div > div {
    width: 50%;
    display: flex;
    justify-content: center;
}
.home-container > div > div > div {
    width: 90%;
}
</style>
