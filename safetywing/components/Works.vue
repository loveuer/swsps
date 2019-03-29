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
                            <el-table :data="this.$store.state.todayMessage.log.logs" style="width:100%;cursor:pointer;" @row-click="openTodayLog">
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
                            <span><b>Opened故障:</b></span>
                            <el-badge :is-dot="this.$store.state.todayMessage.malf.new" class="cards-dot">
                                <el-button circle :type="this.$store.state.todayMessage.malf.type" icon="el-icon-refresh"></el-button>
                            </el-badge>
                        </div>
                    </el-card>
                </div>
            </div>
            <div class="cards-row" style="margin-top:30px;">
                <div class="cards-single">
                    <el-card style="width:80%;">
                        <div slot="header" class="clearfix">
                            <span><b>训练单:</b></span>
                            <el-badge :is-dot="this.$store.state.todayMessage.trainingList.new" class="cards-dot">
                                <el-button circle :type="this.$store.state.todayMessage.trainingList.type" icon="el-icon-refresh"></el-button>
                            </el-badge>
                        </div>
                        <div class="tlist-content">
                            <div class="tlist-one" v-for="(t, index) of this.$store.state.todayMessage.trainingList.list" :key="index">
                                <div class="tone-img" @click="showTimg(t.img)">
                                    <img :src="'http://127.0.0.1:8000' + t.img">
                                </div>
                                <div class="tone-detail" @click="meetTbydate(t.date)">
                                    <div>
                                        <el-row class="tdetail-row">
                                            <div>{{ t.date }}</div>
                                            <div>{{ t.sim }}</div>
                                            <div>{{ t.airline }}</div>
                                        </el-row>
                                        <el-row class="tdetail-row">
                                            <div>{{ t.ts + ' - ' + t.te }}</div>
                                            <div>共计 {{ t.ta }} 小时</div>
                                        </el-row>
                                        <el-row class="tdetail-row">
                                            <div>教员: {{ t.instructor }}</div>
                                            <div>学员: {{ t.students }}</div>
                                        </el-row>
                                        <el-row class="tdetail-row">
                                            <div>训练类型: {{ t.mode }}</div>
                                            <div>机组号: {{ t.crew }}</div>
                                        </el-row>
                                        <el-row class="tdetail-row">
                                            <div>{{ t.auth }}</div>
                                        </el-row>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </el-card>
                </div>
                <div class="cards-single"></div>
            </div>
        </div>
        <div class="light-box">
            <loveuer-lightbox ref="uLightBox"></loveuer-lightbox>
        </div>
        <div style="height:200px;width:100%;zoom:100%;"></div>
    </div>
</template>

<script>
import loveuerMenu from "./uMenu.vue";
import loveuerLightBox from './uLightBox.vue';
import dateformat from "dateformat";

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
        openTodayLog: function() {
            let today = dateformat(Date.now(), "yyyy-mm-dd");
            this.$router.push("/works/logRecorder/check/" + today);
        },
        showTimg: function(imgsrc) {
            this.$refs.uLightBox.ifshow = true;
            this.$refs.uLightBox.imgSrc = "http://127.0.0.1:8000" + imgsrc;
        },
        meetTbydate: function(date) {
            this.$router.push("/works/trainingRecorder/check/" + date);
        },
    },
    mounted() {
        // mock get log
        let newlogs = [
            {time:"06:32:21", auth:"李永翔", sim: "5978", log: "Preflight is OK"},
            {time:"06:49:07", auth:"袁超", sim: "5008", log: "Preflight is OK"},
            {time:"07:32:43", auth:"宋继朋", sim: "5015", log: "Preflight is OK"},
        ];
        this.$store.commit("refreshedLog", newlogs);
        // mock get training list
        let newtraininglist = [
            {date:'2019-03-27',sim:'5989',airline:'西藏航空',session:3,ts:'14:00',te:'16:00',ta:2,instructor:'xxxxxx',students:'yyyyyy, zzzzzz',mode:'aabbccdd',crew:'',auth:'刘博强',img:'/static/img/tlist/mock1.jpg'},
            {date:'2019-03-27',sim:'5989',airline:'西藏航空',session:4,ts:'16:00',te:'20:00',ta:4,instructor:'xxxxxx',students:'yyyyyy, zzzzzz, pppppp, ghjkla, bilibili, alasijia',mode:'aabbccdd',crew:'',auth:'刘博强',img:'/static/img/tlist/mock2.jpg'},
            {date:'2019-03-27',sim:'5989',airline:'西藏航空',session:5,ts:'21:00',te:'01:00',ta:4,instructor:'xxxxxx',students:'yyyyyy, zzzzzz',mode:'aabbccdd',crew:'',auth:'刘博强',img:'/static/img/tlist/mock3.jpg'},
        ];
        this.$store.commit("refreshedTrainingList", newtraininglist);
    },
    components: {
        "loveuer-lightbox": loveuerLightBox,
        "loveuer-menu": loveuerMenu,
    }
    
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
    min-width: 600px;
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
.tlist-one {
    width: 100%;
    height: 220px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-bottom: 1px dashed #ccc;
}
.tone-img {
    height: 200px;
    width:320px;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #eee;
    float: left;
}
.tone-img > img {
    max-height: 210px;
    max-width: 310px;
    cursor: pointer;
}
.tone-detail {
    height: 210px;
    width:100%;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
}
.tone-detail > div {
    width: 100%;
}
.tdetail-row {
    overflow: hidden;
    border-bottom: 1px solid #ccc;
    display: flex;
    height: 40px;
    line-height: 40px;
    align-items: center;
    width: 100%;
}
.tdetail-row > div {
    margin-left: 20px;
    height: 40px;
    overflow: overlay;
}
</style>
