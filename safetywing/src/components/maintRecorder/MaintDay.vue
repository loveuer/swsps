<template>
    <div>
        <div class="my-menu">
            <loveuer-menu ref="uMenu"></loveuer-menu>
        </div>
        <div style="height:30px;width:100%;zoom:100%;"></div>
        <div style="width:100%;display:flex;justify-content: center;">
            <el-card class="box-card">
                <div slot="header" class="clearfix">
                    <div class="card-title">智胜航安模拟机日检工作单</div>
                </div>
                <div>
                    <table>
                        <tr>
                            <th>模拟机编号:</th>
                            <td>5989</td>
                            <th colspan="2">日期</th>
                            <td colspan="9">2019-03-31</td>
                        </tr>
                        <tr>
                            <th>系统</th>
                            <th>维护项目</th>
                            <th>06</th>
                            <th>08</th>
                            <th>10</th>
                            <th>12</th>
                            <th>14</th>
                            <th>16</th>
                            <th>18</th>
                            <th>20</th>
                            <th>22</th>
                            <th>24</th>
                            <th>02</th>
                        </tr>
                        <tr>
                            <th rowspan="5">MICS</th>
                        </tr>
                        <tr>
                            <td>检查计算机柜里的烟雾温度探测器上的LED灯是否正常闪亮</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                        </tr>
                        <tr>
                            <td>观察计算机房消防控制面板指示灯是否正常工作</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                        </tr>
                        <tr>
                            <td>检查设备柜每个服务器电源是否正常工作,每个设备柜包括SIM PDU的冷却风扇是否工作</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                        </tr>
                        <tr>
                            <td>观察计算机房精密空调,座舱空调是否正常工作,压缩机是否有异常声音,制冷供风是否正常.检查除湿机是否漏水和排水管堵塞</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                        </tr>
                        <tr>
                            <th rowspan="3">ACS</th>
                        </tr>
                        <tr>
                            <td>检查smart UPS上的空调出风口出风量是否正常,冷却风扇是否正常运转</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                        </tr>
                        <tr>
                            <td>检查ACS服务器上的空调出风口出风量是否正常,冷却风扇是否正常运转</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                        </tr>
                        <tr>
                            <th rowspan="5">DMS</th>
                        </tr>
                        <tr>
                            <td>检查MCC柜上的冷却风扇是否正常运转</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                        </tr>
                        <tr>
                            <td>检查MCC柜内的运动腿驱动器上的冷却风扇是否正常运转</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                        </tr>
                        <tr>
                            <td>检查CLM PDU冷却风扇是否正常运转.(正面8个,后中部4个,中间顶面3个)</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                        </tr>
                        <tr>
                            <td>检查运动腿注油口和泄油口及HOME SW是否有液压油渗漏</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                        </tr>
                        
                        <tr>
                            <th rowspan="3">DCLS</th>
                        </tr>
                        <tr>
                            <td>检查电源放大器的冷却风扇是否正常运行</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                        </tr>
                        <tr>
                            <td>检查C/L放大器的冷却风扇是否正常运行</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                            <td>+</td>
                        </tr>
                    </table>
                    <div style="width:100%;margin-top:20px;">
                        <div>
                            <font>备注: 日常巡检2小时执行一次,检查同时记录模拟机相关空调工作情况. 如遇系统检查不正常,请将故障详细情况记录当日simlog,并跟踪处理.</font>
                        </div>
                    </div>
                    <div style="width:100%;margin-top:30px;margin-bottom:30px;">
                        <div>
                            <font>维护人员签字: </font>
                        </div>
                    </div>
                </div>
            </el-card>
        </div>
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";

export default {
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/maintRecorder/day";
    },
    components: {
        "loveuer-menu": loveuerMenu,
    },
};
</script>

<style scoped>
.box-card {
    width:90%;
}
.clearfix {
    text-align: center;
}
.clearfix:before,
.clearfix:after {
    display: table;
    content: "";
}
.clearfix:after {
    clear: both
}
.card-title {
    font-size:1.5rem;
    letter-spacing: 0.25rem;
}
table {
    border-collapse: collapse;
}
table td, table th {
    padding: 5px;
    border: 1px solid #aaa;
}
/* 下面是为了打印的样式 */
@media print {
    .my-menu {
        display: none;
    }
    table {
        font-size: 10pt;
    }
}
</style>

