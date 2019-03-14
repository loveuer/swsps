<template>
    <div class="spsContainer">
        <el-row style="min-width:800px">
            <el-col :span="2" class="addbtn-container">
                <div>
                    <el-button type="success" icon="el-icon-plus" @click="dialogFormVisible = true"></el-button>
                </div>
            </el-col>
            <el-col :span="22">
                <el-menu :default-active="this.activeIndex" class="el-menu-demo" mode="horizontal" @select="changeMode">
                    <el-menu-item index="1">查找备件</el-menu-item>
                    <el-menu-item index="2">备件历史</el-menu-item>
                    <el-menu-item index="3">备件统计</el-menu-item>
                </el-menu>
            </el-col>
        </el-row>
        <el-dialog title="新增备件" :visible.sync="dialogFormVisible">
            <el-form>

            </el-form>
            <el-button @click="dialogFormVisible = false">取消</el-button>
            <el-button type="primary" @click="addNewSps">确认</el-button>
        </el-dialog>
        <el-row :style="{marginTop:'20px', marginLeft:'20px'}">
            <component :is='spsMode'></component>
        </el-row>
    </div>
</template>

<script>
import spsSearch from "./SpsSearch.vue";
import spsHistory from "./SpsHistory.vue";
import spsStat from "./SpsStat.vue";

export default {
    data() {
        return {
            activeIndex: "1",
            dialogFormVisible: false,
            spsMode: 'sps-search',
        }
    },
    methods: {
        changeMode: function(index) {
            let modeMap = {"1":"sps-search", '2':'sps-history', '3':'sps-stat'};
            this.spsMode = modeMap[index]

        },
        addNewSps: function() {
            // get data
            // post data
            // then
            this.dialogFormVisible = false;
        },
    },
    components: {
        'sps-search': spsSearch,
        'sps-history': spsHistory,
        'sps-stat': spsStat,
    }
}
</script>

<style scoped>
.addbtn-container{
    height: 61px;
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>

