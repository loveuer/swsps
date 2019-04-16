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

        </div>
    </div>
</template>

<script>
import { clearInterval } from 'timers';
export default {
    data() {
        return {
            socket: null,
            skalive: false,
            skinterval: null,
        };
    },
    methods: {
        doskInterval: function() {
            if (!this.skinterval) {
                this.skinterval = window.setInterval(() => {
                    if(this.skalive) {
                        this.skalive = false;
                        this.socket.send(JSON.stringify({event:"salary",type:"ping",msg:""}));
                    } else {
                        this.socket = new WebSocket("ws://localhost:8000/socket");
                    };
                }, 10000);
            };
        },
        initSK: function() {
            this.socket = new WebSocket("ws://localhost:8000/socket");
            this.socket.onopen = () => { this.skalive = true; };
            this.socket.onclose = () => { this.socket=null; };
            this.socket.onmessage = (event) => {
                let message = JSON.parse(event.data);
                if (message.event === "salary") {
                    switch (message.type) {
                        case "pong":
                            this.skalive = true;
                            break;
                    };
                };
            };
        },
    },
    mounted() {
        this.initSK();
        // ping web socket per 10s
        this.doskInterval();
    },
    beforeDestroy() {
        console.log("before destroy");
        console.log("interval id: ", this.skinterval);
        this.socket.close();
        this.socket = null;
        window.clearInterval(this.skinterval);
        this.skinterval = null;
    },
};
</script>

<style scoped>
.breadcrumb {
    height: 50px;
    width: 100%;
    margin-top: 10px;
    margin-left: 20px;
}
</style>
