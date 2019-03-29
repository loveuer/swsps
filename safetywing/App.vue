<template>
    <div id="app">
        <router-view></router-view>
    </div>
</template>

<script>
export default {
    name: 'app',
    created() {
        // const self = this;
        const socket = new WebSocket("ws://localhost:8000/api/ws");
        socket.addEventListener("open", function() {
            //
        });
        socket.addEventListener("close", function() {
            console.log("ws closed");
        });
        socket.addEventListener("message", (event) => {
            let result = JSON.parse(event.data);
            switch (result.type) {
                case "open":
                    this.$store.commit("setwsid", result.id);
                    break;
                case "newlog":
                    this.$store.commit("getSomeNew", "log", JSON.parse(result.message));
                    break;
                case "modifylog":
                    this.$store.commit("getSomeNew");
                    break;
            };
        });
    },
};
</script>

<style>
#app {
    width: 100%;
    min-height: 100%;
}
.ql-size-small {
    font-size: 0.75em;
}
.ql-size-huge {
    font-size: 2.5em;
}
.ql-size-large {
    font-size: 1.5em;
}
</style>
