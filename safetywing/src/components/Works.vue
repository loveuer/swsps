<template>
    <div>
        <div>
            <loveuer-menu />
        </div>
        <div>
            <el-badge :is-dot="true" style="margin-top:10px;margin-right:40px;">
                <font>Safety</font>
                <font>Wing</font>
            </el-badge>
        </div>
    </div>
</template>

<script>
import loveuerMenu from "./uMenu.vue";

export default {
    components: {
        "loveuer-menu": loveuerMenu,
    },
    created() {
        const self = this;
        const socket = new WebSocket("/api/ws");
        socket.addEventListener("open", function(event) {
            let result = JSON.parse(event.data);
            self.$store.commit("setnewid", result.id);
            console.log("opened ws");
        });
        socket.addEventListener("close", function() {
            console.log("ws closed");
        });
        socket.addEventListener("message", function(event) {
            let result = JSON.parse(event.data);
            switch (result.type) {
                case "newlog":
                    console.log("new log add");
                    that.$store.commit("gotSomeNew");
                case "modifylog":
                    console.log("modified log");
                    that.$store.commit("gotSomeNew");
            }
        });
    },
};
</script>

<style scoped>
.font-logo{
    height: 60px;
    line-height: 60px;
    font-size: 20px;
    font-weight: bold;
}
.font-logo:first-child{
    color: #7ec0d4;
}
.font-logo:last-child{
    color: #85c59a;
}
</style>
