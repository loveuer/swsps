<template>
    <div>
        <el-row style="margin-left:45px;">
            <el-input 
                placeholder="搜索历史记录" 
                prefix-icon="el-icon-search" 
                clearable
                v-model="searchHisKey"
                @keypress.enter.native="doSearchHis"
                :style="{width:'300px'}">
            </el-input>
        </el-row>
        <div style="margin-top:20px;">
            <sps-his-timeline ref="spshistimeline"></sps-his-timeline>
        </div>
    </div>
</template>

<script>
import spsHisTimeline from "./SpsHisTimeline.vue";

export default {
    data() {
        return {
            searchHisKey: '',
            spsHistories: [],
        };
    },
    methods: {
        doSearchHis: function() {
            // this.$refs.spshistimeline.histories = this.$store.state.mockhis;
        },
    },
    mounted() {
        this.$http.get('/api/sphis')
            .then(resp => {
                console.log(resp.data);
                this.$refs.spshistimeline.histories = resp.data;
            })
            .catch(error => { console.log(error.response) });
    },
    components: {
        'sps-his-timeline': spsHisTimeline,
    }
};
</script>

<style scoped>
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
</style>
