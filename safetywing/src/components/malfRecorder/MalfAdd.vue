<template>
    <div>
        <div>
            <loveuer-menu ref="uMenu"></loveuer-menu>
        </div>
        <div style="width:100%;height:30px;zoom:100%;"></div>
        <div class="malfadd-container">
            <div style="margin-left:50px;width:330px;">
                <el-form :model="newmalf" label-position="top" label-width="300" style="width:300px;">
                    <el-form-item label="模拟机" style="width:100%;">
                        <el-select v-model="newmalf.sim" style="width:100%;">
                            <el-option label="5978" value="5978"></el-option>
                            <el-option label="5989" value="5989"></el-option>
                            <el-option label="5008" value="5008"></el-option>
                            <el-option label="5015" value="5015"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="值班人员">
                        <el-popover placement="right" width="200" v-model="showAddCoworkers">
                            <p>添加值班人员</p>
                            <div>
                                <div>
                                    <el-select v-model="prepareCoworkers">
                                        <el-option label="罗宇" value="罗宇"></el-option>
                                        <el-option label="陈本志" value="陈本志"></el-option>
                                        <el-option label="付东明" value="付东明"></el-option>
                                        <el-option label="陈俊峰" value="陈俊峰"></el-option>
                                        <el-option label="郭建" value="郭建"></el-option>
                                        <el-option label="胡彬彬" value="胡彬彬"></el-option>
                                        <el-option label="李津" value="李津"></el-option>
                                        <el-option label="李跃" value="李跃"></el-option>
                                        <el-option label="鲁书贤" value="鲁书贤"></el-option>
                                        <el-option label="张显刚" value="张显刚"></el-option>
                                        <el-option label="杨宏志" value="杨宏志"></el-option>
                                        <el-option label="袁超" value="袁超"></el-option>
                                        <el-option label="张昊宇" value="张昊宇"></el-option>
                                        <el-option label="张伟" value="张伟"></el-option>
                                        <el-option label="凌华南" value="凌华南"></el-option>
                                        <el-option label="刘博强" value="刘博强"></el-option>
                                        <el-option label="郑卿" value="郑卿"></el-option>
                                        <el-option label="赵育鹏" value="赵育鹏"></el-option>
                                        <el-option label="宋继朋" value="宋继朋"></el-option>
                                        <el-option label="李永翔" value="李永翔"></el-option>
                                        <el-option label="姚刚" value="姚刚"></el-option>
                                        <el-option label="江富强" value="江富强"></el-option>
                                        <el-option label="魏宇东" value="魏宇东"></el-option>
                                        <el-option label="黄一林" value="黄一林"></el-option>
                                        <el-option label="杨爱俊" value="杨爱俊"></el-option>
                                        <el-option label="彭伟" value="彭伟"></el-option>
                                    </el-select>
                                </div>
                                <div style="margin-top:12px;">
                                    <el-button type="success">添加</el-button>
                                    <el-button @click="showAddCoworkers=false">取消</el-button>
                                </div>
                            </div>
                            <div slot="reference">
                                <el-input v-model="newmalf.coworkers"></el-input>
                            </div>
                        </el-popover>
                    </el-form-item>
                    <el-form-item label="中断时间">
                        <div style="display:flex;">
                            <div style="width:250px;height:100%;">
                                <el-input></el-input>
                            </div>
                            <div style="width:50px;height:100%;margin-left:auto;text-align:right;">小时</div>
                        </div>
                    </el-form-item>
                    <el-form-item label="等级">
                        <el-select v-model="newmalf.class" style="width:100%;">
                            <el-option label="CAT Ⅰ" value="1"></el-option>
                            <el-option label="CAT Ⅱ" value="2"></el-option>
                            <el-option label="CAT Ⅲ" value="3"></el-option>
                            <el-option label="CAT Ⅳ" value="4"></el-option>
                        </el-select>
                    </el-form-item>
                </el-form>
            </div>
            <div class="newmalf-content">
                <div class="content-zone">
                    <div class="content-label">问题描述</div>
                    <div style="margin-top:22px;">
                        <loveuer-textarea
                            :width="400"
                            :preok="false"
                            @change="updateContent"
                            >
                        </loveuer-textarea>
                    </div>
                </div>
                <div style="margin-left:30px;margin-top:90px;">
                    <el-button type="success">提交</el-button>
                    <el-button>取消</el-button>
                </div>
            </div>  
        </div>
            
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";
import uTextAreaWithSps from "../uTextArea.vue";

export default {
    data() {
        return {
            newmalf: {
                sim: '',
                coworkers: '',
                bt: '', // break time
                class: '',
                content: '',
            },
            showAddCoworkers: false,
            prepareCoworkers: '',
            popoverShow: false,
            searchSpsKey: '',
            searchedSps: [],
            searchedSpsAmount: 0,
        };
    },
    methods: {
        readytoInsert: function() {
            this.findkey = find;
            this.popoverShow = true;
        },
        doSearchSps: function() {
            this.$http.get("/api/sps/search/" + this.searchSpsKey)
                .then(resp => { this.searchedSps = resp.data || []; })
                .catch(error => { console.log('log add search sps error: ', error.response) });
        },
        cancelInsert: function() {
            this.popoverShow = false;
            this.searchedSps = [];
            this.searchSpsKey = '';
        },
        insertSps: function(index, row) {
            this.popoverShow = false;
            this.searchedSps = [];
            this.searchSpsKey = '';
            // replace pn with row
            
            this.$refs.uTextArea.insertSps(row, this.findkey);
        },
        updateContent: function(content) {
            //
        }
    },
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/malfRecorder/add";
    },
    components: {
        "loveuer-textarea": uTextAreaWithSps,
        "loveuer-menu": loveuerMenu,
    },
};
</script>

<style scoped>
.malfadd-container {
    width: 100%;
    height: 100%;
    display: flex;
}
.newmalf-content {
    min-width:700px;
    width: 100%;
    margin-left: 30px;
}
.content-zone {
    margin-left: 30px;
}
.content-label {
    margin-top: 10px;
    font-size: 14px;
}
</style>
