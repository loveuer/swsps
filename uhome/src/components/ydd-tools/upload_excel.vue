<template>
    <div class="container">
        <div>
            <div class="upload-btn">
                <el-button type="primary">选择文件<i class="el-icon-upload el-icon--right"></i></el-button>
            </div>
            <input type="file" id="upload-file" style="display:none">
            <div class="upload-content">
                <div>
                    <div class="upload-process">
                        <el-progress :percentage="upload_percentage" :status="upload_status" :text-inside="true" :stroke-width="18"></el-progress>
                    </div>
                </div>
                <div>
                    <div class="upload-main">
                        <div></div>
                        <div class="top-right">
                            <i class="el-icon-check"></i>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>    
</template>

<script>
export default {
    data() {
        return {
            upload_percentage: 0,
            upload_status: "",
        };
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
    width: 25rem;
    height: 8rem;
    border: 1px solid #409eff;
    border-radius: 5px;
    position: relative;
}
.upload-main > .top-right {
    position: absolute;
    top: 0;
    right: 0;
    width: 2rem;
    height: 2rem;
    background-color: #eee;
}
</style>

