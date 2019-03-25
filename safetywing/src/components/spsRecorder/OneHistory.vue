<template>
    <div>
        <!-- <el-table :data="hiss" style="width:100%">
            <el-table-column type="expand">
                <template slot-scope="props">
                    <el-form label-position="left" class="demo-table-expand" label-width="120px">
                        <el-form-item label="数量变化:">
                            <span>{{ props.row.amount_chg }}</span>
                        </el-form-item>
                        <el-form-item label="状态变化:">
                            <span>{{ `${props.row.last_sts} -> ${props.row.next_sts}` }}</span>
                        </el-form-item>
                        <el-form-item label="模拟机变化">
                            <span>{{ `${props.row.last_sim} -> ${props.row.next_sim}` }}</span>
                        </el-form-item>
                        <el-form-item label="位置变化">
                            <span>{{ `${props.row.last_pos} -> ${props.row.next_pos}` }}</span>
                        </el-form-item>
                    </el-form>
                </template>
            </el-table-column>
            <el-table-column label="日期" prop="simple_date"></el-table-column>
            <el-table-column label="时间" prop="simple_time"></el-table-column>
            <el-table-column label="人员" prop="auth"></el-table-column>
            <el-table-column label="简述" prop="short"></el-table-column>
        </el-table> -->
        <el-collapse v-model="openCollapse" style="max-width:100%;width:700px;">
            <el-collapse-item v-for="(his, index) of hiss" :key="index" :name="index">
                <template slot="title">
                    <div style="display:flex;">
                        <div style="margin-left:5px;">{{ his.simple_date }}</div>
                        <div style="margin-left:30px;">{{ his.simple_time }}</div>
                        <div style="margin-left:30px;">{{ his.auth }}</div>
                        <div style="margin-left:30px;">{{ his.short }}</div>
                    </div>
                </template>
                <div>
                    <div class="detail-label">
                        <div>数量变化:</div>
                        <div>{{ his.amount_chg }}</div>
                    </div>
                    <div class="detail-label">
                        <div>状态变化:</div>
                        <div>{{ `${his.last_sts} -> ${his.next_sts}` }}</div>
                    </div>
                    <div class="detail-label">
                        <div>模拟机变化:</div>
                        <div>{{ `${his.last_sim} -> ${his.next_sim}` }}</div>
                    </div>
                    <div class="detail-label">
                        <div>位置变化:</div>
                        <div>{{ `${his.last_pos} -> ${his.next_pos}` }}</div>
                    </div>
                </div>
            </el-collapse-item>
        </el-collapse>
    </div>
</template>

<script>
export default {
    data() {
        return {
            hiss: [],
            openCollapse: [0,1,2],
        };
    },
    mounted() {
        console.log(this.$route.params.spid);
        this.$http.get('/api/sphis/one/' + this.$route.params.spid)
            .then(resp => {this.hiss = resp.data;console.log(resp.data)})
            .catch(error => console.log('sphis one err: ', error));
    },
};
</script>

<style scoped>
.detail-label {
    display: flex;
    height:40px;
    align-items: center;
}
.detail-label > div:first-child {
    width:150px;
    margin-left: 5px;
}
</style>

