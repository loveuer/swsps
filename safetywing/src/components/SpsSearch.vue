<template>
    <div>
        <el-row>
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
        <el-row>
            <el-table
                :data="spsList"
                @current-change="handleCurrentClick"
                >
                <el-table-column prop="name" label="名称"></el-table-column>
                <el-table-column prop="pn" label="P/N"></el-table-column>
                <el-table-column prop="sn" label="S/N"></el-table-column>
                <el-table-column prop="sim" label="模拟机"></el-table-column>
            </el-table>
        </el-row>
        <el-dialog title="备件详情" :visible.sync="dialogFormVisible">
            <div>
                <sps-detail ></sps-detail>
            </div>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogFormVisible = false">取 消</el-button>
                <el-button type="primary" @click="dialogFormVisible = false">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>
 <script>
import SpsDetail from "./SpsDetail.vue";

 export default {
     data() {
         return {
            searchKeyWord: "",
            spsList: [],
            dialogFormVisible: false,
            ifspsDetailDone: false,
         }
     },
     methods: {
         doSearch: function() {
             // do get
             // then
             console.log("search: " + this.searchKeyWord);
             this.spsList = [
                 {'id':'3',   'name':'sps one',   'pn':'pn-001', 'sn':'sn-1234', 'sim':'5978'},
                 {'id':'23',  'name':'sps two',   'pn':'pn-002', 'sn':'sn-2345', 'sim':'5989'},
                 {'id':'467', 'name':'sps three', 'pn':'pn-003', 'sn':'sn-3456', 'sim':'5015'},
                 {'id':'12',  'name':'sps four',  'pn':'pn-004', 'sn':'sn-4567', 'sim':'5978'},
                 {'id':'39',  'name':'sps five',  'pn':'pn-005', 'sn':'sn-9876', 'sim':'5008'},
             ]
         },
         handleCurrentClick: function(row) {
            this.$store.state.spsOne.id = row['id'];
            setTimeout(() => {
                this.dialogFormVisible = true;
            }, 100);
         },
     },
     components: {
         'sps-detail': SpsDetail,
     },
 };
 </script>

 