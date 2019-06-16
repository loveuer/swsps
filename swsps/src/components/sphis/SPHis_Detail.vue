<template>
    <div class="sphis-content">
        <div>
            <el-card v-for="(item, index) of dealed_sphis" :key="index" shadow="hover" style="margin-top:25px;margin-left:50px;">
                <div slot="header" class="clearfix onehis-header">
                    <div style="width:90px;">{{ item.date }}</div>
                    <div style="width:80px;">{{ item.time }}</div>
                    <div style="width:70px;">{{ item.auth }}</div>
                    <div style="width:300px;" v-bind:title="'NAME: ' + item.name">{{ item.name }}</div>
                    <div style="width:200px;margin-left:30px;" v-bind:title="'P/N: ' + item.pn">{{ item.pn }}</div>
                    <div style="margin-left:auto">
                        <el-tooltip class="item" effect="dark" content="点击查看该备件" placement="top">
                            <i class="el-icon-share" style="cursor:pointer;" @click="meetsp(item.spid)"></i>
                        </el-tooltip>
                    </div>
                </div>
                <div>
                    <div class="onehis-oneline">
                        <div class="oneline-label">数　量变化</div>
                        <div class="oneline-main">
                            {{ item.amount_chg }}
                        </div>
                    </div>
                    <div class="onehis-oneline">
                        <div class="oneline-label">状　态变化</div>
                        <div class="oneline-main">
                            <div style="width:120px;">{{ item.last_sts }}</div>
                            <i class="el-icon-right" style="width:60px;"></i>
                            <div>{{ item.next_sts }}</div>
                        </div>
                    </div>
                    <div class="onehis-oneline">
                        <div class="oneline-label">位　置变化</div>
                        <div class="oneline-main">
                            <div style="width:120px;" v-bind:title="item.last_pos">{{ item.last_pos }}</div>
                            <i class="el-icon-right" style="width:60px;"></i>
                            <div>{{ item.next_pos }}</div>
                        </div>
                    </div>
                    <div class="onehis-oneline">
                        <div class="oneline-label">模拟机变化</div>
                        <div class="oneline-main">
                            <div style="width:120px;">{{ item.last_sim }}</div>
                            <i class="el-icon-right" style="width:60px;"></i>
                            <div>{{ item.next_sim }}</div>
                        </div>
                    </div>
                    <div class="onehis-oneline">
                        <div class="oneline-label">记　　　录</div>
                        <div class="oneline-main">
                            <el-tooltip class="item" effect="dark" :content="item.comment" placement="top">
                                <div>{{ item.comment }}</div>
                            </el-tooltip>
                        </div>
                    </div>
                </div>
            </el-card>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        his: Array,
    },
    computed: {
        dealed_sphis: function() {
            let newsphis = [];
            for(let h of this.his) {
                newsphis.push({name:h.name, pn:h.pn, sn:h.sn, date:h.time.split('T')[0], time:h.time.split('T')[1].split('+')[0],
                    auth:h.auth, last_pos:h.last_pos, next_pos:h.next_pos,
                    last_sim:h.last_sim, next_sim:h.next_sim,
                    last_sts:h.last_sts, next_sts:h.next_sts,
                    short:h.short, spid:h.spid, amount_chg:h.amount_chg, comment:h.comment});
            };
            return newsphis;
        },
    },
    methods: {
        meetsp: function(spid) {
            this.$router.push("/onesp/" + spid);
        },
    },
}
</script>

<style scoped>
.sphis-content {
    width: 100%;
    display: flex;
    flex-wrap: wrap;
}
.sphis-content > div {
    width: 70%;
}
.onehis-header {
    display: flex;
    align-items: center;
}
.onehis-header > div {
    font-size: 0.9rem;
    overflow: hidden;
    white-space: nowrap;
}
.onehis-oneline {
    display: flex;
    align-items: center;
    font-size: 0.9rem;
}
.onehis-oneline:nth-child(n+2) {
    margin-top: 10px;
}
.oneline-label {
    width: 120px;
}
.oneline-main {
    display: flex;
    align-items: center;
    overflow: hidden;
    white-space: nowrap;
    flex: 1;
    text-overflow: ellipsis;
}
</style>
