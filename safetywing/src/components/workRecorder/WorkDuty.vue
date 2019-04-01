<template>
    <div>
        <div>
            <loveuer-menu ref="uMenu"></loveuer-menu>
        </div>
        <div style="height:30px;widht:100%zoom:100%;"></div>
        <div style="margin-left:30px;">
            <el-table
                :data="dutiesWeeked"
                tooltip-effect="dark"
                style="width:75%"
                height="700px"
                @selection-change="handleSelectionChange"
                ref="multipleTable">
                <el-table-column type="selection" width="55"></el-table-column>
                <el-table-column prop="date" label="日期" width="120"></el-table-column>
                <el-table-column prop="week" label="星期" width="120"></el-table-column>
                <el-table-column prop="acday" label="国航白班"></el-table-column>
                <el-table-column prop="suday" label="川大白班"></el-table-column>
                <el-table-column prop="acnight" label="国航夜班"></el-table-column>
                <el-table-column prop="sunight" label="川大夜班"></el-table-column>
            </el-table>
            <div style="margin-top: 20px">
                <el-button @click="copyDuties">复制</el-button>
                <el-button >取消</el-button>
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
            duties: [],
            dutiesSelected: [],
        };
    },
    methods: {
        getDutyByYM: function(ym) {
            this.duties = [
                {date:'2019/03/01',week:3,acday:'陈俊峰，李津',suday:`凌华南，<font style="color:red">江富强(年假)</font>,郑卿`,acnight:'陈本志，李永翔',sunight:'付东明，宋继朋，黄一林'},
                {date:'2019/03/02',week:4,acday:'张伟，郭建',suday:'李跃，赵育鹏',acnight:'陈俊峰，李津',sunight:'凌华南，江富强,郑卿'},
                {date:'2019/03/03',week:5,acday:'胡彬彬,魏宇东',suday:'鲁书贤，张显刚',acnight:'张伟，郭建',sunight:'刘博强（白班换夜班），李跃，赵育鹏'},
                {date:'2019/03/04',week:6,acday:'陈本志，李永翔',suday:'付东明，宋继朋，黄一林',acnight:'胡彬彬，魏宇东',sunight:'刘博强，鲁书贤，张显刚'},
            ];
        },
        handleSelectionChange: function(rows) {
            this.dutiesSelected = rows;
        },
        copyDuties: function() {
            if (!this.dutiesSelected) {
                return;
            };
            let newd = this.dutiesSelected.sort((a,b) => {
                return a.index - b.index;
            });
            for (let d of newd) {
                let lastDuty = this.duties[this.duties.length-1];
                let lastDate = new Date(lastDuty.date.replace(/\//gi, "-"));
                let newDate = dateformat(lastDate.setHours(lastDate.getHours() + 24), "yyyy/mm/dd");
                let newWeek = lastDuty.week + 1 > 7 ? lastDuty.week + 1 - 7 : lastDuty.week + 1;
                let newDuty = {date:newDate, week:newWeek, acday:d.acday, suday:d.suday, acnight:d.acnight, sunight:d.sunight};
                this.duties.push(newDuty);
            };
            
        },
    },
    computed: {
        dutiesWeeked: function() {
            let wmap = {1:'星期一',2:'星期二',3:'星期三',4:'星期四',5:'星期五',6:'星期六',7:'星期天'};
            let dw = [];
            let count = 0;
            for (let d of this.duties) {
                dw.push(
                    {index:count, date:d.date, week:wmap[d.week], acday:d.acday, suday:d.suday, acnight:d.acnight, sunight:d.sunight}
                );
                count++;
            };
            return dw;
            // return this.duties.map(d => {
            //     return {date:d.date, week:wmap[d.week], acday:d.acday, suday:d.suday, acnight:d.acnight, sunight:d.sunight};
            // });
        },
    },
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/workRecorder/duty";
        this.getDutyByYM();
    },
    components: {
        "loveuer-menu": loveuerMenu,
    },
};
</script>

<style scoped>

</style>
