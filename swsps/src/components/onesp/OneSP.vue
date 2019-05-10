<template>
    <div id="container">
        <sps-menu :path="''"></sps-menu>
        <div style="height:50px;width:100%;zoom:100%;"></div>
        <div class="spname-space">
            <div>{{ sp.name }}</div>
        </div>
    </div>
</template>

<script>
import SpsMenu from "../Menu.vue";

export default {
    data() {
        return {
            sp: {name:'RADIO MANAGEMENT PANEL'},
        };
    },
    components: {
        "sps-menu": SpsMenu
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
