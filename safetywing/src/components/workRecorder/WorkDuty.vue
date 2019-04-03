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
                    <div v-for="(p, index) of pendingWorkers" :key="index" class="one-worker">
                        <div class="edit-workers">
                            <div style="display:flex;margin-bottom:5px;">
                                <div>
                                    <el-button @click="p.bold=!p.bold" size="mini" round>B</el-button>
                                </div>
                                <div style="margin-left:10px;">
                                    <el-button @click="p.italic=!p.italic" size="mini" round>I</el-button>
                                </div>
                                <div style="margin-left:10px;">
                                    <el-button @click="p.decoration=!p.decoration" size="mini" round>U</el-button>
                                </div>
                                <div style="margin-left:10px;">
                                    <el-color-picker
                                        v-model="p.color"
                                        size="mini"
                                        :predefine="predefineColors">
                                    </el-color-picker>
                                </div>
                            </div>
                            <div style="height:35px;width:100%;display:flex;align-items: center;">
                                <div
                                    class="everyWorker"
                                    contenteditable="true"
                                    tabindex="-1"
                                    @input="workerinputChange($event, index)"
                                    v-text="p.name"
                                    :style="{color:p.color}"
                                    v-bind:class="{pwisBold:p.bold, pwisItalic:p.italic, pwisDecoration:p.decoration}"
                                    >
                                </div>
                                <div style="height:35px;display:flex;align-items: center;margin-left:10px;">
                                    <el-button type="danger" icon="el-icon-delete" circle size="small" @click="deleteWorker(index)"></el-button>
                                </div>
                            </div>
                            
                        </div>
                    </div>
                    <div>
                        <el-button icon="el-icon-plus" type="success" @click="appendWorker"></el-button>
                    </div>
                </div>
                <div style="margin-top:20px;">
                    <el-button type="primary" @click="changeWorkers">确 定</el-button>
                    <el-button @click="dialogVisible = false">取 消</el-button>
                </div>
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
            pendingWorkers: [],
            pendingWorkerPos: {row:null, col:null},
            predefineColors: [
                '#ff4500',
                '#ff8c00',
                '#ffd700',
                '#90ee90',
                '#00ced1',
                '#1e90ff',
                '#c71585',
            ],
            selColor: "#000000",
        };
    },
    methods: {
        appendWorker: function() {
            this.pendingWorkers.push({
                name:'', color: '#484848', bold: false, italic: false, decoration: false
            });
        },
        deleteWorker: function(index) {
            this.pendingWorkers.splice(index, 1);
        },
        changeWorkers: function() {
            this.dialogVisible = false;
            let list = [];
            for (let w of this.pendingWorkers) {
                list.push({
                    name:w.newname, color:w.color, bold:w.bold, italic:w.italic, decoration:w.decoration
                });
            };
            this.duties[this.pendingWorkerPos.row][this.pendingWorkerPos.col] = list;
        },
        workerinputChange: function(e, i) {
            this.pendingWorkers[i].newname = e.target.innerHTML;
        },
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
                {
                    checked:false,date:'2019/03/01',week:3,
                    acday:[{name:'陈俊峰',color:'#484848',bold:false,italic:false,decoration:false},{name:'李津',color:'#484848'}],
                    suday:[{name:'凌华南',color:'#484848',bold:true,italic:true,decoration:false},{name:'江富强',color:'#ff4500',bold:true,italic:false,decoration:false},{name:'郑卿',color:'#484848',bold:false,italic:false,decoration:false}],
                    acnight:[{name:'陈本志',color:'#484848',bold:false,italic:false,decoration:false},{name:'李永翔',color:'#484848',bold:false,italic:false,decoration:false}],
                    sunight:[{name:'付东明',color:'#484848',bold:false,italic:false,decoration:false},{name:'宋继朋',color:'#484848',bold:false,italic:false,decoration:false},{name:'黄一林',color:'#484848',bold:false,italic:false,decoration:false}],
                },
                {
                    checked:false,date:'2019/03/02',week:4,
                    acday:[{name:'张伟',color:'#484848',bold:false,italic:false,decoration:false},{name:'郭建',color:'#484848',bold:false,italic:false,decoration:false}],
                    suday:[{name:'李跃',color:'#484848',bold:false,italic:false,decoration:false},{name:'赵育鹏',color:'#484848',bold:false,italic:false,decoration:false}],
                    acnight:[{name:'陈俊峰',color:'#484848',bold:false,italic:false,decoration:false},{name:'李津',color:'#484848',bold:false,italic:false,decoration:false}],
                    sunight:[{name:'凌华南',color:'#484848',bold:false,italic:false,decoration:false},{name:'江富强',color:'#484848',bold:false,italic:false,decoration:false},{name:'郑卿',bold:false,italic:false,decoration:false}],
                },
                {
                    checked:false,date:'2019/03/03',week:5,
                    acday:[{name:'胡彬彬',color:'#484848',bold:false,italic:false,decoration:false},{name:'魏宇东',color:'#484848',bold:false,italic:false,decoration:false}],
                    suday:[{name:'鲁书贤',color:'#484848',bold:false,italic:false,decoration:false},{name:'张显刚',color:'#484848',bold:false,italic:false,decoration:false}],
                    acnight:[{name:'张伟',color:'#484848',bold:false,italic:false,decoration:false},{name:'郭建',color:'#484848',bold:false,italic:false,decoration:false}],
                    sunight:[{name:'刘博强（白班换夜班）',color:'#484848',bold:false,italic:false,decoration:false},{name:'李跃',color:'#484848',bold:false,italic:false,decoration:false},{name:'赵育鹏',color:'#484848',bold:false,italic:false,decoration:false}],
                },
                {
                    checked:false,date:'2019/03/04',week:6,
                    acday:[{name:'陈本志',color:'#484848',bold:false,italic:false,decoration:false},{name:'李永翔',color:'#484848',bold:false,italic:false,decoration:false}],
                    suday:[{name:'付东明',color:'#484848',bold:false,italic:false,decoration:false},{name:'宋继朋',color:'#484848',bold:false,italic:false,decoration:false},{name:'黄一林',color:'#484848',bold:false,italic:false,decoration:false}],
                    acnight:[{name:'胡彬彬',color:'#484848',bold:false,italic:false,decoration:false},{name:'魏宇东',color:'#484848',bold:false,italic:false,decoration:false}],
                    sunight:[{name:'刘博强',color:'#484848',bold:false,italic:false,decoration:false},{name:'鲁书贤',color:'#484848',bold:false,italic:false,decoration:false},{name:'张显刚',color:'#484848',bold:false,italic:false,decoration:false}],
                },
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
            this.dialogVisible = true;
            this.pendingWorkerPos.row = row;
            this.pendingWorkerPos.col = col;
            // 
            this.pendingWorkers = this.duties[row][col];
            for (let w of this.pendingWorkers) {
                w.newname = w.name;
            };
        },
        styler: function(ws) {
            let wl = [];
            for (let w of ws) {
                let str = `<font style="color:${w.color};font-weight:${w.bold?'bold':'normal'};font-style:${w.italic?'italic':'normal'};text-decoration:${w.decoration?'underline':'none'}">${w.name}</font>`;
                wl.push(str);
            };
            return wl.join(", ");
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
                    acday:this.styler(d.acday),
                    suday:this.styler(d.suday),
                    acnight:this.styler(d.acnight),
                    sunight:this.styler(d.sunight),
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
.one-worker {
    padding: 20px 0 20px 0;
}
.one-worker:first-child {
    margin-top: -20px;
}
.everyWorker {
    height: 35px;
    line-height: 35px;
    padding: 2px 0 0 2px; 
    text-indent: 10px;
    border: 1px solid rgb(200,200,200);
    border-radius: 5px;
    outline: none;
    width: 50%;
}
.everyWorker:hover {
    border: 1px solid #409EFF;
}
.everyWorker:focus {
    border: 1px solid #409EFF;
    box-shadow: 0 0 5px #409EFF;
}
.pwisBold {
    font-weight: bold;
}
.pwisItalic {
    font-style: italic;
}
.pwisDecoration {
    text-decoration: underline;
}
</style>
