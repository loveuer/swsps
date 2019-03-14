<template>
    <div id="logsManager">
        <el-row :gutter="10" style="width:900px;">
            <el-col :span="2">
                <el-button type="success" icon="el-icon-plus" @click="dialogFormVisible = true"></el-button>
            </el-col>

            <el-col :span="4">
                <el-select v-model="searchlog.sim" placeholder="模拟机">
                    <el-option :key="5978" :value="5978" :label="5978"></el-option>
                    <el-option :key="5989" :value="5989" :label="5989"></el-option>
                    <el-option :key="5008" :value="5008" :label="5008"></el-option>
                    <el-option :key="5015" :value="5015" :label="5015"></el-option>
                </el-select>
            </el-col>

            <el-col :span="6">
                <el-date-picker v-model="searchlog.start" type="date" placeholder="起始日期"></el-date-picker>
            </el-col>

            <el-col :span="6">
                <el-date-picker v-model="searchlog.end" type="date" placeholder="截止日期"></el-date-picker>
            </el-col>

            <el-col :span="4">
                <el-input placeholder="关键字" v-model="searchlog.key"></el-input>
            </el-col>

            <el-col :span="2">
                <el-button type="primary" @click="searchLogs">搜索</el-button>
            </el-col>
            
        </el-row>
        <el-row style="margin-top:30px;"></el-row>
        <el-row>
            <el-timeline :reverse="false">
                <el-timeline-item
                    v-for="(onedaylog, index) of this.mockLogs"
                    :key="index"
                    :timestamp="onedaylog.date"
                >
                    <el-table
                        :data="onedaylog.logs"
                    >
                        <el-table-column prop="time" lable-="time" width="80"></el-table-column>
                        <el-table-column prop="auth" lable-="auth" width="80"></el-table-column>
                        <el-table-column prop="sim" lable-="sim" width="80"></el-table-column>
                        <el-table-column prop="detail" lable-="detail" width="900"></el-table-column>
                    </el-table>
                </el-timeline-item>
            </el-timeline>
        </el-row>

        <el-dialog title="新增Log" :visible.sync="dialogFormVisible">
            <el-form label-position="left" label-width="80px" style="">
                <el-form-item label="模拟机:" style="min-width:100%;">
                    <el-select placeholder="模拟机" v-model="newlog.sim">
                        <el-option label="5978" value="5978"></el-option>
                        <el-option label="5989" value="5989"></el-option>
                        <el-option label="5008" value="5008"></el-option>
                        <el-option label="5015" value="5015"></el-option>
                    </el-select>
                    
                </el-form-item>
            </el-form>
                <div>
                    <u-textarea ref="logDetail"></u-textarea>
                </div>
                <el-button @click="dialogFormVisible = false">取消</el-button>
                <el-button type="primary" @click="addNewLog">确认</el-button>
        </el-dialog>
        
    </div>
</template>

<script>
import dateFormat from "dateformat";
import utextarea from "./Utextarea.vue";

export default {
    data() {
        return {
            logs: [],
            searchlog: {
                sim: "",
                start: "",
                end: "",
                key: "",
            },
            dialogFormVisible: false,
            newlog: {
                sim: "",
                detail: "",
            },
            justStartValue: dateFormat(Date.now(), "yyyy-mm-dd") + " logs:",
            mockLogs: [
                {
                    "date":"2019-03-13", 
                    "logs":[
                        {"time":"17:37:32", "auth":"赵育鹏", "sim":"5978", "detail":"这是一条很长很长的log 17:26，机组训练提前结束，机组反映，在成都机场爬升时ECAM上出现了ELAC 2 fault 故障，且出现自动驾驶接通后自动断开的情况。上模拟机测试，在测试的整个过程30分钟，飞机起飞后接通AP（测试中交替接通了2部AP），均未出现AP自动断开，ECAM也未出ELAC 2 fault 故障信息，在降落和巡航过程的测试结果也未重现该故障。将ELAC1和ELAC2的跳开关重置后，询问后续训练机组是否出现该故障，均答复未出现前述故障。"},
                        {"time":"16:12:34", "auth":"宋继朋", "sim":"5989", "mock": "mock data"},
                        {"time":"14:12:54", "auth":"李永翔", "sim":"5008", "mock": "mock data"},                        
                    ]
                },
                {
                    "date":"2019-03-11", 
                    "logs":[
                        {"time":"16:12:34", "auth":"宋继朋", "sim":"5989", "detail": "mock data"},
                        {"time":"14:12:54", "auth":"李永翔", "sim":"5008", "detail": "mock data"},                        
                    ]
                },
                {
                    "date":"2019-03-01", 
                    "logs":[
                        {"time":"16:12:34", "auth":"宋继朋", "sim":"5989", "detail": "mock data"},
                        {"time":"14:12:54", "auth":"李永翔", "sim":"5008", "detail": "mock data"},                        
                        {"time":"14:12:54", "auth":"刘博强", "sim":"5015", "detail": "mock data"},                        
                        {"time":"14:12:54", "auth":"李跃", "sim":"5008", "detail": "mock data"},                        
                    ]
                },
                {
                    "date":"2019-02-11", 
                    "logs":[
                        {"time":"16:12:34", "auth":"宋继朋", "sim":"5989", "detail": "mock data"},
                        {"time":"14:12:54", "auth":"李永翔", "sim":"5008", "detail": "mock data"},                        
                    ]
                }
            ]
        };
    },
    methods: {
        getLogs: function(option) {
            // option just like = {
            //     sim: "",
            //     start: "",
            //     end: "",
            //     key: "",
            // }

            // do get staff
            // then 
            this.logs = this.mockLogs;
        },
        searchLogs: function() {
            console.log(this.searchlog);
            // do post staff
            // then
            // re mount this component
        },
        addNewLog: function() {
            this.newlog.detail = this.$refs.logDetail.contentCode;
            console.log(this.newlog);
            // do post staff
            // then 
            this.dialogFormVisible = false;
        },
    },
    components: {
        "u-textarea": utextarea,
    },
    mounted() {
        let initOption = {
            sim: "",
            start: dateFormat(Date.now(), "yyyy-mm-dd"),
            end: "",
            key: "",
        }
        this.getLogs(initOption);
    },
}
</script>

<style scoped>
#logsManager{
    margin-top: 20px;
    margin-left: 20px;
    width: 100%;
    min-width: 600px;
}
</style>

