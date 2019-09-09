<template>
    <div class="container">
        <div class="title-zone">
            <font>Auto</font>
            <font>Salary</font>
        </div>
        <div class="upload-zone">
            <el-button type="primary" size="small" @click="handlerChooseFile">上传文件</el-button>
            <input type="file" accept="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" id="excel" @change="UploadExcel">
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            filename: "",
        };
    },
    methods: {
        handlerChooseFile() {
            document.querySelector("#excel").click();
        },
        UploadExcel() {
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
                    this.$confirm('文件上传成功! 是否执行auto saalry?', '提示', {
                        confirmButtonText: '确认',
                        cancelButtonText: '取消',
                        type: 'warning',
                    }).then(() => {
                        // 确认执行
                    }).catch(() => {
                        // 取消执行
                    });
                }
            }).catch(err => {
                console.log(err.response);
            }).finally(() => {
                document.querySelector("#excel").value = null;
            });
        },
    },
    mounted() {
        this.$http.get("/api/salary/status").then(resp => {
            console.log("status: ", resp.data);
        }).catch(err => {
            console.log(err.response);
        });
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
</style>