<template>
    <div>
        <div>
            <loveuer-menu ref="uMenu"></loveuer-menu>
        </div>

        <div style="margin-top:30px;">
            <el-row style="width:100%;">
                <!-- @keypress.enter.native="doSearch" -->
                <span>
                    <el-input
                        placeholder="请输入内容"
                        style="width:300px;margin-left:30px;"
                        @keypress.enter.native="doSearch"
                        clearable
                        v-model="searchKey">
                        <i slot="prefix" class="el-input__icon el-icon-search"></i>
                    </el-input>
                </span>
                <span v-show="sps" style="line-height:40px;margin-left:30px;">
                    <font>共搜索到: </font><font style="color:#F06560" v-text="searchedAmount"></font><font> 个结果</font>
                </span>
            </el-row>
            <el-row class="result-filter">
                <el-checkbox-group v-model="simsList" :min="1">
                    <el-checkbox label="5978"></el-checkbox>
                    <el-checkbox label="5989"></el-checkbox>
                    <el-checkbox label="5008"></el-checkbox>
                    <el-checkbox label="5015"></el-checkbox>
                </el-checkbox-group>
            </el-row>
        </div>
        <!-- content -->
        <div style="width:100%;min-width:1100px;margin-left:30px;">
            <el-table
                :data="filteredSps"
                height="700"
                @row-click="meetSpDetail"
                style="width:100%;">
                <el-table-column type="expand">
                    <template slot-scope="props">
                        <el-form label-position="left" class="demo-table-expand" label-width="120px">
                            <el-form-item label="原始模拟机"><span>{{ props.row.orgsim }}</span></el-form-item>
                            <el-form-item label="备注"><span>{{ props.row.comment }}</span></el-form-item>
                            <el-form-item label="标签"><span>{{ props.row.tags }}</span></el-form-item>
                        </el-form>
                    </template>
                </el-table-column>
                <el-table-column label="名称" prop="name" width="400"></el-table-column>
                <el-table-column label="P/N" prop="pn" width="350"></el-table-column>
                <el-table-column label="S/N" prop="sn" width="250"></el-table-column>
                <el-table-column label="状态" prop="status" width="100"></el-table-column>
                <el-table-column label="模拟机" prop="nowsim" width="100"></el-table-column>
                <el-table-column label="位置" prop="pos" width="150"></el-table-column>
            </el-table>
        </div>
        <div style="width:100%;height:120px;zoom:100%"></div>
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";

export default {
    data() {
        return {
            searchKey: "", 
            simsList: ['5978', '5989', '5008', '5015'],
            sps: [],
        };
    },
    methods: {
        doSearch: function() {
            this.$router.push("/works/spsRecorder/search/" + this.searchKey);
            console.log("search: ", this.searchKey);
            this.$http.get("/api/sps/search/" + this.searchKey)
                .then(resp => {
                    this.sps = resp.data;
                })
                .catch(error => { console.log("sps search http error: ", error.response); });
        },
        meetSpDetail: function(row) {
            this.$router.push("/works/spsRecorder/detail/" + row.id);
        },
    },
    computed: {
        filteredSps: function() {
            if (this.sps) {
                return this.sps.filter(sp => this.simsList.includes(sp.nowsim));
            } else {
                return [];
            };
        },
        searchedAmount: function() {
            if (this.sps) {
                return this.sps.length;
            } else {
                return 0;
            };
        },
    },
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/spsRecorder/search";
        if (this.$route.params.skey) {
            this.searchKey = this.$route.params.skey;
            this.doSearch();
        };
    },
    components: {
        "loveuer-menu": loveuerMenu,
    },
};
</script>

<style scoped>
.result-filter{
    margin-left: 32px;
    margin-top: 20px;
}
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
