<template>
    <div style="min-height:100%;height:100%;">
        <div>
            <loveuer-menu ref="uMenu"></loveuer-menu>
        </div>
        <div class="sps-content">
            <div class="sps-dhm"><!-- detail history modify -->
                <el-menu
                    default-active="1"
                    class="el-menu-vertical-demo"
                    style="height:100%;"
                    text-color="#000"
                    background-color="#fff"
                    @select="spsMenuSelected"
                    active-text-color="#409EFF">
                    <el-menu-item index="1">
                        <i class="el-icon-document"></i>
                        <span slot="title">备件详情</span>
                    </el-menu-item>
                    <el-menu-item index="2">
                        <i class="el-icon-date"></i>
                        <span slot="title">备件历史</span>
                    </el-menu-item>
                    <el-menu-item index="3">
                        <i class="el-icon-edit-outline"></i>
                        <span slot="title">编辑备件</span>
                    </el-menu-item>
                </el-menu>
            </div>
            <div class="sps-main">
                <div style="height:20px;width:100%;zoom:100%;"></div>
                <component :is="spsOneMode"></component>
            </div>
        </div>
        
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";
import spsOneDetail from "./OneDetail.vue";
import spsOneHistory from "./OneHistory.vue";
import spsOneModify from "./OneModify.vue";

export default {
    data() {
        return {
            spsOneMode: "sps-one-detail",
        };
    },
    methods: {
        spsMenuSelected: function(index) {
            let oneModeMap = {"1":"sps-one-detail", "2":"sps-one-history", "3":"sps-one-modify"};
            this.spsOneMode = oneModeMap[index];
        },
    },
    components: {
        "loveuer-menu": loveuerMenu,
        "sps-one-detail": spsOneDetail,
        "sps-one-history": spsOneHistory,
        "sps-one-modify": spsOneModify,
    },
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/spsRecorder/search";
    },
};
</script>

<style scoped>
.sps-content{
    width: 100%;
    height: 100%;
    display: flex;
}
.sps-dhm{
    width: 140px;
    height: 100%;
    background-color: #fff;
}
.sps-main {
    height:100%;
    margin-left: 50px;
    width: 100%;
}
</style>
