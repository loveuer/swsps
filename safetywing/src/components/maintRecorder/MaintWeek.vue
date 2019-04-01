<template>
    <div>
        <div>
            <loveuer-menu ref="uMenu"></loveuer-menu>
        </div>
        <div>maint week</div>
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";

export default {
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/maintRecorder/week";
    },
    components: {
        "loveuer-menu": loveuerMenu,
    }
};
</script>
