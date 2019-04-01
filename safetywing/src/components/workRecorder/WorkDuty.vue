<template>
    <div>
        <div>
            <loveuer-menu ref="uMenu"></loveuer-menu>
        </div>
        <div style="height:20px;widht:100%zoom:100%;"></div>
        <div style="margin-left:30px;width:100%;">
            <table>
                <thead>
                    <tr>
                        <th style="width:55px;">
                            <el-checkbox :indeterminate="isIndeterminate" v-model="checkAll" @change="handleCheckAllChange"></el-checkbox>
                        </th>
                        <th style="width:100px;">日期</th>
                        <th style="width:80px;">星期</th>
                        <th>国航白班</th>
                        <th>川大白班</th>
                        <th>国航夜班</th>
                        <th>川大夜班</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(dt, index) of dutiesWeeked" :key="index" class="tr-contents">
                        <td style="width:55px;text-align:center;">
                            <el-checkbox v-model="duties[index].checked" @change="handleCheckBoxClick(index)"></el-checkbox>
                        </td>
                        <td v-text="dt.date" style="width:100px;"></td>
                        <td v-text="dt.week" style="width:80px;"></td>
                        <td v-html="dt.acday" v-on:dblclick="handleDoubleClick(index,'acday')"></td>
                        <td v-html="dt.suday" v-on:dblclick="handleDoubleClick(index,'suday')"></td>
                        <td v-html="dt.acnight" v-on:dblclick="handleDoubleClick(index,'acnight')"></td>
                        <td v-html="dt.sunight" v-on:dblclick="handleDoubleClick(index,'sunight')"></td>
                    </tr>
                </tbody>
            </table>
            <div style="margin-top:20px;">
                <el-button @click="copyDuties">复制</el-button>
                <el-button >取消</el-button>
            </div>
        </div>
        <div class="chgWorkersPlace">
            <el-dialog 
                :visible.sync="dialogVisible"
                title="修改人员">
                <div>
                    <div></div>
                </div>
                <span slot="footer" class="dialog-footer">
                    <el-button @click="dialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="dialogVisible = false">确 定</el-button>
                </span>
            </el-dialog>
        </div>
        <div style="height:100px;width:100%;zoom:100%;"></div>
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";
import dateformat from "dateformat";

