<template>
    <div>
        <div>
            <loveuer-menu ref="uMenu"></loveuer-menu>
        </div>
        <div>
            <div style="width:100%;height:20px;zoom:100%"></div>
            <div class="container">
                <el-form
                    style="width:100%;"
                    label-width="100px"
                    label-position="left"
                    :model="newlog"
                    ref="form">
                    <el-form-item label="模拟机">
                        <el-select v-model="newlog.sim" placeholder="模拟机">
                            <el-option v-for="sim in selsims" :key="sim.value" :label="sim.label" :value="sim.value"></el-option>
                        </el-select>
                    </el-form-item>
                </el-form>
                <div style="height:100%;width:100%;display:flex;">
                    <div class="textarea-label">描述</div>
                    <el-popover
                        width="700"
                        trigger="manual"
                        v-model="popoverShow"
                        placement="right">
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
                                <el-button size="small" style="margin-left:30px;" @click="cancelInsert">取消</el-button>
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
                        <loveuer-textarea 
                            slot="reference" 
                            ref="uTextArea" 
                            style="width:500px;" 
                            v-on:findpn="readytoInsert">
                        </loveuer-textarea>
                    </el-popover>
                </div>
                <div class="submit-zone">
                    <div>
                        <el-button type="primary" @click="submitNewlog">提交</el-button>
                    </div>
                    <div>
                        <el-button>取消</el-button>
                    </div>
                </div>
            </div>
        </div>
        <div class="showLog" v-html="newlog.logs">

        </div>
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";
import quillTextArea from "../uTextArea.vue";

export default {
    data() {
        return {
            newlog: {
                sim: '',
                logs: '',
            },
            selsims: [
                {value:'5978',label:'5978'},
                {value:'5989',label:'5989'},
                {value:'5008',label:'5008'},
                {value:'5015',label:'5015'},
            ],
            popoverShow: false,
            searchSpsKey: '',
            searchedSps: [],
            findkey: '',
            showInputedLog: '',
        };
    },
    components: {
        "loveuer-menu": loveuerMenu,
        "loveuer-textarea": quillTextArea,
    },
    methods: {
        submitNewlog: function() {
            this.newlog.logs = this.$refs.uTextArea.content;
            console.log(this.newlog);
            if (this.newlog.sim === '' || this.newlog.logs === '') {
                this.$notify.error({
                    title: '警告',
                    message: '内容不完整!',
                    position:'top-left',
                });
                return;
            };
            this.$http.post('/api/log/add')
                .then(resp => {console.log(resp)})
                .catch(error => {console.log("err: "), err});
        },
        readytoInsert: function(find) {
            this.findkey = find;
            this.popoverShow = true;
        },
        insertSps: function(index, row) {
            this.popoverShow = false;
            this.searchedSps = [];
            this.searchSpsKey = '';
            // replace pn with row
            
            this.$refs.uTextArea.insertSps(row, this.findkey);
        },
        cancelInsert: function() {
            this.popoverShow = false;
            this.searchedSps = [];
            this.searchSpsKey = '';
            // replace pn
        },
        doSearchSps: function() {
            this.$http.get("/api/sps/search/" + this.searchSpsKey)
                .then(resp => { this.searchedSps = resp.data; })
                .catch(error => { console.log('log add search sps error: ', error.response) });
        },
    },
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/logRecorder/add";
    },
    computed: {
        
    },
};
</script>

<style scoped>
.container{
    margin-left: 30px;
    width: 80%;
}
.textarea-label{
    height:100%;
    width:100px;
    font-size:14px;
    color:#666;
    display: flex;
    align-items: center;
    margin-top: 10px;
}
.submit-zone{
    width:100%;
    display: flex;
    margin-top: 20px;
}
.submit-zone > div:first-child{
    margin-left:100px;
    margin-right: 30px;
}

.showLog{
    height:200px;
    width:50%;
    zoom:100%;
    background-color:#eee;
    margin-top:30px;
    margin-left:30px;
}
</style>
