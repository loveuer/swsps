<template>
    <div id="container">
        <sps-menu :path="'/search'"></sps-menu>
        <div class="search-main">
            <el-card style="width:90%;margin-top:30px;">
                <div slot="header" class="clearfix">
                    <span>
                        <el-autocomplete
                            ref="searchkey_input"
                            style="width:400px;"
                            :fetch-suggestions="getSearchSuggestions"
                            placeholder="搜索备件"
                            @keypress.enter.native="doSearch"
                            @select="searchSelectEvent"
                            clearable
                            v-model="search_key"
                        >
                            <i slot="prefix" class="el-input__icon el-icon-search"></i>
                        </el-autocomplete>
                    </span>
                    <span style="margin-left:20px;">
                        <el-button type="primary" @click="doSearch">搜索</el-button>
                    </span>
                    <span style="margin-left:50px;">
                        <el-checkbox v-model="foldResult">折叠结果</el-checkbox>
                    </span>
                    <span style="margin-left:50px;">
                        <el-popover
                            trigger="hover"
                            placement="bottom"
                            width="400"
                        >
                            <div class="pop-filter">
                                <div class="filter-side">
                                    <div style="width:100%;height:30px;">模拟机</div>
                                    <div style="width:100%;"><el-checkbox v-model="filter.sim5978">5978</el-checkbox></div>
                                    <div style="width:100%;margin-top:5px;"><el-checkbox v-model="filter.sim5989">5989</el-checkbox></div>
                                    <div style="width:100%;margin-top:5px;"><el-checkbox v-model="filter.sim5008">5008</el-checkbox></div>
                                    <div style="width:100%;margin-top:5px;"><el-checkbox v-model="filter.sim5015">5015</el-checkbox></div>
                                </div>
                                <div class="filter-side">
                                    <div style="width:100%;height:30px;">状态</div>
                                    <div style="width:100%"><el-checkbox v-model="filter.sts_bj">备件</el-checkbox></div>
                                    <div style="width:100%;margin-top:5px;"><el-checkbox v-model="filter.sts_sy">使用</el-checkbox></div>
                                    <div style="width:100%;margin-top:5px;"><el-checkbox v-model="filter.sts_wx">维修</el-checkbox></div>
                                    <div style="width:100%;margin-top:5px;"><el-checkbox v-model="filter.sts_fq">废弃</el-checkbox></div>
                                </div>
                                <div style="width:50%;text-align:center;margin-top:10px;">
                                    <el-button size="small" @click="filter2Default">恢复默认</el-button>
                                </div>
                            </div>
                            <el-button slot="reference" size="small">过滤器</el-button>
                        </el-popover>
                    </span>
                    
                </div>
                <div style="width:100%;">
                    <ext-table :sps="filteredSPS" v-if="!foldResult"></ext-table>
                    <folded-table :sps="filteredSPS" v-if="foldResult"></folded-table>
                </div>
            </el-card>
        </div>
        <div style="height:100px;width:100%;zoom:100%;"></div>
    </div>
</template>

<script>
import SpsMenu from "../Menu.vue";
import foldedTable from "./Search_fold.vue";
import extTable from "./Search_ext.vue";

export default {
    data() {
        return {
            foldResult: false,
            filter: {sim5978: true, sim5989: true, sim5008: true, sim5015: true, sts_bj: true, sts_sy: true, sts_wx: true, sts_fq: false},
            search_key: "",
            searchedSps: []
        };
    },
    computed: {
        'filteredSPS': function() {
            let newsps = this.searchedSps.filter(one => {
                if (this.filterArray.sim.includes(one.nowsim) && this.filterArray.sts.includes(one.status)) {
                    return one;
                }
            });
            return newsps;
        },
        'filterArray': function() {
            let fltary = {sim: [], sts: []};
            if (this.filter.sim5978) {fltary.sim.push('5978')};
            if (this.filter.sim5989) {fltary.sim.push('5989')};
            if (this.filter.sim5008) {fltary.sim.push('5008')};
            if (this.filter.sim5015) {fltary.sim.push('5015')};
            if (this.filter.sts_bj) {fltary.sts.push('备件')};
            if (this.filter.sts_sy) {fltary.sts.push('使用')};
            if (this.filter.sts_wx) {fltary.sts.push('维修')};
            if (this.filter.sts_fq) {fltary.sts.push('废弃')};
            return fltary;
        },
    },
    methods: {
        getSearchSuggestions(str, cb) {
            let suggestions = [];
            this.$http.get("/api/user/searchhis/last")
                .then(resp => {
                    for(let h of resp.data) {
                        suggestions.push({value:h});
                    };
                })
                .catch(err => {
                    switch (err.response.status) {
                        case 401:
                            this.$router.push("/login");
                            break;
                        default:
                            console.log(err.response);
                            break;
                    };
                });
            cb(suggestions);
            return;
        },
        doSearch: function() {
            this.$refs.searchkey_input.close();
            if (this.search_key === "") {
                this.$message({showClose: true,message: "搜索关键字不能为空",type: "warning"});
                return;
            };
            
            this.get_sps_by_key();
            console.log(this.search_key);
            this.$router.replace("/search/" + this.search_key);
        },
        get_sps_by_key: function() {
            if (!this.search_key) {
                return;
            };
            this.$http.get(`/api/sps/search/${this.search_key}`)
                .then(resp => {
                    if (resp.data instanceof Array) {
                        this.searchedSps = resp.data;
                    } else {
                        this.$message({showClose: true,message: "发生错误,请重试",type: "warning"});
                        console.log("wrong resp data type: ", typeof(resp.data));
                    };
                })
                .catch(err => {
                    switch (err.response.status) {
                        case 401:
                            this.$router.push("/login");
                            break;
                    };
                });
        },
        searchSelectEvent: function() {
            this.doSearch();
        },
        
        filter2Default: function() {
            this.filter = {sim5978: true, sim5989: true, sim5008: true, sim5015: true, sts_bj: true, sts_sy: true, sts_wx: true, sts_fq: false};
        },
    },
    mounted() {
        if (this.$route.params.key !== "") {
            this.search_key = this.$route.params.key;
            this.get_sps_by_key();
        };
    },
    components: {
        "sps-menu": SpsMenu,
        "folded-table": foldedTable,
        "ext-table": extTable,
    }
};
</script>

<style>
.search-main { 
    width: 100%;
    display: flex;
    justify-content: center;
}

.pop-filter {
    width:100%;
    display: flex;
    flex-wrap: wrap;
}
.filter-side {
    width:49%;
    text-align: center;
}
</style>