export default {
    data() {
        return {
            isIndeterminate: false,
            checkAll: false,
            duties: [],
            dialogVisible: false,
            pendingWorkersStyled: [],
        };
    },
    methods: {
        handleCheckAllChange: function(val) {
            if (val) {
                for (let d of this.duties) {
                    d.checked = true;
                };
            } else {
                for (let d of this.duties) {
                    d.checked = false;
                };
            };
            this.isIndeterminate = false;
        },
        handleCheckBoxClick: function(row) {
            this.checkAll = this.dutiesCheckedIndexs.length === this.duties.length;
            this.isIndeterminate = this.dutiesCheckedIndexs.length > 0 && this.dutiesCheckedIndexs.length < this.duties.length;
        },
        getDutyByYM: function(ym) {
            this.duties = [
                {checked:false,date:'2019/03/01',week:3,acday:'<font style="font-weight:bold;">陈俊峰</font> | <font>李津</font>',suday:`<font style="font-style:italic;">凌华南</font> | <font style="color:#ff4500">江富强(年假)</font> | <font>郑卿</font>`,acnight:'<font>陈本志</font> | <font>李永翔</font>',sunight:'<font>付东明</font> | <font>宋继朋</font> | <font>黄一林</font>'},
                {checked:false,date:'2019/03/02',week:4,acday:'<font>张伟</font> | <font>郭建</font>',suday:'<font>李跃</font> | <font>赵育鹏</font>',acnight:'<font>陈俊峰</font> | <font>李津</font>',sunight:'<font>凌华南</font> | <font>江富强</font> | <font>郑卿</font>'},
                {checked:false,date:'2019/03/03',week:5,acday:'<font>胡彬彬</font> | <font>魏宇东</font>',suday:'<font>鲁书贤</font> | <font>张显刚</font>',acnight:'<font>张伟</font> | <font>郭建</font>',sunight:'<font>刘博强（白班换夜班）</font> | <font>李跃</font> | <font>赵育鹏</font>'},
                {checked:false,date:'2019/03/04',week:6,acday:'<font>陈本志</font> | <font>李永翔</font>',suday:'<font>付东明</font> | <font>宋继朋</font> | <font>黄一林</font>',acnight:'<font>胡彬彬</font> | <font>魏宇东</font>',sunight:'<font>刘博强</font> | <font>鲁书贤</font> | <font>张显刚</font>'},
            ];
        },
        copyDuties: function() {
            if (!this.dutiesCheckedIndexs) {
                return;
            };
            if (this.dutiesCheckedIndexs.length > 99 || this.duties.length > 199) {
                this.$alert("请一次不要复制太多内容", "警告", {
                    confirmButtonText: '确定',
                });
                return;
            };
            let newd = this.dutiesCheckedIndexs.sort((a,b) => {
                return a - b;
            });
            for (let i of this.dutiesCheckedIndexs) {
                let lastDuty = this.duties[this.duties.length-1];
                let lastDate = new Date(lastDuty.date.replace(/\//gi, "-"));
                let newDate = dateformat(lastDate.setHours(lastDate.getHours() + 24), "yyyy/mm/dd");
                let newWeek = lastDuty.week + 1 > 7 ? lastDuty.week + 1 - 7 : lastDuty.week + 1;
                let newDuty = {date:newDate, week:newWeek, acday:this.duties[i].acday, suday:this.duties[i].suday,acnight:this.duties[i].acnight,sunight:this.duties[i].sunight};
                this.duties.push(newDuty);
            };
            
        },
        handleDoubleClick: function(row, col) {
            this.pendingWorkersStyled = this.duties[row][col].split(" | ");
            this.dialogVisible = true;
            console.log(this.pendingWorkers);
        },
    },
    computed: {
        dutiesWeeked: function() {
            let wmap = {1:'星期一',2:'星期二',3:'星期三',4:'星期四',5:'星期五',6:'星期六',7:'星期天'};
            let dw = [];
            let count = 0;
            for (let d of this.duties) {
                dw.push({
                    index:count, 
                    date:d.date, 
                    week:wmap[d.week], 
                    acday:d.acday.replace(/ \| /gi, ", "), 
                    suday:d.suday.replace(/ \| /gi, ", "), 
                    acnight:d.acnight.replace(/ \| /gi, ", "), 
                    sunight:d.sunight.replace(/ \| /gi, ", "), 
                });
                count++;
            };
            return dw;
        },
        dutiesCheckedIndexs: function() {
            let list = [];
            let count = 0;
            for (let d of this.duties) {
                if (d.checked) {
                    list.push(count);
                };
                count++;
            };
            return list;
        },
        pendingWorkers: function() {
            let list = [];
            for (let w of this.pendingWorkersStyled) {
                if (!w.match(/^<font.*>.+<\/font>$/)) {
                    this.$notify.error({
                        title:'错误',
                        message:'发生了未知错误, 请联系管理员',
                        position:'top-left',
                    });
                    this.dialogVisible = false;
                    console.log(`人员格式不对, match font 失败`);
                    return;
                };
                let inner = w.replace(/^<font.*?>/, "").replace("</font>", "");
                let color = null;
                let bold = false;
                let decoration = false;
                let italic = false;
                // find color
                let colorIndex = w.search("color");
                if (colorIndex !== -1) {
                    color = w.slice(colorIndex+6, colorIndex+13);
                };
                // find bold
                if (w.match("font-weight")) {
                    bold = true;
                };
                // find decoration
                if (w.match("text-decoration")) {
                    decoration = true;
                };
                // find italic
                if (w.match("font-style")) {
                    italic = true;
                };

                list.push({
                    inner:inner, color:color, bold:bold, decoration:decoration, italic:italic,
                });
            };
            return list;
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
table {
    border-collapse: collapse;
    width:90%;
    height: 100%;
    display: flex;
    flex-flow: column;
}
table thead {
    flex: 0 0 auto;
    width: calc(100% - 0.9em);
}
table tbody {
    height: 600px;
    flex: 1 1 auto;
    display: block;
    overflow: auto;
}
table tbody tr {
    width: 100%;
}
table thead, table tbody tr {
    display: table;
    table-layout: fixed;
}
table thead tr,table tbody tr {
    height: 45px;
    width:100%;
    line-height: 45px;
    border-bottom: 1px solid rgb(232,232,232);
}
table thead tr th:nth-child(n+2), table tbody tr td:nth-child(n+2) {
    padding: 0 20px 0 20px;
}
table tbody tr:hover {
    background-color: rgb(245,247,250);
}
table thead tr th:first-child {
    text-align: center;
}
table thead tr th:nth-child(n+2) {
    text-align: left;
}
table tbody tr td:first-child {
    text-align: center;
}
table tbody tr td:nth-child(n+2) {
    text-align: left;
}
table th {
    font-size: 14px;
    color: rgb(144,147,153);
}
table td {
    font-size: 14px;
    color: rgb(72,72,72);
    cursor: pointer;
}
</style>
