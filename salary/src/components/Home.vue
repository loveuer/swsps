<template>
    <div class="container">
        <div class="title-zone">
            <font>Auto</font>
            <font>Salary</font>
        </div>
        <div class="upload-zone" >
            <el-button type="primary" size="small" @click="handlerChooseFile">上传文件</el-button>
            <input type="file" accept="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" id="excel" @change="UploadExcel">
        </div>
        <div class="showfile-zone" >
            <div>
                <div class="fileflag-zone">
                    <svg viewBox="0 0 1498 1024" version="1.1"  p-id="5204"  width="14" height="14" style="transform: rotate(-45deg);margin-top:10px;"><path d="M618.396098 1024L0 403.605854l140.862439-141.861464 477.533659 479.531708L1357.674146 0 1498.536585 140.862439l-880.140487 883.137561z" p-id="5205" fill="#ffffff"></path></svg>
                </div>
                <div class="fileflag-delete">
                    <svg viewBox="0 0 1024 1024" version="1.1" p-id="3697" width="12" height="12"><path d="M886.528 908.032c-28.096 28.096-73.856 28.096-102.016 0L138.304 261.824c-28.096-28.16-28.16-73.856 0-102.016 28.032-28.16 73.792-28.16 102.08 0l646.144 646.144C914.624 834.24 914.752 879.872 886.528 908.032L886.528 908.032zM885.76 261.504 239.616 907.648c-28.224 28.224-73.92 28.224-102.08 0-28.16-28.096-28.16-73.728 0.064-102.016L783.744 159.552c28.224-28.16 73.984-28.16 102.016-0.064C913.984 187.648 913.856 233.344 885.76 261.504L885.76 261.504z" p-id="3698" fill="#888"></path></svg>
                </div>
                <div>
                    <svg viewBox="0 0 1024 1024" version="1.1" p-id="2039" width="60" height="60"><path d="M728.345 64.223H224.056V321.34h68.687V133.019h378.844v187.78h186.917v569.777H292.743v-84.271h-0.715v-68.796h0.715v-53.046h-68.687v274.909h703.135V256.486z" fill="#819292" p-id="2040"></path><path d="M660.276 418.403h129.217v67.077H660.276zM660.276 578.386h129.217v67.077H660.276zM660.276 738.369h129.217v67.077H660.276z" fill="#819292" p-id="2041"></path><path d="M586.15 408.784h-68.687v328.725H292.028v68.796H586.15z" fill="#50A135" p-id="2042"></path><path d="M460.55 684.463V321.34H98.004v363.123H460.55z m-234.043-70.305l-85.787 0.722 93.843-110.912-93.843-113.045 81.805 0.225 53.747 63.523 53.747-63.523 81.805-0.225-93.843 113.044 93.843 110.912-85.787-0.722-49.765-59.948-49.765 59.949z" fill="#50A135" p-id="2043"></path><path d="M326.038 614.158l85.787 0.722-93.843-110.913 93.843-113.044-81.805 0.225-53.747 63.524-53.748-63.524-81.805-0.225 93.844 113.044L140.72 614.88l85.787-0.722 49.766-59.948z" fill="#FFFFFF" p-id="2044"></path></svg>
                </div>
                <div style="margin-left:10px;">
                    <font v-text="filename" style="font-size:0.85rem;"></font>
                </div>
            </div>
        </div>
        <div class="message-zone">
            <el-alert :title="messagebox.title" :type="messagebox.type" style="max-width:700px;" center :closable='false'></el-alert>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            connection: true,
            filename: "2019年06月工资表.xlsx",
            runningStatus: false,
            doneLine: 2,
            messagebox: {
                type: 'success',
                title: 'Welcome to Auto Salary',
            },
        };
    },
    methods: {
        handlerChooseFile() {
            if (!this.connection) {
                return;
            };
            document.querySelector("#excel").click();
        },
        UploadExcel() {
            if (!this.connection) {
                return;
            };
            let file = new FormData();
            let uploaded = document.querySelector("#excel").files[0];
            if(!uploaded) {
                console.log("empty choose");
                return;
            };
            console.log('filename: ', uploaded.name);
            console.log("filename type: ", typeof(uploaded.name));
            if (!uploaded.name.match(/20(\d{2})年(0[1-9])|(1[0-2])月工资条.xlsx/)) {
                console.log("filename not match regex");
                // this.$message({type: 'warning', showClose: true, message: '文件命不符合! (20xx年xx月工资表.xlsx)'});
                this.messagebox.title = '文件命不符合! (20xx年xx月工资表.xlsx)'
                this.messagebox.type = 'error'
                return;
            };
            file.append("excel", uploaded);
            this.$http.post("/api/salary/upload", file, {
                method: 'post',
                headers: {'Content-Type': 'multipart/form-data'},
            }).then(resp => {
                if(resp.data) {
                    // upload file success
                    this.filename = uploaded.name;
                    this.$confirm('文件上传成功! 是否执行auto salary?', '提示', {
                        confirmButtonText: '确认',
                        cancelButtonText: '取消',
                        type: 'warning',
                    }).then(() => {
                        this.startScript();
                    }).catch(() => {
                        // 取消执行
                    });
                }
            }).catch(err => {
                console.log(err.response);
                this.$message({type:'error', showClose:true, message: ':( 上传文件失败了~'});
            }).finally(() => {
                document.querySelector("#excel").value = null;
            });
        },
        getRunningStatus() {
            this.$http.get("/api/salary/status").then(resp => {
                if(resp.data) {
                    this.runningStatus = true;
                    this.filename = resp.data.filename;
                } else {
                    this.runningStatus = false;
                };
                this.messagebox.title = this.messagebox.title.split("(")[0] + '(链接服务器成功!)';
                console.log('running status: ', resp.data);
            }).catch(err => {
                console.log(err.response);
                this.messagebox.title = this.messagebox.title.split("(")[0] + '(链接服务器失败!)';
                this.messagebox.type = 'error';
                this.connection = false; 
            });
        },
        startScript() {
            this.$http.post(`/api/salary/start/${this.filename}/${this.doneLine}`).then(resp => {
                if (resp.data) {
                    this.messagebox.title = '执行任务成功,开始发送...'
                    this.messagebox.type = 'success';
                    this.runningStatus = true;
                }
            }).catch(err => {
                console.log(err.response);
            });
        },
        getLines(pos) {
            this.$http.get(`/api/salary/lines/${this.filename}/${pos}`).then(resp => {
                console.log(resp.data);
                switch (pos) {
                    case 0:
                        
                        break;
                    case -1:

                        break;
                
                    default:
                        break;
                }
            }).catch(err => {
                console.log(err.response);
            });
        },
    },
    watch: {
        runningStatus: function() {
            if (this.runningStatus && this.filename !== "") {
                setTimeout(() => {
                    //get the log file head
                    this.getLines(0)
                }, 2000);
                setInterval(() => {
                    this.getLines(-1);
                }, 3000);
            }
        }  
    },
    mounted() {
        this.messagebox.title = this.messagebox.title +'(链接服务器...)';
        this.getRunningStatus();
    },
}
</script>

