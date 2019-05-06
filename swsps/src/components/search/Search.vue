<template>
    <div id="container">
        <sps-menu :path="'/search'"></sps-menu>
        <div class="search-main">
            <el-card style="width:90%;margin-top:30px;">
                <div slot="header" class="clearfix">
                    <span>
                        <el-autocomplete
                            style="width:400px;"
                            :fetch-suggestions="getSearchSuggestions"
                            placeholder="搜索备件"
                            @keypress.enter.native="doSearch"
                            v-model="search_key">
                            <i slot="prefix" class="el-input__icon el-icon-search"></i>
                        </el-autocomplete>
                    </span>
                    <span style="margin-left:20px;">
                        <el-button type="primary" @click="doSearch">搜索</el-button>
                    </span>
                </div>
                <div style="width:100%;">
                    <el-table :data="searchedSps" style="font-size:14px;width:100%;cursor:pointer" @row-click="handleRowClick">
                        <el-table-column type="expand">
                            <template slot-scope="props">
                                <el-form label-position="left" class="demo-table-expand">
                                    <el-form-item label="位置">
                                        <span>{{ props.row.pos }}</span>
                                    </el-form-item>
                                    <el-form-item label="备注">
                                        <span>{{ props.row.comment }}</span>
                                    </el-form-item>
                                </el-form>
                            </template>
                        </el-table-column>
                        <el-table-column prop="name" label="Name"></el-table-column>
                        <el-table-column prop="pn" label="P/N"></el-table-column>
                        <el-table-column prop="sn" label="S/N"></el-table-column>
                        <el-table-column prop="sim" label="模拟机"></el-table-column>
                        <el-table-column prop="amount" label="数量"></el-table-column>
                    </el-table>
                </div>
            </el-card>
        </div>
    </div>
</template>

<script>
import SpsMenu from '../Menu.vue';

export default {
    data() {
        return {
            search_key: '',
            searchedSps: [],
        };
    },
    methods: {
        getSearchSuggestions(str, cb) {
            let mock = [
                {value:'suggestion 1'},
                {value:'suggestion 2'},
                {value:'suggestion 3'},
            ];
            cb(mock);
            return;
        },
        doSearch: function() {
            if (this.search_key === '') {
                this.$message({
                    showClose: true,
                    message: '搜索关键字不能为空', 
                    type: 'warning',
                });
                return;
            };
            let mocksps = [
                {id: 1, name:'mock name 1', pn:'mock pn 1', sn: 'mock sn 1', sim: '5978', amount:'1', pos: 'A-1-1', comment: 'mock comment 1'},
                {id: 2, name:'mock name 2', pn:'mock pn 2', sn: 'mock sn 2', sim: '5978', amount:'2', pos: 'A-1-2', comment: 'mock comment 2'},
                {id: 3, name:'mock name 3', pn:'mock pn 3', sn: 'mock sn 3', sim: '5978', amount:'3', pos: 'A-1-3', comment: 'mock comment 3'},
                {id: 4, name:'mock name 4', pn:'mock pn 4', sn: 'mock sn 4', sim: '5978', amount:'1', pos: 'A-2-1', comment: 'mock comment 4'},
            ];
            this.searchedSps = mocksps;
        },
        handleRowClick: function(row) {
            console.log("clicked sp id: ", row.id);
        },
    },
    components: {
        'sps-menu': SpsMenu,
    },
};
</script>

<style scoped>

.search-main {
    width: 100%;
    display: flex;
    justify-content: center;
}
</style>
