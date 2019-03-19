<template>
    <div>
        <el-row style="margin-left:45px;">
            <el-input
                placeholder="请输入关键字"
                prefix-icon="el-icon-search"
                v-model="searchKeyWord"
                :style="{width:'300px'}"
                @keypress.enter.native="doSearch"
                clearable
                >
            </el-input>
        </el-row>
        <!-- sps 模拟机 filter -->
        <el-row style="margin-top:20px;margin-left:45px;">
            <el-checkbox-group v-model="simsList">
                <el-checkbox label="5978">5978</el-checkbox>
                <el-checkbox label="5989">5989</el-checkbox>
                <el-checkbox label="5008">5008</el-checkbox>
                <el-checkbox label="5015">5015</el-checkbox>
            </el-checkbox-group>
        </el-row>
        <!-- 展示sps -->
        <el-row style="margin-top:10px;margin-left:45px;">
            <el-collapse v-model="activeCollapses">
                <el-collapse-item title="搜索名称的结果: " name="1">
                    <el-table
                        :data="filteredSpsListByName"
                        @row-click='handleCurrentClick'
                        style="width:100%;">
                        <el-table-column prop="name" label="名称"></el-table-column>
                        <el-table-column prop="pn" label="P/N"></el-table-column>
                        <el-table-column prop="sn" label="S/N"></el-table-column>
                        <el-table-column prop="pos" label="位置"></el-table-column>
                        <el-table-column prop="sim" label="模拟机"></el-table-column>
                    </el-table>
                </el-collapse-item>
                <el-collapse-item title="搜索P/N的结果: " name="2">
                    <el-table
                        :data="filteredSpsListByPN"
                        @row-click='handleCurrentClick'
                        style="width:100%;">
                        <el-table-column prop="name" label="名称"></el-table-column>
                        <el-table-column prop="pn" label="P/N"></el-table-column>
                        <el-table-column prop="sn" label="S/N"></el-table-column>
                        <el-table-column prop="pos" label="位置"></el-table-column>
                        <el-table-column prop="sim" label="模拟机"></el-table-column>
                    </el-table>
                </el-collapse-item>
                <el-collapse-item title="搜索S/N的结果: " name="3">
                    <el-table
                        :data="filteredSpsListBySN"
                        @row-click='handleCurrentClick'
                        style="width:100%;">
                        <el-table-column prop="name" label="名称"></el-table-column>
                        <el-table-column prop="pn" label="P/N"></el-table-column>
                        <el-table-column prop="sn" label="S/N"></el-table-column>
                        <el-table-column prop="pos" label="位置"></el-table-column>
                        <el-table-column prop="sim" label="模拟机"></el-table-column>
                    </el-table>
                </el-collapse-item>
                <el-collapse-item title="搜索备注的结果: " name="4">
                    <el-table
                        :data="filteredSpsListByComment"
                        @row-click='handleCurrentClick'
                        style="width:100%;">
                        <el-table-column prop="name" label="名称"></el-table-column>
                        <el-table-column prop="pn" label="P/N"></el-table-column>
                        <el-table-column prop="sn" label="S/N"></el-table-column>
                        <el-table-column prop="pos" label="位置"></el-table-column>
                        <el-table-column prop="sim" label="模拟机"></el-table-column>
                    </el-table>
                </el-collapse-item>
            </el-collapse>
        </el-row>
        <!-- 以上是展示sps -->
        <el-dialog title="备件详情" :visible.sync="dialogFormVisible">
            <div>
                <sps-detail></sps-detail>
            </div>
        </el-dialog>
        <div style="width:100%;height:200px;zoom:100%;"></div>
    </div>
</template>
 <script>
import SpsDetail from "./SpsDetail.vue";

 export default {
    data() {        
        return {
            searchKeyWord: "",
            spsListByName: [],
            spsListByPN: [],
            spsListBySN: [],
            spsListByComment: [],
            dialogFormVisible: false,
            ifspsDetailDone: false,
            simsList: ['5978', '5989', '5008', '5015'],
            activeCollapses: ['1', '2', '3', '4'],
        };
    },
    computed: {
        filteredSpsListByName: function() {
            // return this.spsList.filter(sp => this.simsList.includes(sp.sps.nowsim));
            return this.spsListByName.filter(sp => this.simsList.includes(sp.nowsim));
        },
        filteredSpsListByPN: function() {
            return this.spsListByPN.filter(sp => this.simsList.includes(sp.nowsim));
        },
        filteredSpsListBySN: function() {
            return this.spsListBySN.filter(sp => this.simsList.includes(sp.nowsim));
        },
        filteredSpsListByComment: function() {
            return this.spsListByComment.filter(sp => this.simsList.includes(sp.nowsim));
        }
    },
    methods: {
        doSearch: function() {
             // do get
             // then
            console.log("search: " + this.searchKeyWord);
            let spsList = this.$store.state.mockSearchSps;
            this.spsListByName = spsList.byname;
            this.spsListByPN = spsList.bypn;
            this.spsListBySN = spsList.bysn;
            this.spsListByComment = spsList.bycomment;
        },
        handleCurrentClick: function(row) {
            this.$store.state.spsOne.id = row['id'];
            this.dialogFormVisible = true;
        },
    },
    components: {
        'sps-detail': SpsDetail,
    },
    mounted() {
        
    },
    watch: {
        "$store.state.spsOne.id": function() {
            console.log("watched");
            this.$http.get('/api/sps/' + this.$store.state.spsOne.id)
                .then(resp => {
                    this.$store.mutations('updateSpsOne', JSON.parse(resp.data.val));
                    console.log(JSON.parse(resp.data.val));
                    console.log(this.$store.state.spsOne);
                })
                .catch(error => {
                    console.log(error.response);
                });
        },
    },
};
 </script>

 