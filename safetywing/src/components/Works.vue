<template>
    <div>
        <loveuer-menu />
    </div>
</template>

<script>
import loveuerMenu from "./uMenu.vue";

export default {
    components: {
        "loveuer-menu": loveuerMenu,
    }
};
</script>

<style scoped>
.font-logo{
    height: 60px;
    line-height: 60px;
    font-size: 20px;
    font-weight: bold;
}
.font-logo:first-child{
    color: #7ec0d4;
}
.font-logo:last-child{
    color: #85c59a;
}
</style>
