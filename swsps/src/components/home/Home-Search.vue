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
                        clearable
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
                <el-table :data="mostSearchedSps" style="font-size:12px;width:100%;cursor:pointer;" @row-click="handleRowClick">
                    <el-table-column prop="name" label="Name" width="300" :show-overflow-tooltip="true"></el-table-column>
                    <el-table-column prop="pn" label="P/N" width="300" :show-overflow-tooltip="true"></el-table-column>
                    <el-table-column prop="nowsim" label="现模拟机" width="90"></el-table-column>
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
                this.$message({
                    showClose: true,
                    message: '搜索关键字不能为空', 
                    type: 'warning',
                });
                return;
            };
            this.$router.push(`/search/${this.search_key}`);
        },
        searchSelectEvent() {
            this.searchSps();
        },
        getSearchSuggestions(str, cb) {
            let suggestions = [];
            this.$http.get(`/api/user/searchhis/most`)
                .then(resp => {
                    resp.data.forEach(one => {
                        suggestions.push({value:one.key});
                    });
                })
                .catch(err => {
                    switch (err.response.status) {
                        case 401:
                            this.$router.push('/login');
                            break;
                        case 500:
                            this.$message({showClose: true, message: "服务器发生未知错误", type: "warning"});
                            break;
                        default:
                            console.log(err.response);
                            break;
                };
                });
            cb(suggestions);            
            return;
        },
        handleRowClick: function(row) {
            this.$router.push(`/onesp/${row.id}`);
        },
    },
    mounted() {
        this.$http.get("/api/sps/clicked/most")
            .then(resp => {
                this.mostSearchedSps = resp.data;
            })
            .catch(err => {
                switch (err.response.status) {
                    case 401:
                        this.$router.push('/login');
                        break;
                    case 500:
                        this.$message({showClose: true, message: "服务器发生未知错误", type: "warning"});
                        break;
                    default:
                        console.log(err.response);
                        break;
                };
            })
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
