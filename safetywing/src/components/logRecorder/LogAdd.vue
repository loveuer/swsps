<template>
    <div>
        <loveuer-menu ref="uMenu"></loveuer-menu>
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";

export default {
    components: {
        "loveuer-menu": loveuerMenu,
    },
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/logRecorder/add";
    },
};
</script>
