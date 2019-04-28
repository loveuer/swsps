<template>
    <div id="container">
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <span>
                    <el-autocomplete 
                        v-model="search_key" 
                        placeholder="搜索备件" 
                        @keypress.enter.native="searchSps"
                        :fetch-suggestions="getSearchSuggestions"
                        @select="searchSelectEvent"
                        style="width:300px;">
                        <i slot="prefix" class="el-input__icon el-icon-search"></i>
                    </el-autocomplete>
                </span>
                <span>
                    <el-button type="primary" style="margin-left:20px;" @click="searchSps">搜 索</el-button>
                </span>
            </div>
            <div>
                <!-- 经常搜索的10个备件信息 -->
                <div class="mostsearch-notify">搜索次数最多:</div>
                <el-table :data="mostSearchedSps" style="font-size:12px;width:100%;">
                    <el-table-column prop="name" label="Name" width="300"></el-table-column>
                    <el-table-column prop="pn" label="P/N" width="300"></el-table-column>
                    <el-table-column prop="amount" label="数量" width="50"></el-table-column>
                </el-table>
            </div>
        </el-card>
    </div>
</template>

<script>
export default {
    data() {
        return {
            search_key: '',
            mostSearchedSps: [],
        };
    },
    methods: {
        searchSps() {
            console.log("search sps: ", this.search_key);
            if (this.search_key === "") {
                this.$notify({
                    title: '警告',
                    message: '搜索关键字不能为空', 
                    type: 'warning',
                    position: 'top-left',
                    duration: 2000,
                });
            };
        },
        searchSelectEvent() {
            this.searchSps();
        },
        getSearchSuggestions(str, cb) {
            let mock = [
                {value:'suggestion 1'},
                {value:'suggestion 2'},
                {value:'suggestion 3'},
            ];
            cb(mock);
            return;
        },
    },
    mounted() {
        this.mostSearchedSps = [
            {name:'MFD assembly', pn:'60001STM073-501', amount:'7'},
            {name:'MCDU - MULTI FUNCTION CTRL DISP UNIT', pn:'F4023-ASM0CDU-000', amount:'3'},
            {name:'MFD assembly', pn:'60001STM073-501', amount:'7'},
            {name:'MFD assembly', pn:'60001STM073-501', amount:'7'},
            {name:'MFD assembly', pn:'60001STM073-501', amount:'7'},
            {name:'MFD assembly', pn:'60001STM073-501', amount:'7'},
            {name:'MFD assembly', pn:'60001STM073-501', amount:'7'},
            {name:'MFD assembly', pn:'60001STM073-501', amount:'7'},
            {name:'MFD assembly', pn:'60001STM073-501', amount:'7'},
            {name:'MFD assembly', pn:'60001STM073-501', amount:'7'},
        ];
    },
};
</script>

<style scoped>
.mostsearch-notify {
    font-size: 14px;
    text-indent: 10px;
    font-weight: bold;
    margin-top: -5px;
}
</style>