<style scoped>
.container {
    width: 100%;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}
#excel {
    display: none;
}
div.title-zone {
    width: 100%;
    height: 100%;
    margin-top: 20px;
    margin-bottom: 30px;
    display: flex;
    justify-content: center;
}
div.title-zone font {
    font-size: 40px;
    font-weight: bold;
    font-family: monospace;
}
.container.upload-zone {
    height: 100px;
    width: 100%;
    background: #ddd;
}
#excel {
    display: none;
}
.message-zone {
    width: 100%;
    display: flex;
    justify-content: center;
    margin-top: 20px;
}
.showfile-zone {
    height: 100px;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}
.showfile-zone > div {
    height: 80px;
    width: 380px;
    border: 1px solid #13ce66;
    border-radius: 3px;
    display: flex;
    align-items: center;
    overflow: hidden;
    position: relative;
}
.showfile-zone > div:hover .fileflag-zone {
    display: none;
}
.showfile-zone > div:hover .fileflag-delete {
    display: flex;
}
.fileflag-zone {
    position: absolute;
    right: -17px;
    top: -7px;
    width: 46px;
    height: 26px;
    background: #13ce66;
    text-align: center;
    transform: rotate(45deg);
    box-shadow: 0 1px 1px #ccc;
}
.fileflag-delete {
    position: absolute;
    right: 0px;
    top: 0px;
    width: 30px;
    height: 25px; 
    cursor: pointer;
    display: none;
    align-items: center;
    justify-content: center;
}
</style>