<template>
    <div>
        <div>
            <loveuer-menu />
        </div>
        <div style="height:30px;width:100%;zoom:100%;"></div>
        <div style="width:100%;">
            <div class="cards-row">
                <div class="cards-single">
                    <el-card style="width:80%;">
                        <div slot="header" class="clearfix">
                            <span><b>今日log:</b></span>
                            <el-badge :is-dot="this.$store.state.todayMessage.log.new" class="cards-dot">
                                <el-button @click="refreshLog" :type="this.$store.state.todayMessage.log.type" circle icon="el-icon-refresh"></el-button>
                            </el-badge>
                        </div>
                        <div class="cards-content">
                            <el-table :data="this.$store.state.todayMessage.log.logs" style="width:100%;" >
                                <el-table-column prop="time" label="时间" width="100"></el-table-column>
                                <el-table-column prop="auth" label="人员" width="100"></el-table-column>
                                <el-table-column prop="sim" label="模拟机" width="80"></el-table-column>
                                <el-table-column prop="log" label="内容" :show-overflow-tooltip="true"></el-table-column>
                            </el-table>
                        </div>
                    </el-card>
                </div>
                <div class="cards-single">
                    <el-card style="width:80%;">
                        <div slot="header" class="clearfix">
                            <span><b>今日log:</b></span>
                            <el-badge :is-dot="this.$store.state.todayMessage.log.new" class="cards-dot">
                                <el-button circle :type="this.$store.state.todayMessage.log.type" icon="el-icon-refresh"></el-button>
                            </el-badge>
                        </div>
                    </el-card>
                </div>
            </div>
            <div>
                <div></div>
                <div></div>
            </div>
        </div>
    </div>
</template>

<script>
import loveuerMenu from "./uMenu.vue";

export default {
    components: {
        "loveuer-menu": loveuerMenu,
    },
    methods: {
        refreshLog: function() {
            let logs = [
                {time:"06:32:21", auth:"李永翔", sim: "5978", log: "Preflight is OK"},
                {time:"06:49:07", auth:"袁超", sim: "5008", log: "Preflight is OK"},
                {time:"07:32:43", auth:"宋继朋", sim: "5015", log: "Preflight is OK"},
                {time:"07:52:35", auth:"赵育鹏", sim: "5989", log: "我来测试一下,这是一条很长的log,甚至还有什么什么插入备件在里面什么的,反正就是很多很多"},
            ];
            this.$store.commit("refreshedLog", logs);
        },
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
.cards-row {
    width: 100%;
    display: flex;
    align-items: center;
}
.cards-single {
    width: 50%;
    display:flex;
    justify-content: center;
    /* align-items: center; */
}
.cards-dot {
    float: right;
}
.clearfix {
    height:30px;
    line-height: 30px;
}
.clearfix:before,
.clearfix:after {
    display: table;
    content: "";
}
.clearfix:after {
    clear: both
}
</style>
