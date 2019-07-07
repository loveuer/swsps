<template>
    <div id="container">
        <sps-menu :path="'/addsp'"></sps-menu>
        <div style="width:100%;height:30px;zoom:100%;"></div>
        <div class="main-container">
            <div style="width:65%;min-width:700px;">
                <el-card class="box-card">
                    <div slot="header" class="clearfix">
                        <span>新增备件</span>
                        <el-tooltip content="重置表单" placement="top">
                            <el-button icon="el-icon-refresh" circle style="float:right;margin-top:-8px;" @click="resetForm"></el-button>
                        </el-tooltip>
                    </div>
                    <div>
                        <el-form ref="form" :model="newsp" label-width="120px" style="width:100%;">
                            <el-form-item label="名称">
                                <el-autocomplete
                                    style="width:400px;"
                                    class="inline-input"
                                    v-model="newsp.name"
                                    required
                                    :fetch-suggestions="nameSuggestion"
                                >
                                </el-autocomplete>
                            </el-form-item>
                            <el-form-item label="P/N">
                                <el-autocomplete
                                    style="width:400px;"
                                    class="inline-input"
                                    v-model="newsp.pn"
                                    :fetch-suggestions="pnSuggestion"
                                >
                                </el-autocomplete>
                            </el-form-item>
                            <el-form-item label="S/N">
                                <el-input style="width:400px;" v-model="newsp.sn"></el-input>
                            </el-form-item>
                            <el-form-item label="原模拟机">
                                <el-select v-model="newsp.orgsim" style="width:400px;">
                                    <el-option v-for="item of sims" :key="item.index" :value="item.sim"></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item label="模拟机">
                                <el-select v-model="newsp.nowsim" style="width:400px;">
                                    <el-option v-for="item of sims" :key="item.index" :value="item.sim"></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item label="位置">
                                <el-input v-model="newsp.pos" style="width:400px;"></el-input>
                            </el-form-item>
                            <el-form-item label="状态">
                                <el-select v-model="newsp.status" style="width:400px;">
                                    <el-option value="备件"></el-option>
                                    <el-option value="使用"></el-option>
                                    <el-option value="维修"></el-option>
                                    <el-option value="废弃" style="color:#F56C6C;"></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item label="数量">
                                <el-input-number v-model="newsp.num" :min="0"></el-input-number>
                            </el-form-item>
                            <el-form-item label="备注">
                                <el-input type="textarea" style="width:400px;" :rows="3" v-model="newsp.comment"></el-input>
                            </el-form-item>
                            <el-form-item label="图片(最多两张)">
                                <el-button type="primary" size="small" @click="handleUploadImg">点击上传</el-button>
                                <input type="file" accept="image/jpeg,image/jpg" style="display:none;" id="newsp-img1" @change="img1Uploaded">
                                <input type="file" accept="image/jpeg,image/jpg" style="display:none;" id="newsp-img2" @change="img2Uploaded">
                                <div class="upload-img" v-show="newsp.img1">
                                    <div class="place-img">
                                        <img :src="newsp.img1" alt="">
                                    </div>
                                    <div class="place-describe">
                                        <div>IMG 1</div>
                                    </div>
                                    <label class="place-corner">
                                        <div class="upload-success">
                                            <i class="el-icon-check corner-check"></i>
                                        </div>
                                        <div class="delete-upload"  @click="deleteUpload('img1')">
                                            <i class="el-icon-close corner-close"></i>
                                        </div>
                                    </label>
                                </div>
                                <div class="upload-img" v-show="newsp.img2">
                                    <div class="place-img">
                                        <img :src="newsp.img2" alt="">
                                    </div>
                                    <div class="place-describe">
                                        <div>IMG 2</div>
                                    </div>
                                    <label class="place-corner">
                                        <div class="upload-success">
                                            <i class="el-icon-check corner-check"></i>
                                        </div>
                                        <div class="delete-upload" @click="deleteUpload('img2')">
                                            <i class="el-icon-close corner-close"></i>
                                        </div>
                                    </label>
                                </div>
                            </el-form-item>
                            <el-form-item>
                                <el-button type="success" @click="doUpload" :loading="uploading">新增备件</el-button>
                            </el-form-item>
                        </el-form>
                    </div>
                </el-card>
            </div>
        </div>
        <div style="width:100%;height:100px;zoom:100%;"></div>
    </div>
