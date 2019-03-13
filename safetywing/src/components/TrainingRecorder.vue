<template>
    <div id="training-recoder">
        <div class="tr-main">
            <el-form ref="form" :model="form" label-width="100px">

                <el-form-item label="训练日期:">
                        <el-date-picker type="date" placeholder="选择日期" v-model="form.date" style="width: 100%;"></el-date-picker>
                </el-form-item>

                <el-form-item label="模拟机编号:">
                    <el-select v-model="form.sim" placeholder="" style="width: 100%;">
                        <el-option label="5978" value="5978"></el-option>
                        <el-option label="5989" value="5989"></el-option>
                        <el-option label="5008" value="5008"></el-option>
                        <el-option label="5015" value="5015"></el-option>
                    </el-select>
                </el-form-item>

                <el-form-item label="航空公司:">
                    <el-select v-model="form.airline" placeholder="" style="width: 100%;">
                        <el-option label="国　　航" value="国　　航"></el-option>
                        <el-option label="藏　　航" value="藏　　航"></el-option>
                        <el-option label="重庆航空" value="重庆航空"></el-option>
                        <el-option label="成都航空" value="成都航空"></el-option>
                        <el-option label="东方航空" value="东方航空"></el-option>
                        <el-option label="试飞中心" value="试飞中心"></el-option>
                        <el-option label="华夏航空" value="华夏航空"></el-option>
                        <el-option label="藏航理论" value="藏航理论"></el-option>
                        <el-option label="天府航空" value="天府航空"></el-option>
                        <el-option label="其　　他" value="其　　他"></el-option>
                    </el-select>
                </el-form-item>

                <el-form-item label="训练场次:">
                    <el-select v-model="form.airline" placeholder="" style="width: 100%;">
                        <el-option label="第 1 场" value="1"></el-option>
                        <el-option label="第 2 场" value="2"></el-option>
                        <el-option label="第 3 场" value="3"></el-option>
                        <el-option label="第 4 场" value="4"></el-option>
                        <el-option label="第 5 场" value="5"></el-option>
                        <el-option label="第 6 场" value="6"></el-option>
                        <el-option label="第 7 场" value="7"></el-option>
                        <el-option label="第 8 场" value="8"></el-option>
                        <el-option label="第 9 场" value="9"></el-option>
                    </el-select>
                </el-form-item>

                <el-form-item label="训练时间:">
                    <el-col :span="8">
                        <el-select v-model="form.start" placeholder="" style="width: 100%;">
                            <el-option label="00 : 00" value="0"></el-option>
                            <el-option label="01 : 00" value="1"></el-option>
                            <el-option label="02 : 00" value="2"></el-option>
                            <el-option label="03 : 00" value="3"></el-option>
                            <el-option label="04 : 00" value="4"></el-option>
                            <el-option label="05 : 00" value="5"></el-option>
                            <el-option label="06 : 00" value="6"></el-option>
                            <el-option label="07 : 00" value="7"></el-option>
                            <el-option label="08 : 00" value="8"></el-option>
                            <el-option label="09 : 00" value="9"></el-option>
                            <el-option label="10 : 00" value="10"></el-option>
                            <el-option label="11 : 00" value="11"></el-option>
                            <el-option label="12 : 00" value="12"></el-option>
                            <el-option label="13 : 00" value="13"></el-option>
                            <el-option label="14 : 00" value="14"></el-option>
                            <el-option label="15 : 00" value="15"></el-option>
                            <el-option label="16 : 00" value="16"></el-option>
                            <el-option label="17 : 00" value="17"></el-option>
                            <el-option label="18 : 00" value="18"></el-option>
                            <el-option label="19 : 00" value="19"></el-option>
                            <el-option label="20 : 00" value="20"></el-option>
                            <el-option label="21 : 00" value="21"></el-option>
                            <el-option label="22 : 00" value="22"></el-option>
                            <el-option label="23 : 00" value="23"></el-option>
                            <el-option label="24 : 00" value="24"></el-option>
                        </el-select>
                    </el-col>
                    <el-col class="line" :span="1" style="text-align:center;">-</el-col>
                    <el-col :span="8">
                        <el-select v-model="form.end" placeholder="" style="width: 100%;">
                            <el-option label="00 : 00" value="0"></el-option>
                            <el-option label="01 : 00" value="1"></el-option>
                            <el-option label="02 : 00" value="2"></el-option>
                            <el-option label="03 : 00" value="3"></el-option>
                            <el-option label="04 : 00" value="4"></el-option>
                            <el-option label="05 : 00" value="5"></el-option>
                            <el-option label="06 : 00" value="6"></el-option>
                            <el-option label="07 : 00" value="7"></el-option>
                            <el-option label="08 : 00" value="8"></el-option>
                            <el-option label="09 : 00" value="9"></el-option>
                            <el-option label="10 : 00" value="10"></el-option>
                            <el-option label="11 : 00" value="11"></el-option>
                            <el-option label="12 : 00" value="12"></el-option>
                            <el-option label="13 : 00" value="13"></el-option>
                            <el-option label="14 : 00" value="14"></el-option>
                            <el-option label="15 : 00" value="15"></el-option>
                            <el-option label="16 : 00" value="16"></el-option>
                            <el-option label="17 : 00" value="17"></el-option>
                            <el-option label="18 : 00" value="18"></el-option>
                            <el-option label="19 : 00" value="19"></el-option>
                            <el-option label="20 : 00" value="20"></el-option>
                            <el-option label="21 : 00" value="21"></el-option>
                            <el-option label="22 : 00" value="22"></el-option>
                            <el-option label="23 : 00" value="23"></el-option>
                            <el-option label="24 : 00" value="24"></el-option>
                        </el-select>
                    </el-col>
                    <el-col class="line" :span="2" style="text-align:center;float:right;">
                        <span v-text="form.totalTime"></span>
                        <span> Hs</span>
                    </el-col>
                </el-form-item>

                <el-form-item label="教员姓名:">
                    <el-row class="demo-autocomplete" style="width: 100%;">
                        <el-col style="width: 100%;">
                            <el-autocomplete
                                class="inline-input"
                                v-model="form.instructor"
                                :fetch-suggestions="fetchMostInstructors"
                                placeholder=""
                                @select="handleSelect"
                            ></el-autocomplete>
                        </el-col>
                    </el-row>
                </el-form-item>

            </el-form>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            form: {
                date:　"",

                totalTime: 0,
                instructorName: "",
            }
        };
    },
    methods: {
        fetchMostInstructors: () => {

        },
    },
}
</script>

<style scoped>
#training-recoder{
    height: 100%;
    width: 100%;
    display: flex;
}
.tr-main{
    margin-top: 30px;
    margin-left: 30px;
}
</style>

