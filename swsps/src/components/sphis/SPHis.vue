<template>
    <div id="container">
        <sps-menu :path="'/sphis'"></sps-menu>
        <div class="sphis-main">
            <div class="search-line">
                <div>
                    <el-date-picker
                        v-model="filter_date"
                        type="daterange"
                        align="right"
                        unlink-panels
                        range-separator="至"
                        start-placeholder="开始月份"
                        end-placeholder="结束月份"
                        @change="filter_sphis"
                        :picker-options="pickerOptions">
                    </el-date-picker>
                </div>
            </div>
            <div class="sphis-content">
                <div>
                    <el-card v-for="(item, index) of dealed_sphis" :key="index" shadow="hover" style="margin-top:25px;margin-left:50px;">
                        <div slot="header" class="clearfix">
                            <span>{{ item.date }}</span>
                            <span style="margin-left:10px;">{{ item.time }}</span>
                            <span style="margin-left:20px;width:80px;display:inline-block;">{{ item.auth }}</span>
                            <span style="margin-left:20px;display:inline-block;width:200px;">{{ item.last_pos }}</span>
                        </div>
                        <div>
                            <div>{{ item.amount_chg }}</div>
                            <div>{{ item.comment }}</div>
                        </div>
                    </el-card>
                </div>
            </div>
        </div>
        <div style="height:100px;width:100%;zoom:100%;"></div>
    </div>
</template>

<script>
import spsMenu from '../Menu.vue';
import dateformat from 'dateformat';

export default {
    data() {
        return {
            sphis: [],
            filter_date: '',
            pickerOptions: {
                shortcuts: [{
                    text: '最近一周',
                    onClick(picker) {
                        const end = new Date();
                        const start = new Date();
                        start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
                        picker.$emit('pick', [start, end]);
                    }
                }, 
                {
                    text: '最近一个月',
                    onClick(picker) {
                        const end = new Date();
                        const start = new Date();
                        start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
                        picker.$emit('pick', [start, end]);
                    }
                }, 
                {
                    text: '最近三个月',
                    onClick(picker) {
                    const end = new Date();
                    const start = new Date();
                    start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
                    picker.$emit('pick', [start, end]);
                    }
                }],
            },
        };
    },
    methods: {
        showfilterdate: function() {
            console.log(this.filter_date);
        },
        filter_sphis: function() {
            let start = dateformat(this.filter_date[0], 'yyyy-mm-dd');
            let end = dateformat(this.filter_date[1], 'yyyy-mm-dd');
            this.$http.get(`/api/sphis/filter/date?start=${start}&end=${end}`)
                .then(resp => {
                    console.log(resp.data);
                })
                .catch(err => {
                    console.log(err.response);
                    switch (err.response.status) {
                        case 401:
                            this.$router.push('/login');
                            break;
                        case 500:
                            this.$message({showClose: true, message: '发生未知错误', type: 'warning',});
                            break;
                    };
                });
        },

    },
    computed: {
        dealed_sphis: function() {
            let newsphis = [];
            for(let h of this.sphis) {
                newsphis.push({name:h.name, pn:h.pn, sn:h.sn, date:h.time.split('T')[0], time:h.time.split('T')[1].split('+')[0],
                    auth:h.auth, last_pos:h.last_pos, next_pos:h.next_pos,
                    last_sim:h.last_sim, next_sim:h.next_sim,
                    last_sts:h.last_sts, next_sts:h.next_sts,
                    short:h.short, spid:h.spid, amount_chg:h.amount_chg, comment:h.comment});
            };
            return newsphis;
        },
    },
    mounted() {
        this.$http.get('/api/sphis/all/0')
            .then(resp => {
                this.sphis = resp.data;
            })
            .catch(err => {
                console.log(err.response);
                switch (err.response.status) {
                    case 401:
                        this.$router.push('/login');
                        break;
                }
            })
    },
    components: {
        'sps-menu': spsMenu,
    },
};
</script>

<style scoped>
.sphis-main {
    width: 100%;
    margin-top: 30px;
}
.search-line {
    width: 100%;
    display: flex;
}
.search-line > div {
    margin-left: 50px;
}
.sphis-content {
    width: 100%;
    display: flex;
    flex-wrap: wrap;
}
.sphis-content > div {
    width: 70%;
}
</style>
