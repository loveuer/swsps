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
                            <div style="width:120px;text-indent:5px;">{{ his.date }}</div>
                            <div style="width:100px;">{{ his.time }}</div>
                            <div style="width:100px;">{{ his.auth }}</div>
                            <div style="width:100px;">{{ his.short }}</div>
                        </template>
                        <div style="text-indent:5px;">
                            <div>数　量 变化:　　{{ his.amount_chg }}
                                <el-tooltip class="item" effect="dark" content="点击查看备件" placement="top">
                                    <i class="el-icon-more" style="float:right;padding:0 8px 0 10px;cursor:pointer;" @click="meetSps(his.id)"></i>
                                </el-tooltip>
                            </div>
                            <div>模拟机 变化:　　{{ his.last_sim }} <i class="el-icon-right" style="padding:0 10px 0 10px;"></i> {{ his.next_sim }}</div>
                            <div>状　态 变化:　　{{ his.last_sts }} <i class="el-icon-right" style="padding:0 10px 0 10px;"></i> {{ his.next_sts }}</div>
                            <div>位　置 变化:　　{{ his.last_pos }} <i class="el-icon-right" style="padding:0 10px 0 10px;"></i> {{ his.next_pos }}</div>
                            <div>备　注 　　:　　{{ his.comment }}</div>
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
        this.$http.get("/api/sphis/0", {withCredentials: true})
            .then(resp => {
                let count = 1;
                for(let h of resp.data) {
                    if (count > 10) {
                        break;
                    };
                    h.date = h.time.split("T")[0];
                    h.time = h.time.split("T")[1].split("+")[0];
                    this.last10histories.push(h);
                    count++;
                };
            })
            .catch(err => {
                console.log(err);
                switch (err.response.status) {
                    case 401:
                        this.$router.push("/login");
                        break;
                };
            });
    },
};
</script>

<style scoped>
.his-title {
    display: flex;
}
</style>
