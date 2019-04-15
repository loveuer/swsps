<template>
    <div id="container">
        <u-menu></u-menu>
        <div>this is home</div>
    </div>
</template>

<script>
import uMenu from "./u-cmpnts/u_menu.vue";

export default {
    data() {
        return {

        };
    },
    components: {
        "u-menu": uMenu,
    },
};
</script>

<style scoped>
#container {
    height: 100%;
    min-height: 100%;
    width: 100%;
}
</style>
