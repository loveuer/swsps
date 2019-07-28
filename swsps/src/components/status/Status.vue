<template>
    <div class="container">
        <div id="echarts-zone"></div>
    </div>
</template>

<script>

// example html;
// https://echarts.baidu.com/examples/editor.html?c=bar-label-rotation

var mockdata = {
    5978: {
        count: 1523,
        using: [
            {name:'u_5978_01', pn:'pn_01', sn:'sn_01', amount:2},
            {name:'u_5978_02', pn:'pn_02', sn:'sn_11', amount:1},
            {name:'u_5978_03', pn:'pn_03', sn:'sn_34', amount:1},
        ],
        fixing: [
            {name:'f_5978_01', pn:'pn_f', sn:'sn_f', amount:1},
        ],
        abandon: [],
    },
    5989: {
        count: 1073,
        using: [
            {name:'u_5989_01', pn:'pn_01', sn:'sn_01', amount:2},
            {name:'u_5989_02', pn:'pn_02', sn:'sn_11', amount:1},
            {name:'u_5989_03', pn:'pn_03', sn:'sn_34', amount:1},
            {name:'u_5989_04', pn:'pn_07', sn:'sn_99', amount:3},
        ],
        fixing: [
            {name:'f_5989_01', pn:'pn_f', sn:'sn_f', amount:1},
        ],
        abandon: [],
    },
    5008: {
        count: 524,
        using: [
            {name:'u_5008_01', pn:'pn_01', sn:'sn_01', amount:2},
            {name:'u_5008_02', pn:'pn_02', sn:'sn_11', amount:1},
        ],
        fixing: [
            {name:'f_5008_01', pn:'pn_f', sn:'sn_f', amount:1},
        ],
        abandon: [],
    },
    5015: {
        count: 817,
        using: [
            {name:'u_5015_01', pn:'pn_01', sn:'sn_01', amount:2},
            {name:'u_5015_02', pn:'pn_02', sn:'sn_11', amount:1},
            {name:'u_5015_03', pn:'pn_03', sn:'sn_34', amount:1},
            {name:'u_5015_04', pn:'pn_07', sn:'sn_99', amount:3},
            {name:'u_5015_05', pn:'pn_09', sn:'sn_02', amount:3},
        ],
        fixing: [
            {name:'f_5015_01', pn:'pn_f', sn:'sn_f', amount:1},
        ],
        abandon: [],
    },
};
import echars from "echarts";

export default {
    mounted() {

        var mainCharts = echars.init(document.getElementById("echarts-zone"));
        var mainOption = {
            color: ['#67C23A', '#409EFF', '#E6A23C', '#F56C6C'],
            tooltip: {
                trigger: 'axis',
                axisPointer: {type: 'shadow'}
            },
            legend: {
                data: ['spare', 'using', 'fixing', 'abandon'],
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
                    name: 'spare',
                    type: 'bar',
                    barGap: 0,
                    data: [1234, 1132, 675, 893],
                },
                {
                    name: 'using',
                    type: 'bar',
                    barGap: 0,
                    data: [16, 12, 4, 9],
                },
                {
                    name: 'fixing',
                    type: 'bar',
                    barGap: 0,
                    data: [3, 2, 0, 1],
                },
                {
                    name: 'abandon',
                    type: 'bar',
                    barGap: 0,
                    data: [4, 8, 0, 0],
                },
            ]
        };
        mainCharts.setOption(mainOption);
    },
};
</script>

<style scoped>
.container {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}
#echarts-zone {
    width: 800px;
    height: 400px;
    background-color: #efefef;
}
</style>
