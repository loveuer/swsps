<template>
    <div>
        <div>
            <loveuer-menu />
        </div>
        <div>
            <el-badge :is-dot="true" style="margin-top:10px;margin-right:40px;">
                <font>Safety</font>
                <font>Wing</font>
            </el-badge>
        </div>
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
