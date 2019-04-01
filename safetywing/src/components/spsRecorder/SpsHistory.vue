<template>
    <div>
        <div>
            <loveuer-menu ref="uMenu"></loveuer-menu>
        </div>
        <div style="height:30px;width:100%;zoom:100%;"></div>
        <div>
            <div class="datePicker">
                <el-date-picker v-model="seld.start"></el-date-picker>
                <el-date-picker v-model="seld.end" style="margin-left:10px;"></el-date-picker>
                <el-button style="margin-left:30px;" type="primary" @click="getHistory">查看</el-button>
            </div>
        </div>
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";
import dateformat from "dateformat";

export default {
    data() {
        return {
            seld: {start:null, end: null},
        };
    },
    methods: {
        getHistory: function() {
            if(this.seld.start > this.seld.end) {
                this.$notify.error({
                    title:'错误',
                    message: '开始日期应在前!',
                    position:'top-left',
                });
                return;
            };
            let start = dateformat(this.seld.start, "yyyy-mm-dd");
            let end = dateformat(this.seld.end, "yyyy-mm-dd");
            console.log(start, end);
        },
    },
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/spsRecorder/history";
        this.seld.start = Date.now();
        this.seld.end = Date.now();
        this.getHistory();
    },
    components: {
        "loveuer-menu": loveuerMenu,
    },
};
</script>

<style scoped>
.datePicker {
    margin-left: 30px;
}
</style>
