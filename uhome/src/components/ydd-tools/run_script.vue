<template>
    <div id="container">
        <div class="input-script">
            <el-input placeholder="请输入年月(如: 201904)" suffix-icon="el-icon-edit" v-model="script_name" style="margin:10px;"></el-input>
            <el-button style="margin:10px;" type="primary" @click="doScript">执行</el-button>
        </div>
        <div class="run-btn"></div>
    </div>    
</template>

<script>
export default {
    data() {
        return {
            script_name: "",
        };
    },
    props: {
        "filename": String,  
    },
    methods: {
        validateFileName() {
            if (this.$props.filename === "") {
                return;
            };
            this.script_name = this.$props.filename.slice(0,4) + this.$props.filename.slice(5,7);
            this.$confirm(`是否执行上传的<< ${this.filename} >>?`, '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'info',
            }).then(() => {
                this.doScript()
            });
        },
        doScript() {
            if (!this.script_name.match(/20\d{2}(0[1-9]|1[0-2])/)) {
                this.$alert(`输入的年月 ( ${this.script_name} 不符合格式)`, "提示", {
                    confirmButtonText: '确定',
                    type: "error",
                });
                this.script_name = "";
                return;
            };
            // do post script to run
        },
    },
    computed: {
        
    },
    mounted() {
        this.validateFileName();
    },
};
</script>

<style scoped>
.input-script {
    width: 100%;
    display: flex;
}
</style>