</template>

<script>
import spsMenu from "../Menu.vue";

export default {
    data() {
        return {
            newsp: { name: '', pn: '', sn: '', orgsim: '', nowsim: '', pos: '', status: '', num: 0, comment: '', img1: null, img2: null, },
            sims: [{index:1, sim:'5978'}, {index:2, sim:'5989'}, {index:3, sim:'5008'}, {index:4, sim:'5015'}],
            uploading: false,
        };
    },
    methods: {
        nameSuggestion: function(inputedStr, cb) {
            if (inputedStr.length < 3) {
                cb([]);
                return;
            };
            this.$http.get("/api/sps/autocomplete/name/" + inputedStr)
                .then(resp => {
                    let data = [];
                    if (resp.data) {
                        for(let i of resp.data) {
                            data.push({value: i.name});
                        };
                    };
                    cb(data);
                })
                .catch(err => {
                    switch (err.response.status) {
                        case 401:
                            this.$router.push("/login");
                            break;
                    
                        default:
                            console.log(err.response);
                            break;
                    };
                });
        },
        pnSuggestion: function(inputedStr, cb) {
            if (inputedStr.length < 4) {
                cb([]);
                return;
            };
            this.$http.get("/api/sps/autocomplete/pn/" + inputedStr)
                .then(resp => {
                    let data = [];
                    if (resp.data) {
                        for(let i of resp.data) {
                            data.push({value:i.pn});
                        };
                    };
                    cb(data);
                })
                .catch(err => {
                    switch (err.response.status) {
                        case 401:
                            this.$router.push("/login");
                            break;
                    
                        default:
                            console.log(err.response);
                            break;
                    };
                });
        },
        handleUploadImg: function() {
            if (this.newsp.img1 && this.newsp.img2) {
                this.$message({showClose: true, message: "已经选择了2张图片", type: "warning"});
                return;
            };
            if (this.newsp.img1) {
                document.querySelector('#newsp-img2').click();
            } else {
                document.querySelector('#newsp-img1').click();
            };
        },
        img1Uploaded: function() {
            let img1file = document.querySelector('#newsp-img1').files[0];
            let imgReader = new FileReader();
            imgReader.onload = (event) => {
                this.newsp.img1 = event.target.result;
            };
            imgReader.readAsDataURL(img1file);
        },
        img2Uploaded: function() {
            let img2file = document.querySelector('#newsp-img2').files[0];
            let imgReader = new FileReader();
            imgReader.onload = (event) => {
                this.newsp.img2 = event.target.result;
            };
            imgReader.readAsDataURL(img2file);
        },
        deleteUpload: function(which) {
            this.newsp[which] = null;
            document.querySelector(`#newsp-${which}`).value = '';
        },
        resetForm: function() {
            this.$confirm('此操作将重置你已经输入的表单, 是否继续?', '提示', {confirmButtonText: '确认', cancleButtonText: '取消', type: 'warning'})
                .then(() => {
                    let emptysp = { name: '', pn: '', sn: '', orgsim: '', nowsim: '', pos: '', status: '', num: 0, comment: '', img1: null, img2: null, };
                    this.newsp = emptysp;
                    document.querySelector('#newsp-img1').value = '';
                    document.querySelector('#newsp-img2').value = '';
                })
        },
        check_upload_data: function() {
            if (this.newsp.name === "") {
                return false;
            };
            if (this.newsp.pn === "") {
                return false;
            };
            if (this.newsp.sn === "") {
                return false;
            };
            if (this.newsp.orgsim === "") {
                return false;
            };
            if (this.newsp.nowsim === "") {
                return false;
            };
            if (this.newsp.pos === "") {
                return false;
            };
            if (this.newsp.status === "") {
                return false;
            };
            return true;
        },
        doUpload: function() {
            this.uploading = true;
            if (!this.check_upload_data()) {
                this.$message({showClose: true, message: "信息不完整", type: "warning"});
                this.uploading = false;
                return;
            };
            let uploadData = new FormData();
            uploadData.append('name', this.newsp.name);
            uploadData.append('pn', this.newsp.pn);
            uploadData.append('sn', this.newsp.sn);
            uploadData.append('orgsim', this.newsp.orgsim);
            uploadData.append('nowsim', this.newsp.nowsim);
            uploadData.append('pos', this.newsp.pos);
            uploadData.append('status', this.newsp.status);
            uploadData.append('amount', this.newsp.num);
            uploadData.append('comment', this.newsp.comment);
            if (this.newsp.img1) {
                uploadData.append('img1', document.querySelector('#newsp-img1').files[0]);
            };
            if (this.newsp.img2) {
                uploadData.append('img2', document.querySelector('#newsp-img2').files[0]);
            };
            this.$http.post('/api/sps/add', uploadData, {
                method: 'post',
                headers: {'Content-Type': 'multipart/form-data'},
                onUploadProgress: function(event) {
                    // console.log(event.loaded / event.total);
                },
            }).then( resp => {
                if (resp.data.id === 0) {
                    this.$message({showClose: true, message: "发生未知错误, 请联系管理员", type: "warning"});
                    let emptysp = { name: '', pn: '', sn: '', orgsim: '', nowsim: '', pos: '', status: '', num: 0, comment: '', img1: null, img2: null, };
                    this.newsp = emptysp;
                    document.querySelector('#newsp-img1').value = '';
                    document.querySelector('#newsp-img2').value = '';
                    return;
                };
                this.$router.push(`/onesp/${resp.data.id}`);
            }).catch(err => {
                switch (err.response.status) {
                    case 401:
                        this.$router.push("/login");
                        break;
                    default:
                        console.log(err.response);
                        this.$message({showClose: true, message: "发生未知错误, 请联系管理员", type: "warning"});
                        this.uploading = false;
                        break;
                };
            });
        },
    },
    computed: {
        
    },
    components: {
        'sps-menu': spsMenu,
    }
};
</script>

