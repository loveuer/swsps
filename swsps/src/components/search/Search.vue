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
                </div>
                <div style="width:100%;">
                    <el-table
                        :data="searchedSps"
                        style="font-size:14px;width:100%;cursor:pointer"
                        @row-click="handleRowClick"
                    >
                        <el-table-column type="expand">
                            <template slot-scope="props">
                                <el-form label-position="left" class="demo-table-expand" label-width="100px">
                                    <el-form-item label="原模拟机" style="height:30px;">
                                        <span>{{ props.row.orgsim }}</span>
                                    </el-form-item>
                                    <el-form-item label="位　　置" style="height:30px;">
                                        <span>{{ props.row.pos }}</span>
                                    </el-form-item>
                                    <el-form-item label="状　　态" style="height:30px;">
                                        <span>{{ props.row.status }}</span>
                                    </el-form-item>
                                    <el-form-item label="备　　注" style="height:30px;">
                                        <span>{{ props.row.comment }}</span>
                                    </el-form-item>
                                </el-form>
                            </template>
                        </el-table-column>
                            <el-table-column prop="name" label="Name" min-width="300"></el-table-column>
                            <el-table-column prop="pn" label="P/N" min-width="250"></el-table-column>
                            <el-table-column prop="sn" label="S/N" min-width="200"></el-table-column>
                            <el-table-column prop="nowsim" label="模拟机" min-width="75"></el-table-column>
                            <el-table-column prop="amount" label="数量" min-width="75"></el-table-column>
                    </el-table>
                </div>
            </el-card>
        </div>
        <div style="height:100px;width:100%;zoom:100%;"></div>
    </div>
</template>

<script>
import SpsMenu from "../Menu.vue";

export default {
    data() {
        return {
            search_key: "",
            searchedSps: []
        };
    },
    methods: {
        getSearchSuggestions(str, cb) {
            let suggestions = [];
            this.$http.get("/api/user/searchhismost")
                .then(resp => {
                    for(let h of resp.data) {
                        suggestions.push({value:h.key});
                    };
                })
                .catch(err => {
                    console.log(err.response);
                });
            cb(suggestions);
            return;
        },
        doSearch: function() {
            this.$refs.searchkey_input.close();
                if (this.search_key === "") {
                    this.$message({
                        showClose: true,
                        message: "搜索关键字不能为空",
                        type: "warning"
                    });
                    return;
                }

            this.$http.get(`/api/search/${this.search_key}`)
                .then(resp => {
                    this.searchedSps = resp.data;
                })
                .catch(err => {
                    console.log(err.response);
                });
        },
        searchSelectEvent: function() {
            this.doSearch();
        },
        handleRowClick: function(row) {
            console.log("clicked sp id: ", row.id);
        }
    },
    mounted() {
        if (this.$route.params.key !== "") {
            this.search_key = this.$route.params.key;
            this.doSearch();
        }
    },
    components: {
        "sps-menu": SpsMenu
    }
};
</script>

<style scoped>
.search-main {
    width: 100%;
    display: flex;
    justify-content: center;
}
</style>
