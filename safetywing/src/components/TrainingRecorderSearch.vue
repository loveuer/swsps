<template>
    <div id="t-add">
        <el-row>
            <el-date-picker v-model="searchDate" type="date" placeholder="Pick a day" @change="getScanedListByDate"> </el-date-picker>
        </el-row>
        <el-row>
            
        </el-row>
    </div>
</template>

<script>
var dateFormat = require('dateformat');

export default {
    data() {
        return {
            searchDate: '',
        };
    },
    methods: {
        getScanedListByDate: function() {
            let d = dateFormat(this.searchDate, "yyyy-mm-dd");
            console.log("get the images by date: ", d);
        },
    },
    mounted() {
        this.searchDate = Date.now();
        this.getScanedListByDate();
    },
};
</script>

<style scoped>
#t-add{
    margin-top: 20px;
    margin-left: 20px;
}
</style>
