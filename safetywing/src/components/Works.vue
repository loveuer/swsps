<template>
    <div>
        <div>
            <loveuer-menu />
        </div>
        <div style="height:30px;width:100%;zoom:100%;"></div>
        <div class="works-content">
            <!-- left side -->
            <div style="width:50%;min-height:100%;height:100%;">
                <!-- log single -->
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
                <!-- training single -->
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
            </div>
            <!-- right side -->
            <div style="width:50%;min-height:100%;height:100%;">
                <div class="cards-single">
                    <el-card style="width:80%;">
                        <div slot="header" class="clearfix">
                            <span><b>Opened故障:</b></span>
                            <el-badge :is-dot="this.$store.state.todayMessage.malf.new" class="cards-dot">
                                <el-button circle :type="this.$store.state.todayMessage.malf.type" icon="el-icon-refresh"></el-button>
                            </el-badge>
                        </div>
                        <div class="malfs-content">
                            <div class="malfs-one" v-for="(malf, index) of this.$store.state.todayMessage.malf.malfs" :key="index">
                                <el-card>
                                    <div slot="header" class="mone-title">
                                        <div style="width:150px;">{{ malf.drid }}</div>
                                        <div style="width:180px;">{{ malf.time }}</div>
                                        <div style="width:60px;">{{ malf.sim }}</div>
                                        <div style="width:70px;">{{ malf.auth }}</div>
                                        <div>Cat: {{ malf.class }}</div>
                                    </div>
                                    <div class="mone-content">
                                        <div>{{ malf.content }}</div>
                                    </div>
                                    <div class="mone-commits">
                                        <div class="mone-onecommit">
                                            ......
                                        </div>
                                        <div v-for="(c, ci) of malf.last2commits" :key="ci">
                                            <div class="mone-onecommit"><font style="text-decoration: underline;font-weight:bold;font-style: italic;">{{ c.time }}　{{ c.auth }}</font> :　{{ c.content }}</div>
                                        </div>
                                    </div>
                                </el-card>
                            </div>
                        </div>
                    </el-card>
                </div>
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
        // mock get malfs
        let newmalfs = [
            {drid:'20190318093030',time:"2019-03-02 13:32:47",sim:'5978',cap:'陈本志',auth:'李永翔',coworkers:'陈本志,李永翔',support:'xxxyyy',bt:0.5,class:3,tl:3,content:'18:00-22:00场次教员反映飞行中，飞机自动低头下坠，无法控制；为弄清楚具体状况，和机组一起飞了半小时，后来重现了故障，具体情况为：ZUUU 02L 、Approach 20NM, HYD G+B LO PRESSURE, 机组飞行过程中超速，飞机直接低头下坠，俯仰配平DN 4.0且无法拉动配平轮，配平轮一直往前打。机组拉不起来，冻结拍照发现没有ELAC故障，和其无关。 晚上训练结束后测试该故障： 1、 ZUUU 02L 、Approach 20NM， 配平正常然后解冻，飞机突然低头下坠，查看F/CTL页面发现和之前一样俯仰配平DN 4.0且无法拉动配平轮，配平轮一直往前打，冻结拍照。 2、ZUUU 02L 、Approach 20NM, HYD G+B LO PRESSURE, 一超速飞机直接低头下坠，和机组遇到的一样，且也没有ELAC故障。怀疑是缝翼卡组引起的。 3、ZUUU 02L 、Approach 20NM，设置SLATS LOCKED， 解冻飞行一切正常，即使超速也未出异常。 4、再次尝试重现第二次的故障，没能重现，需继续排查原因。',last2commits:[{auth:'郑卿',time:'2019-03-03 01:34:27',content:'在5989和5015上测试ZUUU 02L 、Approach 20NM, HYD G+B LO PRESSURE，模拟机正常飞行，没有出现下坠的情况。'},{auth:'付东明',time:'2019-03-04 12:13:25',content:'为保证追踪系统数据链完整性，不提供删除，对于不恰当的地方可以追加信息进行描述，请大家仔细斟酌，认真填写!'}]},
            {drid:'20190314091905',time:"2019-03-03 15:32:47",sim:'5989',cap:'付东明',auth:'宋继朋',coworkers:'黄一林',support:'xxxyyy',bt:0.5,class:3,tl:3,content:'18:00-22:00场次教员反映飞行中，飞机自动低头下坠，无法控制；为弄清楚具体状况，和机组一起飞了半小时，后来重现了故障，具体情况为：ZUUU 02L 、Approach 20NM, HYD G+B LO PRESSURE, 机组飞行过程中超速，飞机直接低头下坠，俯仰配平DN 4.0且无法拉动配平轮，配平轮一直往前打。机组拉不起来，冻结拍照发现没有ELAC故障，和其无关。 晚上训练结束后测试该故障： 1、 ZUUU 02L 、Approach 20NM， 配平正常然后解冻，飞机突然低头下坠，查看F/CTL页面发现和之前一样俯仰配平DN 4.0且无法拉动配平轮，配平轮一直往前打，冻结拍照。 2、ZUUU 02L 、Approach 20NM, HYD G+B LO PRESSURE, 一超速飞机直接低头下坠，和机组遇到的一样，且也没有ELAC故障。怀疑是缝翼卡组引起的。 3、ZUUU 02L 、Approach 20NM，设置SLATS LOCKED， 解冻飞行一切正常，即使超速也未出异常。 4、再次尝试重现第二次的故障，没能重现，需继续排查原因。',last2commits:[{auth:'郑卿',time:'2019-03-03 01:34:27',content:'在5989和5015上测试ZUUU 02L 、Approach 20NM, HYD G+B LO PRESSURE，模拟机正常飞行，没有出现下坠的情况。'},{auth:'付东明',time:'2019-03-04 12:13:25',content:'为保证追踪系统数据链完整性，不提供删除，对于不恰当的地方可以追加信息进行描述，请大家仔细斟酌，认真填写!'}]},
            {drid:'20190218085722',time:"2019-03-04 17:32:47",sim:'5015',cap:'李跃',auth:'李跃',coworkers:'赵育鹏',support:'xxxyyy',bt:0.5,class:3,tl:3,content:'视景柜VCC面板风扇图形的红色指示灯闪烁并发出报警声，视景整体功能没有影响，利用机组休息时间将VCC完全断电后重启，该问题消除，保持观察。',last2commits:[{auth:'郑卿',time:'2019-03-03 01:34:27',content:'该问题没有继续出现，偶发故障，可以关闭'},{auth:'付东明',time:'2019-03-04 12:13:25',content:'为保证追踪系统数据链完整性，不提供删除，对于不恰当的地方可以追加信息进行描述，请大家仔细斟酌，认真填写!'}]},
        ];
        this.$store.commit("refreshedMalf", newmalfs);
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
.works-content {
    width:100%;
    min-height: 100%;
    height: 100%;
    display: flex;
}
.cards-single {
    width: 100%;
    display:flex;
    justify-content: center;
    min-width: 600px;
}
.cards-single:nth-child(n+2) {
    margin-top: 30px;
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
/* training recorder  */
.tlist-one {
    width: 100%;
    height: 220px;
    display: flex;
    align-items: center;
    justify-content: center;
    padding-bottom: 15px;
}
.tlist-one:not(:last-child) {
    border-bottom: 1px dashed #ded;
}
.tone-img {
    height: 200px;
    width:320px;
    display: flex;
    align-items: center;
    justify-content: center;
    float: left;
    border: 1px dashed #fff;
    border-radius: 10px;
}
.tone-img:hover {
    border: 1px dashed #409EFF;
}
.tone-img > img {
    max-height: 195px;
    max-width: 315px;
    cursor: zoom-in;
    border-radius: 10px;
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
    border-bottom: 1px solid #ded;
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
/* cards - malfs */
.malfs-one:nth-child(n+2) {
    margin-top: 20px;
}
.mone-title {
    display: flex;
    color: #409EFF;
}
.mone-title > div:nth-child(n+2) {
    margin-left: 10px;
}
.mone-title:hover > div {
    text-decoration: underline;
    cursor: pointer;
}
.mone-content {
    padding-bottom: 20px;
}
.mone-onecommit {
    padding: 10px 0 10px 0;
    border-top: 1px solid #ded;
    line-height: 30px;
}
</style>
