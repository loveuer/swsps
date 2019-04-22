<template>
    <div id="container">
        <div class="breadcrumb">
            <el-breadcrumb separator="/">
                <el-breadcrumb-item :to="{path: '/'}">首页</el-breadcrumb-item>
                <el-breadcrumb-item>余丹丹</el-breadcrumb-item>
                <el-breadcrumb-item>工资条</el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="content">
            <div class="process-step">
                <div>
                    <el-steps :active="step" align-center>
                        <el-step title="上传文件" icon="el-icon-upload"></el-step>
                        <el-step title="发送邮件" icon="el-icon-message"></el-step>
                        <el-step title="实时结果" icon="el-icon-view"></el-step>
                    </el-steps>
                </div>
            </div>
            <div class="platform">
                <component :is="platform" v-on:update="handleUpdate" v-on:error="handleError" :filename="filename"></component>
            </div>
        </div>
    </div>
</template>

<script>
import uploadExcel from "./upload_excel.vue";
import runScript from "./run_script.vue";

export default {
    data() {
        return {
            socket: null,
            step: 1,
            filename: '',
            run_script: '',
            // skinterval: null,
            // skalive: false,
        };
    },
    computed: {
        platform: function() {
            let pltfmMap = {1:'upload-excel', 2:'run-script', 3:''};
            return pltfmMap[this.step];
        },
    },
    methods: {
        // 每10秒 ping 通过socket ping一次服务器, 暂时取消,想通过服务器主动来ping的方式
        // doskInterval: function() {
        //     if (!this.skinterval) {
        //         this.skinterval = window.setInterval(() => {
        //             if(this.skalive) {
        //                 this.skalive = false;
        //                 this.socket.send(JSON.stringify({Event:"ping",Type:"",Data:""}));
        //             } else {
        //                 this.socket = new WebSocket("ws://localhost:8000/api/socket");
        //             };
        //         }, 10000);
        //     };
        // },
        initSK: function() {
            this.socket = new WebSocket("ws://localhost:8000/api/socket");
            this.socket.onopen = () => { console.log("web socket opened") };
            // this.socket.onclose = () => { this.socket=null; };
            this.socket.onclose = () => {
                console.log("web socket closed, reconnecting...");
                setTimeout(() => {
                    this.initSK();
                }, 2000);
            };
            this.socket.onerror = () => {
                console.log("web socket err");
                this.socket.close();
            };
            this.socket.onmessage = (event) => {
                let message = JSON.parse(event.data);
                switch (message.event) {
                    case "salary":
                        console.log("web_socket msg: ", message);
                        // let data = JSON.parse(message.data);
                        break;
                };
            };
        },
        handleUpdate: function(kws) {
            switch (kws.event) {
                case "upload":
                    if (kws.status === "success") {
                        this.step = 2;
                        this.filename = kws.argv;
                    };
                    break;
                case 'run_script':
                    this.run_script = kws.filename;
                    this.step = 3;
            }
        },
        handleError: function(errCode) {
            switch (errCode) {
                case 201:
                    this.step = 1;
                    this.filename = "";
                    break;
            }
        },
    },
    mounted() {
        this.initSK();
        // this.doskInterval();
    },
    beforeDestroy() {
        this.socket.close();
        this.socket = null;
    },
    components: {
        "upload-excel": uploadExcel,
        "run-script": runScript,
    },
};
</script>

<style scoped>
.container {
    width: 100%;
}
.breadcrumb {
    height: 50px;
    width: 95%;
    margin-top: 10px;
    margin-left: 20px;
}
.content {
    width: 100%;
}
.content > .process-step {
    width: 100%;
    display: flex;
    justify-content: center;
}
.content > .process-step > div {
    width: 80%;
}
.content > .platform {
    width: 100%;
    margin-top: 20px;
    display: flex;
    justify-content: center;
}
</style>
