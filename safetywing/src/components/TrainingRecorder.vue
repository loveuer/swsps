<template>
    <div>
        <el-menu :default-active="this.activeIndex" class="el-menu-demo" mode="horizontal" @select="changeMode">
            <el-menu-item index="1">上传训练单</el-menu-item>
            <el-menu-item index="2">查看训练单</el-menu-item>
            <el-menu-item index="3">修改训练单</el-menu-item>
        </el-menu>
        <component :is="this.trainingRecorderMode"></component>
    </div>
</template>

<script>
import trainingRecorderAdd from "./TrainingRecorderAdd.vue";
import trainingRecorderSearch from "./TrainingRecorderSearch.vue";
import trainingRecorderModify from "./TrainingRecorderModify.vue";

export default {
    data() {
        return {
            activeIndex: '1',
            trainingRecorderMode: "t-add",
        }
    },
    methods: {
        changeMode: function(index) {
            let modeMap = {"1":"t-add", "2":"t-search", "3":"t-modify"};
            this.trainingRecorderMode = modeMap[index];
        },
    },
    components: {
        "t-add": trainingRecorderAdd,
        "t-search": trainingRecorderSearch,
        "t-modify": trainingRecorderModify,
    },
};
</script>

<style scoped>

</style>
