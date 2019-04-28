<template>
    <div id="container">
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <el-tooltip class="item" effect="dark" content="查看全部历史" placement="top">
                    <el-link href="/history/all" style="font-size:16px;font-weight:bold;">备件历史</el-link>
                </el-tooltip>
            </div>
            <div>
                <el-collapse>
                    <el-collapse-item 
                        v-for="(his, index) of last10histories"
                        :key="index"
                        :name="index">
                        <template slot="title" class="his-title">
                            <div style="width:120px;">{{ his.date }}</div>
                            <div style="width:100px;">{{ his.detail.time }}</div>
                            <div style="width:100px;">{{ his.detail.auth }}</div>
                            <div style="width:100px;">{{ his.detail.short }}</div>
                        </template>
                        <div>
                            <div>数　量 变化:　　{{ his.detail.amount }}
                                <el-tooltip class="item" effect="dark" content="点击查看备件" placement="top">
                                    <i class="el-icon-more" style="float:right;padding:0 8px 0 10px;cursor:pointer;" @click="meetSps(his.detail.id)"></i>
                                </el-tooltip>
                            </div>
                            <div>模拟机 变化:　　{{ his.detail.oldsim }} <i class="el-icon-right" style="padding:0 10px 0 10px;"></i> {{ his.detail.newsim }}</div>
                            <div>状　态 变化:　　{{ his.detail.oldsts }} <i class="el-icon-right" style="padding:0 10px 0 10px;"></i> {{ his.detail.newsts }}</div>
                            <div>位　置 变化:　　{{ his.detail.oldpos }} <i class="el-icon-right" style="padding:0 10px 0 10px;"></i> {{ his.detail.newpos }}</div>
                        </div>
                    </el-collapse-item>
                </el-collapse>
            </div>
        </el-card>
    </div>
</template>

<script>
export default {
    data() {
        return {
            last10histories: [],
        };
    },
    methods: {
        meetSps(spid){
            console.log("meet sps index: ", spid);
        },
    },
    mounted() {
        let mockhis = {date: '2019/04/28', detail: {
                time: "12:33:21", auth: '赵育鹏', short: '更换备件 出库',
                amount: 0, oldsim: '5978', newsim: '5978',
                oldsts: '备件', newsts: '使用',
                oldpos: 'b-01-1', newpos: '用于5978 SR2500备件服务器',
                recorder: '用于5978 SR2500备件服务器',
                id: '2233',
            }};
        for (let i=0;i<10;i++) {
            this.last10histories.push(mockhis);
        };
    },
};
</script>

<style scoped>
.his-title {
    display: flex;
}
</style>
