<template>
    <div id="logsManager">
        <el-row :gutter="10" style="width:900px;">
            <el-col :span="2">
                <el-button type="success" icon="el-icon-plus" @click="dialogFormVisible = true"></el-button>
            </el-col>

            <el-col :span="4">
                <el-select v-model="this.psim" placeholder="模拟机">
                    <el-option :key="5978" :value="5978" :label="5978"></el-option>
                    <el-option :key="5989" :value="5989" :label="5989"></el-option>
                    <el-option :key="5008" :value="5008" :label="5008"></el-option>
                    <el-option :key="5015" :value="5015" :label="5015"></el-option>
                </el-select>
            </el-col>

            <el-col :span="6">
                <el-date-picker v-model="pdate.start" type="date" placeholder="起始日期"></el-date-picker>
            </el-col>

            <el-col :span="6">
                <el-date-picker v-model="pdate.end" type="date" placeholder="截止日期"></el-date-picker>
            </el-col>

            <el-col :span="4">
                <el-input placeholder="关键字" v-model="this.pkey"></el-input>
            </el-col>

            <el-col :span="2">
                <el-button type="primary" @click="searchLogs">搜索</el-button>
            </el-col>
            
        </el-row>
        <el-row style="margin-top:20px;"></el-row>
        <el-row v-show="this.justStart" v-text="justStartValue" style="text-align:center;font-size:20px;font-weight:bold"></el-row>
        <el-row>
            <el-table :data="this.logs" height="700">
                <el-table-column prop="date" label="date" width="120"></el-table-column>
                <el-table-column prop="time" label="time" width="120"></el-table-column>
                <el-table-column prop="auth" label="auth" width="120"></el-table-column>
                <el-table-column prop="sim" label="sim" width="120"></el-table-column>
                <el-table-column prop="detail" label="log" width="900"></el-table-column>
            </el-table>
        </el-row>

        <el-dialog title="新增Log" :visible.sync="dialogFormVisible">
            <el-form label-position="left" label-width="80px" style="min-width:100%;">
                <el-form-item label="模拟机:" style="min-width:100%;">
                    <el-select placeholder="模拟机" v-model="newlog.sim">
                        <el-option label="5978" value="5978"></el-option>
                        <el-option label="5989" value="5989"></el-option>
                        <el-option label="5008" value="5008"></el-option>
                        <el-option label="5015" value="5015"></el-option>
                    </el-select>
                    
                </el-form-item>
                <el-form-item label="具体内容:">
                        <!-- <el-input type="textarea" :rows="4" style="font-size:16px;" v-model="newlog.detail">
                        </el-input> -->
                    <tinyMceEditor api-key="API_KEY" :init="{plugins: 'wordcount'}" v-model="newlog.detail"></tinyMceEditor>
                </el-form-item>
            </el-form>
                <el-button @click="dialogFormVisible = false">Cancel</el-button>
                <el-button type="primary" @click="dialogFormVisible = false">Confirm</el-button>
        </el-dialog>
    </div>
</template>

<script>
import dateFormat from "dateformat";
import tinyMceEditor from '@tinymce/tinymce-vue';

export default {
    data() {
        return {
            logs: {

            },
            psim: "",
            pdate: {
                start: "",
                end: "",
            },
            pkey: "",
            dialogFormVisible: false,
            newlog: {
                sim: "",
                detail: "",
            },
            justStart: true,
            justStartValue: dateFormat(Date.now(), "yyyy-mm-dd") + " logs:",
            mockLogs: [
                {"date":"2019-03-13", "time":"17:37:32", "auth":"赵育鹏", "sim":"5978", "detail":"这是一条很长很长的log 17:26，机组训练提前结束，机组反映，在成都机场爬升时ECAM上出现了ELAC 2 fault 故障，且出现自动驾驶接通后自动断开的情况。上模拟机测试，在测试的整个过程30分钟，飞机起飞后接通AP（测试中交替接通了2部AP），均未出现AP自动断开，ECAM也未出ELAC 2 fault 故障信息，在降落和巡航过程的测试结果也未重现该故障。将ELAC1和ELAC2的跳开关重置后，询问后续训练机组是否出现该故障，均答复未出现前述故障。"},
                {"date":"2019-03-12", "time":"15:24:12", "auth":"宋继朋", "sim":"5989", "detail":"mock log 2"},
                {"date":"2019-03-11", "time":"15:35:32", "auth":"李永翔", "sim":"5978", "detail":"mock log 3"},
                {"date":"2019-03-09", "time":"12:45:09", "auth":"李跃", "sim":"5008", "detail":"mock log 4"},
                {"date":"2019-03-01", "time":"23:57:46", "auth":"付东明", "sim":"5015", "detail":"mock log 5"},
            ],
        };
    },
    methods: {
        getLogs: function(option) {
            this.logs = this.mockLogs;
        },
        searchLogs: function() {

        },
    },
    components: {
        "tinyMceEditor": tinyMceEditor,
    },
    mounted() {
        let initOption = {
            sim: "",
            startDate: dateFormat(Date.now(), "yyyy-mm-dd"),
            endDate: "",
            keyWord: "",
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

