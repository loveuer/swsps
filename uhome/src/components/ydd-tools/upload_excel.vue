<template>
    <div class="container">
        <div>
            <div class="upload-btn">
                <el-button type="primary" @click="doSelectFile">选择文件<i class="el-icon-upload el-icon--right"></i></el-button>
            </div>
            <input type="file" id="upload-file" style="display:none" @change="selectedFileChanged" accept="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet">
            <div class="upload-content" v-show="fileSelected">
                <div>
                    <div class="upload-main" tabindex="0">
                        <div class="excel-icon">
                            <div>
                                <svg t="1555482128821" style="height:90%;width:90%;" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1161" xmlns:xlink="http://www.w3.org/1999/xlink">
                                    <path d="M728.345 64.223H224.056V321.34h68.687V133.019h378.844v187.78h186.917v569.777H292.743v-84.271h-0.715v-68.796h0.715v-53.046h-68.687v274.909h703.135V256.486z" fill="#819292" p-id="1162">
                                    </path>
                                    <path d="M660.276 418.403h129.217v67.077H660.276zM660.276 578.386h129.217v67.077H660.276zM660.276 738.369h129.217v67.077H660.276z" fill="#819292" p-id="1163">
                                    </path>
                                    <path d="M586.15 408.784h-68.687v328.725H292.028v68.796H586.15z" fill="#50A135" p-id="1164">
                                    </path>
                                    <path d="M460.55 684.463V321.34H98.004v363.123H460.55z m-234.043-70.305l-85.787 0.722 93.843-110.912-93.843-113.045 81.805 0.225 53.747 63.523 53.747-63.523 81.805-0.225-93.843 113.044 93.843 110.912-85.787-0.722-49.765-59.948-49.765 59.949z" fill="#50A135" p-id="1165">
                                    </path>
                                    <path d="M326.038 614.158l85.787 0.722-93.843-110.913 93.843-113.044-81.805 0.225-53.747 63.524-53.748-63.524-81.805-0.225 93.844 113.044L140.72 614.88l85.787-0.722 49.766-59.948z" fill="#FFFFFF" p-id="1166">
                                    </path>
                                </svg>
                            </div>
                            <div v-show="uploading">
                                <el-progress type="circle" :percentage="upload_percentage" :width="80" style="margin-top:5px;"></el-progress>
                            </div>
                        </div>
                        <div style="margin-left:10px;">
                            <div class="upload-file-name" v-text="upload_file_name"></div>
                        </div>
                        <label class="top-right">
                            <div class="top-right-success">
                                <i class="el-icon-check"></i>
                            </div>
                        </label>
                        <div class="top-right-close" @click="deleteSelectedFile">
                            <i class="el-icon-close"></i>
                        </div>
                    </div>
                </div>
                
                <div>
                    <el-button @click="do_upload" :disabled="uploading">上传文件</el-button>
                </div>
            </div>
        </div>
    </div>    
</template>

<script>
import { clearInterval } from 'timers';
export default {
    data() {
        return {
            upload_percentage: 0,
            upload_status: "",
            upload_file: null,
            upload_file_name: '',
            fileSelected: false,
            uploading: false,
        };
    },
    methods: {
        deleteSelectedFile: function() {
            this.upload_file = null;
            this.upload_file_name = "";
            this.fileSelected = false;
        },
        doSelectFile: function() {
            console.log("a?");
            document.querySelector("#upload-file").click();
        },
        do_upload: function() {
            this.uploading = true;
            // 模拟上传
            var upload_interval = window.setInterval(() => {
                this.upload_percentage += 5;
                if (this.upload_percentage === 100) {
                    console.log("arrive 100");
                    window.clearInterval(upload_interval);
                    this.doNext();
                }
            }, 100);
            // 真的上传
            // var xlsx = new FormData();
            // xlsx.append('xlsx', this.upload_file);
            // this.$http.post("/api/salary/upload", xlsx, {
            //     headers: {'Content-Type':'multipart/form-data'}, 
            //     onUploadProgress: e => {
            //         if (e.loaded >= e.total) {
            //             this.uploading = false;
            //             this.uploaded = true;
            //         };
            //     },
            // }).then(resp => {
            //     console.log(resp.data);
            // }).catch(err => { console.log(err.response)} );
        },
        selectedFileChanged: function() {
            let filename = document.querySelector("#upload-file").files[0].name;
            console.log(!filename.split(".")[0].match(/20\d{2}年(0[1-9]|1[0-2])月工资条/));
            if (filename.split(".").length !== 2 || !filename.split(".")[0].match(/20\d{2}年(0[1-9]|1[0-2])月工资条/)) {
                this.$alert("文件名称不匹配", "警告", {
                    confirmButtonText: '确定',
                    type: 'warning ',
                });
                document.querySelector("#upload-file").files = null;
                return;
            };

            let filetype = filename.split(".")[1];
            if (filetype !== 'xlsx') {
                this.$alert("只能上传Excel (xlsx) 文件", "警告", {
                    confirmButtonText: '确定',
                    type: 'warning ',
                });
                return;
            };

            this.upload_file_name = filename;
            this.upload_file = document.querySelector("#upload-file").files[0];
            this.fileSelected = true;
        },
        doNext: function() {
            this.$confirm("上传文件成功, 是否继续?", "提示", {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                this.$emit("update", {event:"upload", status:"success", argv:this.upload_file_name});
            }).catch(() => {
                console.log("clicked 取消");
            });
        },
    },
};
</script>

<style scoped>
.container {
    width: 100%;
}
.upload-btn {
    width: 100%;
    display: flex;
    justify-content: center;
}
.upload-content {
    width: 100%;
}
.upload-content > div {
    width: 100%;
    display: flex;
    margin-top: 20px;
    justify-content: center;
}
.upload-process {
    width: 20rem;
}
.upload-main {
    width: 342px;
    height: 87px;
    border: 1px solid #8b9baa;
    border-radius: 5px;
    position: relative;
    background: #fff;
    overflow: hidden;
    outline: none;
    display: flex;
    align-items: center;
}
.upload-main:hover {
    border: 1px solid #409eff;
}
.upload-main > .top-right {
    position: absolute;
    top: -7px;
    right: -17px;
    width: 46px;
    height: 26px;
    background: #13ce66;
    transform: rotate(45deg);
    text-align: center;
    box-shadow: 0 1px 1px #ccc;
    display: inline-block;
}
.upload-main:hover > .top-right {
    background: #fff;
    box-shadow: none;
    cursor: pointer;
}
.upload-main > .top-right > div {
    transform: rotate(-45deg);
    color: #fff;
    margin-top: 12px;
    font-size: 12px;
}
.upload-main:hover > .top-right-sucess {
    display: none;
}
.top-right-close {
    color: #606266;
    display: none;
    font-size: 12px;
    position: absolute;
    top: 0;
    right: 0;
    padding: 5px;
    cursor: pointer;
}
.upload-main:hover > .top-right-close {
    display: inline-block;
}
.excel-icon {
    height: 80px;
    width: 80px;
    margin-left: .5rem;
    display: flex;
    align-items: center;
}
.upload-file-name {
    font-size:14px;
    color:#606266;
}
.upload-main:hover .upload-file-name {
    color: #409eff;
}
</style>

