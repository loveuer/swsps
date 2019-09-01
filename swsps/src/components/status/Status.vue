<template>
    <div class="container">
        <div style="width:100%;">
            <sps-menu :path="'/status'"></sps-menu>
        </div>
        <div id="echarts-zone"></div>
        <div class="detail-table" style="width:100%;margin0-top:100px;">
            <status-table :showingMap="oneTable"></status-table>
        </div>
    </div>
</template>

<script>

// example html;
// https://echarts.baidu.com/examples/editor.html?c=bar-label-rotation

import echars from "echarts";
import spsMenu from "../Menu.vue";
import statusTable from "./Sts_Table.vue";

export default {
    data() {
        return {
            mainCharts: null,
            req: null,
            mainOption: {},
            detailTable: null,
            oneTable: {},
        };
    },
    
    components: {
        'sps-menu': spsMenu,
        'status-table': statusTable,
    },

    methods: {
        'getReq': function() {
            this.$http.get("/api/sps/status")
                .then(resp => {
                    this.req = resp.data;
                    this.init_detailTable();
                    this.init_echarts();
                    this.ebind_click();
                })
                .catch(err => {
                    switch(err.response.status) {
                        case 401:
                            this.router.push("/login");
                            break;
                        case 500:
                            this.$message({showClose: true, message: '服务器发生故障', type: 'warning',});
                            console.log(err.response);
                            break;
                        default:
                            console.log(err.response);
                            break;
                    };
                });
        },
        'init_detailTable': function() {
            this.detailTable = {
                s5978: {
                    total: this.req['5978_num'],
                    using: this.req['5978_using'] ? this.req['5978_using'] : [],
                    fixing: this.req['5978_fixing'] ? this.req['5978_fixing'] : [],
                    abandon: this.req['5978_abandon'] ? this.req['5978_abandon'] : [],
                },
                s5989: {
                    total: this.req['5989_num'],
                    using: this.req['5989_using'] ? this.req['5989_using'] : [],
                    fixing: this.req['5989_fixing'] ? this.req['5989_fixing'] : [],
                    abandon: this.req['5989_abandon'] ? this.req['5989_abandon'] : [],
                },
                s5008: {
                    total: this.req['5008_num'],
                    using: this.req['5008_using'] ? this.req['5008_using'] : [],
                    fixing: this.req['5008_fixing'] ? this.req['5008_fixing'] : [],
                    abandon: this.req['5008_abandon'] ? this.req['5008_abandon'] : [],
                },
                s5015: {
                    total: this.req['5015_num'],
                    using: this.req['5015_using'] ? this.req['5015_using'] : [],
                    fixing: this.req['5015_fixing'] ? this.req['5015_fixing'] : [],
                    abandon: this.req['5015_abandon'] ? this.req['5015_abandon'] : [],
                },
            };
        },
        'init_echarts': function() {
            this.mainOption = {
                color: ['#67C23A', '#409EFF', '#E6A23C', '#F56C6C'],
                tooltip: {
                    trigger: 'axis',
                    axisPointer: {type: 'shadow'}
                },
                legend: {
                    data: ['备件', '使用', '维修', '废弃'],
                    selected: {
                        '备件': false,
                        '使用': true,
                        '维修': true,
                        '废弃': true,
                    },
                },
                toolbox: {
                    show: true,
                    orient: 'right',
                    left: 'right',
                    top: 'center',
                    feature: {
                        mark: {show: true},
                        dataView: {show: true, readOnly: false},
                        magicType: {show: true, type: ['line', 'bar', 'stack', 'tiled']},
                        restore: {show: true},
                        saveAsImage: {show: true}
                    },
                },
                calculable: true,
                xAxis: [
                    {
                        type: 'category',
                        axisTick: {show: false},
                        data: ['5978', '5989', '5008', '5015'],
                    }
                ],
                yAxis: [
                    {
                        type: 'value',
                    }
                ],
                series: [
                    {
                        name: '备件',
                        type: 'bar',
                        barGap: 0,
                        // data: [1234, 1132, 675, 893],
                        data: [
                            this.detailTable.s5978.total - this.detailTable.s5978.using.length - this.detailTable.s5978.fixing.length - this.detailTable.s5978.abandon.length,
                            this.detailTable.s5989.total - this.detailTable.s5989.using.length - this.detailTable.s5989.fixing.length - this.detailTable.s5989.abandon.length,
                            this.detailTable.s5008.total - this.detailTable.s5008.using.length - this.detailTable.s5008.fixing.length - this.detailTable.s5008.abandon.length,
                            this.detailTable.s5015.total - this.detailTable.s5015.using.length - this.detailTable.s5015.fixing.length - this.detailTable.s5015.abandon.length,
                        ],
                    },
                    {
                        name: '使用',
                        type: 'bar',
                        barGap: 0,
                        // data: [16, 12, 4, 9],
                        data: [
                            this.detailTable.s5978.using.length, this.detailTable.s5989.using.length, this.detailTable.s5008.using.length, this.detailTable.s5015.using.length,
                        ],
                    },
                    {
                        name: '维修',
                        type: 'bar',
                        barGap: 0,
                        // data: [3, 2, 0, 1],
                        data: [
                            this.detailTable.s5978.fixing.length, this.detailTable.s5989.fixing.length, this.detailTable.s5008.fixing.length, this.detailTable.s5015.fixing.length,
                        ],
                    },
                    {
                        name: '废弃',
                        type: 'bar',
                        barGap: 0,
                        // data: [4, 8, 0, 0],
                        data: [
                            this.detailTable.s5978.abandon.length, this.detailTable.s5989.abandon.length, this.detailTable.s5008.abandon.length, this.detailTable.s5015.abandon.length,
                        ],
                    },
                ],
            };
            this.mainCharts = echars.init(document.getElementById("echarts-zone"));
            this.mainCharts.setOption(this.mainOption);
        },
        'ebind_click': function() {
            this.mainCharts.on('click', (params) => {
                switch (params.name) {
                    case '5978':
                        this.oneTable = this.detailTable.s5978;
                        break;
                    case '5989':
                        this.oneTable = this.detailTable.s5989;
                        break;
                    case '5008':
                        this.oneTable = this.detailTable.s5008;
                        break;
                    case '5015':
                        this.oneTable = this.detailTable.s5015;
                        break;
                    default:
                        break;
                };
            });
        },
    },
    
    mounted() {
        this.getReq();
    },
};
</script>

<style scoped>
.container {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
}
#echarts-zone {
    width: 800px;
    height: 400px;
    background-color: #efefef;
    margin-top: 50px;
    border-radius: 5px;
}
.detail-table {
    margin-top: 100px;
    width: 100%;
    display: flex;
    justify-content: center;
}
</style>
