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
        <el-row style="margin-top:10px;margin-left:45px;max-width:1150px;width:1100px;">
            <el-collapse :value="activeCollapses">
                <el-collapse-item :title="'搜索名称的结果:  ' + filteredSpsListByName.length" name="1" v-show="spsListByName">
                    <el-table
                        :data="filteredSpsListByName"
                        @row-click='handleCurrentClick'
                        style="width:100%;">
                        <el-table-column prop="name" label="名称" width="350"></el-table-column>
                        <el-table-column prop="pn" label="P/N" width="250"></el-table-column>
                        <el-table-column prop="sn" label="S/N" width="200"></el-table-column>
                        <el-table-column prop="pos" label="位置" width="200"></el-table-column>
                        <el-table-column prop="nowsim" label="模拟机" width="100"></el-table-column>
                    </el-table>
                </el-collapse-item>
                <el-collapse-item :title="'搜索P/N的结果:  ' + filteredSpsListByPN.length" name="2" v-show="spsListByPN">
                    <el-table
                        :data="filteredSpsListByPN"
                        @row-click='handleCurrentClick'
                        style="width:100%;">
                        <el-table-column prop="name" label="名称" width="350"></el-table-column>
                        <el-table-column prop="pn" label="P/N" width="250"></el-table-column>
                        <el-table-column prop="sn" label="S/N" width="200"></el-table-column>
                        <el-table-column prop="pos" label="位置" width="200"></el-table-column>
                        <el-table-column prop="sim" label="模拟机" width="100"></el-table-column>
                    </el-table>
                </el-collapse-item>
                <el-collapse-item :title="'搜索S/N的结果:  ' + filteredSpsListBySN.length" name="3" v-show="spsListBySN">
                    <el-table
                        :data="filteredSpsListBySN"
                        @row-click='handleCurrentClick'
                        style="width:100%;">
                        <el-table-column prop="name" label="名称" width="350"></el-table-column>
                        <el-table-column prop="pn" label="P/N" width="250"></el-table-column>
                        <el-table-column prop="sn" label="S/N" width="200"></el-table-column>
                        <el-table-column prop="pos" label="位置" width="200"></el-table-column>
                        <el-table-column prop="sim" label="模拟机" width="100"></el-table-column>
                    </el-table>
                </el-collapse-item>
                <el-collapse-item :title="'搜索备注的结果:  ' + filteredSpsListByComment.length" name="4" v-show="spsListByComment">
                    <el-table
                        :data="filteredSpsListByComment"
                        @row-click='handleCurrentClick'
                        style="width:100%;">
                        <el-table-column prop="name" label="名称" width="350"></el-table-column>
                        <el-table-column prop="pn" label="P/N" width="250"></el-table-column>
                        <el-table-column prop="sn" label="S/N" width="200"></el-table-column>
                        <el-table-column prop="pos" label="位置" width="200"></el-table-column>
                        <el-table-column prop="sim" label="模拟机" width="100"></el-table-column>
                    </el-table>
                </el-collapse-item>
            </el-collapse>
        </el-row>
        <!-- 以上是展示sps -->
        <el-dialog title="备件详情" :visible.sync="dialogFormVisible" style="min-width:1100px;">
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
            // activeCollapses: ['1', '2', '3', '4'],
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
        },
        activeCollapses: function() {
            let activecollapsesList = [];
            if (this.spsListByName.length > 0) {
                activecollapsesList.push("1");
            };
            if (this.spsListByPN.length > 0) {
                activecollapsesList.push('2');
            };
            if (this.spsListBySN.length > 0) {
                activecollapsesList.push('3');
            };
            if (this.spsListByComment.length > 0) {
                activecollapsesList.push('4');
            };
            return activecollapsesList;
        },
    },
    methods: {
        doSearch: function() {
            let skey = this.searchKeyWord.trim();
            if (skey === '') {
                console.log("empty search key: " + this.searchKeyWord);
                return;
            };
            this.$http.get('/api/sps/search/' + skey)
                .then(resp => {
                    this.spsListByName = resp.data.byname || [];
                    this.spsListByPN = resp.data.bypn || [];
                    this.spsListBySN = resp.data.bysn || [];
                    this.spsListByComment = resp.data.bycomment || [];
                })
                .catch(error => { console.log(error.response) });
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
            this.$http.get('/api/sps/one/' + this.$store.state.spsOne.id)
                .then(resp => { this.$store.commit('updateSpsOne', resp.data); })
                .catch(error => { console.log(error.response); });
        },
    },
};
 </script>

 