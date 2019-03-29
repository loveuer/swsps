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
                    <div style="margin-left:-30px;">
                        <loveuer-textarea 
                            :width="500"
                            :preok="true"
                            @change="updateContent">
                        </loveuer-textarea>
                    </div>
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
        <div class="showLog" v-html="newlog.logs.html">

        </div>
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";
import uTextArea from "../uTextArea.vue";
import qs from "qs";

export default {
    data() {
        return {
            newlog: {
                sim: '',
                logs: {},
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
        "loveuer-textarea": uTextArea,
    },
    methods: {
        submitNewlog: function() {
            this.newlog.logs = this.$refs.uTextArea.content;
            if (this.newlog.sim === '' || this.newlog.logs === '') {
                this.$notify.error({
                    title: '警告',
                    message: '内容不完整!',
                    position:'top-left',
                });
                return;
            };
            let postData = {
                wsid: this.$store.state.ws.id,
                sim: this.newlog.sim,
                log: this.newlog.logs,
            };
            this.$http.post('/api/log/add', qs.stringify(postData))
                .then(resp => {console.log("post new log resp: ", resp)})
                .catch(error => {console.log("err: "), error});
        },
        updateContent: function(content) {
            this.newlog.logs = content;
            console.log(content);
        },
    },
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/logRecorder/add";
    },
    computed: {
        searchedSpsAmount: function() {
            return this.searchedSps.length;
        },
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