<style scoped>
.main-container{
    width: 100%;
    display: flex;
    justify-content: center;
}
.upload-img {
    border: 1px solid rgb(220,223,230);
    border-radius: 5px;
    padding: 10px;
    width: 380px;
    height: 100px;
    margin-top: 10px;
    position: relative;
    overflow: hidden;
    display: flex;
}
.upload-img:hover {
    border: 1px solid rgb(192,196,204);
}
.upload-img:hover .upload-success {
    display: none;
}
.upload-img:hover .place-corner {
    background: #fff;
    box-shadow: none;
}
.upload-img:hover .delete-upload {
    display: block;
}
.place-img {
    height: 100%;
    width: 150px;
}
.place-img > img {
    width: 150px;
    height: 100%;
}
.place-describe {
    width: 100px;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}
.place-corner {
    position: absolute;
    top: -7px;
    right: -17px;
    width: 46px;
    height: 26px;
    background: #13ce66;
    transform: rotate(45deg);
    text-align: center;
    box-shadow: 0 1px 1px #ccc;
    display: block;
}
.upload-success{
    margin-top: -2px;
}
.corner-check {
    transform: rotate(-45deg);
    color: #fff;
    margin-top: 12px;
    font-size: 12px;
}
.delete-upload {
    display: none;
}
.corner-close {
    transform: rotate(-45deg);
    margin-top: 12px;
    font-size: 12px;
    cursor: pointer;
}
</style>
