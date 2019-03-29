<template>
    <div class="container">
        <el-popover
            trigger="manual"
            v-model="popoverShow"
            placement="right"
            width="700">
            <div>
                <el-row>
                    <el-input
                        @keypress.enter.native="doSearchSps"
                        size="small"
                        prefix-icon="el-icon-search"
                        clearable
                        placeholder="搜索备件用以插入log"
                        style="width:300px;"
                        v-model="searchSpsKey">
                    </el-input>
                    <span style="margin-left:20px;">
                        <font>搜索到 "</font><font style="color:#F06560;">{{ searchedSpsAmount }}</font><font>" 个备件</font>
                    </span>
                    <el-button size="small" style="margin-left:30px;float:right;" @click="cancelInsert">取消</el-button>
                </el-row>
                <el-row>
                    <el-table
                        height="300"
                        style="width:100%;"
                        :data="searchedSps">
                        <el-table-column width="200" label="名称" prop="name"></el-table-column>
                        <el-table-column width="180" label="P/N" prop="pn"></el-table-column>
                        <el-table-column width="120" label="S/N" prop="sn"></el-table-column>
                        <el-table-column width="70" label="模拟机" prop="nowsim"></el-table-column>
                        <el-table-column align="right">
                            <template slot-scope="scope">
                                <el-button size="mini" @click="insertSps(scope.$index, scope.row)">插入</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-row>
            </div>
            <quill-textarea
                ref="quill"
                :style="{width:width+'px'}"
                :preok="preok"
                v-on:findpn="readytoInsert"
                @change="updateContent"
                slot="reference">
            </quill-textarea>
        </el-popover>
    </div>
</template>

<script>
import quillTextArea from "./uPackedQuill.vue";

export default {
    data() {
        return {
            popoverShow: false,
            searchSpsKey: '',
            searchedSps: [],
            findkey: '',
        };
    },
    props: {
        'width': Number,
        'preok': Boolean,
    },
    computed: {
        searchedSpsAmount: function() {
            return this.searchedSps.length;
        },
    },
    methods: {
        cancelInsert: function() {
            this.popoverShow = false;
            this.searchedSps = [];
            this.searchSpsKey = '';

        },
        readytoInsert: function(find) {
            this.findkey = find;
            this.popoverShow = true;
        },
        doSearchSps: function() {
            this.$http.get("/api/sps/search/" + this.searchSpsKey)
                .then(resp => { this.searchedSps = resp.data || []; })
                .catch(error => { console.log('log add search sps error: ', error.response) });
        },
        insertSps: function(index, row) {
            this.popoverShow = false;
            this.searchedSps = [];
            this.searchSpsKey = '';
            // replace pn with row
            this.$refs.quill.insertSps(row, this.findkey);
        },
        updateContent: function(c) {
            this.$emit('change', c);
        }
    },
    components: {
        "quill-textarea": quillTextArea,
    },
};
</script>

<style scoped>
.container {
    width: 1200px;
}
</style>

