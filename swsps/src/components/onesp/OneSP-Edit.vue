<template>
    <div id="container">
        <div class="editsp-form">
            <div class="one-row">
                <div class="label">Name</div>
                <div class="input">
                    <el-input v-model="editsp.name" style="width:600px;" clearable></el-input>
                </div>
            </div>
            <div class="one-row">
                <div class="label">P/N</div>
                <div class="input">
                    <el-input v-model="editsp.pn" style="width:600px;" clearable></el-input>
                </div>
            </div>
            <div class="one-row">
                <div class="label">S/N</div>
                <div class="input">
                    <el-input v-model="editsp.sn" style="width:600px;" clearable></el-input>
                </div>
            </div>
            <div class="one-row">
                <div class="label">现模拟机</div>
                <div class="input">
                    <el-select v-model="editsp.nowsim" style="width:600px;">
                        <el-option label="5978" value="5978"></el-option>
                        <el-option label="5989" value="5989"></el-option>
                        <el-option label="5008" value="5008"></el-option>
                        <el-option label="5015" value="5015"></el-option>
                    </el-select>
                </div>
            </div>
            <div class="one-row">
                <div class="label">位置</div>
                <div class="input">
                    <el-input v-model="editsp.pos" style="width:600px;" clearable></el-input>
                </div>
            </div>
            <div class="one-row">
                <div class="label">状态</div>
                <div class="input">
                    <el-select v-model="editsp.status" style="width:600px;">
                        <el-option label="备件" value="备件"></el-option>
                        <el-option label="使用" value="使用"></el-option>
                        <el-option label="故障" value="故障"></el-option>
                        <el-option label="维修" value="维修"></el-option>
                        <el-option label="废弃" value="废弃" style="color:#F56C6C;"></el-option>
                    </el-select>
                </div>
            </div>
            <div class="one-row">
                <div class="label">数量</div>
                <div class="input">
                    <el-input-number v-model="editsp.amount" :min="0" ></el-input-number>
                </div>
            </div>
            <div class="row-comment">
                <div class="comment-label">备注</div>
                <div class="input">
                    <el-input v-model="editsp.comment" type="textarea" :rows="3" style="width:600px;"></el-input>
                </div>
            </div>
            <div class="row-comment">
                <div class="comment-label">记录</div>
                <div class="input">
                    <el-input v-model="recoder" type="textarea" :rows="3" style="width:600px;" placeholder="该次备件更改的 备注 或者 记录"></el-input>
                </div>
            </div>
            <div class="row-imgs">
                <div class="imgs-label">图片</div>
                <div class="imgs-zone">
                    <el-button type="primary" size="small" style="float:left;" @click="handleUploadBtnClick">点击上传</el-button>
                    <br><br>
                    <div>
                        <div class="imgs-div imgs-img1" v-show="imgshow.img1">
                            <input type="file" accept="image/jpg,image/jpeg" id="input-img1" style="display:none;" @change="handleImgUploadChg('img1')">
                            <div class="place-img">
                                <img :src="editsp.img1" alt="">
                            </div>
                            <div class="place-describe">
                                <div>IMG-1</div>
                            </div>
                            <div class="img-corner">
                                <div class="upload-success">
                                    <i class="el-icon-check"></i>
                                </div>
                                <div class="delete-upload" @click="handlerDeleteUpload('img1')">
                                    <i class="el-icon-close"></i>
                                </div>
                            </div>
                        </div>
                        <div class="imgs-div imgs-img2" v-show="imgshow.img2">
                            <input type="file" accept="image/jpg,image/jpeg" id="input-img2" style="display:none;" @change="handleImgUploadChg('img2')">
                            <div class="place-img">
                                <img :src="editsp.img2" alt="">
                            </div>
                            <div class="place-describe">
                                <div>IMG-2</div>
                            </div>
                            <div class="img-corner">
                                <div class="upload-success">
                                    <i class="el-icon-check"></i>
                                </div>
                                <div class="delete-upload" @click="handlerDeleteUpload('img2')">
                                    <i class="el-icon-close"></i>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div style="display:flex;width:100%;margin-top:30px;">
                <div style="width:150px;"></div>
                <div style="width:100%;">
                    <el-button type="primary" @click="comfirmEdit">确认</el-button> 
                    <el-button @click="cancelEdit">取消</el-button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            editsp: {},
            recoder: '',
        };
    },
    props: {
        onesp: Object,
    },
    methods: {
        cancelEdit: function() {
            this.$emit('chgmod', 'sp-imf');
        },
        handleUploadBtnClick: function() {
            if (!this.imgshow.img1) {
                document.querySelector('#input-img1').click();
                return;
            };
            if (!this.imgshow.img2) {
                document.querySelector('#input-img2').click();
                return;
            };
            if (this.imgshow.img1 && this.imgshow.img2) {
                this.$message({showClose:true, message:'最多可保存两张图片, 已存在两张图片!', type:'warning'});
                return;
            };
        },
        handleImgUploadChg: function(which) {
            let imgfile = document.querySelector('#input-' + which).files[0];
            let imgReader = new FileReader();
            imgReader.onload = (event) => {
                this.editsp[which] = event.target.result;
            };
            imgReader.readAsDataURL(imgfile);
        },
        handlerDeleteUpload: function(which) {
            this.editsp[which] = 'None';
        },
        cancelEdit: function() {
            this.$emit('chgmod', 'sp-imf');
        },
        comfirmEdit: function() {
            let uploadData = new FormData();
            uploadData.append('id', this.editsp.id);
            uploadData.append('name', this.editsp.name);
            uploadData.append('pn', this.editsp.pn);
            uploadData.append('sn', this.editsp.sn);
            uploadData.append('nowsim', this.editsp.nowsim);
            uploadData.append('orgsim', this.editsp.orgsim);
            uploadData.append('pos', this.editsp.pos);
            uploadData.append('status', this.editsp.status);
            uploadData.append('amount', this.editsp.amount);
            uploadData.append('comment', this.editsp.comment);
            uploadData.append('recoder', this.recoder);
            if (this.imgshow.img1) {
                uploadData.append('img1', document.querySelector('#input-img1').files[0]);
            };
            if (this.imgshow.img2) {
                uploadData.append('img2', document.querySelector('#input-img2').files[0]);
            };
            console.log(uploadData);
            this.$http.post('/api/sps/update', uploadData, {
                method: 'post',
                headers: {'Content-Type': 'multipart/form-data'},
                onUploadProgress: function(event) {
                    // console.log(event.loaded / event.total);
                },
            }).then( resp => {
                if (resp.data === 'Done') {
                    this.$router.go(0);
                } else {
                    this.$message({showClose:true, message:'发生未知错误, 即将刷新页面', type:'warning'});
                    setTimeout(() => {
                        this.router.go(0);
                    }, 2000);
                };
            }).catch(err => {
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
    },
    computed: {
        imgshow: function() {
            let show = {img1: false,img2: false};
            if (this.editsp.img1 !== 'None') {
                show.img1 = true;
            };
            if (this.editsp.img2 !== 'None') {
                show.img2 = true;
            };
            return show;
        },
    },
    mounted() {
        this.editsp = JSON.parse(JSON.stringify(this.onesp));
    },
}
</script>

<style scoped>
@font-face {
    font-family: hack;
    src: url("~@/assets/fonts/Hack-BoldItalic.ttf");
}
#container {
    width: 100%;
    display: flex;
    justify-content: center;
}
.editsp-form {
    width: 80%;
    margin-top: 50px;
    margin-bottom: 100px;
}
.one-row {
    display: flex;
    height: 40px;
    width: 100%;
    line-height: 40px;
}
.one-row:nth-child(n+2) {
    margin-top: 10px;
}
.one-row > .label {
    width: 150px;
    font-family: hack;
    font-weight: bold;
    font-size: 0.85rem;
    text-indent: 5px;
    color: #666;
}
.one-row > .input {
    width: 100%;
}
.row-comment {
    display: flex;
    height: 100px;
    width: 100%;
    line-height: 100px;
    align-items: center;
}
.row-comment > .comment-label {
    width: 150px;
    font-family: hack;
    font-weight: bold;
    font-size: 0.85rem;
    text-indent: 5px;
    color: #666;
}
.row-comment > .input {
    width: 100%;
    height: 100%;
}
.row-imgs {
    display: flex;
    width: 100%;
    margin-top: 20px;
}
.row-imgs > .imgs-label {
    width: 150px;
    font-family: hack;
    font-weight: bold;
    font-size: 0.85rem;
    text-indent: 5px;
    color: #666;
    margin-top: 5px;
}
.row-imgs > .imgs-zone {
    width: 100%;
}
.imgs-div {
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
.imgs-div:hover {
    border: 1px solid rgb(192,196,204);
}
.imgs-div:hover .img-corner {
    background-color: #fff;
    box-shadow: none;
}
.imgs-div:hover .upload-success {
    display: none;
}
.imgs-div:hover .delete-upload {
    display: block;
}
.imgs-div > .img-corner {
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
.imgs-div > .img-corner > .upload-success {
    transform: rotate(-45deg);
    color: #fff;
    margin-top: 12px;
    font-size: 12px;
}
.delete-upload {
    transform: rotate(-45deg);
    margin-top: 12px;
    font-size: 12px;
    cursor: pointer;
    display: none;
}
.place-img {
    display: flex;
    align-items: center;
}
.place-img > img {
    width: 150px;
    height: auto;
    cursor: pointer;
}
.place-describe {
    height: 100%;
    display: flex;
    align-items: center;
    margin-left: 5px;
}
</style>

