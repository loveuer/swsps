<template>
    <div>
        <el-row :style="{marginTop:'-30px'}">
            <el-menu :default-active="this.activeIndex" class="el-menu-demo" mode="horizontal" @select="changeMode">
                    <el-menu-item index="1">备件详情</el-menu-item>
                    <el-menu-item index="2">备件历史</el-menu-item>
                    <el-menu-item index="3">编辑备件</el-menu-item>
            </el-menu>
        </el-row>
        <el-row>
            <component :is="spsDetailMode"></component>
        </el-row>
    </div>
</template>

<script>
import spD_Detail from './SpsD_Detail.vue';
import spD_History from './SpsD_History.vue';
import spD_Modify from './SpsD_Modify.vue';

export default {
    data() {
        return {
            activeIndex: '1',
            spsDetailMode: 'sps-d-detail',
        };
    },
    methods: {
        changeMode: function(index) {
            let modeMap = {'1':'sps-d-detail', '2':'sps-d-history', '3':'sps-d-modify'};
            this.spsDetailMode = modeMap[index];
        },
        
    },
    components: {
        'sps-d-detail': spD_Detail,
        'sps-d-history': spD_History, 
        'sps-d-modify': spD_Modify,
    },
    mounted() {
        // get sps detail by id: this.$store.state.spsOne.id
        let mock_sps = {
            id: this.$store.state.spsOne.id, name: '我是一个备件', pn: '我是pn', sn: '我不是sn!', sim: '6666',
        };
        this.$store.commit('updateSpsOne', mock_sps);
    },
};
</script>
